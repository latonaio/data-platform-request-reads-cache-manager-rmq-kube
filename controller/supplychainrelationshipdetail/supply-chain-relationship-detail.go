package supplychainrelationshipdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	supplychainrelationshipdetail "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/supplychainrelationshipdetaill"
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

type SupplyChainRelationshipDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewSupplyChainRelationshipDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *SupplyChainRelationshipDetailCtrl {
	return &SupplyChainRelationshipDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *SupplyChainRelationshipDetailCtrl) SupplyChainRelationshipDetail(
	msg rabbitmq.RabbitmqMessage,
	l *logger.Logger,
) error {
	start := time.Now()
	params := extractSupplyChainRelationshipDetailParam(msg)
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

	scrRes, err := c.supplyChainRelationshipRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	c.fin(
		params,
		scrRes,
		reqKey,
		"SupplyChainRelationshipDetail",
		&cacheResult,
	)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func extractSupplyChainRelationshipDetailParam(
	msg rabbitmq.RabbitmqMessage,
) *dpfm_api_input_reader.SupplyChainRelationshipDetail {
	data := dpfm_api_input_reader.ReadSupplyChainRelationshipDetail(msg)
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

func (c *SupplyChainRelationshipDetailCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailParams,
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

func (c *SupplyChainRelationshipDetailCtrl) supplyChainRelationshipRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.SupplyChainRelationshipRes, error) {
	defer recovery(c.log)
	req := supplychainrelationshipdetail.CreateSupplyChainRelationshipRequest(
		params,
		sID,
		[]string{
			"Generals",
			"DeliveryRelationsBySCRID",
			"DeliveryPlantRelationsBySCRID",
		},
		c.log,
	)
	res, err := c.request("data-platform-api-supply-chain-rel-master-reads-queue", req, sID, reqKey, "SupplyChainRelationship", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("supply chain relationship cache set error: %w", err)
	}
	resBody, err := apiresponses.CreateSupplyChainRelationshipRes(res)
	if err != nil {
		return nil, xerrors.Errorf("supply chain relationship response parse error: %w", err)
	}
	return resBody, nil
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

func (c *SupplyChainRelationshipDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *SupplyChainRelationshipDetailCtrl) fin(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetail,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	url,
	api string,
	setFlag *RedisCacheApiName,
) error {
	general := dpfm_api_output_formatter.SupplyChainRelationshipDetailContent{
		Content: "General",
		Param:   scrRes.Message.General,
	}

	delivery := dpfm_api_output_formatter.SupplyChainRelationshipDetailContent{
		Content: "Delivery",
		Param:   scrRes.Message.DeliveryRelation,
	}

	deliveryPlant := dpfm_api_output_formatter.SupplyChainRelationshipDetailContent{
		Content: "DeliveryPlant",
		Param:   scrRes.Message.DeliveryPlantRelation,
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
	data := dpfm_api_output_formatter.SupplyChainRelationshipDetail{
		Header: dpfm_api_output_formatter.SupplyChainRelationshipDetailHeader{
			Index: idx,
			Key:   key,
		},
		Contents: []dpfm_api_output_formatter.SupplyChainRelationshipDetailContent{
			general,
			delivery,
			deliveryPlant,
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

func (c *SupplyChainRelationshipDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := supplychainrelationshipdetail.CreateBusinessPartnerReq(params, scrRes, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpReq, sID, reqKey, "BusinessPartner", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("business partner cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("business partner response parse error: %w", err)
	}
	return bpRes, nil
}
