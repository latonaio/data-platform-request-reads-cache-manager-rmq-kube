package deliverydocumentdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/deliverydocumentdetail"
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

type DeliveryDocumentDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewDeliveryDocumentDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *DeliveryDocumentDetailCtrl {
	return &DeliveryDocumentDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *DeliveryDocumentDetailCtrl) DeliveryDocumentDetail(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractDeliveryDocumentDetailParam(msg)
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

	var ddRes *apiresponses.DeliveryDocumentRes
	var pmRes *apiresponses.ProductMasterRes
	var pmdRes *apiresponses.ProductMasterDocRes
	var pgRes *apiresponses.ProductGroupRes
	var accountingRes *apiresponses.ProductMasterRes
	var oRes *apiresponses.OrdersRes
	var bpRes *apiresponses.BusinessPartnerRes
	var stockToRes *apiresponses.ProductStockRes
	var toPlantRes *apiresponses.PlantRes
	var tagRes *apiresponses.ProductTagRes
	var bmRes *apiresponses.BillOfMaterialRes
	var bmiRes *apiresponses.BillOfMaterialRes
	var materialProductRes *apiresponses.ProductMasterRes
	var sbRes *apiresponses.StorageBinRes

	defer func() {
		err = c.fin(params, ddRes, pmRes, pmdRes, pgRes, accountingRes, oRes, bpRes, stockToRes, toPlantRes, sbRes, materialProductRes, tagRes, reqKey, "DeliveryDocumentDetail", &cacheResult)
	}()

	ddRes, err = c.deliveryDocumentRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("deliveryDocumentRequest error: %w", err)
	}
	pmRes, err = c.productRequest(&params.Params, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return xerrors.Errorf("search header error: %w", err)
	}
	pmdRes, err = c.productMasterDocRequest(&params.Params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes)
		return err
	}
	pgRes, err = c.productGroupRequest(&params.Params, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes)
		return err
	}
	accountingRes, err = c.accountingRequest(&params.Params, pmRes, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes)
		return err
	}
	oRes, err = c.ordersRequest(&params.Params, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}
	stockToRes, err = c.productStockToRequest(&params.Params, ddRes, accountingRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes, accountingRes)
		return err
	}
	bpRes, err = c.businessPartnerRequest(&params.Params, pmRes, oRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, oRes)
		return err
	}

	bmRes, err = c.materialHeaderRequest(&params.Params, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes)
		return err
	}
	toPlantRes, err = c.deliverToPlantRequest(&params.Params, pmRes, ddRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(pmRes, ddRes)
		return err
	}

	tagRes, err = c.tagRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
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
	sbRes, err = c.storageBinRequest(&params.Params, ddRes, pmRes, sID, reqKey, &cacheResult)
	if err != nil {
		c.Log(ddRes, pmRes)
		return err
	}

	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())

	return nil
}

func (c *DeliveryDocumentDetailCtrl) materialProductRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	bmRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	pmReq := deliverydocumentdetail.CreateProductMaterialReq(params, bmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "MaterialProduct", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialProduct cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialProduct response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *DeliveryDocumentDetailCtrl) materialHeaderRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BillOfMaterialRes, error) {
	defer recovery(c.log)
	bmReq := deliverydocumentdetail.CreateBillOfMaterialHeaderReq(params, ddRes, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", bmReq, sID, reqKey, "MaterialHeader", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialHeader cache set error: %w", err)
	}
	bmRes, err := apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialHeader response parse error: %w", err)
	}
	return bmRes, nil
}

func (c *DeliveryDocumentDetailCtrl) materialItemRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	bmRes *apiresponses.BillOfMaterialRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BillOfMaterialRes, error) {
	defer recovery(c.log)
	bmReq := deliverydocumentdetail.CreateBillOfMaterialItemsReq(params, bmRes, sID, c.log)
	res, err := c.request("data-platform-api-bill-of-material-reads-queue", bmReq, sID, reqKey, "MaterialItem", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("MaterialItem cache set error: %w", err)
	}
	miRes, err := apiresponses.CreateBillOfMaterialRes(res)
	if err != nil {
		return nil, xerrors.Errorf("MaterialItem response parse error: %w", err)
	}
	return miRes, nil
}

func (c *DeliveryDocumentDetailCtrl) storageBinRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	bpPlantRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.StorageBinRes, error) {
	defer recovery(c.log)
	sbReq := deliverydocumentdetail.CrerateStorageBinReq(params, ddRes, bpPlantRes, sID, c.log)
	res, err := c.request("data-platform-api-storage-bin-reads-queue", sbReq, sID, reqKey, "StorageBin", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("StorageBin cache set error: %w", err)
	}
	sbRes, err := apiresponses.CreateStorageBinRes(res)
	if err != nil {
		return nil, xerrors.Errorf("StorageBin response parse error: %w", err)
	}
	return sbRes, nil
}

