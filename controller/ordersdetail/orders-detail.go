package ordersdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/ordersdetail"
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

type OrdersDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewOrdersDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *OrdersDetailCtrl {
	return &OrdersDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *OrdersDetailCtrl) OrdersDetail(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractOrderDetailParam(msg)
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

	var oiRes *apiresponses.OrdersRes
	var pmRes *apiresponses.ProductMasterRes
	var pgRes *apiresponses.ProductGroupRes
	var bpPlantRes *apiresponses.ProductMasterRes
	var accountingRes *apiresponses.ProductMasterRes
	var bpRes *apiresponses.BusinessPartnerRes
	var pmdRes *apiresponses.ProductMasterDocRes
	var stockToRes *apiresponses.ProductStockRes
	var fromPlantRes *apiresponses.PlantRes
	var toPlantRes *apiresponses.PlantRes
	var stockFromRes *apiresponses.ProductStockRes
	var bmRes *apiresponses.BillOfMaterialRes
	var bmiRes *apiresponses.BillOfMaterialRes
	var materialProductRes *apiresponses.ProductMasterRes
	var tagRes *apiresponses.ProductTagRes
	var scRes *apiresponses.SupplyChainRelationshipRes

	defer func() {
		err = c.fin(
			params, pmRes, pmdRes, pgRes, bpRes, accountingRes, fromPlantRes, toPlantRes, stockFromRes, stockToRes, oiRes, materialProductRes, bpPlantRes, tagRes, scRes, reqKey, "OrdersDetail", &cacheResult,
		)
	}()

	oiRes, err = c.orderItemRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}
	pmRes, err = c.productRequest(&params.Params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(oiRes)
		return err
	}
	pgRes, err = c.productGroupRequest(&params.Params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes)
		return err
	}
	bpPlantRes, err = c.bpPlantRequest(&params.Params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(oiRes)
		return err
	}
	accountingRes, err = c.accountingRequest(&params.Params, pmRes, bpPlantRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, bpPlantRes)
		return err
	}
	bpRes, err = c.businessPartnerRequest(&params.Params, pmRes, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, oiRes)
		return err
	}
	pmdRes, err = c.productMasterDocRequest(&params.Params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes)
		return err
	}
	stockToRes, err = c.productStockToRequest(&params.Params, oiRes, accountingRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(oiRes, accountingRes)
		return err
	}

	scRes, err = c.supplyChainListRequest(&params.Params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	fromPlantRes, err = c.deliverFromPlantRequest(&params.Params, pmRes, oiRes, scRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, oiRes)
		return err
	}
	toPlantRes, err = c.deliverToPlantRequest(&params.Params, pmRes, oiRes, scRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, oiRes)
		return err
	}
	stockFromRes, err = c.productStockFromRequest(&params.Params, oiRes, accountingRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(oiRes, accountingRes)
		return err
	}
	bmRes, err = c.materialHeaderRequest(&params.Params, oiRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(oiRes)
		return err
	}
	bmiRes, err = c.materialItemRequest(&params.Params, bmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(bmRes)
		return err
	}
	materialProductRes, err = c.materialProductRequest(&params.Params, bmiRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(bmiRes)
		return err
	}
	tagRes, err = c.tagRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return err
	}

	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())

	return nil
}

func (c *OrdersDetailCtrl) supplyChainListRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (scRes *apiresponses.SupplyChainRelationshipRes, err error) {
	defer recovery(c.log, err)
	scReq := ordersdetail.CreateSupplyChainReq(params, oiRes, sID, c.log)
	res, err := c.request("data-platform-api-supply-chain-rel-master-reads-queue", scReq, sID, reqKey, "SupplyChainList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument cache set error: %w", err)
	}
	scRes, err = apiresponses.CreateSupplyChainRelationshipRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliveryDocument response parse error: %w", err)
	}
	return scRes, nil
}

