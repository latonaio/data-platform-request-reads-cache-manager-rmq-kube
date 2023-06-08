package supplychainrelationshiplist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	supplychainrelationshiplist "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/supplychainrelationshiplist"
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

// CreateSupplyChainRelationshipItemsReq
type SupplyChainRelationshipListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewSupplyChainRelationshipListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *SupplyChainRelationshipListCtrl {
	return &SupplyChainRelationshipListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *SupplyChainRelationshipListCtrl) SupplyChainRelationshipList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractSupplyChainRelationshipListParam(msg)
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

	scrRes, err := c.supplyChainRelationshipRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	bpRes, err := c.businessPartnerRequest(&params.Params, scrRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	c.fin(params, scrRes, bpRes, reqKey, "SupplyChainRelationshipList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *SupplyChainRelationshipListCtrl) supplyChainRelationshipRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.SupplyChainRelationshipRes, error) {
	defer recovery(c.log)
	scrReq := supplychainrelationshiplist.CreateSupplyChainRelationshipRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-supply-chain-rel-master-reads-queue", scrReq, sID, reqKey, "SupplyChainRelationship", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("supply chain relationship cache set error: %w", err)
	}
	scrRes, err := apiresponses.CreateSupplyChainRelationshipRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Supply Chain Relationship master response parse error: %w", err)
	}
	return scrRes, nil
}

func (c *SupplyChainRelationshipListCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := supplychainrelationshiplist.CreateBusinessPartnerReq(params, scrRes, sID, c.log)
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

func (c *SupplyChainRelationshipListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractSupplyChainRelationshipListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.SupplyChainRelationshipList {
	data := dpfm_api_input_reader.ReadSupplyChainRelationshipList(msg)
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

func (c *SupplyChainRelationshipListCtrl) fin(
	params *dpfm_api_input_reader.SupplyChainRelationshipList,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	bpRes *apiresponses.BusinessPartnerRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	type bpIDRel struct {
		orderID        int
		sellerID       int
		buyerID        int
		deliveryStatus string
	}
	data := dpfm_api_output_formatter.SupplyChainRelationshipList{
		SupplyChainRelationship: make([]dpfm_api_output_formatter.SupplyChainRelationship, 0),
	}
	bpMapper := map[int]apiresponses.BPGeneral{}
	for _, v := range *bpRes.Message.Generals {
		bpMapper[v.BusinessPartner] = v
	}

	for _, info := range *scrRes.Message.General {
		if *params.Params.User == "Buyer" {
			if *info.Buyer != *params.Params.Buyer {
				continue
			}
		} else if *params.Params.User == "Seller" {
			if *info.Seller != *params.Params.Seller {
				continue
			}
		}
		buyerName := ""
		sellerName := ""
		buyer, ok := bpMapper[*info.Buyer]
		if ok {
			buyerName = *buyer.BusinessPartnerFullName
		}
		seller, ok := bpMapper[*info.Seller]
		if ok {
			sellerName = *seller.BusinessPartnerFullName
		}

		data.SupplyChainRelationship = append(data.SupplyChainRelationship,
			dpfm_api_output_formatter.SupplyChainRelationship{
				SupplyChainRelationshipID: info.SupplyChainRelationshipID,
				SellerName:                sellerName,
				Seller:                    info.Seller,
				BuyerName:                 buyerName,
				Buyer:                     info.Buyer,
				IsMarkedForDeletion:       info.IsMarkedForDeletion,
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
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *SupplyChainRelationshipListCtrl) finEmptyProcess(
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

func (c *SupplyChainRelationshipListCtrl) Log(args ...interface{}) {
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
