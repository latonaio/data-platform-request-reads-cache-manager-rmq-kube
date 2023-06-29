package operationsdetaillist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	operationsdetaillist "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/operationsdetaillist"
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

type OperationsDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOperationsDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OperationsDetailListCtrl {
	return &OperationsDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OperationsDetailListCtrl) OperationsDetailList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractOperationsDetailListParam(msg)
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

	operationsRes, err := c.operationsRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("operationsRequest error: %w", err)
	}

	plRes, err := c.plantRequest(&params.Params, operationsRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("plantRequest error: %w", err)
	}

	err = c.addHeaderInfo(&params.Params, sID, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}

	c.fin(params, operationsRes, plRes, reqKey, "OperationsDetailList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *OperationsDetailListCtrl) plantRequest(
	params *dpfm_api_input_reader.OperationsDetailListParams,
	pvRes *apiresponses.OperationsRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := operationsdetaillist.CreatePlantReq(params, pvRes, sID, c.log)
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

func (c *OperationsDetailListCtrl) operationsRequest(
	params *dpfm_api_input_reader.OperationsDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.OperationsRes, error) {
	defer recovery(c.log)
	operationsReq := operationsdetaillist.CreateOperationsRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-operations-reads-queue", operationsReq, sID, reqKey, "OperationsDetailList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("OperationsDetail cache set error: %w", err)
	}
	operationsRes, err := apiresponses.CreateOperationsRes(res)
	if err != nil {
		return nil, xerrors.Errorf("OperationsDetail response parse error: %w", err)
	}
	return operationsRes, nil
}

func (c *OperationsDetailListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.OperationsDetailListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	//202@gmail.com/priceMaster/list/user=OwnerProductionPlantBusinessPartner/OperationsList
	key := fmt.Sprintf(`%s/operations/list/user=%s/OperationsList`,
		params.UserID, params.User)
	api := "OperationsDetailListHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	opList := dpfm_api_output_formatter.OperationsList{}
	err = json.Unmarshal(b, &opList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range opList.Operations {
		if v.Operations == params.Operations {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *OperationsDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *OperationsDetailListCtrl) fin(
	params *dpfm_api_input_reader.OperationsDetailList,
	operationsRes *apiresponses.OperationsRes,
	plRes *apiresponses.PlantRes,
	url,
	api string,
	setFlag *RedisCacheApiName,
) error {
	key := fmt.Sprintf(`%s/operations/list/user=%s/OperationsList`,
		params.Params.UserID, params.Params.User)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	operationsList := dpfm_api_output_formatter.OperationsList{}
	err = json.Unmarshal(b, &operationsList)
	if err != nil {
		return err
	}
	idx := 0
	for i, v := range operationsList.Operations {
		if v.Operations == params.Params.Operations {
			idx = i
			break
		}
	}

	header := dpfm_api_output_formatter.OperationsDetailHeader{
		Index: idx,
		Key:   key,
	}

	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	details := make([]dpfm_api_output_formatter.OperationsDetail, 0, len(*operationsRes.Message.Item))
	for _, v := range *operationsRes.Message.Item {

		details = append(details, dpfm_api_output_formatter.OperationsDetail{
			OperationsItem:          v.Operations,
			OperationsText:          v.OperationsText,
			ProductionPlantName:     plantMapper[*v.ProductionPlant].PlantName,
			StandardLotSizeQuantity: v.StandardLotSizeQuantity,
			OperationsUnit:          v.OperationsUnit,
			ValidityStartDate:       v.ValidityStartDate,
			IsMarkedForDeletion:     v.IsMarkedForDeletion,
		})
	}

	data := dpfm_api_output_formatter.OperationsDetailList{
		OperationsDetailHeader: header,
		OperationsDetail:       details,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	redisKey := strings.Join([]string{url, api}, "/")
	b, _ = json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *OperationsDetailListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.OperationsDetailList{
		OperationsDetail: make([]dpfm_api_output_formatter.OperationsDetail, 0),
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

func extractOperationsDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OperationsDetailList {
	data := dpfm_api_input_reader.ReadOperationsDetailList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *OperationsDetailListCtrl) Log(args ...interface{}) {
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
