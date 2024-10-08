package routers

import (
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	controllersAccountInfo "data-platform-request-reads-cache-manager-rmq-kube/controllers/account-info"
	controllersAfterLoginUserInfo "data-platform-request-reads-cache-manager-rmq-kube/controllers/after-login/user-info"
	controllersAfterPointAcquisition "data-platform-request-reads-cache-manager-rmq-kube/controllers/point-acquisition/after-point-acquisition"
	controllersArticleCreatesSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/article/creates-single-unit"
	controllersArticleListForOwners "data-platform-request-reads-cache-manager-rmq-kube/controllers/lists-for-owners/article-list"
	controllersArticleListForPointUsers "data-platform-request-reads-cache-manager-rmq-kube/controllers/lists-for-point-users/article-list"
	controllersArticleSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/article/single-unit"
	controllersAttendanceList "data-platform-request-reads-cache-manager-rmq-kube/controllers/attendance/list"
	controllersAttendanceListForOwnersByEvent "data-platform-request-reads-cache-manager-rmq-kube/controllers/attendance/list-for-owners-by-event"
	controllersBatchMasterRecordList "data-platform-request-reads-cache-manager-rmq-kube/controllers/batch-master-record/list"
	controllersBillOfMaterialDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialHeaderSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/header-single-unit"
	controllersBillOfMaterialList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/list"
	controllersBusinessPartnerDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/detail-general"
	controllersBusinessPartnerList "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/list"
	controllersCertificateAuthorityChainWithUsageControlChain "data-platform-request-reads-cache-manager-rmq-kube/controllers/certificate-authority-chain-with-usage-control-chain/certificate-authority-chain-with-usage-control-chain"
	controllersContentListForOwners "data-platform-request-reads-cache-manager-rmq-kube/controllers/lists-for-owners/content-list"
	controllersContentListForPointUsers "data-platform-request-reads-cache-manager-rmq-kube/controllers/lists-for-point-users/content-list"
	controllersDeliveryDocumentDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/detail-list"
	controllersDeliveryDocumentDetailListForADeliveryInstruction "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/detail-list-for-a-delivery-instruction"
	controllersDeliveryDocumentItem "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/item"
	controllersDeliveryDocumentSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/item-single-unit"
	controllersDeliveryDocumentList "data-platform-request-reads-cache-manager-rmq-kube/controllers/delivery-document/list"
	controllersEquipmentMasterDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/detail-general"
	controllersEquipmentMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/list"
	controllersEventCreatesSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/event/creates-single-unit"
	controllersEventSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/event/single-unit"
	controllersFriendList "data-platform-request-reads-cache-manager-rmq-kube/controllers/friend/list"
	controllersInspectionLotComponentComposition "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/component-composition"
	controllersInspectionLotSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/header-single-unit"
	controllersInspectionLotSingleUnitMillBoxInterface "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/header-single-unit-mill-box-interface"
	controllersInspectionLotSingleUnitMillSheet "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/header-single-unit-mill-sheet"
	controllersInspectionLotInspection "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/inspection"
	controllersInspectionLotList "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/list"
	controllersInspectionLotSpecDetail "data-platform-request-reads-cache-manager-rmq-kube/controllers/inspection-lot/spec-detail"
	controllersInstagramAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/controllers/instagram/authenticator"
	controllersInvoiceDocumentDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/invoice-document/detail-list"
	controllersInvoiceDocumentList "data-platform-request-reads-cache-manager-rmq-kube/controllers/invoice-document/list"
	controllersLoginAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/controllers/login/authenticator"
	controllersLoginAuthenticatorInitialEmailAuth "data-platform-request-reads-cache-manager-rmq-kube/controllers/login/authenticator-initial-email-auth"
	controllersLoginAuthenticatorInitialGoogleAccountAuth "data-platform-request-reads-cache-manager-rmq-kube/controllers/login/authenticator-initial-google-account-auth"
	controllersLoginAuthenticatorInitialSMSAuth "data-platform-request-reads-cache-manager-rmq-kube/controllers/login/authenticator-initial-sms-auth"
	controllersUserInfoCreates "data-platform-request-reads-cache-manager-rmq-kube/controllers/login/user-info-creates"
	controllersMessageInteractionWithFriend "data-platform-request-reads-cache-manager-rmq-kube/controllers/message/interaction-with-friend"
	controllersMessageList "data-platform-request-reads-cache-manager-rmq-kube/controllers/message/list"
	controllersOperationsDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/detail-list"
	controllersOperationsList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/list"
	controllersOrdersDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/detail-list"
	controllersOrdersDetailListForAnOrderDocument "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/detail-list-for-an-order-document"
	controllersOrdersItem "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item"
	controllersOrdersItemPricingElement "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-pricing-element"
	controllersOrdersItemScheduleLine "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-schedule-line"
	controllersOrdersSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-single-unit"
	controllersOrdersItemSingleUnitMillSheet "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/item-single-unit-mill-sheet"
	controllersOrdersList "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/list"
	controllersOrdersPartnersWithAddress "data-platform-request-reads-cache-manager-rmq-kube/controllers/orders/partners-with-address"
	controllersParticipationList "data-platform-request-reads-cache-manager-rmq-kube/controllers/participation/list"
	controllersParticipationListForOwnersByEvent "data-platform-request-reads-cache-manager-rmq-kube/controllers/participation/list-for-owners-by-event"
	controllersParticipationSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/participation/single-unit"
	controllersPlantDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/plant/detail-list"
	controllersPlantList "data-platform-request-reads-cache-manager-rmq-kube/controllers/plant/list"
	controllersPointBalanceSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/point-balance/single-unit"
	controllersPointConsumption "data-platform-request-reads-cache-manager-rmq-kube/controllers/point-consumption"
	controllersPointTransactionList "data-platform-request-reads-cache-manager-rmq-kube/controllers/point-transaction/list"
	controllersPointTransactionSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/point-transaction/single-unit"
	controllersPostListByFriend "data-platform-request-reads-cache-manager-rmq-kube/controllers/post/list-by-friend"
	controllersPostListMe "data-platform-request-reads-cache-manager-rmq-kube/controllers/post/list-me"
	controllersPostSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/post/single-unit"
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
	controllersProductionOrderHeaderSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/header-single-unit"
	controllersProductionOrderItemComponentList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/item-component-list"
	controllersProductionOrderItemOperationList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/item-operation-list"
	controllersProductionOrderItemSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/item-single-unit"
	controllersProductionOrderList "data-platform-request-reads-cache-manager-rmq-kube/controllers/production-order/list"
	controllersPurchaseRequisitionList "data-platform-request-reads-cache-manager-rmq-kube/controllers/purchase-requisition/list"
	controllersQRCodeListForOwners "data-platform-request-reads-cache-manager-rmq-kube/controllers/qr-code/list-for-owners"
	controllersQuotationsList "data-platform-request-reads-cache-manager-rmq-kube/controllers/quotations/list"
	controllersSiteCreatesSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/site/creates-single-unit"
	controllersSiteList "data-platform-request-reads-cache-manager-rmq-kube/controllers/site/list"
	controllersSiteSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/site/single-unit"
	controllersShopCreatesSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/shop/creates-single-unit"
	controllersShopList "data-platform-request-reads-cache-manager-rmq-kube/controllers/shop/list"
	controllersShopListForOwners "data-platform-request-reads-cache-manager-rmq-kube/controllers/lists-for-owners/shop-list"
	controllersShopSingleUnit "data-platform-request-reads-cache-manager-rmq-kube/controllers/shop/single-unit"
	controllersStorageBinList "data-platform-request-reads-cache-manager-rmq-kube/controllers/storage-bin/list"
	controllersSupplyChainRelationshipDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/supply-chain-relationship/detail-general"
	controllersSupplyChainRelationshipList "data-platform-request-reads-cache-manager-rmq-kube/controllers/supply-chain-relationship/list"
	controllersUserProfile "data-platform-request-reads-cache-manager-rmq-kube/controllers/user-profile"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"net/http"
)

