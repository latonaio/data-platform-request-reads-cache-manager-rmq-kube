package pricemasterdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	pricemasterdetaillist "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/pricemasterdetaillist"
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

// CreatePriceMasterDetailItemsReq
type PriceMasterDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewPriceMasterDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *PriceMasterDetailListCtrl {
	return &PriceMasterDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *PriceMasterDetailListCtrl) PriceMasterDetailList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractPriceMasterDetailListParam(msg)
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

	pmdRes, err := c.priceMasterDetailRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	drRes, err := c.descriptionRequest(&params.Params, pmdRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	err = c.addHeaderInfo(&params.Params, sID, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}

	c.fin(params, pmdRes, drRes, reqKey, "PriceMasterDetailList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *PriceMasterDetailListCtrl) descriptionRequest(
	params *dpfm_api_input_reader.PriceMasterDetailListParams,
	pmdRes *apiresponses.PriceMasterDetailRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	drReq := pricemasterdetaillist.CreateDescriptionReq(params, pmdRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", drReq, sID, reqKey, "ProductionVersionList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("description cache set error: %w", err)
	}
	drRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}

	return drRes, nil
}

func (c *PriceMasterDetailListCtrl) priceMasterDetailRequest(
	params *dpfm_api_input_reader.PriceMasterDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.PriceMasterDetailRes, error) {
	defer recovery(c.log)
	pmdReq := pricemasterdetaillist.CreatePriceMasterDetailReq(params, sID, c.log)
	res, err := c.request("data-platform-api-price-master-reads-queue", pmdReq, sID, reqKey, "PriceMasterDetail", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("price master cache set error: %w", err)
	}
	pmdRes, err := apiresponses.CreatePriceMasterDetailRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Price Master master response parse error: %w", err)
	}
	return pmdRes, nil
}

func (c *PriceMasterDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.PriceMasterDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	//101@gmail.com/priceMaster/list/user=buyer
	//101@gmail.com/priceMaster/list/user=Buyer/PriceMasterList
	key := fmt.Sprintf(`%s/priceMaster/list/user=%s/PriceMasterList`,
		params.UserID, params.User)
	api := "PriceMasterDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	pList := dpfm_api_output_formatter.PriceMasterList{}
	err = json.Unmarshal(b, &pList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range pList.PriceMasters {
		if v.SupplyChainRelationshipID == params.SupplyChainRelationshipID {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *PriceMasterDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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
	// 	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return res, nil
}

func extractPriceMasterDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.PriceMasterDetailList {
	data := dpfm_api_input_reader.ReadPriceMasterDetailList(msg)
	return data
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

type RedisCacheApiName map[string]map[string]interface{}

func (c *PriceMasterDetailListCtrl) fin(
	params *dpfm_api_input_reader.PriceMasterDetailList,
	pmRes *apiresponses.PriceMasterDetailRes,
	pmasRes *apiresponses.ProductMasterRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	type bpIDRel struct {
		orderID        int
		sellerID       int
		buyerID        int
		deliveryStatus string
	}

	//101@gmail.com/priceMaster/list/user=buyer
	key := fmt.Sprintf(`%s/priceMaster/list/user=%s/PriceMasterList`,
		params.Params.UserID, params.Params.User)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	pList := dpfm_api_output_formatter.PriceMasterList{}
	err = json.Unmarshal(b, &pList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range pList.PriceMasters {
		if v.SupplyChainRelationshipID == params.Params.SupplyChainRelationshipID {
			idx = i
			break
		}
	}

	header := dpfm_api_output_formatter.PriceMasterDetailHeader{
		Index: idx,
		Key:   key,
	}

	descriptionMapper := map[string]apiresponses.ProductDescByBP{}
	for _, v := range *pmasRes.Message.ProductDescByBP {
		descriptionMapper[v.Product] = v
	}

	details := make([]dpfm_api_output_formatter.PriceMasterDetail, 0, len(pmRes.Message.PriceMasterDetail))
	for _, v := range pmRes.Message.PriceMasterDetail {
		details = append(details, dpfm_api_output_formatter.PriceMasterDetail{
			Product:                   *v.Product,
			ProductionDescription:     *descriptionMapper[*v.Product].ProductDescription,
			ConditionRateValue:        *v.ConditionRateValue,
			ConditionScaleQuantity:    v.ConditionScaleQuantity,
			ConditionRateValueUnit:    v.ConditionRateValueUnit,
			ConditionType:             v.ConditionType,
			ConditionCurrency:         v.ConditionCurrency,
			ConditionRecord:           v.ConditionRecord,
			ConditionSequentialNumber: v.ConditionSequentialNumber,
			IsMarkedForDeletion:       v.IsMarkedForDeletion,
		})
	}

	data := dpfm_api_output_formatter.PriceMasterDetailList{
		PriceMasterDetailHeader: header,
		PriceMasterDetail:       details,
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

func (c *PriceMasterDetailListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.DeliveryDocumentList{
		DeliveryDocuments: make([]dpfm_api_output_formatter.DeliveryDocument, 0),
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

func orderAsc[T any](d map[int]T) []T {
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

func orderDesc[T any](d map[int]T) []T {
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

func (c *PriceMasterDetailListCtrl) Log(args ...interface{}) {
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