func (c *DeliveryDocumentDetailCtrl) accountingRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	bpPlantRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	accountingReq := deliverydocumentdetail.CrerateAccountingReq(params, pmRes, bpPlantRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", accountingReq, sID, reqKey, "Accounting", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Accounting cache set error: %w", err)
	}
	accountingRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Accounting response parse error: %w", err)
	}
	return accountingRes, nil
}

func (c *DeliveryDocumentDetailCtrl) productGroupRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductGroupRes, error) {
	defer recovery(c.log)
	pgReq := deliverydocumentdetail.CreateProductGroupReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-group-reads-queue", pgReq, sID, reqKey, "ProductGroup", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup cache set error: %w", err)
	}
	pgRes, err := apiresponses.CreateProductGroupRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup response parse error: %w", err)
	}
	return pgRes, nil
}

func (c *DeliveryDocumentDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	oiRes *apiresponses.OrdersRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpGeneralReq := deliverydocumentdetail.CreateBusinessPartnerReq(params, pmRes, oiRes, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-general-queue", bpGeneralReq, sID, reqKey, "BusinessPartnerGeneral", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartnerGeneral response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *DeliveryDocumentDetailCtrl) tagRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductTagRes, error) {
	defer recovery(c.log)
	ptReq := deliverydocumentdetail.CreateProductTagReq(params, sID, c.log)
	res, err := c.request("data-platform-api-product-tag-reads-queue", ptReq, sID, reqKey, "ProductTag", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	ptRes, err := apiresponses.CreateProductTagRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return ptRes, nil
}

func (c *DeliveryDocumentDetailCtrl) ordersRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.OrdersRes, error) {
	defer recovery(c.log)
	psReq := deliverydocumentdetail.CrerateOrdersReq(ddRes, sID, c.log)
	res, err := c.request("data-platform-api-orders-reads-queue", psReq, sID, reqKey, "OrdersItem", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("OrdersItem cache set error: %w", err)
	}
	oRes, err := apiresponses.CreateOrdersRes(res)
	if err != nil {
		return nil, xerrors.Errorf("OrdersItem response parse error: %w", err)
	}
	return oRes, nil
}

func (c *DeliveryDocumentDetailCtrl) productStockToRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	accountingRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductStockRes, error) {
	defer recovery(c.log)
	psReq := deliverydocumentdetail.CreateProductStockToReq(params, ddRes, accountingRes, sID, c.log)
	res, err := c.request("data-platform-api-product-stock-reads-queue", psReq, sID, reqKey, "ProductStockTo", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock cache set error: %w", err)
	}
	psRes, err := apiresponses.CreateProductStockRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductStock response parse error: %w", err)
	}
	return psRes, nil
}

func (c *DeliveryDocumentDetailCtrl) deliverToPlantRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	dtpReq := deliverydocumentdetail.CreateDeliverToPlantReq(params, ddRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", dtpReq, sID, reqKey, "DeliverToPlant", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant cache set error: %w", err)
	}
	dtpRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("DeliverToPlant response parse error: %w", err)
	}
	return dtpRes, nil
}

func (c *DeliveryDocumentDetailCtrl) productMasterDocRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	pmDocReq := deliverydocumentdetail.CreateProductMasterDocReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pmDocReq, sID, reqKey, "ProductMasterDoc", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc cache set error: %w", err)
	}
	pmdRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductMasterDoc response parse error: %w", err)
	}
	return pmdRes, nil
}

