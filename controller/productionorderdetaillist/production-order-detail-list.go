package productionorderdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/productionorderdetaillist"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ProductionOrderDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewProductionOrderDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *ProductionOrderDetailListCtrl {
	return &ProductionOrderDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *ProductionOrderDetailListCtrl) ProductionOrderDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractProductionOrderDetailListParam(msg)
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
		if err != nil {
			return
		}
		b, _ := json.Marshal(cacheResult)
		err = c.cache.Set(c.ctx, reqKey, b, 0)
		if err != nil {
			c.log.Error("cache set error: %w", err)
		}
	}()

	poRes, err := c.productionOrderRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("productionOrderRequest error: %w", err)
	}

	pRes, err := c.productRequest(&params.Params, poRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("productRequest error: %w", err)
	}

	err = c.addHeaderInfo(&params.Params, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}
	err = c.fin(params, poRes, pRes, reqKey, "ProductionOrderDetailList", &cacheResult)
	if err != nil {
		return err
	}
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())

	return nil
}

func (c *ProductionOrderDetailListCtrl) productRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailListParams,
	poRes *apiresponses.ProductionOrderRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	pmReq := productionorderdetaillist.CreateProductRequest(params, poRes, sID, c.log)
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

func (c *ProductionOrderDetailListCtrl) productionOrderRequest(
	params *dpfm_api_input_reader.ProductionOrderDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductionOrderRes, error) {
	defer recovery(c.log)
	bpReq := productionorderdetaillist.CreateProductionOrderReq(params, sID, c.log)
	res, err := c.request("data-platform-api-production-order-reads-queue", bpReq, sID, reqKey, "ProductionOrder", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductionOrder cache set error: %w", err)
	}
	poRes, err := apiresponses.CreateProductionOrderRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductionOrder response parse error: %w", err)
	}
	return poRes, nil
}

func (c *ProductionOrderDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *ProductionOrderDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.ProductionOrderDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	// 101@gmail.com/production-order/list/user=OwnerProductionPlantBusinessPartner/headerIsMarkedForDeletion=false
	key := fmt.Sprintf(`%s/productionOrder/list/user=%s/headerIsMarkedForDeletion=%v/ProductionOrder`,
		*params.UserID, *params.User, *params.IsMarkedForDeletion)
	api := "ProductionOrderDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ddList := dpfm_api_output_formatter.ProductionOrderList{}
	err = json.Unmarshal(b, &ddList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range ddList.ProductionOrders {
		if v.ProductionOrder == *params.ProductionOrder {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *ProductionOrderDetailListCtrl) fin(
	params *dpfm_api_input_reader.ProductionOrderDetailList,
	poRes *apiresponses.ProductionOrderRes,
	pRes *apiresponses.ProductMasterRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	// 101@gmail.com/production-order/list/user=OwnerProductionPlantBusinessPartner/headerIsMarkedForDeletion=false
	key := fmt.Sprintf(`%s/productionOrder/list/user=%s/headerIsMarkedForDeletion=%v/ProductionOrder`,
		*params.Params.UserID, *params.Params.User, *params.Params.IsMarkedForDeletion)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ddList := dpfm_api_output_formatter.ProductionOrderList{}
	err = json.Unmarshal(b, &ddList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range ddList.ProductionOrders {
		if v.ProductionOrder == *params.Params.ProductionOrder {
			idx = i
			break
		}
	}
	header := dpfm_api_output_formatter.ProductionOrderDetailHeader{
		Index: idx,
		Key:   key,
	}
	data := dpfm_api_output_formatter.ProductionOrderDetailList{
		Header:  header,
		Details: make([]dpfm_api_output_formatter.ProductionOrderItemSummary, 0, len(*poRes.Message.Item)),
	}

	pNames := make(map[string]*string)
	for _, v := range *pRes.Message.ProductDescByBP {
		pNames[v.Product] = v.ProductDescription
	}

	for _, v := range *poRes.Message.Item {
		data.Details = append(data.Details,
			dpfm_api_output_formatter.ProductionOrderItemSummary{
				ProductionOrderItem:      v.ProductionOrderItem,
				Product:                  *v.Product,
				ProductName:              pNames[*v.Product],
				MRPArea:                  v.MRPArea,
				TotalQuantity:            &v.TotalQuantity,
				ConfirmedYieldQuantity:   v.ConfirmedYieldQuantity,
				ItemIsConfirmed:          v.ItemIsConfirmed,
				ItemIsPartiallyConfirmed: v.ItemIsPartiallyConfirmed,
				ItemIsReleased:           v.ItemIsReleased,
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
	b, _ = json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]string{
		"keyName": redisKey,
	}
	return nil
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

func extractProductionOrderDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.ProductionOrderDetailList {
	data := dpfm_api_input_reader.ReadProductionOrderDetailList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *ProductionOrderDetailListCtrl) Log(args ...interface{}) {
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
