package operationslist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/operationslist"
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

// CreateOperationsItemsReq
type OperationsListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOperationsListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OperationsListCtrl {
	return &OperationsListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OperationsListCtrl) OperationsList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractOperationsListParam(msg)
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

	pRes, err := c.operationsRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	drRes, err := c.descriptionRequest(&params.Params, pRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	plRes, err := c.plantRequest(&params.Params, pRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	opdRes, err := c.operationsDocRequest(&params.Params, pRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	c.fin(params, pRes, drRes, plRes, opdRes, reqKey, "OperationsList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *OperationsListCtrl) plantRequest(
	params *dpfm_api_input_reader.OperationsListParams,
	opRes *apiresponses.OperationsRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := operationslist.CreatePlantReq(params, opRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", plReq, sID, reqKey, "OperationsList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Plant cache set error: %w", err)
	}
	plRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Plant response parse error: %w", err)
	}
	return plRes, nil
}

func (c *OperationsListCtrl) operationsGroupRequest(
	params *dpfm_api_input_reader.OperationsListParams,
	pmRes *apiresponses.OperationsRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.OperationsRes, error) {
	defer recovery(c.log)
	pgReq := operationslist.CreateOperationsRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-operations-group-reads-queue", pgReq, sID, reqKey, "OperationsGroup", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("OperationsGroup cache set error: %w", err)
	}
	pgRes, err := apiresponses.CreateOperationsRes(res)
	if err != nil {
		return nil, xerrors.Errorf("OperationsGroup response parse error: %w", err)
	}
	return pgRes, nil
}

func (c *OperationsListCtrl) operationsDocRequest(
	params *dpfm_api_input_reader.OperationsListParams,
	opRes *apiresponses.OperationsRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	opdReq := operationslist.CreateProductMasterDocReq(params, opRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", opdReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	opdRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("product response parse error: %w", err)
	}
	return opdRes, nil
}

func (c *OperationsListCtrl) operationsRequest(
	params *dpfm_api_input_reader.OperationsListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.OperationsRes, error) {
	defer recovery(c.log)
	pmReq := operationslist.CreateOperationsRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-operations-reads-queue", pmReq, sID, reqKey, "Operations", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("operations cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateOperationsRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Operations master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *OperationsListCtrl) descriptionRequest(
	params *dpfm_api_input_reader.OperationsListParams,
	opRes *apiresponses.OperationsRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	drReq := operationslist.CreateDescriptionReq(params, opRes, sID, c.log)
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

func (c *OperationsListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractOperationsListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OperationsList {
	data := dpfm_api_input_reader.ReadOperationsList(msg)
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

func (c *OperationsListCtrl) fin(
	params *dpfm_api_input_reader.OperationsList,
	opRes *apiresponses.OperationsRes,
	pmRes *apiresponses.ProductMasterRes,
	plRes *apiresponses.PlantRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	/*docs := make(map[string]*dpfm_api_output_formatter.OperationsImage)
	for _, v := range *pmdRes.Message.HeaderDoc {
		docs[v.Operations] = &dpfm_api_output_formatter.OperationsImage{
			BusinessPartnerID: v.DocIssuerBusinessPartner,
			DocID:             v.DocID,
			FileExtension:     v.FileExtension,
		}
	}*/

	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	imgMapper := map[string]apiresponses.PMDHeaderDoc{}
	for _, v := range *pmdRes.Message.HeaderDoc {
		imgMapper[v.Product] = v
	}

	descriptionMapper := map[string]apiresponses.ProductDescByBP{}
	for _, v := range *pmRes.Message.ProductDescByBP {
		descriptionMapper[v.Product] = v
	}

	operationsionOrders := make([]dpfm_api_output_formatter.Operations, 0)
	for _, v := range *opRes.Message.Header {
		img := &dpfm_api_output_formatter.ProductImage{}
		if i, ok := imgMapper[v.Product]; ok {
			img = &dpfm_api_output_formatter.ProductImage{
				BusinessPartnerID: i.DocIssuerBusinessPartner,
				DocID:             i.DocID,
				FileExtension:     i.FileExtension,
			}
		}

		operationsionOrders = append(operationsionOrders, dpfm_api_output_formatter.Operations{
			Operations:               v.Operations,
			Product:                  &v.Product,
			ProductDescription:       descriptionMapper[v.Product].ProductDescription,
			OwnerProductionPlantName: plantMapper[v.OwnerProductionPlant].PlantName,
			ValidityStartDate:        v.ValidityStartDate,
			IsMarkedForDeletion:      v.IsMarkedForDeletion,
			Images: dpfm_api_output_formatter.Images{
				Operations: img,
			},
		},
		)
	}

	data := dpfm_api_output_formatter.OperationsList{
		Operations: operationsionOrders,
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

func (c *OperationsListCtrl) finEmptyProcess(
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

func (c *OperationsListCtrl) Log(args ...interface{}) {
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