func (c *DeliveryDocumentDetailCtrl) deliveryDocumentRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.DeliveryDocumentRes, error) {
	defer recovery(c.log)
	ddReq := deliverydocumentdetail.CreateDeliveryDocumentReq(params, sID, c.log)
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

func (c *DeliveryDocumentDetailCtrl) productRequest(
	params *dpfm_api_input_reader.DeliveryDocumentDetailParams,
	ddRes *apiresponses.DeliveryDocumentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	pmReq := deliverydocumentdetail.CreateProductRequest(params, ddRes, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *DeliveryDocumentDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *DeliveryDocumentDetailCtrl) fin(
	params *dpfm_api_input_reader.DeliveryDocumentDetail,
	ddRes *apiresponses.DeliveryDocumentRes,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,

	pgRes *apiresponses.ProductGroupRes,
	aRes *apiresponses.ProductMasterRes,
	oRes *apiresponses.OrdersRes,
	bpRes *apiresponses.BusinessPartnerRes,

	stockToRes *apiresponses.ProductStockRes,
	toPlantRes *apiresponses.PlantRes,

	sbRes *apiresponses.StorageBinRes,
	materialProductRes *apiresponses.ProductMasterRes,

	tagRes *apiresponses.ProductTagRes,

	url, api string, setFlag *RedisCacheApiName,
) (err error) {
	data := dpfm_api_output_formatter.DeliveryDocumentDetail{}

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

	itemText := (*ddRes.Message.Item)[0].DeliveryDocumentItemTextByBuyer
	if params.Params.User == "DeliverToParty" {
		itemText = (*ddRes.Message.Item)[0].DeliveryDocumentItemTextBySeller
	}
	tags := c.CreateProductTag(tagRes)
	pInfo := c.CreateProductInfo(pgRes, aRes, bpRes, oRes, materialProductRes, pmRes)
	expirationDate := (*oRes.Message.Item)[0].DeliverFromPlantBatchValidityEndDate
	if params.Params.User == "DeliverToParty" {
		expirationDate = (*oRes.Message.Item)[0].DeliverToPlantBatchValidityEndDate
	}
	dItem := (*ddRes.Message.Item)[0]
	sbin := (*sbRes.Message.General)[0].StorageBin

	data.DeliveryDocument = params.Params.DeliveryDocument
	data.DeliveryDocumentItem = params.Params.DeliveryDocumentItem
	data.PlannedGoodsReceiptDate = dItem.PlannedGoodsReceiptDate
	data.PlannedGoodsReceiptTime = dItem.PlannedGoodsReceiptTime
	data.ActualGoodsReceiptDate = dItem.ActualGoodsReceiptDate
	data.ActualGoodsReceiptTime = dItem.ActualGoodsReceiptTime
	data.ProductName = *itemText
	data.ProductCode = (*pmRes.Message.General)[0].Product
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
	data.OrderQuantityInDelivery = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.ActualGoodsIssueQuantity,
		Unit:     *dItem.DeliveryUnit,
	}
	data.OrderQuantityInBase = dpfm_api_output_formatter.OrderQuantityInBase{
		Quantity: dItem.ActualGoodsIssueQtyInBaseUnit,
		Unit:     *dItem.BaseUnit,
	}
	data.ProductTag = *tags
	data.Stock = dpfm_api_output_formatter.Stock{
		ProductStock:    int(*stockToRes.Message.ProductStock.ProductStock),
		StorageLocation: *toPlantRes.Message.General.PlantFullName,
	}
	data.StorageLocationFullName = toPlantRes.Message.StorageLocation.StorageLocationFullName
	data.StorageBin = sbin
	data.BestByDate = nil
	data.ExpirationDate = expirationDate
	data.ProductInfo = pInfo
	data.BatchMgmtPolicyInDeliverToPlant = dItem.BatchMgmtPolicyInDeliverToPlant
	data.DeliverToPlantBatch = dItem.DeliverToPlantBatch
	data.DeliverToPlantBatchValidityStartDate = dItem.DeliverToPlantBatchValidityStartDate
	data.DeliverToPlantBatchValidityStartTime = dItem.DeliverToPlantBatchValidityStartTime
	data.DeliverToPlantBatchValidityEndDate = dItem.DeliverToPlantBatchValidityEndDate
	data.DeliverToPlantBatchValidityEndTime = dItem.DeliverToPlantBatchValidityEndTime
	data.BatchMgmtPolicyInDeliverFromPlant = dItem.BatchMgmtPolicyInDeliverFromPlant
	data.DeliverFromPlantBatch = dItem.DeliverFromPlantBatch
	data.DeliverFromPlantBatchValidityStartDate = dItem.DeliverFromPlantBatchValidityStartDate
	data.DeliverFromPlantBatchValidityStartTime = dItem.DeliverFromPlantBatchValidityStartTime
	data.DeliverFromPlantBatchValidityEndDate = dItem.DeliverFromPlantBatchValidityEndDate
	data.DeliverFromPlantBatchValidityEndTime = dItem.DeliverFromPlantBatchValidityEndTime
	data.PlannedGoodsIssueDate = dItem.PlannedGoodsIssueDate
	data.PlannedGoodsIssueTime = dItem.PlannedGoodsIssueTime
	data.PlannedGoodsIssueQuantity = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.PlannedGoodsIssueQuantity,
		Unit:     *dItem.DeliveryUnit,
	}
	data.PlannedGoodsIssueQtyInBaseUnit = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.PlannedGoodsIssueQtyInBaseUnit,
		Unit:     *dItem.BaseUnit,
	}
	data.PlannedGoodsReceiptQuantity = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.PlannedGoodsReceiptQuantity,
		Unit:     *dItem.DeliveryUnit,
	}
	data.PlannedGoodsReceiptQtyInBaseUnit = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.PlannedGoodsReceiptQtyInBaseUnit,
		Unit:     *dItem.BaseUnit,
	}
	data.ActualGoodsIssueDate = dItem.ActualGoodsIssueDate
	data.ActualGoodsIssueTime = dItem.ActualGoodsIssueTime
	data.ActualGoodsIssueQuantity = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.ActualGoodsIssueQuantity,
		Unit:     *dItem.DeliveryUnit,
	}
	data.ActualGoodsIssueQtyInBaseUnit = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.ActualGoodsIssueQtyInBaseUnit,
		Unit:     *dItem.BaseUnit,
	}
	data.ActualGoodsReceiptQuantity = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.ActualGoodsReceiptQuantity,
		Unit:     *dItem.DeliveryUnit,
	}
	data.ActualGoodsReceiptQtyInBaseUnit = dpfm_api_output_formatter.OrderQuantityInDelivery{
		Quantity: dItem.ActualGoodsReceiptQtyInBaseUnit,
		Unit:     *dItem.BaseUnit,
	}
	data.OrderDetailJumpReq = dpfm_api_output_formatter.OrderDetailJumpReq{
		OrderID:   (*ddRes.Message.Item)[0].OrderID,
		OrderItem: (*ddRes.Message.Item)[0].OrderItem,
		Product:   (*ddRes.Message.Item)[0].Product,
		Payer:     (*ddRes.Message.Item)[0].Payer,
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	return nil
}

