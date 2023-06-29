package productionversionlist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/productionversionlist"
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

type ProductionVersionListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewProductionVersionListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *ProductionVersionListCtrl {
	return &ProductionVersionListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *ProductionVersionListCtrl) ProductionVersionList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractProductionVersionListParam(msg)
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

	pvRes, err := c.productionVersionRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("productionVersionRequest error: %w", err)
	}
	if pvRes.Message.Header == nil || len(*pvRes.Message.Header) == 0 {
		c.finEmptyProcess(params, reqKey, "ProductionVersionList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	plRes, err := c.plantRequest(&params.Params, pvRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("plantRequest error: %w", err)
	}

	drRes, err := c.descriptionRequest(&params.Params, pvRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	pvdRes, err := c.productionVersionDocRequest(&params.Params, pvRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	c.fin(params, pvRes, drRes, pvdRes, plRes, reqKey, "ProductionVersionList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *ProductionVersionListCtrl) productionVersionDocRequest(
	params *dpfm_api_input_reader.ProductionVersionListParams,
	pvRes *apiresponses.ProductionVersionRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	pvReq := productionversionlist.CreateProductMasterDocReq(params, pvRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pvReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pvdRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("product response parse error: %w", err)
	}
	return pvdRes, nil
}

func (c *ProductionVersionListCtrl) productionVersionRequest(
	params *dpfm_api_input_reader.ProductionVersionListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductionVersionRes, error) {
	defer recovery(c.log)
	pvReq := productionversionlist.CreateProductionVersionReq(params, sID, c.log)
	res, err := c.request("data-platform-api-production-version-reads-queue", pvReq, sID, reqKey, "ProductionVersionList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductionVersion cache set error: %w", err)
	}
	pvRes, err := apiresponses.CreateProductionVersionRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductionVersion response parse error: %w", err)
	}
	return pvRes, nil
}

func (c *ProductionVersionListCtrl) descriptionRequest(
	params *dpfm_api_input_reader.ProductionVersionListParams,
	pvRes *apiresponses.ProductionVersionRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	drReq := productionversionlist.CreateDescriptionReq(params, pvRes, sID, c.log)
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

func (c *ProductionVersionListCtrl) plantRequest(
	params *dpfm_api_input_reader.ProductionVersionListParams,
	pvRes *apiresponses.ProductionVersionRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := productionversionlist.CreatePlantReq(params, pvRes, sID, c.log)
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

func (c *ProductionVersionListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *ProductionVersionListCtrl) fin(
	params *dpfm_api_input_reader.ProductionVersionList,
	pvRes *apiresponses.ProductionVersionRes,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	plRes *apiresponses.PlantRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	descriptionMapper := map[string]apiresponses.ProductDescByBP{}
	for _, v := range *pmRes.Message.ProductDescByBP {
		descriptionMapper[v.Product] = v
	}

	imgMapper := map[string]apiresponses.PMDHeaderDoc{}
	for _, v := range *pmdRes.Message.HeaderDoc {
		imgMapper[v.Product] = v
	}

	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	data := dpfm_api_output_formatter.ProductionVersionList{}
	for _, v := range *pvRes.Message.Header {
		img := &dpfm_api_output_formatter.ProductImage{}
		if i, ok := imgMapper[v.Product]; ok {
			img = &dpfm_api_output_formatter.ProductImage{
				BusinessPartnerID: i.DocIssuerBusinessPartner,
				DocID:             i.DocID,
				FileExtension:     i.FileExtension,
			}
		}

		data.ProductionVersions = append(data.ProductionVersions,
			dpfm_api_output_formatter.ProductionVersion{
				Product:             &v.Product,
				ProductionVersion:   v.ProductionVersion,
				ProductDescription:  descriptionMapper[v.Product].ProductDescription,
				OwnerPlant:          &v.OwnerPlant,
				OwnerPlantName:      plantMapper[v.OwnerPlant].PlantName,
				BillOfMaterial:      &v.BillOfMaterial,
				Operations:          &v.Operations,
				ValidityStartDate:   v.ValidityStartDate,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
				Images: dpfm_api_output_formatter.Images{
					ProductionVersion: img,
				},
			},
		)
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	redisKey := strings.Join([]string{url, api}, "/")
	// redisKey := strings.Join([]string{url, api, params.User}, "/")
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *ProductionVersionListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.ProductionVersionList{
		ProductionVersions: make([]dpfm_api_output_formatter.ProductionVersion, 0),
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

func extractProductionVersionListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.ProductionVersionList {
	data := dpfm_api_input_reader.ReadProductionVersionList(msg)
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

func (c *ProductionVersionListCtrl) Log(args ...interface{}) {
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
