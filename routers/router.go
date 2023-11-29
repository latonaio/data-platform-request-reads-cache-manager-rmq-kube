package routers

import (
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	controllersBatchMasterRecordList "data-platform-request-reads-cache-manager-rmq-kube/controllers/batch-master-record/list"
	controllersBillOfMaterialDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/list"
	controllersBusinessPartnerDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/detail-general"
	controllersBusinessPartnerList "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/list"
	controllersDeliveryDocumentDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/detail-list"
	controllersDeliveryDocumentList "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/list"
	controllersEquipmentMasterDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/detail-general"
	controllersEquipmentMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/list"
	controllersInvoiceDocumentDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/invoice-document/detail-list"
	controllersInvoiceDocumentList "data-platform-request-reads-cache-manager-rmq-kube/controllers/invoice-document/list"
	controllersOperationsDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/detail-list"
	controllersOperationsList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/list"
	controllersOrdersDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/detail-list"
	controllersOrdersItemPricingElement "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-pricing-element"
	controllersOrdersItemScheduleLine "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-schedule-line"
	controllersOrdersSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-single-unit"
	controllersOrdersList "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/list"
	controllersPlantDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/plant/detail-list"
	controllersPlantList "data-platform-request-reads-cache-manager-rmq-kube/controllers/plant/list"
	controllersPriceMasterDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/price-master/detail-list"
	controllersPriceMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/price-master/list"
	controllersProductMasterDetailBPPlant "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/detail-bp-plant"
	controllersProductMasterDetailBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/detail-business-partner"
	controllersProductMasterDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/detail-general"
	controllersProductMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/list"
	controllersProductSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/product-master-single-unit"
	controllersProductStockAvailabilityDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-availability-detail-list"
	controllersProductStockAvailabilityList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-availability-list"
	controllersProductStockByStorageBinByBatchList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-by-storage-bin-by-batch-list"
	controllersProductStockDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-detail-list"
	controllersProductStockList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-list"
	controllersProductStockSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-stock/product-stock-single-unit"
	controllersProductionOrderConfHeaderSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order-conf/header-single-unit"
	controllersProductionOrderDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/detail-list"
	controllersProductionOrderItemOperationList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/item-operation-list"
	controllersProductionOrderItemSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/item-single-unit"
	controllersProductionOrderList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/list"
	controllersPurchaseRequisitionList "data-platform-request-reads-cache-manager-rmq-kube/controllers/purchase-requisition/list"
	controllersQuotationsList "data-platform-request-reads-cache-manager-rmq-kube/controllers/quotations/list"
	controllersStorageBinList "data-platform-request-reads-cache-manager-rmq-kube/controllers/storage-bin/list"
	controllersSupplyChainRelationshipDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/supply-chain-relationship/detail-general"
	controllersSupplyChainRelationshipList "data-platform-request-reads-cache-manager-rmq-kube/controllers/supply-chain-relationship/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()

	redisCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 0, nil)
	//_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, "token")

	redisTokenCacheKeyPrefix := "tokens"

	_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, &redisTokenCacheKeyPrefix)
	//redisTokenCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1)

	quotationsListController := &controllersQuotationsList.QuotationsListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	purchaseRequisitionListController := &controllersPurchaseRequisitionList.PurchaseRequisitionListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersListController := &controllersOrdersList.OrdersListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersDetailListController := &controllersOrdersDetailList.OrdersDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersSingleUnit := &controllersOrdersSingleUnit.OrdersSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersItemScheduleLineController := &controllersOrdersItemScheduleLine.OrdersItemScheduleLineController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersItemPricingElementController := &controllersOrdersItemPricingElement.OrdersItemPricingElementController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	deliveryDocumentListController := &controllersDeliveryDocumentList.DeliveryDocumentListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	deliveryDocumentDetailListController := &controllersDeliveryDocumentDetailList.DeliveryDocumentDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	invoiceDocumentListController := &controllersInvoiceDocumentList.InvoiceDocumentListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	invoiceDocumentDetailListController := &controllersInvoiceDocumentDetailList.InvoiceDocumentDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	billOfMaterialListController := &controllersBillOfMaterialList.BillOfMaterialListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	billOfMaterialDetailListController := &controllersBillOfMaterialDetailList.BillOfMaterialDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	operationsListController := &controllersOperationsList.OperationsListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	operationsDetailListController := &controllersOperationsDetailList.OperationsDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterListController := &controllersProductMasterList.ProductMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterDetailGeneralController := &controllersProductMasterDetailGeneral.ProductMasterDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterDetailBusinessPartnerController := &controllersProductMasterDetailBusinessPartner.ProductMasterDetailBusinessPartnerController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterDetailBPPlantController := &controllersProductMasterDetailBPPlant.ProductMasterDetailBPPlantController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterSingleUnitController := &controllersProductSingleUnit.ProductMasterSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	priceMasterListController := &controllersPriceMasterList.PriceMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	priceMasterDetailListController := &controllersPriceMasterDetailList.PriceMasterDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	batchMasterRecordListController := &controllersBatchMasterRecordList.BatchMasterRecordListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockListController := &controllersProductStockList.ProductStockListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockDetailListController := &controllersProductStockDetailList.ProductStockDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockByStorageBinByBatchListController := &controllersProductStockByStorageBinByBatchList.ProductStockByStorageBinByBatchListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockSingleUnitController := &controllersProductStockSingleUnit.ProductStockSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockAvailabilityListController := &controllersProductStockAvailabilityList.ProductStockAvailabilityListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productStockAvailabilityDetailListController := &controllersProductStockAvailabilityDetailList.ProductStockAvailabilityDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	supplyChainRelationshipListController := &controllersSupplyChainRelationshipList.SupplyChainRelationshipListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	supplyChainRelationshipDetailGeneralController := &controllersSupplyChainRelationshipDetailGeneral.SupplyChainRelationshipDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	businessPartnerListController := &controllersBusinessPartnerList.BusinessPartnerListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	businessPartnerDetailGeneralController := &controllersBusinessPartnerDetailGeneral.BusinessPartnerDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	equipmentMasterListController := &controllersEquipmentMasterList.EquipmentMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	equipmentMasterDetailGeneralController := &controllersEquipmentMasterDetailGeneral.EquipmentMasterDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	plantListController := &controllersPlantList.PlantListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	plantDetailListController := &controllersPlantDetailList.PlantDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	storageBinListController := &controllersStorageBinList.StorageBinListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderDetailListController := &controllersProductionOrderDetailList.ProductionOrderDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderItemOperationListController := &controllersProductionOrderItemOperationList.ProductionOrderItemOperationListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderConfHeaderSingleUnitController := &controllersProductionOrderConfHeaderSingleUnit.ProductionOrderConfHeaderSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderListController := &controllersProductionOrderList.ProductionOrderListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderItemSingleUnitController := &controllersProductionOrderItemSingleUnit.ProductionOrderItemSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	quotations := beego.NewNamespace(
		"/quotations",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", quotationsListController),
	)

	purchaseRequisition := beego.NewNamespace(
		"/purchase-requisition",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", purchaseRequisitionListController),
	)

	orders := beego.NewNamespace(
		"/orders",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", ordersListController),
		beego.NSRouter("/detail/list/:userType", ordersDetailListController),
		beego.NSRouter("/item-single-unit/:userType", ordersSingleUnit),
		beego.NSRouter("/item-schedule-line/:userType", ordersItemScheduleLineController),
		beego.NSRouter("/item-pricing-element/:userType", ordersItemPricingElementController),
	)

	deliveryDocument := beego.NewNamespace(
		"/delivery-document",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", deliveryDocumentListController),
		beego.NSRouter("/detail/list/:userType", deliveryDocumentDetailListController),
	)

	invoiceDocument := beego.NewNamespace(
		"/invoice-document",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", invoiceDocumentListController),
		beego.NSRouter("/detail/list/:userType", invoiceDocumentDetailListController),
	)

	billOfMaterial := beego.NewNamespace(
		"/bill-of-material",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", billOfMaterialListController),
		beego.NSRouter("/detail/list/:userType", billOfMaterialDetailListController),
	)

	operations := beego.NewNamespace(
		"/operations",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", operationsListController),
		beego.NSRouter("/detail/list/:userType", operationsDetailListController),
	)

	productMaster := beego.NewNamespace(
		"/product-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", productMasterListController),
		beego.NSRouter("/list/:userType", productMasterDetailGeneralController),
		beego.NSRouter("/list/:userType", productMasterDetailBusinessPartnerController),
		beego.NSRouter("/list/:userType", productMasterDetailBPPlantController),
		beego.NSRouter("/product-single-unit/:userType", productMasterSingleUnitController),
	)

	priceMaster := beego.NewNamespace(
		"/price-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", priceMasterListController),
		beego.NSRouter("/detail/list/:userType", priceMasterDetailListController),
	)

	batchMasterRecord := beego.NewNamespace(
		"/batch-master-record",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", batchMasterRecordListController),
	)

	productStock := beego.NewNamespace(
		"/product-stock",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/product-stock-list/:userType", productStockListController),
		beego.NSRouter("/product-stock-detail/list/:userType", productStockDetailListController),
		beego.NSRouter("/product-stock-by-storage-bin-by-batch-list/:userType", productStockByStorageBinByBatchListController),
		beego.NSRouter("/product-stock-single-unit/:userType", productStockSingleUnitController),
		beego.NSRouter("/product-stock-availability-list/:userType", productStockAvailabilityListController),
		beego.NSRouter("/product-stock-availability-detail/list/:userType", productStockAvailabilityDetailListController),
	)

	supplyChainRelationship := beego.NewNamespace(
		"/supply-chain-relationship",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", supplyChainRelationshipListController),
		beego.NSRouter("/list/:userType", supplyChainRelationshipDetailGeneralController),
	)

	businessPartner := beego.NewNamespace(
		"/business-partner",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", businessPartnerListController),
		beego.NSRouter("/list/:userType", businessPartnerDetailGeneralController),
	)

	equipmentMaster := beego.NewNamespace(
		"/equipment-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", equipmentMasterListController),
		beego.NSRouter("/list/:userType", equipmentMasterDetailGeneralController),
	)

	plant := beego.NewNamespace(
		"/plant",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", plantListController),
		beego.NSRouter("/detail/list/:userType", plantDetailListController),
	)

	storageBin := beego.NewNamespace(
		"/storage-bin",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", storageBinListController),
	)

	productionOrder := beego.NewNamespace(
		"/production-order",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", productionOrderListController),
		beego.NSRouter("/item-single-unit/:userType", productionOrderItemSingleUnitController),
		beego.NSRouter("/detail/list/:userType", productionOrderDetailListController),
		beego.NSRouter("/item-operation/list", productionOrderItemOperationListController),
	)

	productionOrderConf := beego.NewNamespace(
		"/production-order-conf",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/header-single-unit/:userType", productionOrderConfHeaderSingleUnitController),
	)

	beego.AddNamespace(
		businessPartner,
		productMaster,
		quotations,
		purchaseRequisition,
		orders,
		deliveryDocument,
		invoiceDocument,
		billOfMaterial,
		operations,
		supplyChainRelationship,
		priceMaster,
		batchMasterRecord,
		productStock,
		equipmentMaster,
		plant,
		storageBin,
		productionOrder,
		productionOrderConf,
	)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		_ = ctx.Input.Header("Authorization")
		//jwtToken := ctx.Input.Header("Authorization")
		//
		//trimJwtToken := strings.TrimPrefix(jwtToken, "Bearer ")
		//
		//token, err := redisTokenCache.GetRaw(trimJwtToken)
		//
		//fmt.Sprintf("token: %v", token)
		//
		//if err == nil {
		//	return
		//}

		//services.VerifyToken(ctx, l, jwtToken)
	})

	//	beego.AddNamespace(billOfMaterial)

	//beego.Router("/:aPIServiceName/:aPIType", &controllers.APIModuleRuntimesController{})
	//beego.Router("/register", &controllers.RegisterController{})
	//beego.Router("/login", &controllers.LoginController{})
}
