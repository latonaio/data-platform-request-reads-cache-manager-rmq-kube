package invoicedetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/invoicedetaillist"
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

type InvoiceDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewInvoiceDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *InvoiceDetailListCtrl {
	return &InvoiceDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *InvoiceDetailListCtrl) InvoiceDetailList(msg rabbitmq.RabbitmqMessage) error {
	// start := time.Now()
	params := extractInvoiceDetailListParam(msg)
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

	iRes, err := c.invoiceRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("failed to invoice request : %w", err)
	}
	err = c.addHeaderInfo(&params.Params, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}
	c.fin(params, iRes, reqKey, "InvoiceDocumentList", &cacheResult)
	return nil
}

func (c *InvoiceDetailListCtrl) invoiceRequest(
	params *dpfm_api_input_reader.InvoiceDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.InvoiceRes, error) {
	defer recovery(c.log)
	iReq := invoicedetaillist.CreateInvoiceReq(params, sID, c.log)
	res, err := c.request("data-platform-api-invoice-document-reads-queue", iReq, sID, reqKey, "Orders", setFlag)

	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateInvoiceRes(res)
	if err != nil {
		return nil, xerrors.Errorf("orders response parse error: %w", err)
	}
	return oiRes, nil

}

func (c *InvoiceDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.InvoiceDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	//201@gmail.com/invoiceDocument/list/user=BillToParty/businessPartner=201/headerPaymentBlockStatus=false
	key := fmt.Sprintf(`%s/invoiceDocument/list/user=%s/businessPartner=%d/headerPaymentBlockStatus=%v/InvoiceDocumentList`,
		params.UserID, params.User, params.BusinessPartner, false)
	api := "InvoiceDocumentDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	iList := dpfm_api_output_formatter.InvoiceList{}
	err = json.Unmarshal(b, &iList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range iList.Invoices {
		if v.InvoiceDocument == params.InvoiceDocument {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *InvoiceDetailListCtrl) fin(
	params *dpfm_api_input_reader.InvoiceDetailList,
	iRes *apiresponses.InvoiceRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	//201@gmail.com/invoiceDocument/list/user=BillToParty/businessPartner=201/headerPaymentBlockStatus=false
	key := fmt.Sprintf(`%s/invoiceDocument/list/user=%s/businessPartner=%d/headerPaymentBlockStatus=%v/InvoiceDocumentList`,
		params.Params.UserID, params.Params.User, params.Params.BusinessPartner, false)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	iList := dpfm_api_output_formatter.InvoiceList{}
	err = json.Unmarshal(b, &iList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range iList.Invoices {
		if v.InvoiceDocument == params.Params.InvoiceDocument {
			idx = i
			break
		}
	}

	header := dpfm_api_output_formatter.InvoiceDocumentDetailHeader{
		Index: idx,
		Key:   key,
	}

	details := make([]dpfm_api_output_formatter.InvoiceDocumentDetailSummary, 0, len(*iRes.Message.Item))
	for _, v := range *iRes.Message.Item {
		details = append(details, dpfm_api_output_formatter.InvoiceDocumentDetailSummary{
			InvoiceDocument:           v.InvoiceDocument,
			InvoiceDocumentItem:       v.InvoiceDocumentItem,
			Product:                   v.Product,
			InvoiceDocumentItemText:   v.InvoiceDocumentItemText,
			InvoiceQuantityInBaseUnit: v.InvoiceQuantityInBaseUnit,
			InvoiceQuantityUnit:       v.InvoiceQuantityUnit,
			ActualGoodsIssueDate:      v.ActualGoodsIssueDate,
			ActualGoodsIssueTime:      v.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:    v.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:    v.ActualGoodsReceiptTime,
			ItemBillingIsConfirmed:    v.ItemBillingIsConfirmed,
			IsCancelled:               v.IsCancelled,
			// IsMarkedForDeletion: v.IsMarkedForDeletion,
			OrdersDetailJumpReq: dpfm_api_output_formatter.OrdersDetailJumpReq{
				OrderID:   v.OrderID,
				OrderItem: v.OrderItem,
				Product:   v.Product,
				Buyer:     v.Buyer,
			},
			DeliveryDetailJumpReq: dpfm_api_output_formatter.DeliveryDetailJumpReq{
				DeliveryDocument:     *v.DeliveryDocument,
				DeliveryDocumentItem: *v.DeliveryDocumentItem,
				DeliverToParty:       v.DeliverToParty,
				DeliverFromParty:     v.DeliverFromParty,
				Product:              v.Product,
				Buyer:                v.Buyer,
			},
		})
	}
	data := dpfm_api_output_formatter.InvoiceDocumentDetailList{
		InvoiceDocumentDetailHeader: header,
		InvoiceDetail:               details,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	redisKey := strings.Join([]string{url, api}, "/")
	b, _ = json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *InvoiceDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractInvoiceDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.InvoiceDetailList {
	data := dpfm_api_input_reader.ReadInvoiceDetailList(msg)
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

func (c *InvoiceDetailListCtrl) Log(args ...interface{}) {
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