func (c *OrdersDetailCtrl) materialProductRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	bmRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (pmRes *apiresponses.ProductMasterRes, err error) {
	defer recovery(c.log, err)
	pmReq := ordersdetail.CreateProductMaterialReq(params, bmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "MaterialProduct", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialProduct cache set error: %w", err)
	}
	pmRes, err = apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialProduct response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *OrdersDetailCtrl) materialHeaderRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (bmRes *apiresponses.BillOfMaterialRes, err error) {
	defer recovery(c.log, err)
	bmReq := ordersdetail.CreateBillOfMaterialHeaderReq(params, oRes, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", bmReq, sID, reqKey, "MaterialHeader", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialHeader cache set error: %w", err)
	}
	bmRes, err = apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialHeader response parse error: %w", err)
	}
	return bmRes, nil
}

func (c *OrdersDetailCtrl) materialItemRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	bmRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (miRes *apiresponses.BillOfMaterialRes, err error) {
	defer recovery(c.log, err)
	bmReq := ordersdetail.CreateBillOfMaterialItemsReq(params, bmRes, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", bmReq, sID, reqKey, "MaterialItem", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialItem cache set error: %w", err)
	}
	miRes, err = apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialItem response parse error: %w", err)
	}
	return miRes, nil
}

func (c *OrdersDetailCtrl) tagRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (ptRes *apiresponses.ProductTagRes, err error) {
	defer recovery(c.log, err)
	ptReq := ordersdetail.CreateProductTagReq(params, sID, c.log)
	res, err := c.request("data-platform-api-product-tag-reads-queue", ptReq, sID, reqKey, "ProductTag", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	ptRes, err = apiresponses.CreateProductTagRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return ptRes, nil
}

func (c *OrdersDetailCtrl) productRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (pmRes *apiresponses.ProductMasterRes, err error) {
	defer recovery(c.log, err)
	pmReq := ordersdetail.CreateProductRequest(params, oiRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err = apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *OrdersDetailCtrl) productGroupRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (pgRes *apiresponses.ProductGroupRes, err error) {
	defer recovery(c.log, err)
	pgReq := ordersdetail.CreateProductGroupReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-group-reads-queue", pgReq, sID, reqKey, "ProductGroup", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup cache set error: %w", err)
	}
	pgRes, err = apiresponses.CreateProductGroupRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup response parse error: %w", err)
	}
	return pgRes, nil
}

func (c *OrdersDetailCtrl) bpPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (bpPlantRes *apiresponses.ProductMasterRes, err error) {
	defer recovery(c.log, err)
	bpPlantReq := ordersdetail.CreateBPPlantReq(params, oRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", bpPlantReq, sID, reqKey, "BPPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BPPlant cache set error: %w", err)
	}
	bpPlantRes, err = apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BPPlant response parse error: %w", err)
	}
	return bpPlantRes, nil
}

func (c *OrdersDetailCtrl) accountingRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	bpPlantRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (accountingRes *apiresponses.ProductMasterRes, err error) {
	defer recovery(c.log, err)
	accountingReq := ordersdetail.CreateAccountingReq(params, pmRes, bpPlantRes, []string{
		"Accounting",
	}, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", accountingReq, sID, reqKey, "Accounting", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Accounting cache set error: %w", err)
	}
	accountingRes, err = apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Accounting response parse error: %w", err)
	}
	return accountingRes, nil
}

func (c *OrdersDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (bpRes *apiresponses.BusinessPartnerRes, err error) {
	defer recovery(c.log, err)
	bpGeneralReq := ordersdetail.CreateBusinessPartnerReq(params, pmRes, oiRes, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpGeneralReq, sID, reqKey, "BusinessPartnerGeneral", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral cache set error: %w", err)
	}
	bpRes, err = apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *OrdersDetailCtrl) productMasterDocRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (pmdRes *apiresponses.ProductMasterDocRes, err error) {
	defer recovery(c.log, err)
	pmDocReq := ordersdetail.CreateProductMasterDocReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pmDocReq, sID, reqKey, "ProductMasterDoc", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc cache set error: %w", err)
	}
	pmdRes, err = apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc response parse error: %w", err)
	}
	return pmdRes, nil
}

