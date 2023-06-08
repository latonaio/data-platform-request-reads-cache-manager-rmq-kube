package billofmateriallist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/billofmateriallist"
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

type BillOfMaterialListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewBillOfMaterialListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *BillOfMaterialListCtrl {
	return &BillOfMaterialListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *BillOfMaterialListCtrl) BillOfMaterialList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractBillOfMaterialListParam(msg)
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
		c.finEmptyProcess(params, reqKey, "BillOfMaterialList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	drRes, err := c.descriptionRequest(&params.Params, bomRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	pvdRes, err := c.billOfMaterialDocRequest(&params.Params, bomRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	plRes, err := c.plantRequest(&params.Params, bomRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("plantRequest error: %w", err)
	}

	c.fin(params, bomRes, drRes, pvdRes, plRes, reqKey, "BillOfMaterialList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *BillOfMaterialListCtrl) billOfMaterialDocRequest(
	params *dpfm_api_input_reader.BillOfMaterialListParams,
	pvRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	pvReq := billofmateriallist.CreateProductMasterDocReq(params, pvRes, sID, c.log)
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

func (c *BillOfMaterialListCtrl) billOfMaterialRequest(
	params *dpfm_api_input_reader.BillOfMaterialListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BillOfMaterialRes, error) {
	defer recovery(c.log)
	pvReq := billofmateriallist.CreateBillOfMaterialReq(params, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", pvReq, sID, reqKey, "BillOfMaterialList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BillOfMaterial cache set error: %w", err)
	}
	pvRes, err := apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BillOfMaterial response parse error: %w", err)
	}
	return pvRes, nil
}

func (c *BillOfMaterialListCtrl) plantRequest(
	params *dpfm_api_input_reader.BillOfMaterialListParams,
	pvRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := billofmateriallist.CreatePlantReq(params, pvRes, sID, c.log)
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

func (c *BillOfMaterialListCtrl) descriptionRequest(
	params *dpfm_api_input_reader.BillOfMaterialListParams,
	pvRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	drReq := billofmateriallist.CreateDescriptionReq(params, pvRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", drReq, sID, reqKey, "BillOfMaterialList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("description cache set error: %w", err)
	}
	drRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}

	return drRes, nil
}

func (c *BillOfMaterialListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *BillOfMaterialListCtrl) fin(
	params *dpfm_api_input_reader.BillOfMaterialList,
	bomRes *apiresponses.BillOfMaterialRes,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	plRes *apiresponses.PlantRes,
	url,
	api string,
	setFlag *RedisCacheApiName,
) error {
	descriptionMapper := map[string]apiresponses.ProductDescByBP{}
	for _, v := range *pmRes.Message.ProductDescByBP {
		descriptionMapper[v.Product] = v
	}

	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	data := dpfm_api_output_formatter.BillOfMaterialList{}

	for _, v := range *bomRes.Message.Header {
		img := &dpfm_api_output_formatter.ProductImage{}

		for _, pmdResHeaderV := range *pmdRes.Message.HeaderDoc {
			if &pmdResHeaderV.DocIssuerBusinessPartner != nil &&
				pmdResHeaderV.DocIssuerBusinessPartner == *v.OwnerBusinessPartner &&
				&v.Product != nil &&
				pmdResHeaderV.Product == *v.Product {
				img = &dpfm_api_output_formatter.ProductImage{
					BusinessPartnerID: pmdResHeaderV.DocIssuerBusinessPartner,
					DocID:             pmdResHeaderV.DocID,
					FileExtension:     pmdResHeaderV.FileExtension,
				}
			}
		}

		data.BillOfMaterials = append(data.BillOfMaterials,
			dpfm_api_output_formatter.BillOfMaterial{
				Product:             v.Product,
				BillOfMaterial:      v.BillOfMaterial,
				ProductDescription:  descriptionMapper[*v.Product].ProductDescription,
				OwnerPlant:          v.OwnerPlant,
				OwnerPlantName:      plantMapper[*v.OwnerPlant].PlantName,
				ValidityStartDate:   v.ValidityStartDate,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
				Images: dpfm_api_output_formatter.Images{
					Product: img,
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

func (c *BillOfMaterialListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.BillOfMaterialList{
		BillOfMaterials: make([]dpfm_api_output_formatter.BillOfMaterial, 0),
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

func extractBillOfMaterialListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.BillOfMaterialList {
	data := dpfm_api_input_reader.ReadBillOfMaterialList(msg)
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

func (c *BillOfMaterialListCtrl) Log(args ...interface{}) {
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
