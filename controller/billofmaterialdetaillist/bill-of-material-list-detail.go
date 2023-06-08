package billofmaterialdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	billofmaterialdetail "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/billofmaterialdetaillist"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type BillOfMaterialDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewBillOfMaterialDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *BillOfMaterialDetailListCtrl {
	return &BillOfMaterialDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *BillOfMaterialDetailListCtrl) BillOfMaterialDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractBillOfMaterialDetailListParam(msg)
	reqKey, err := getRequestKey(msg.Data())
	if err != nil {
		return xerrors.Errorf("reqKey error: %w", err)
	}
	sID, err := getSessionID(msg.Data())
	if err != nil {
		return xerrors.Errorf("session ID error: %w", err)
	}
	cacheResult := RedisCacheApiName{
		"redisCacheApiName": map[string]interface{}{},
	}
	defer func() {
		b, _ := json.Marshal(cacheResult)
		err = c.cache.Set(c.ctx, reqKey, b, 0)
		if err != nil {
			c.log.Error("cache set error: %w", err)
		}
	}()

	bomRes, err := c.billOfMaterialRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("billOfMaterialRequest error: %w", err)
	}
	if bomRes.Message.Header == nil || len(*bomRes.Message.Header) == 0 {
		c.finEmptyProcess(params, reqKey, "BillOfMaterialDetailList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	err = c.addHeaderInfo(&params.Params, sID, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}

	c.fin(params, bomRes, reqKey, "BillOfMaterialDetailList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *BillOfMaterialDetailListCtrl) billOfMaterialRequest(
	params *dpfm_api_input_reader.BillOfMaterialDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BillOfMaterialRes, error) {
	defer recovery(c.log)
	bomReq := billofmaterialdetail.CreateBillOfMaterialReq(params, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", bomReq, sID, reqKey, "BillOfMaterialDetailList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BillOfMaterial cache set error: %w", err)
	}
	bomRes, err := apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BillOfMaterialDetail response parse error: %w", err)
	}
	return bomRes, nil
}

func (c *BillOfMaterialDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.BillOfMaterialDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	//201@gmail.com/billOfMaterial/list/user=BillToParty/businessPartner=201/headerBillOfMaterial=false/headerValidityStartDate=false/headerOwnerPlantName=false/headerProductDescription=false=/headerPaymentBlockStatus=false/BillOfMaterialList
	key := fmt.Sprintf(`%s/billOfMaterial/list/user=%s/businessPartner=%d/headerBillOfMaterial=%d/headerValidityStartDate=%d/headerOwnerPlantName=%d/headerProductDescription=%d/headerPaymentBlockStatus=%v/BillOfMaterialList`,
		params.UserID, params.User, params.BusinessPartner, params.BillOfMaterial, params.ValidityStartDate, params.OwnerPlantName, params.ProductDescription, false)
	api := "BillOfMaterialDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	iList := dpfm_api_output_formatter.BillOfMaterialList{}
	err = json.Unmarshal(b, &iList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range iList.BillOfMaterials {
		if v.BillOfMaterial == params.BillOfMaterial {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *BillOfMaterialDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
	resFunc := c.rmq.SessionRequest(queue, req, sID)
	res := resFunc()
	if res == nil {
		return nil, xerrors.Errorf("receive nil response")
	}
	// redisKey := strings.Join([]string{url, api}, "/")
	// err := c.cache.Set(c.ctx, redisKey, res.Raw(), 1*time.Hour)
	// if err != nil {
	// 	return nil, xerrors.Errorf("cache set error: %w", err)
	// }
	// (*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return res, nil
}

func getRequestKey(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawReqID, ok := m["ui_key_function_url"]
	if !ok {
		return "", xerrors.Errorf("keyName not included")
	}

	return fmt.Sprintf("%v", rawReqID), nil
}

func getSessionID(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawSID, ok := m["runtime_session_id"]
	if !ok {
		return "", xerrors.Errorf("runtime_session_id not included")
	}

	return fmt.Sprintf("%v", rawSID), nil
}

func (c *BillOfMaterialDetailListCtrl) fin(
	params *dpfm_api_input_reader.BillOfMaterialDetailList,
	bomRes *apiresponses.BillOfMaterialRes,
	url,
	api string,
	setFlag *RedisCacheApiName,
) error {
	key := fmt.Sprintf(`%s/billOfMaterial/list/user=%s/businessPartner=%d/headerBillOfMaterial=%d/headerValidityStartDate=%d/headerOwnerPlantName=%d/headerProductDescription=%d/headerPaymentBlockStatus=%v/BillOfMaterialList`,
		params.Params.UserID, params.Params.User, params.Params.BusinessPartner, params.Params.BillOfMaterial, params.Params.ValidityStartDate, params.Params.OwnerPlantName, params.Params.ProductDescription, false)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	bomList := dpfm_api_output_formatter.BillOfMaterialList{}
	err = json.Unmarshal(b, &bomList)
	if err != nil {
		return err
	}
	idx := -1
	for i, v := range bomList.BillOfMaterials {
		if v.BillOfMaterial == params.Params.BillOfMaterial {
			idx = i
			break
		}
	}

	header := dpfm_api_output_formatter.BillOfMaterialDetailHeader{
		Index: idx,
		Key:   key,
	}

	details := make([]dpfm_api_output_formatter.BillOfMaterialDetail, 0, len(*bomRes.Message.Item))
	for _, v := range *bomRes.Message.Item {
		details = append(details, dpfm_api_output_formatter.BillOfMaterialDetail{
			ComponentProduct:          *v.ComponentProduct,
			BillOfMaterialItemText:    *v.BillOfMaterialItemText,
			StockConfirmationPlant:    *v.StockConfirmationPlant,
			BOMItemQuantityInBaseUnit: *v.BOMItemQuantityInBaseUnit,
			BOMItemBaseUnit:           *v.BOMItemBaseUnit,
			ValidityStartDate:         *v.ValidityStartDate,
			IsMarkedForDeletion:       *v.IsMarkedForDeletion,
		})
	}

	data := dpfm_api_output_formatter.BillOfMaterialDetailList{
		BillOfMaterialDetailHeader: header,
		BillOfMaterialDetail:       details,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	redisKey := strings.Join([]string{url, api}, "/")
	// redisKey := strings.Join([]string{url, api, params.User}, "/")
	b, _ = json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *BillOfMaterialDetailListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.BillOfMaterialDetailList{
		BillOfMaterialDetail: make([]dpfm_api_output_formatter.BillOfMaterialDetail, 0),
	}

	redisKey := strings.Join([]string{url, api}, "/")
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func extractBillOfMaterialDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.BillOfMaterialDetailList {
	data := dpfm_api_input_reader.ReadBillOfMaterialDetailList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func billOfMaterialAsc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for _, i := range ids {
		sli = append(sli, d[i])
	}
	return sli
}

func billOfMaterialDesc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for i := len(ids) - 1; i >= 0; i-- {
		sli = append(sli, d[ids[i]])
	}
	return sli
}

func (c *BillOfMaterialDetailListCtrl) Log(args ...interface{}) {
	for _, v := range args {
		b, _ := json.Marshal(v)
		c.log.Error("%s", string(b))
	}
}

func recovery(l *logger.Logger) {
	if e := recover(); e != nil {
		l.Error("%+v", e)
		return
	}
}