func (c *OrdersDetailCtrl) productStockFromRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	accountingRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (psRes *apiresponses.ProductStockRes, err error) {
	defer recovery(c.log, err)
	psReq := ordersdetail.CreateProductStockFromReq(oiRes, accountingRes, sID, c.log)
	res, err := c.request("data-platform-api-product-stock-reads-queue", psReq, sID, reqKey, "ProductStockFrom", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock cache set error: %w", err)
	}
	psRes, err = apiresponses.CreateProductStockRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock response parse error: %w", err)
	}
	return psRes, nil
}
func (c *OrdersDetailCtrl) productStockToRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	oiRes *apiresponses.OrdersRes,
	accountingRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (psRes *apiresponses.ProductStockRes, err error) {
	defer recovery(c.log, err)
	psReq := ordersdetail.CreateProductStockToReq(oiRes, accountingRes, sID, c.log)
	res, err := c.request("data-platform-api-product-stock-reads-queue", psReq, sID, reqKey, "ProductStockTo", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock cache set error: %w", err)
	}
	psRes, err = apiresponses.CreateProductStockRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock response parse error: %w", err)
	}
	return psRes, nil
}

func (c *OrdersDetailCtrl) orderItemRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (oiRes *apiresponses.OrdersRes, err error) {
	defer recovery(c.log, err)
	oiReq := ordersdetail.CreateOrdersItemReq(params, sID, c.log)
	res, err := c.request("data-platform-api-orders-reads-queue", oiReq, sID, reqKey, "OrderItem", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("OrderItem cache set error: %w", err)
	}
	oiRes, err = apiresponses.CreateOrdersRes(res)
	if err != nil {
		return nil, xerrors.Errorf("OrderItem response parse error: %w", err)
	}
	return oiRes, nil
}

func (c *OrdersDetailCtrl) deliverFromPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (dfpRes *apiresponses.PlantRes, err error) {
	defer recovery(c.log, err)
	dfpReq := ordersdetail.CreateDeliverFromPlantReq(params, oiRes, scRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", dfpReq, sID, reqKey, "DeliverFromPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant cache set error: %w", err)
	}
	dfpRes, err = apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverFromPlant response parse error: %w", err)
	}
	return dfpRes, nil
}

func (c *OrdersDetailCtrl) deliverToPlantRequest(
	params *dpfm_api_input_reader.OrdersDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	scRes *apiresponses.SupplyChainRelationshipRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (dtpRes *apiresponses.PlantRes, err error) {
	defer recovery(c.log, err)
	dtpReq := ordersdetail.CreateDeliverToPlantReq(params, oiRes, scRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", dtpReq, sID, reqKey, "DeliverToPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant cache set error: %w", err)
	}
	dtpRes, err = apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant response parse error: %w", err)
	}
	return dtpRes, nil
}

