package deliverydocumentdetailpagination

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/deliverydocumentdetailpagination"
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

type DeliveryDocumentDetailPaginationCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewDeliveryDocumentDetailPaginationCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *DeliveryDocumentDetailPaginationCtrl {
	return &DeliveryDocumentDetailPaginationCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *DeliveryDocumentDetailPaginationCtrl) Pagination(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractDeliveryDocumentDetailPaginationParam(msg)
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

	ddRes, err := c.deliveryDocumentRequest(params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("order Item Request error: %w", err)
	}

	err = c.fin(params, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("cache set error: %w", err)
	}
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *DeliveryDocumentDetailPaginationCtrl) deliveryDocumentRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailPagination,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.DeliveryDocumentRes, error) {
	defer recovery(c.log)
	ddReq := deliverydocumentdetailpagination.CreateDeliveryDocumentReq(&params.Params, sID, c.log)
	res, err := c.request("data-platform-api-delivery-document-reads-queue", ddReq, sID, reqKey, "DeliveryDocument", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument cache set error: %w", err)
	}
	ddRes, err := apiresponses.CreateDeliveryDocumentRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument response parse error: %w", err)
	}
	return ddRes, nil
}

func (c *DeliveryDocumentDetailPaginationCtrl) fin(
	params *dpfm_api_input_reader.DeliveryDocumentDetailPagination,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	url string,
	setFlag *RedisCacheApiName,
) error {
	api := "DeliveryDocumentDetailPagination"
	p := make([]dpfm_api_output_formatter.DeliveryDocumentDetailPage, 0, len(*ddRes.Message.Item))

	for _, v := range *ddRes.Message.Item {
		p = append(p, dpfm_api_output_formatter.DeliveryDocumentDetailPage{
			DeliveryDocument:     v.DeliveryDocument,
			DeliveryDocumentItem: v.DeliveryDocumentItem,
			Product:              *v.Product,
		})
	}

	data := dpfm_api_output_formatter.DeliveryDocumentDetailPagination{
		Paginations: p,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
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

func extractDeliveryDocumentDetailPaginationParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.DeliveryDocumentDetailPagination {
	data := dpfm_api_input_reader.ReadDeliveryDocumentDetailPagination(msg)
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

func (c *DeliveryDocumentDetailPaginationCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

type RedisCacheApiName map[string]map[string]interface{}

func (c *DeliveryDocumentDetailPaginationCtrl) Log(args ...interface{}) {
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
