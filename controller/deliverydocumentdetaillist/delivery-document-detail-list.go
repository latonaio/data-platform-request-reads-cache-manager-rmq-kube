package deliverydocumentdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/deliverydocumentdetaillist"
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

type DeliveryDocumentDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewDeliveryDocumentDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *DeliveryDocumentDetailListCtrl {
	return &DeliveryDocumentDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *DeliveryDocumentDetailListCtrl) DeliveryDocumentDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractDeliveryDocumentDetailListParam(msg)
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

	ddRes, err := c.deliveryDocumentRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("deliveryDocumentRequest error: %w", err)
	}
	err = c.addHeaderInfo(&params.Params, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}
	err = c.fin(params, ddRes, reqKey, "DeliveryDocumentDetailList", &cacheResult)
	if err != nil {
		return err
	}
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())

	return nil
}

func (c *DeliveryDocumentDetailListCtrl) deliveryDocumentRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.DeliveryDocumentRes, error) {
	defer recovery(c.log)
	ddReq := deliverydocumentdetaillist.CreateDeliveryDocumentReq(params, sID, c.log)
	res, err := c.request("data-platform-api-delivery-document-reads-queue", ddReq, sID, reqKey, "DeliveryDocumentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument cache set error: %w", err)
	}
	ddRes, err := apiresponses.CreateDeliveryDocumentRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument response parse error: %w", err)
	}
	return ddRes, nil
}

func (c *DeliveryDocumentDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *DeliveryDocumentDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.DeliveryDocumentDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	// "101@gmail.com/deliveryDocument/list/user=DeliverToParty/headerBillingStatusException=CL"
	key := fmt.Sprintf(`%s/deliveryDocument/list/user=%s/headerBillingStatusException=%v/DeliveryDocumentList`,
		params.UserID, params.User, "CL")
	api := "DeliveryDocumentDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ddList := dpfm_api_output_formatter.DeliveryDocumentList{}
	err = json.Unmarshal(b, &ddList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range ddList.DeliveryDocuments {
		if v.DeliveryDocument == params.DeliveryDocument {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *DeliveryDocumentDetailListCtrl) fin(
	params *dpfm_api_input_reader.DeliveryDocumentDetailList,
	ddRes *apiresponses.DeliveryDocumentRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	// "101@gmail.com/deliveryDocument/list/user=DeliverToParty/headerBillingStatusException=CL"
	key := fmt.Sprintf(`%s/deliveryDocument/list/user=%s/headerBillingStatusException=%v/DeliveryDocumentList`,
		params.Params.UserID, params.Params.User, "CL")
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	ddList := dpfm_api_output_formatter.DeliveryDocumentList{}
	err = json.Unmarshal(b, &ddList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range ddList.DeliveryDocuments {
		if v.DeliveryDocument == params.Params.DeliveryDocument {
			idx = i
			break
		}
	}
	header := dpfm_api_output_formatter.DeliveryDocumentDetailHeader{
		Index: idx,
		Key:   key,
	}

	data := dpfm_api_output_formatter.DeliveryDocumentDetailList{
		DeliveryDocumentDetailHeader: header,
		DeliveryDocumentDetail:       make([]dpfm_api_output_formatter.DeliveryDocumentDetailSummary, 0, len(*ddRes.Message.Item)),
	}

	for _, v := range *ddRes.Message.Item {
		itemText := v.DeliveryDocumentItemTextByBuyer
		if params.Params.User == "DeliverToParty" {
			itemText = v.DeliveryDocumentItemTextBySeller
		}

		quantity := 0
		if v.PlannedGoodsIssueQuantity != nil {
			quantity = int(*v.PlannedGoodsIssueQuantity)
		}

		data.DeliveryDocumentDetail = append(data.DeliveryDocumentDetail,
			dpfm_api_output_formatter.DeliveryDocumentDetailSummary{
				DeliveryDocumentItem:           v.DeliveryDocumentItem,
				Product:                        v.Product,
				DeliveryDocumentItemText:       *itemText,
				OriginalQuantityInDeliveryUnit: quantity,
				DeliveryUnit:                   v.DeliveryUnit,
				ActualGoodsIssueDate:           v.ActualGoodsIssueDate,
				ActualGoodsIssueTime:           v.ActualGoodsIssueTime,
				ActualGoodsReceiptDate:         v.ActualGoodsReceiptDate,
				ActualGoodsReceiptTime:         v.ActualGoodsReceiptTime,
				IsCancelled:                    v.IsCancelled,
				IsMarkedForDeletion:            v.IsMarkedForDeletion,
				OrderDetailJumpReq: dpfm_api_output_formatter.OrderDetailJumpReq{
					OrderID:   v.OrderID,
					OrderItem: v.OrderItem,
					Product:   v.Product,
					Payer:     v.Payer,
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

func extractDeliveryDocumentDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.DeliveryDocumentDetailList {
	data := dpfm_api_input_reader.ReadDeliveryDocumentDetailList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *DeliveryDocumentDetailListCtrl) Log(args ...interface{}) {
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
