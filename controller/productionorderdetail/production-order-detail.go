package productionorderdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/productionorderdetail"
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

// CreateProductionOrderItemsReq
type ProductionOrderDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewProductionOrderDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *ProductionOrderDetailCtrl {
	return &ProductionOrderDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *ProductionOrderDetailCtrl) ProductionOrderDetail(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractProductionOrderDetailParam(msg)
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

	poRes, err := c.productionOrderRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	if poRes.Message.Item == nil || len(*poRes.Message.Item) == 0 {
		c.finEmptyProcess(params, reqKey, "ProductionOrderDetail", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	bpRes, err := c.businessPartnerRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	pRes, err := c.productRequest(&params.Params, poRes, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	plantRes, err := c.plantRequest(&params.Params, poRes, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	pmdRes, err := c.productMasterDocRequest(&params.Params, pRes, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	c.fin(params, poRes, bpRes, pRes, plantRes, pmdRes, reqKey, "ProductionOrderDetail", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *ProductionOrderDetailCtrl) productionOrderRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductionOrderRes, error) {
	defer recovery(c.log)
	oiReq := productionorderdetail.CreateProductionOrderReq(params, sID, c.log)
	res, err := c.request("data-platform-api-production-order-reads-queue", oiReq, sID, reqKey, "ProductionOrder", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateProductionOrderRes(res)
	if err != nil {
		return nil, xerrors.Errorf("orders response parse error: %w", err)
	}
	return oiRes, nil
}

func (c *ProductionOrderDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := productionorderdetail.CreateBusinessPartnerReq(params, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpReq, sID, reqKey, "BusinessPartnerGeneral", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *ProductionOrderDetailCtrl) plantRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailParams,
	poRes *apiresponses.ProductionOrderRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	dtpReq := productionorderdetail.CreateDeliverToPlantReq(params, poRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", dtpReq, sID, reqKey, "DeliveryDocumentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryToPlant cache set error: %w", err)
	}
	dtpRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant response parse error: %w", err)
	}
	return dtpRes, nil
}

func (c *ProductionOrderDetailCtrl) productRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailParams,
	poRes *apiresponses.ProductionOrderRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	pmReq := productionorderdetail.CreateProductRequest(params, poRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *ProductionOrderDetailCtrl) productMasterDocRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	pmDocReq := productionorderdetail.CreateProductMasterDocReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pmDocReq, sID, reqKey, "ProductMasterDoc", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc cache set error: %w", err)
	}
	pmdRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc response parse error: %w", err)
	}
	return pmdRes, nil
}

func (c *ProductionOrderDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractProductionOrderDetailParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.ProductionOrderDetail {
	data := dpfm_api_input_reader.ReadProductionOrderDetail(msg)
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

func (c *ProductionOrderDetailCtrl) fin(
	params *dpfm_api_input_reader.ProductionOrderDetail,
	poRes *apiresponses.ProductionOrderRes,
	bpRes *apiresponses.BusinessPartnerRes,
	pRes *apiresponses.ProductMasterRes,
	plantRes *apiresponses.PlantRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	pNames := make(map[string]*string)
	for _, v := range *pRes.Message.ProductDescByBP {
		pNames[v.Product] = v.ProductDescription
	}

	components := make([]dpfm_api_output_formatter.ComponentItem, 0)
	for _, v := range *poRes.Message.ItemComponent {
		components = append(components, dpfm_api_output_formatter.ComponentItem{
			ComponentProduct:                v.ComponentProduct,
			ComponentProductRequirementDate: v.ComponentProductRequirementDate,
			ComponentProductRequirementTime: v.ComponentProductRequirementTime,
			RequiredQuantity:                v.RequiredQuantity,
			WithdrawnQuantity:               v.WithdrawnQuantity,
			BaseUnit:                        v.BaseUnit,
			CostingPolicy:                   v.CostingPolicy,
			StandardPrice:                   v.StandardPrice,
			MovingAveragePrice:              v.MovingAveragePrice,
		})
	}

	operations := make([]dpfm_api_output_formatter.Operation, 0)
	for _, v := range *poRes.Message.ItemOperation {
		operations = append(operations, dpfm_api_output_formatter.Operation{
			OperationText:                        v.OperationsText,
			WorkCenter:                           v.WorkCenter,
			OperationPlannedTotalQuantity:        v.OperationPlannedTotalQuantity,
			OperationTotalConfirmedYieldQuantity: v.OperationTotalConfirmedYieldQuantity,
			OperationErlstSchedldExecStrtDte:     v.OperationErlstSchedldExecStrtDte,
			OperationErlstSchedldExecStrtTme:     v.OperationErlstSchedldExecStrtTme,
			OperationErlstSchedldExecEndDate:     v.OperationErlstSchedldExecEndDate,
			OperationErlstSchedldExecEndTime:     v.OperationErlstSchedldExecEndTme,
			OperationActualExecutionStartDate:    v.OperationActualExecutionStartDate,
			OperationActualExecutionStartTime:    v.OperationActualExecutionStartTime,
			OperationActualExecutionEndDate:      v.OperationActualExecutionEndDate,
			OperationActualExecutionEndTime:      v.OperationActualExecutionEndTime,
		})
	}

	item := (*poRes.Message.Item)[0]
	pmd := (*pmdRes.Message.HeaderDoc)[0]

	data := dpfm_api_output_formatter.ProductionOrderDetail{
		ProductionOrder:                 item.ProductionOrder,
		ProductionOrderItem:             item.ProductionOrderItem,
		OrderItemText:                   item.ProductionOrderItemText,
		Product:                         item.Product,
		ProductName:                     pNames[*item.Product],
		MRPArea:                         item.MRPArea,
		ProductionVersion:               item.ProductionVersion,
		MinimumLotSizeQuantity:          item.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:          item.MaximumLotSizeQuantity,
		StandardLotSizeQuantity:         item.StandardLotSizeQuantity,
		LotSizeRoundingQuantity:         item.LotSizeRoundingQuantity,
		ProductionOrderPlannedStartDate: item.ProductionOrderPlannedStartDate,
		ProductionOrderPlannedStartTime: item.ProductionOrderPlannedStartTime,
		ProductionOrderPlannedEndDate:   item.ProductionOrderPlannedEndDate,
		ProductionOrderPlannedEndTime:   item.ProductionOrderPlannedEndTime,
		ProductionOrderActualStartDate:  item.ProductionOrderActualStartDate,
		ProductionOrderActualStartTime:  item.ProductionOrderActualStartTime,
		ProductionOrderActualEndDate:    item.ProductionOrderActualEndDate,
		ProductionOrderActualEndTime:    item.ProductionOrderActualEndTime,
		TotalQuantity:                   &item.TotalQuantity,
		PlannedScrapQuantity:            item.PlannedScrapQuantity,
		ConfirmedYieldQuantity:          item.ConfirmedYieldQuantity,
		ProductionUnit:                  item.ProductionUnit,
		ProductionPlant:                 plantRes.Message.General.PlantName,
		ProductionPlantStorageLocation:  plantRes.Message.StorageLocation.StorageLocationName,
		// BillOfMaterialItem:              item.BillOfMaterialItem,
		Components: components,
		Operations: operations,
		Images: dpfm_api_output_formatter.Images{
			Product: &dpfm_api_output_formatter.ProductImage{
				BusinessPartnerID: *params.Params.BusinessPartner,
				DocID:             pmd.DocID,
				FileExtension:     pmd.FileExtension,
			},
		},
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

func (c *ProductionOrderDetailCtrl) finEmptyProcess(
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

func (c *ProductionOrderDetailCtrl) Log(args ...interface{}) {
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