func (c *OrdersDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *OrdersDetailCtrl) fin(
	params *dpfm_api_input_reader.OrdersDetail,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	pgRes *apiresponses.ProductGroupRes,
	bpRes *apiresponses.BusinessPartnerRes,
	accountingRes *apiresponses.ProductMasterRes,
	fromPlantRes *apiresponses.PlantRes,
	toPlantRes *apiresponses.PlantRes,
	stockFromRes *apiresponses.ProductStockRes,
	stockToRes *apiresponses.ProductStockRes,
	oiRes *apiresponses.OrdersRes,
	materialProductRes *apiresponses.ProductMasterRes,
	bpPlantRes *apiresponses.ProductMasterRes,
	tagRes *apiresponses.ProductTagRes,
	scRes *apiresponses.SupplyChainRelationshipRes,

	url, api string, setFlag *RedisCacheApiName,
) (err error) {
	data := dpfm_api_output_formatter.OrdersDetail{}

	defer func() {
		redisKey := strings.Join([]string{url, api}, "/")
		b, _ := json.Marshal(data)
		err = c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
		if err != nil {
			return
		}
		(*setFlag)["redisCacheApiName"][api] = map[string]string{
			"keyName": redisKey,
		}
	}()

	// fromPlantMapper := map[int]apiresponses.PlantGeneral{}
	// toPlantMapper := map[int]apiresponses.PlantGeneral{}
	// {
	// 	v := *fromPlantRes.Message.General
	// 	fromPlantMapper[v.BusinessPartner] = v
	// 	v = *toPlantRes.Message.General
	// 	toPlantMapper[v.BusinessPartner] = v
	// }

	fromPlantMapper := map[int]apiresponses.PlantGeneral{}
	for _, v := range *fromPlantRes.Message.Generals {
		fromPlantMapper[v.BusinessPartner] = v
	}
	toPlantMapper := map[int]apiresponses.PlantGeneral{}
	for _, v := range *toPlantRes.Message.Generals {
		toPlantMapper[v.BusinessPartner] = v
	}

	deliver := map[int]dpfm_api_output_formatter.Deliver{}
	for _, dr := range *scRes.Message.DeliveryRelation {
		if _, ok := deliver[dr.SupplyChainRelationshipID]; !ok {
			deliver[dr.SupplyChainRelationshipID] = dpfm_api_output_formatter.Deliver{
				DeliverFromParty: make([]dpfm_api_output_formatter.DeliveryPlant, 0),
				DeliverToParty:   make([]dpfm_api_output_formatter.DeliveryPlant, 0),
			}
		}

		if _, ok := fromPlantMapper[dr.DeliverFromParty]; ok {
			tmp := deliver[dr.SupplyChainRelationshipID]
			tmp.DeliverFromParty = append(tmp.DeliverFromParty, dpfm_api_output_formatter.DeliveryPlant{
				DeliverFromPlant:     fromPlantMapper[dr.DeliverFromParty].Plant,
				DeliverFromPlantName: fromPlantMapper[dr.DeliverFromParty].PlantName,
				DeliverFromParty:     dr.DeliverFromParty,
				DefaultRelation:      *dr.DefaultRelation,
			},
			)
			deliver[dr.SupplyChainRelationshipID] = tmp
		}

		if _, ok := toPlantMapper[dr.DeliverToParty]; ok {
			tmp := deliver[dr.SupplyChainRelationshipID]
			tmp.DeliverToParty = append(tmp.DeliverToParty, dpfm_api_output_formatter.DeliveryPlant{
				DeliverToPlant:     toPlantMapper[dr.DeliverToParty].Plant,
				DeliverToPlantName: toPlantMapper[dr.DeliverToParty].PlantName,
				DeliverToParty:     dr.DeliverToParty,
				DefaultRelation:    *dr.DefaultRelation,
			},
			)
			deliver[dr.SupplyChainRelationshipID] = tmp
		}
	}

	data.Pulldown = dpfm_api_output_formatter.OrdersDetailPullDown{
		SupplyChains: deliver,
	}

	tags := c.CreateProductTag(tagRes)
	pInfo := c.CreateProductInfo(pgRes, accountingRes, oiRes, bpRes, materialProductRes, pmRes)

	oItem := (*oiRes.Message.Item)[0]

	data.ProductName = *(*pmRes.Message.ProductDescByBP)[0].ProductDescription
	data.ProductCode = (*pmRes.Message.ProductDescByBP)[0].Product
	data.ProductInfo = pInfo
	data.ProductTag = tags
	pmd := (*pmdRes.Message.HeaderDoc)[0]
	data.Images = dpfm_api_output_formatter.Images{
		Product: &dpfm_api_output_formatter.ProductImage{
			BusinessPartnerID: (*pmRes.Message.BusinessPartner)[0].BusinessPartner,
			DocID:             pmd.DocID,
			FileExtension:     pmd.FileExtension,
		},
		Barcord: &dpfm_api_output_formatter.BarcordImage{
			ID:          *(*pmRes.Message.General)[0].ProductStandardID,
			Barcode:     "",
			BarcodeType: *(*pmRes.Message.General)[0].BarcodeType,
		},
	}
	data.Stock = dpfm_api_output_formatter.Stock{
		ProductStock:    int(*stockToRes.Message.ProductStock.ProductStock),
		StorageLocation: *toPlantRes.Message.General.PlantFullName,
	}
	data.AvailabilityStock = dpfm_api_output_formatter.Stock{
		ProductStock:    int(*stockFromRes.Message.ProductStock.ProductStock),
		StorageLocation: *fromPlantRes.Message.General.PlantFullName,
	}
	data.OrderQuantityInDelivery = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: oItem.OrderQuantityInDeliveryUnit,
		Unit:     *oItem.DeliveryUnit,
	}
	data.OrderQuantityInBase = dpfm_api_output_formatter.OrderQuantityInBase{
		Quantity: oItem.OrderQuantityInBaseUnit,
		Unit:     *oItem.BaseUnit,
	}
	schedule := (*oiRes.Message.ItemScheduleLine)[0]
	data.ConfirmedOrderQuantityByPDTAvailCheck = dpfm_api_output_formatter.ConfirmedOrderQuantityByPDTAvailCheck{
		Quantity: schedule.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
		Unit:     *oItem.BaseUnit,
	}
	data.Popup = dpfm_api_output_formatter.Popup{
		RequestedDeliveryDate: *schedule.RequestedDeliveryDate,
		RequestedDeliveryTime: *schedule.RequestedDeliveryTime,
		ConfirmedDeliveryDate: *schedule.ConfirmedDeliveryDate,
		// ConfirmedDeliveryTime:       *schedule.ConfirmedDeliveryTime,
		OrderQuantityInBaseUnit:     int(*oItem.OrderQuantityInBaseUnit),
		BaseUnit:                    *oItem.BaseUnit,
		OrderQuantityInDeliveryUnit: int(*oItem.OrderQuantityInDeliveryUnit),
		DeliveryUnit:                *oItem.DeliveryUnit,
		ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: int(*schedule.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit),
		DeliverToPlantBatch:                             oItem.DeliverToPlantBatch,
		BatchMgmtPolicyInDeliverToPlant:                 oItem.BatchMgmtPolicyInDeliverToPlant,
		DeliverToPlantBatchValidityStartDate:            oItem.DeliverToPlantBatchValidityStartDate,
		DeliverToPlantBatchValidityStartTime:            oItem.DeliverToPlantBatchValidityStartTime,
		DeliverToPlantBatchValidityEndDate:              oItem.DeliverToPlantBatchValidityEndDate,
		DeliverToPlantBatchValidityEndTime:              oItem.DeliverToPlantBatchValidityEndTime,
		DeliverFromPlantBatch:                           oItem.DeliverFromPlantBatch,
		BatchMgmtPolicyInDeliverFromPlant:               oItem.BatchMgmtPolicyInDeliverFromPlant,
		DeliverFromPlantBatchValidityStartDate:          oItem.DeliverFromPlantBatchValidityStartDate,
		DeliverFromPlantBatchValidityStartTime:          oItem.DeliverFromPlantBatchValidityStartTime,
		DeliverFromPlantBatchValidityEndDate:            oItem.DeliverFromPlantBatchValidityEndDate,
		DeliverFromPlantBatchValidityEndTime:            oItem.DeliverFromPlantBatchValidityEndTime,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	return nil
}