func (c *DeliveryDocumentDetailCtrl) CreateProductTag(ptRes *apiresponses.ProductTagRes) *[]dpfm_api_output_formatter.ProductTag {
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

func (c *DeliveryDocumentDetailCtrl) CreateProductInfo(
	pgRes *apiresponses.ProductGroupRes,
	aRes *apiresponses.ProductMasterRes,
	bpRes *apiresponses.BusinessPartnerRes,
	oRes *apiresponses.OrdersRes,
	materialProductRes *apiresponses.ProductMasterRes,
	pmRes *apiresponses.ProductMasterRes,
) (d []dpfm_api_output_formatter.ProductInfo) {
	defer recovery(c.log)
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

func (c *DeliveryDocumentDetailCtrl) CreateProductInfoProductGroup(pgRes *apiresponses.ProductGroupRes) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "ProductGroupName",
		Key:     "商品分類",
		Value:   *(*pgRes.Message.ProductGroupText)[0].ProductGroupName,
	}
}

func (c *DeliveryDocumentDetailCtrl) CreateProductInfoPriceUnitQty(
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "PriceUnitQty",
		Key:     "価格単位",
		Value:   (*aRes.Message.Accounting)[0].PriceUnitQty,
	}
}

func (c *DeliveryDocumentDetailCtrl) CreateProductInfoBPName(
	bpRes *apiresponses.BusinessPartnerRes,
) dpfm_api_output_formatter.ProductInfo {
	return dpfm_api_output_formatter.ProductInfo{
		KeyName: "BusinessPartnerName",
		Key:     "製造者",
		Value:   bpRes.Message.General.BusinessPartnerFullName,
	}
}

func (c *DeliveryDocumentDetailCtrl) CreateProductInfoPrice(
	oRes *apiresponses.OrdersRes,
	aRes *apiresponses.ProductMasterRes,
) dpfm_api_output_formatter.ProductInfo {
	netAmount := 0
	grossAmount := 0

	for _, v := range *oRes.Message.ItemPricingElement {
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

func (c *DeliveryDocumentDetailCtrl) CreateInternalCapacityQuantity(
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
func (c *DeliveryDocumentDetailCtrl) CreateProductInfoMaterial(pmRes *apiresponses.ProductMasterRes) dpfm_api_output_formatter.ProductInfo {
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

func (c *DeliveryDocumentDetailCtrl) CreateProductInfoAllergen(pmRes *apiresponses.ProductMasterRes) dpfm_api_output_formatter.ProductInfo {
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

func extractDeliveryDocumentDetailParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.DeliveryDocumentDetail {
	data := dpfm_api_input_reader.ReadDeliveryDocumentDetail(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *DeliveryDocumentDetailCtrl) Log(args ...interface{}) {
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
