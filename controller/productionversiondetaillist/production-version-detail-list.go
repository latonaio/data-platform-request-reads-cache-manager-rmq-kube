package productionversiondetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/productionversiondetaillist"
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

type ProductionVersionDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewProductionVersionDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *ProductionVersionDetailListCtrl {
	return &ProductionVersionDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *ProductionVersionDetailListCtrl) ProductionVersionDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractProductionVersionDetailListParam(msg)
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

	pvdRes, err := c.productionVersionDetailRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("productionVersionRequest error: %w", err)
	}
	/*if pvdRes.Message.Header == nil || len(*pvdRes.Message.Header) == 0 {
		c.finEmptyProcess(params, reqKey, "ProductionVersionDetailList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}*/

	plRes, err := c.plantRequest(&params.Params, pvdRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("plantRequest error: %w", err)
	}

	drRes, err := c.descriptionRequest(&params.Params, pvdRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	err = c.addHeaderInfo(&params.Params, sID, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}

	c.fin(params, pvdRes, drRes, plRes, reqKey, "ProductionVersionDetailList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *ProductionVersionDetailListCtrl) productionVersionDetailRequest(
	params *dpfm_api_input_reader.ProductionVersionDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductionVersionRes, error) {
	defer recovery(c.log)
	pvdReq := productionversiondetaillist.CreateProductionVersionDetailReq(params, sID, c.log)
	res, err := c.request("data-platform-api-production-version-reads-queue", pvdReq, sID, reqKey, "ProductionVersionDetailList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductionVersion cache set error: %w", err)
	}
	pvdRes, err := apiresponses.CreateProductionVersionRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductionVersion response parse error: %w", err)
	}
	return pvdRes, nil
}

func (c *ProductionVersionDetailListCtrl) descriptionRequest(
	params *dpfm_api_input_reader.ProductionVersionDetailListParams,
	pvRes *apiresponses.ProductionVersionRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	drReq := productionversiondetaillist.CreateDescriptionReq(params, pvRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", drReq, sID, reqKey, "ProductionVersionDetailList", setFlag)
	//c.log.JsonParseOut(drReq)
	if err != nil {
		return nil, xerrors.Errorf("description cache set error: %w", err)
	}
	drRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}

	return drRes, nil
}

func (c *ProductionVersionDetailListCtrl) plantRequest(
	params *dpfm_api_input_reader.ProductionVersionDetailListParams,
	pvRes *apiresponses.ProductionVersionRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := productionversiondetaillist.CreatePlantReq(params, pvRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", plReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Plant cache set error: %w", err)
	}
	plRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Plant response parse error: %w", err)
	}
	return plRes, nil
}

func (c *ProductionVersionDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.ProductionVersionDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	//201@gmail.com/productionVersion/list/user=OwnerProductionPlantBusinessPartner/headerIsMarkedForDeletion=false/ProductionVersionList
	key := fmt.Sprintf(`%s/productionVersion/list/user=%s/ProductionVersionList`,
		params.UserID, params.User)
	api := "ProductionVersionDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	pvList := dpfm_api_output_formatter.ProductionVersionList{}
	err = json.Unmarshal(b, &pvList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range pvList.ProductionVersions {
		if v.ProductionVersion == *params.ProductionVersion {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *ProductionVersionDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *ProductionVersionDetailListCtrl) fin(
	params *dpfm_api_input_reader.ProductionVersionDetailList,
	pvRes *apiresponses.ProductionVersionRes,
	pmRes *apiresponses.ProductMasterRes,
	plRes *apiresponses.PlantRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	key := fmt.Sprintf(`%s/productionVersion/list/user=%s/ProductionVersionList`,
		params.Params.UserID, params.Params.User)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	pvList := dpfm_api_output_formatter.ProductionVersionList{}
	err = json.Unmarshal(b, &pvList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range pvList.ProductionVersions {
		if v.ProductionVersion == *params.Params.ProductionVersion {
			idx = i
			break
		}
	}
	header := dpfm_api_output_formatter.ProductionVersionDetailHeader{
		Index: idx,
		Key:   key,
	}

	descriptionMapper := map[string]apiresponses.ProductDescByBP{}
	for _, v := range *pmRes.Message.ProductDescByBP {
		descriptionMapper[v.Product] = v
	}

	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	details := make([]dpfm_api_output_formatter.ProductionVersionDetail, 0, len(*pvRes.Message.Item))
	for _, v := range *pvRes.Message.Item {
		details = append(details, dpfm_api_output_formatter.ProductionVersionDetail{
			ProductionVersion:     &v.ProductionVersion,
			ProductionVersionItem: &v.ProductionVersionItem,
			Product:               &v.Product,
			ProductDescription:    descriptionMapper[v.Product].ProductDescription,
			Plant:                 &v.Plant,
			PlantName:             plantMapper[v.Plant].PlantName,
			BillOfMaterial:        &v.BillOfMaterial,
			Operations:            &v.Operations,
			ValidityStartDate:     v.ValidityStartDate,
			IsMarkedForDeletion:   v.IsMarkedForDeletion,
		})
	}

	data := dpfm_api_output_formatter.ProductionVersionDetailList{
		Header:             header,
		ProductionVersions: details,
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

func (c *ProductionVersionDetailListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.ProductionVersionDetailList{
		ProductionVersions: make([]dpfm_api_output_formatter.ProductionVersionDetail, 0),
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

func extractProductionVersionDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.ProductionVersionDetailList {
	data := dpfm_api_input_reader.ReadProductionVersionDetailList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func productionVersionAsc[T any](d map[int]T) []T {
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

func productionVersionDesc[T any](d map[int]T) []T {
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

func (c *ProductionVersionDetailListCtrl) Log(args ...interface{}) {
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