func (c *OrdersDetailCtrl) CreateProductTag(ptRes *apiresponses.ProductTagRes) *[]dpfm_api_output_formatter.ProductTag {
	if ptRes == nil || ptRes.Message.ProductTag == nil {
		return &[]dpfm_api_output_formatter.ProductTag{}
	}
	tags := make([]dpfm_api_output_formatter.ProductTag, 0, len(*ptRes.Message.ProductTag))
	for _, v := range *ptRes.Message.ProductTag {
		tags = append(tags, dpfm_api_output_formatter.ProductTag{
			Key:      v.Key,
			DocCount: v.DocCount,
		},
		)
	}
	return &tags
}

func (c *OrdersDetailCtrl) CreateProductInfo(
	pgRes *apiresponses.ProductGroupRes,
	aRes *apiresponses.ProductMasterRes,
	oRes *apiresponses.OrdersRes,
	bpRes *apiresponses.BusinessPartnerRes,
	materialProductRes *apiresponses.ProductMasterRes,
	pmRes *apiresponses.ProductMasterRes,
) (d []dpfm_api_output_formatter.ProductInfo) {
	defer recovery(c.log, nil)
	d = make([]dpfm_api_output_formatter.ProductInfo, 0, 4)
	d = append(d, c.CreateProductInfoProductGroup(pgRes))
	// d = append(d, c.CreateProductInfoPriceUnitQty(aRes))
	d = append(d, c.CreateInternalCapacityQuantity(oRes))
	d = append(d, c.CreateProductInfoPrice(oRes, aRes))
	d = append(d, c.CreateProductInfoBPName(bpRes))
	d = append(d, c.CreateProductInfoMaterial(materialProductRes))
	d = append(d, c.CreateProductInfoAllergen(pmRes))
	return d
}

