package invoicedetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/invoicelist"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type InvoiceDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewInvoiceDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *InvoiceDetailCtrl {
	return &InvoiceDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *InvoiceDetailCtrl) InvoiceDetail(msg rabbitmq.RabbitmqMessage) error {
	// start := time.Now()
	params := extractInvoiceListParam(msg)
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
	bpRes, err := c.businessPartnerRequest(&params.Params, iRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(iRes)
		return xerrors.Errorf("failed to invoice request : %w", err)
	}

	c.fin(params, iRes, bpRes, reqKey, "InvoiceDocumentList", &cacheResult)
	return nil
}

func (c *InvoiceDetailCtrl) invoiceRequest(
	params *dpfm_api_input_reader.InvoiceListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.InvoiceRes, error) {
	defer recovery(c.log)
	oiReq := invoicelist.CreateInvoiceReq(params, sID, c.log)
	res, err := c.request("data-platform-api-invoice-document-reads-queue", oiReq, sID, reqKey, "Orders", setFlag)

	if err != nil {
		return nil, xerrors.Errorf("orders cache set error: %w", err)
	}
	oiRes, err := apiresponses.CreateInvoiceRes(res)
	if err != nil {
		return nil, xerrors.Errorf("orders response parse error: %w", err)
	}
	return oiRes, nil

}

func (c *InvoiceDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.InvoiceListParams,
	iRes *apiresponses.InvoiceRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := invoicelist.CreateBusinessPartnerReq(params, iRes, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpReq, sID, reqKey, "BusinessPartner", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartner response parse error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("business partner response parse error: %w", err)
	}
	return bpRes, nil
}

// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
// func (c *InvoiceListCtrl)
func (c *InvoiceDetailCtrl) fin(
	params *dpfm_api_input_reader.InvoiceList,
	iRes *apiresponses.InvoiceRes,
	bpRes *apiresponses.BusinessPartnerRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	bpMapper := map[int]apiresponses.BPGeneral{}
	for _, v := range *bpRes.Message.Generals {
		bpMapper[v.BusinessPartner] = v
	}

	invoices := make([]dpfm_api_output_formatter.Invoices, 0, 0)
	for _, v := range *iRes.Message.Header {
		invoices = append(invoices, dpfm_api_output_formatter.Invoices{
			InvoiceDocument:          v.InvoiceDocument,
			BillToParty:              bpMapper[*v.BillToParty].BusinessPartnerName,
			BillFromParty:            bpMapper[*v.BillFromParty].BusinessPartnerName,
			InvoiceDocumentDate:      v.InvoiceDocumentDate,
			PaymentDueDate:           v.PaymentDueDate,
			HeaderBillingIsConfirmed: v.HeaderBillingIsConfirmed,
		})
	}
	data := dpfm_api_output_formatter.InvoiceList{
		Invoices: invoices,
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

func (c *InvoiceDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractInvoiceListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.InvoiceList {
	data := dpfm_api_input_reader.ReadInvoiceList(msg)
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

func (c *InvoiceDetailCtrl) Log(args ...interface{}) {
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
