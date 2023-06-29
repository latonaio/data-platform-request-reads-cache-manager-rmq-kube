package supplychainrelationshipexconflist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/supplychainrelationshipdetaillist"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

type SupplyChainRelationshipExconfListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewSupplyChainRelationshipExconfListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *SupplyChainRelationshipExconfListCtrl {
	return &SupplyChainRelationshipExconfListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *SupplyChainRelationshipExconfListCtrl) SupplyChainRelationshipExconfList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractSupplyChainRelationshipExconfListParam(msg)
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

	err = c.addHeaderInfo(&params.Params, sID, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}

	scrRes, err := c.supplyChainRelationshipGeneralRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	c.fin(params, scrRes, reqKey, "SupplyChainRelationshipExconfList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *SupplyChainRelationshipExconfListCtrl) supplyChainRelationshipGeneralRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.SupplyChainRelationshipExconfRes, error) {
	defer recovery(c.log)
	scrReq := supplychainrelationshipdetaillist.CreateSupplyChainRelationshipGeneralRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-supply-chain-relationship-exconf-queue", scrReq, sID, reqKey, "SupplyChainRelationship", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("supply chain relationship exconf cache set error: %w", err)
	}
	scrRes, err := apiresponses.CreateSupplyChainRelationshipExconfRes(res)
	if err != nil {
		return nil, xerrors.Errorf("supply chain relationship exconf response parse error: %w", err)
	}
	return scrRes, nil
}

func (c *SupplyChainRelationshipExconfListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractSupplyChainRelationshipExconfListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.SupplyChainRelationshipExconfList {
	data := dpfm_api_input_reader.ReadSupplyChainRelationshipExconfList(msg)
	return data
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

func (c *SupplyChainRelationshipExconfListCtrl) fin(
	params *dpfm_api_input_reader.SupplyChainRelationshipExconfList,
	scrRes *apiresponses.SupplyChainRelationshipExconfRes,
	url, api string, setFlag *RedisCacheApiName,
) error {

	generalExconf := dpfm_api_output_formatter.SupplyChainRelationshipExconfList{
		Content: "General",
		Exist:   scrRes.SupplyChainRelationshipGeneral.ExistenceConf,
		Param:   scrRes.SupplyChainRelationshipGeneral,
	}

	// 201@gmail.com/SupplyChainRelationship/list/user=BusinessPartner/SupplyChainRelationshipList
	key := fmt.Sprintf("%s/supplyChainRelationship/list/user=%s/SupplyChainRelationshipList",
		*params.Params.UserID, *params.Params.User)
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	scrList := dpfm_api_output_formatter.SupplyChainRelationshipList{}
	err = json.Unmarshal(b, &scrList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range scrList.SupplyChainRelationship {
		if v.SupplyChainRelationshipID == *params.Params.SupplyChainRelationshipID {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	data := dpfm_api_output_formatter.SupplyChainRelationshipGeneralExconfList{
		General: dpfm_api_output_formatter.SupplyChainRelationshipExconfGeneral{
			Index: idx,
			Key:   key,
		},
		Existences: []dpfm_api_output_formatter.SupplyChainRelationshipExconfList{
			generalExconf,
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
	b, _ = json.Marshal(data)
	err = c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *SupplyChainRelationshipExconfListCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	url string, setFlag *RedisCacheApiName,
) error {
	// 201@gmail.com/supplyChainRelationship/list/user=Buyer/SupplyChainRelationshipList
	key := fmt.Sprintf(`%s/supplyChainRelationship/list/user=%s/SupplyChainRelationshipList`,
		*params.UserID, *params.User)
	api := "SupplyChainRelationshipGeneral"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	srcList := dpfm_api_output_formatter.SupplyChainRelationshipList{}
	err = json.Unmarshal(b, &srcList)
	if err != nil {
		return err
	}

	src := dpfm_api_output_formatter.SupplyChainRelationship{}

	idx := -1
	for i, v := range srcList.SupplyChainRelationship {
		if v.SupplyChainRelationshipID == *params.SupplyChainRelationshipID {
			data, err := json.Marshal(v)
			if err != nil {
				return err
			}
			err = json.Unmarshal(data, &src)
			if err != nil {
				return err
			}
			idx = i
			break
		}
	}

	params.Buyer = src.Buyer
	params.Seller = src.Seller

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func isExist[T any](d interface{}) *bool {
	tr := true
	fa := false

	if d == nil {
		return &fa
	}

	switch d := d.(type) {
	case T:
		return &tr
	case *T:
		if d != nil {
			return &tr
		}
	case []T:
		if len(d) > 0 {
			return &tr
		}
	case *[]T:
		if d == nil {
			return &fa
		}
		if len(*d) > 0 {
			return &tr
		}
	}
	return &fa
}

func recovery(l *logger.Logger) {
	if e := recover(); e != nil {
		l.Error("%+v", e)
		return
	}
}
