package deliverydocumentlist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/deliverydocumentlist"
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

type DeliveryDocumentListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewDeliveryDocumentListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *DeliveryDocumentListCtrl {
	return &DeliveryDocumentListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *DeliveryDocumentListCtrl) DeliveryDocumentList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractDeliveryDocumentListParam(msg)
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

	ddRes, err := c.deliveryDocumentRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("deliveryDocumentRequest error: %w", err)
	}
	if ddRes.Message.Header == nil || len(*ddRes.Message.Header) == 0 {
		c.finEmptyProcess(params, reqKey, "DeliveryDocumentList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	scRes, err := c.supplyChainListRequest(&params.Params, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}

	bpRes, err := c.businessPartnerRequest(&params.Params, ddRes, scRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}

	fromPlantRes, err := c.deliverFromPlantRequest(&params.Params, ddRes, scRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}
	toPlantRes, err := c.deliverToPlantRequest(&params.Params, ddRes, scRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}

	c.fin(params, ddRes, bpRes, fromPlantRes, toPlantRes, scRes, reqKey, "DeliveryDocumentList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *DeliveryDocumentListCtrl) supplyChainListRequest(
	params *dpfm_api_input_reader.DeliveryDocumentListParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.SupplyChainRelationshipRes, error) {
	defer recovery(c.log)
	scReq := deliverydocumentlist.CreateSupplyChainReq(params, ddRes, sID, c.log)
	res, err := c.request("data-platform-api-supply-chain-rel-master-reads-queue", scReq, sID, reqKey, "SupplyChainList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument cache set error: %w", err)
	}
	scRes, err := apiresponses.CreateSupplyChainRelationshipRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument response parse error: %w", err)
	}
	return scRes, nil
}