func (c *OrdersDetailCtrl) CreateProductInfoProductGroup(pgRes *apiresponses.ProductGroupRes) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "ProductGroupName",
		Key:     "商品分類",
		Value:   *(*pgRes.Message.ProductGroupText)[0].ProductGroupName,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoPriceUnitQty(
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "PriceUnitQty",
		Key:     "価格単位",
		Value:   (*aRes.Message.Accounting)[0].PriceUnitQty,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoBPName(
	bpRes *apiresponses.BusinessPartnerRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "BusinessPartnerName",
		Key:     "製造者",
		Value:   bpRes.Message.General.BusinessPartnerFullName,
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoPrice(
	oRes *apiresponses.OrdersRes,
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	netAmount := 0
	grossAmount := 0

	for _, v := range *oRes.Message.ItemPricingElement {
		if v.ConditionType == nil {
			continue
		}
		switch *v.ConditionType {
		case "PR00":
			netAmount = int(*v.ConditionRateValue)
		case "MWST":
			grossAmount = int(*v.ConditionRateValue)
		}
	}

	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "Price",
		Key:     "価格",
		Value: map[string]interface{}{
			"PriceWithoutTax": 0,
			"NetAmount":       netAmount,
			"PriceUnitQty":    (*aRes.Message.Accounting)[0].PriceUnitQty,
			"GrossAmount":     netAmount + grossAmount,
		},
	}
}

func (c *OrdersDetailCtrl) CreateInternalCapacityQuantity(
	oRes *apiresponses.OrdersRes,
) dpfm_api_output_formatter.ProductInfo {

	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "InternalCapacityQuantity",
		Key:     "内容量",
		Value: map[string]interface{}{
			"InternalCapacityQuantity":     (*oRes.Message.Item)[0].InternalCapacityQuantity,
			"InternalCapacityQuantityUnit": (*oRes.Message.Item)[0].InternalCapacityQuantityUnit,
		},
	}
}

func (c *OrdersDetailCtrl) CreateProductInfoMaterial(pmRes *apiresponses.ProductMasterRes) dpfm_api_output_formatter.ProductInfo {
	materials := make([]string, 0, len(*pmRes.Message.ProductDescByBP))
	for _, v := range *pmRes.Message.ProductDescByBP {
		materials = append(materials, *v.ProductDescription)
	}
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "Material",
		Key:     "原材料",
		Value:   materials,
	}
}
func (c *OrdersDetailCtrl) CreateProductInfoAllergen(pmRes *apiresponses.ProductMasterRes) dpfm_api_output_formatter.ProductInfo {
	allergens := make([]dpfm_api_output_formatter.Allergen, 0, len(*pmRes.Message.Allergen))
	for _, v := range *pmRes.Message.Allergen {
		allergens = append(allergens, dpfm_api_output_formatter.Allergen{
			Name:    v.Allergen,
			Contain: v.AllergenIsContained,
		})
	}
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "Allergen",
		Key:     "アレルゲン",
		Value:   allergens,
	}
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

func extractOrderDetailParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.OrdersDetail {
	data := dpfm_api_input_reader.ReadOrdersDetail(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *OrdersDetailCtrl) Log(args ...interface{}) {
	for _, v := range args {
		b, _ := json.Marshal(v)
		c.log.Error(string(b))
	}
}

func recovery(l *logger.Logger, err error) {
	if e := recover(); e != nil {
		err = xerrors.Errorf("recovered: %+v", e)
		return
	}
}