type HealthCheckController struct {
	beego.Controller
}

func (c *HealthCheckController) Get() {
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]string{"status": "healthy"}
	c.ServeJSON()
}

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()

	redisCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 0, nil)
	//_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, "token")

	redisTokenCacheKeyPrefix := "tokens"

	_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, &redisTokenCacheKeyPrefix)
	//redisTokenCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1)

	loginAuthenticatorController := &controllersLoginAuthenticator.LoginAuthenticatorController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	loginAuthenticatorInitialEmailAuthController := &controllersLoginAuthenticatorInitialEmailAuth.LoginAuthenticatorInitialEmailAuthController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	loginAuthenticatorInitialSMSAuthController := &controllersLoginAuthenticatorInitialSMSAuth.LoginAuthenticatorInitialSMSAuthController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	loginAuthenticatorInitialGoogleAccountAuthController := &controllersLoginAuthenticatorInitialGoogleAccountAuth.LoginAuthenticatorInitialGoogleAccountAuthController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	instagramAuthenticatorController := &controllersInstagramAuthenticator.InstagramAuthenticatorController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	afterLoginUserInfoController := &controllersAfterLoginUserInfo.AfterLoginUserInfoController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	contentListForPointUsersController := &controllersContentListForPointUsers.ContentListForPointUsersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	contentListForOwnersController := &controllersContentListForOwners.ContentListForOwnersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	articleListForPointUsersController := &controllersArticleListForPointUsers.ArticleListForPointUsersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	articleListForOwnersController := &controllersArticleListForOwners.ArticleListForOwnersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}
	
	shopListForOwnersController := &controllersShopListForOwners.ShopListForOwnersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	eventSingleUnitController := &controllersEventSingleUnit.EventSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	eventCreatesSingleUnitController := &controllersEventCreatesSingleUnit.EventCreatesSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	articleSingleUnitController := &controllersArticleSingleUnit.ArticleSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	articleCreatesSingleUnitController := &controllersArticleCreatesSingleUnit.ArticleCreatesSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	siteListController := &controllersSiteList.SiteListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	siteSingleUnitController := &controllersSiteSingleUnit.SiteSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	siteCreatesSingleUnitController := &controllersSiteCreatesSingleUnit.SiteCreatesSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	shopListController := &controllersShopList.ShopListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	shopSingleUnitController := &controllersShopSingleUnit.ShopSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	shopCreatesSingleUnitController := &controllersShopCreatesSingleUnit.ShopCreatesSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	participationListController := &controllersParticipationList.ParticipationListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	participationListForOwnersByEventController := &controllersParticipationListForOwnersByEvent.ParticipationListForOwnersByEventController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	participationSingleUnitController := &controllersParticipationSingleUnit.ParticipationSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	attendanceListController := &controllersAttendanceList.AttendanceListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	attendanceListForOwnersByEventController := &controllersAttendanceListForOwnersByEvent.AttendanceListForOwnersByEventController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	pointTransactionListController := &controllersPointTransactionList.PointTransactionListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	pointTransactionSingleUnitController := &controllersPointTransactionSingleUnit.PointTransactionSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	pointBalanceSingleUnitController := &controllersPointBalanceSingleUnit.PointBalanceSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	pointConsumptionController := &controllersPointConsumption.PointConsumptionController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	postListByFriendController := &controllersPostListByFriend.PostListByFriendController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	postListMeController := &controllersPostListMe.PostListMeController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	postSingleUnitController := &controllersPostSingleUnit.PostSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	afterPointAcquisitionController := &controllersAfterPointAcquisition.AfterPointAcquisitionController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	userInfoCreatesController := &controllersUserInfoCreates.UserInfoCreatesController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	userProfileController := &controllersUserProfile.UserProfileController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	accountInfoController := &controllersAccountInfo.AccountInfoController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	friendListController := &controllersFriendList.FriendListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	messageListController := &controllersMessageList.MessageListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	messageInteractionWithFriendController := &controllersMessageInteractionWithFriend.MessageInteractionWithFriendController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	qRCodeListForOwnersController := &controllersQRCodeListForOwners.QRCodeListForOwnersController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

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

	ordersDetailListForAnOrderDocumentController := &controllersOrdersDetailListForAnOrderDocument.OrdersDetailListForAnOrderDocumentController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersSingleUnit := &controllersOrdersSingleUnit.OrdersSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersItemSingleUnitMillSheet := &controllersOrdersItemSingleUnitMillSheet.OrdersItemSingleUnitMillSheetController{
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

	ordersItemController := &controllersOrdersItem.OrdersItemController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	ordersPartnersWithAddressController := &controllersOrdersPartnersWithAddress.OrdersPartnersWithAddressController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotSingleUnitController := &controllersInspectionLotSingleUnit.InspectionLotSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotSingleUnitMillSheetController := &controllersInspectionLotSingleUnitMillSheet.InspectionLotSingleUnitMillSheetController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotSingleUnitMillBoxInterfaceController := &controllersInspectionLotSingleUnitMillBoxInterface.InspectionLotSingleUnitMillBoxInterfaceController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotListController := &controllersInspectionLotList.InspectionLotListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotSpecDetailController := &controllersInspectionLotSpecDetail.InspectionLotSpecDetailController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotComponentCompositionController := &controllersInspectionLotComponentComposition.InspectionLotComponentCompositionController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	inspectionLotInspectionController := &controllersInspectionLotInspection.InspectionLotInspectionController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	certificateAuthorityChainWithUsageControlChainController := &controllersCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainController{
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

	deliveryDocumentDetailListForADeliveryInstructionController := &controllersDeliveryDocumentDetailListForADeliveryInstruction.DeliveryDocumentDetailListForADeliveryInstructionController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	deliveryDocumentSingleUnitController := &controllersDeliveryDocumentSingleUnit.DeliveryDocumentSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	deliveryDocumentItemController := &controllersDeliveryDocumentItem.DeliveryDocumentItemController{
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

	billOfMaterialHeaderSingleUnitController := &controllersBillOfMaterialHeaderSingleUnit.BillOfMaterialHeaderSingleUnitController{
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

	productionOrderListController := &controllersProductionOrderList.ProductionOrderListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderDetailListController := &controllersProductionOrderDetailList.ProductionOrderDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderHeaderSingleUnitController := &controllersProductionOrderHeaderSingleUnit.ProductionOrderHeaderSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderItemSingleUnitController := &controllersProductionOrderItemSingleUnit.ProductionOrderItemSingleUnitController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productionOrderItemComponentListController := &controllersProductionOrderItemComponentList.ProductionOrderItemComponentListController{
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

	loginAuthenticator := beego.NewNamespace(
		"/login-authenticator",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/login-authenticator", loginAuthenticatorController),
		beego.NSRouter("/login-authenticator-initial-email-auth", loginAuthenticatorInitialEmailAuthController),
		beego.NSRouter("/login-authenticator-initial-sms-auth", loginAuthenticatorInitialSMSAuthController),
		beego.NSRouter("/login-authenticator-initial-google-account-auth", loginAuthenticatorInitialGoogleAccountAuthController),
	)

	instagram := beego.NewNamespace(
		"/instagram",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/authenticator", instagramAuthenticatorController),
	)

	afterLoginUserInfo := beego.NewNamespace(
		"/after-login-user-info",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/after-login-user-info", afterLoginUserInfoController),
	)

	listsForPointUsers := beego.NewNamespace(
		"/lists-for-point-users",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/content-list-for-point-users", contentListForPointUsersController),
		beego.NSRouter("/article-list-for-point-users", articleListForPointUsersController),
	)

	listsForOwners := beego.NewNamespace(
		"/lists-for-owners",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/content-list-for-owners", contentListForOwnersController),
		beego.NSRouter("/article-list-for-owners", articleListForOwnersController),
		beego.NSRouter("/shop-list-for-owners", shopListForOwnersController),
	)

	event := beego.NewNamespace(
		"/event",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/single-unit", eventSingleUnitController),
		beego.NSRouter("/creates-single-unit", eventCreatesSingleUnitController),
	)

	article := beego.NewNamespace(
		"/article",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/single-unit", articleSingleUnitController),
		beego.NSRouter("/creates-single-unit", articleCreatesSingleUnitController),
	)

	site := beego.NewNamespace(
		"/site",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", siteListController),
		beego.NSRouter("/single-unit", siteSingleUnitController),
		beego.NSRouter("/creates-single-unit", siteCreatesSingleUnitController),
	)

	shop := beego.NewNamespace(
		"/shop",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", shopListController),
		beego.NSRouter("/single-unit", shopSingleUnitController),
		beego.NSRouter("/creates-single-unit", shopCreatesSingleUnitController),
	)

	participation := beego.NewNamespace(
		"/participation",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", participationListController),
		beego.NSRouter("/list-for-owners-by-event", participationListForOwnersByEventController),
		beego.NSRouter("/single-unit", participationSingleUnitController),
	)

	attendance := beego.NewNamespace(
		"/attendance",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", attendanceListController),
		beego.NSRouter("/list-for-owners-by-event", attendanceListForOwnersByEventController),
	)

	pointTransaction := beego.NewNamespace(
		"/point-transaction",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", pointTransactionListController),
		beego.NSRouter("/single-unit", pointTransactionSingleUnitController),
	)

	pointBalance := beego.NewNamespace(
		"/point-balance",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/single-unit", pointBalanceSingleUnitController),
	)

	pointAcquisition := beego.NewNamespace(
		"/point-acquisition",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/after-point-acquisition", afterPointAcquisitionController),
	)

	pointConsumption := beego.NewNamespace(
		"/point-consumption",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/point-consumption", pointConsumptionController),
	)

	post := beego.NewNamespace(
		"/post",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list-by-friend", postListByFriendController),
		beego.NSRouter("/list-me", postListMeController),
		beego.NSRouter("/single-unit", postSingleUnitController),
	)

	userInfo := beego.NewNamespace(
		"/user-info",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/user-info-creates", userInfoCreatesController),
	)

	userProfile := beego.NewNamespace(
		"/user-profile",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/:businessPartner", userProfileController),
	)

	accountInfo := beego.NewNamespace(
		"/account-info",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/account-info", accountInfoController),
	)

	message := beego.NewNamespace(
		"/message",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", messageListController),
		beego.NSRouter("/interaction-with-friend", messageInteractionWithFriendController),
		//		beego.NSRouter("/single-unit", messageSingleUnitController),
	)

	friend := beego.NewNamespace(
		"/friend",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list", friendListController),
	)

	qRCode := beego.NewNamespace(
		"/message",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list-for-owners", qRCodeListForOwnersController),
	)

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
		beego.NSRouter("/detail/list-for-an-order-document", ordersDetailListForAnOrderDocumentController),
		beego.NSRouter("/partners-with-address", ordersPartnersWithAddressController),
		beego.NSRouter("/item-single-unit/:userType", ordersSingleUnit),
		beego.NSRouter("/item-single-unit-mill-sheet/:userType", ordersItemSingleUnitMillSheet),
		beego.NSRouter("/item-schedule-line/:userType", ordersItemScheduleLineController),
		beego.NSRouter("/item-pricing-element/:userType", ordersItemPricingElementController),
		beego.NSRouter("/item/:userType", ordersItemController),
	)

	inspectionLot := beego.NewNamespace(
		"/inspection-lot",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/header-single-unit", inspectionLotSingleUnitController),
		beego.NSRouter("/header-single-unit-mill-sheet", inspectionLotSingleUnitMillSheetController),
		beego.NSRouter("/header-single-unit-mill-box-interface", inspectionLotSingleUnitMillBoxInterfaceController),
		beego.NSRouter("/list", inspectionLotListController),
		beego.NSRouter("/spec-detail", inspectionLotSpecDetailController),
		beego.NSRouter("/component-composition", inspectionLotComponentCompositionController),
		beego.NSRouter("/inspection", inspectionLotInspectionController),
	)

	certificateAuthorityChainWithUsageControlChain := beego.NewNamespace(
		"/certificate-authority-chain-with-usage-control-chain",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/certificate-authority-chain-with-usage-control-chain", certificateAuthorityChainWithUsageControlChainController),
	)

	deliveryDocument := beego.NewNamespace(
		"/delivery-document",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", deliveryDocumentListController),
		beego.NSRouter("/detail/list/:userType", deliveryDocumentDetailListController),
		beego.NSRouter("/detail/list-for-a-delivery-instruction", deliveryDocumentDetailListForADeliveryInstructionController),
		beego.NSRouter("/item-single-unit/:userType", deliveryDocumentSingleUnitController),
		beego.NSRouter("/item/:userType", deliveryDocumentItemController),
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
		beego.NSRouter("/header-single-unit", billOfMaterialHeaderSingleUnitController),
		beego.NSRouter("/detail/list", billOfMaterialDetailListController),
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
		beego.NSRouter("/detail/list/:userType", productionOrderDetailListController),
		beego.NSRouter("/header-single-unit/:userType", productionOrderHeaderSingleUnitController),
		beego.NSRouter("/item-single-unit/:userType", productionOrderItemSingleUnitController),
		beego.NSRouter("/item-component/list", productionOrderItemComponentListController),
		beego.NSRouter("/item-operation/list", productionOrderItemOperationListController),
	)

	productionOrderConf := beego.NewNamespace(
		"/production-order-conf",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/header-single-unit/:userType", productionOrderConfHeaderSingleUnitController),
	)

	beego.AddNamespace(
		beego.NewNamespace("/api/reads").
			Namespace(
				loginAuthenticator,
				afterLoginUserInfo,
				listsForPointUsers,
				listsForOwners,
				event,
				article,
				site,
				shop,
				participation,
				attendance,
				pointTransaction,
				pointBalance,
				pointAcquisition,
				pointConsumption,
				post,
				instagram,
				userInfo,
				userProfile,
				accountInfo,
				message,
				friend,
				qRCode,
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
				inspectionLot,
				certificateAuthorityChainWithUsageControlChain,
			),
	)

	beego.Router("/", &HealthCheckController{})

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
