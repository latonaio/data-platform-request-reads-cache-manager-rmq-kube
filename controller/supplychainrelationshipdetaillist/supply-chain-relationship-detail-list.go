package supplychainrelationshipdetaillist

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
	"sort"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

// CreateProductItemsReq
type SupplyChainRelationshipDetailListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewSupplyChainRelationshipDetailListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *SupplyChainRelationshipDetailListCtrl {
	return &SupplyChainRelationshipDetailListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *SupplyChainRelationshipDetailListCtrl) SupplyChainRelationshipDetailList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractSupplyChainRelationshipDetailListParam(msg)
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
	err = c.pushSupplyChainRelationshipHeader(params, sID, &cacheResult)
	if err != nil {
		return err
	}
	c.fin(params, scrRes, reqKey, "SupplyChainRelationshipDetailExconfList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *SupplyChainRelationshipDetailListCtrl) supplyChainRelationshipRequest(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.SupplyChainRelationshipRes, error) {
	defer recovery(c.log)
	scrReq := supplychainrelationshipdetaillist.CreateSupplyChainRelationshipRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-supply-chain-rel-master-reads-queue", scrReq, sID, reqKey, "SupplyChainRelationship", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	scrRes, err := apiresponses.CreateSupplyChainRelationshipRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return scrRes, nil
}

// func (c *SupplyChainRelationshipDetailListCtrl) productExconfRequest(
// 	params *dpfm_api_input_reader.SupplyChainRelationshipDetailListParams,
// 	pmRes *apiresponses.ProductMasterRes,
// 	sID string,
// 	reqKey string,
// 	setFlag *RedisCacheApiName,
// 	l *logger.Logger,
// ) (*apiresponses.ProductMasterRes, error) {
// 	defer recovery(c.log)
// 	pmReq := supplychainrelationshipdetaillist.CreateProductRequest(params, sID, c.log)
// 	c.rmq.SessionRequest()
// 	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
// 	if err != nil {
// 		return nil, xerrors.Errorf("product cache set error: %w", err)
// 	}
// 	pmRes, err := apiresponses.CreateProductMasterRes(res)
// 	if err != nil {
// 		return nil, xerrors.Errorf("Product master response parse error: %w", err)
// 	}
// 	return pmRes, nil
// }

func (c *SupplyChainRelationshipDetailListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractSupplyChainRelationshipDetailListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.SupplyChainRelationshipDetailList {
	data := dpfm_api_input_reader.ReadSupplyChainRelationshipDetailList(msg)
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

func (c *SupplyChainRelationshipDetailListCtrl) pushSupplyChainRelationshipHeader(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailList,
	url string, setFlag *RedisCacheApiName,
) error {
	api := "SupplyChainRelationshipDetailExconfListHeader"
	// 201@gmail.com/supplyChainRelationship/list/user=BusinessPartner/SupplyChainRelationshipList
	key := fmt.Sprintf("%s/supplychainrelationship/list/user=%s/isMarkedForDeletion=%v/SupplyChainRelationshipList",
		*params.Params.UserID, *params.Params.User, false)
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
	return nil
}

func (c *SupplyChainRelationshipDetailListCtrl) fin(
	params *dpfm_api_input_reader.SupplyChainRelationshipDetailList,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	url, api string, setFlag *RedisCacheApiName,
) error {

	generalExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "General",
		Exist:   isExist[apiresponses.SCRGeneral](scrRes.Message.General),
		Param:   scrRes.Message.General,
	}
	deliveryExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "Delivery",
		Exist:   isExist[apiresponses.DeliveryRelation](scrRes.Message.DeliveryRelation),
		Param:   scrRes.Message.DeliveryRelation,
	}
	deliveryPlantExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "DeliveryPlant",
		Exist:   isExist[apiresponses.DeliveryPlantRelation](scrRes.Message.DeliveryPlantRelation),
		Param:   scrRes.Message.DeliveryPlantRelation,
	}
	billingExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "Billing",
		Exist:   isExist[apiresponses.BillingRelation](scrRes.Message.BillingRelation),
		Param:   scrRes.Message.BillingRelation,
	}
	paymentExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "Payment",
		Exist:   isExist[apiresponses.PaymentRelation](scrRes.Message.PaymentRelation),
		Param:   scrRes.Message.PaymentRelation,
	}
	transactionExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
		Content: "Transaction",
		Exist:   isExist[apiresponses.Transaction](scrRes.Message.Transaction),
		Param:   scrRes.Message.Transaction,
	}

	// textExconf := dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
	// 	Content: "Text",
	// 	Exist:   isExist[apiresponses.ProductDescription](pmRes.Message.ProductDescription),
	// 	Param:   pmRes.Message.ProductDescription,
	// }
	// bpExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "BP",
	// 	Exist:   isExist[apiresponses.BusinessPartner](pmRes.Message.BusinessPartner),
	// 	Param:   pmRes.Message.BusinessPartner,
	// }

	// bpTextExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "BPText",
	// 	Exist:   isExist[apiresponses.ProductDescByBP](pmRes.Message.ProductDescByBP),
	// 	Param:   pmRes.Message.ProductDescByBP,
	// }
	// bpPlantExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "BPPlant",
	// 	Exist:   isExist[apiresponses.BPPlant](pmRes.Message.BPPlant),
	// 	Param:   pmRes.Message.BPPlant,
	// }
	// storageLocationExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "StorageLocation",
	// 	Exist:   isExist[apiresponses.StorageLocation](pmRes.Message.StorageLocation),
	// 	Param:   pmRes.Message.StorageLocation,
	// }

	// storageBinExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "StorageBin",
	// 	Exist:   isExist[apiresponses.PMStorageBin](pmRes.Message.StorageBin),
	// 	Param:   pmRes.Message.StorageBin,
	// }
	// mrpAreaExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "MRPArea",
	// 	Exist:   isExist[apiresponses.MRPArea](pmRes.Message.MRPArea),
	// 	Param:   pmRes.Message.MRPArea,
	// }
	// workScheduleExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "WorkSchedule",
	// 	Exist:   isExist[apiresponses.WorkScheduling](pmRes.Message.WorkScheduling),
	// 	Param:   pmRes.Message.WorkScheduling,
	// }

	// qualityExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "Quality",
	// 	Exist:   isExist[apiresponses.Quality](pmRes.Message.Quality),
	// 	Param:   pmRes.Message.Quality,
	// }
	// taxExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "Tax",
	// 	Exist:   isExist[apiresponses.Tax](pmRes.Message.Tax),
	// 	Param:   pmRes.Message.Tax,
	// }
	// accountingExconf := dpfm_api_output_formatter.Exconf{
	// 	Content: "Accounting",
	// 	Exist:   isExist[apiresponses.PMAccounting](pmRes.Message.Accounting),
	// 	Param:   pmRes.Message.Accounting,
	// }

	// 201@gmail.com/SupplyChainRelationship/list/user=BusinessPartner/SupplyChainRelationshipList
	key := fmt.Sprintf("%s/supplychainrelationship/list/user=%s/isMarkedForDeletion=%v/SupplyChainRelationshipList",
		*params.Params.UserID, *params.Params.User, false)
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
	data := dpfm_api_output_formatter.SupplyChainRelationshipDetailList{
		Header: dpfm_api_output_formatter.SupplyChainRelationshipDetailHeader{
			Index: idx,
			Key:   key,
		},
		Existences: []dpfm_api_output_formatter.SupplyChainRelationshipDetailExconf{
			generalExconf, deliveryExconf, deliveryPlantExconf,
			billingExconf, paymentExconf, transactionExconf,
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

func (c *SupplyChainRelationshipDetailListCtrl) finEmptyProcess(
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

func (c *SupplyChainRelationshipDetailListCtrl) Log(args ...interface{}) {
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