func (c *DeliveryDocumentListCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.DeliveryDocumentListParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := deliverydocumentlist.CreateBusinessPartnerReq(params, ddRes, scRes, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpReq, sID, reqKey, "DeliveryDocumentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("business-partner cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("business-partner response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *DeliveryDocumentListCtrl) deliveryDocumentRequest(
	params *dpfm_api_input_reader.DeliveryDocumentListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.DeliveryDocumentRes, error) {
	defer recovery(c.log)
	ddReq := deliverydocumentlist.CreateDeliveryDocumentReq(params, sID, c.log)
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

func (c *DeliveryDocumentListCtrl) deliverFromPlantRequest(
	params *dpfm_api_input_reader.DeliveryDocumentListParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	dfpReq := deliverydocumentlist.CreateDeliverFromPlantReq(params, ddRes, scRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", dfpReq, sID, reqKey, "DeliveryDocumentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant cache set error: %w", err)
	}
	dfpRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant response parse error: %w", err)
	}
	return dfpRes, nil
}

func (c *DeliveryDocumentListCtrl) deliverToPlantRequest(
	params *dpfm_api_input_reader.DeliveryDocumentListParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	dtpReq := deliverydocumentlist.CreateDeliverToPlantReq(params, ddRes, scRes, sID, c.log)
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

func (c *DeliveryDocumentListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *DeliveryDocumentListCtrl) fin(
	params *dpfm_api_input_reader.DeliveryDocumentList,
	ddRes *apiresponses.DeliveryDocumentRes,
	bpRes *apiresponses.BusinessPartnerRes,
	dfpRes *apiresponses.PlantRes,
	dtpRes *apiresponses.PlantRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	fromPlantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *dfpRes.Message.Generals {
		fromPlantMapper[v.Plant] = v
	}
	toPlantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *dtpRes.Message.Generals {
		toPlantMapper[v.Plant] = v
	}
	bpMapper := map[int]*string{}
	for _, v := range *bpRes.Message.Generals {
		bpMapper[v.BusinessPartner] = v.BusinessPartnerName
	}

	deliver := map[int]dpfm_api_output_formatter.Deliver{}
	for _, dd := range *ddRes.Message.Header {
		deliver[dd.DeliveryDocument] = dpfm_api_output_formatter.Deliver{
			DeliverFromParty: make([]dpfm_api_output_formatter.DeliveryPlant, 0),
			DeliverToParty:   make([]dpfm_api_output_formatter.DeliveryPlant, 0),
		}

		buyer := dd.Buyer
		seller := dd.Seller
		toParty := dd.DeliverToParty
		fromParty := dd.DeliverFromParty
		toPlant := dd.DeliverToPlant
		fromPlant := dd.DeliverFromPlant
		scrdID := dd.SupplyChainRelationshipDeliveryID

		for _, sc := range *scRes.Message.DeliveryPlantRelation {
			if sc.SupplyChainRelationshipDeliveryID != *scrdID {
				continue
			}
			if true &&
				sc.Buyer == *buyer &&
				sc.Seller == *seller &&
				sc.DeliverFromParty == *fromParty &&
				sc.DeliverFromPlant == *fromPlant {
				tmp := deliver[dd.DeliveryDocument]
				tmp.DeliverToParty = append(tmp.DeliverToParty, dpfm_api_output_formatter.DeliveryPlant{
					DeliverToPlant:                         toPlantMapper[sc.DeliverToPlant].Plant,
					DeliverToPlantName:                     toPlantMapper[sc.DeliverToPlant].PlantName,
					DeliverToParty:                         sc.DeliverToParty,
					DeliverToPartyName:                     bpMapper[sc.DeliverToParty],
					DefaultRelation:                        *sc.DefaultRelation,
					SupplyChainRelationshipDeliveryID:      sc.SupplyChainRelationshipDeliveryID,
					SupplyChainRelationshipDeliveryPlantID: sc.SupplyChainRelationshipDeliveryPlantID,
				},
				)
				deliver[dd.DeliveryDocument] = tmp
			}
			if true &&
				sc.Buyer == *buyer &&
				sc.Seller == *seller &&
				sc.DeliverToParty == *toParty &&
				sc.DeliverToPlant == *toPlant {
				tmp := deliver[dd.DeliveryDocument]
				tmp.DeliverFromParty = append(tmp.DeliverFromParty, dpfm_api_output_formatter.DeliveryPlant{
					DeliverFromPlant:                       fromPlantMapper[sc.DeliverFromPlant].Plant,
					DeliverFromPlantName:                   fromPlantMapper[sc.DeliverFromPlant].PlantName,
					DeliverFromParty:                       sc.DeliverFromParty,
					DeliverFromPartyName:                   bpMapper[sc.DeliverFromParty],
					DefaultRelation:                        *sc.DefaultRelation,
					SupplyChainRelationshipDeliveryID:      sc.SupplyChainRelationshipDeliveryID,
					SupplyChainRelationshipDeliveryPlantID: sc.SupplyChainRelationshipDeliveryPlantID,
				},
				)
				deliver[dd.DeliveryDocument] = tmp
			}
		}

	}

	deliveryDocumentInfo := map[int]apiresponses.DeliveryDocumentHeader{}
	for _, v := range *ddRes.Message.Header {
		deliveryDocumentInfo[v.DeliveryDocument] = v
	}
	data := dpfm_api_output_formatter.DeliveryDocumentList{
		DeliveryDocuments: make([]dpfm_api_output_formatter.DeliveryDocument, 0),
		Pulldown: dpfm_api_output_formatter.DeliveryDocumentListPullDown{
			SupplyChains: deliver,
		},
	}
	infos := deliveryDocumentDesc(deliveryDocumentInfo)
	for _, info := range infos {
		if params.Params.User == "DeliverToParty" {
			if *info.Buyer != *params.Params.DeliverToParty {
				continue
			}
		} else if params.Params.User == "DeliverFromParty" {
			if *info.Seller != *params.Params.DeliverFromParty {
				continue
			}
		}

		data.DeliveryDocuments = append(data.DeliveryDocuments,
			dpfm_api_output_formatter.DeliveryDocument{
				DeliverToParty:     *info.DeliverToParty,
				DeliverToPartyName: bpMapper[*info.DeliverToParty],
				DeliverToPlant:     *info.DeliverToPlant,
				DeliverToPlantName: *toPlantMapper[*info.DeliverToPlant].PlantName,

				DeliverFromParty:     *info.DeliverFromParty,
				DeliverFromPartyName: bpMapper[*info.DeliverFromParty],
				DeliverFromPlant:     *info.DeliverFromPlant,
				DeliverFromPlantName: *fromPlantMapper[*info.DeliverFromPlant].PlantName,

				DeliveryDocument:          info.DeliveryDocument,
				SupplyChainRelationshipID: info.DeliveryDocument,
				HeaderDeliveryStatus:      info.HeaderDeliveryStatus,
				HeaderBillingStatus:       info.HeaderBillingStatus,
				PlannedGoodsReceiptDate:   info.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:   info.PlannedGoodsReceiptTime,

				IsCancelled:         info.IsCancelled,
				IsMarkedForDeletion: info.IsMarkedForDeletion,
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
		return err
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *DeliveryDocumentListCtrl) finEmptyProcess(
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

func extractDeliveryDocumentListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.DeliveryDocumentList {
	data := dpfm_api_input_reader.ReadDeliveryDocumentList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func deliveryDocumentAsc[T any](d map[int]T) []T {
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

func deliveryDocumentDesc[T any](d map[int]T) []T {
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

func (c *DeliveryDocumentListCtrl) Log(args ...interface{}) {
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
