package controller

import (
	"context"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/operationsdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/supplychainrelationshipexconflist"

	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/billofmaterialdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/billofmateriallist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/businesspartnerlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/deliverydocumentdetail"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/deliverydocumentdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/deliverydocumentdetailpagination"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/deliverydocumentlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/equipmentlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/invoicedetail"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/invoicedetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/invoicelist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/operationslist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/ordersdetail"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/ordersdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/ordersdetailpagination"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/orderslist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/pricemasterdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/pricemasterlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionorderdetail"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionorderdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionorderdetailpagination"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionorderlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionversiondetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productionversionlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/productlist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/supplychainrelationshipdetaillist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/supplychainrelationshiplist"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/update"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller/workcenterlist"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type Controller struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger

	OrdersDetail                     *ordersdetail.OrdersDetailCtrl
	OrdersList                       *orderslist.OrdersListCtrl
	OrdersDetailList                 *ordersdetaillist.OrdersDetailListCtrl
	OrdersDetailPagination           *ordersdetailpagination.OrdersDetailPaginationCtrl
	DeliveryDocumentList             *deliverydocumentlist.DeliveryDocumentListCtrl
	DeliveryDocumentDetailList       *deliverydocumentdetaillist.DeliveryDocumentDetailListCtrl
	DeliveryDocumentDetail           *deliverydocumentdetail.DeliveryDocumentDetailCtrl
	DeliveryDocumentDetailPagination *deliverydocumentdetailpagination.DeliveryDocumentDetailPaginationCtrl
	InvoiceList                      *invoicelist.InvoiceListCtrl
	InvoiceDetailList                *invoicedetaillist.InvoiceDetailListCtrl
	InvoiceDetail                    *invoicedetail.InvoiceDetailCtrl
	ProductionOrderList              *productionorderlist.ProductionOrderListCtrl
	ProductionOrderDetailList        *productionorderdetaillist.ProductionOrderDetailListCtrl
	ProductionOrderDetail            *productionorderdetail.ProductionOrderDetailCtrl
	ProductionOrderDetailPagination  *productionorderdetailpagination.ProductionOrderDetailPaginationCtrl
	ProductList                      *productlist.ProductListCtrl
	ProductDetailList                *productdetaillist.ProductDetailListCtrl
	BillOfMaterialList               *billofmateriallist.BillOfMaterialListCtrl
	BillOfMaterialDetailList         *billofmaterialdetaillist.BillOfMaterialDetailListCtrl
	// ProductionVersionList            *productionversionlist.ProductionVersionListCtrl
	ProductionVersionDetailList *productionversiondetaillist.ProductionVersionDetailListCtrl
	OperationsList              *operationslist.OperationsListCtrl
	OperationsDetailList        *operationsdetaillist.OperationsDetailListCtrl
	EquipmentList               *equipmentlist.EquipmentListCtrl
	// EquipmentDetailList              *equipmentdetaillist.EquipmentDetailListCtrl
	WorkCenterList *workcenterlist.WorkCenterListCtrl
	// WorkCenterDetailList             *workcenterdetaillist.WorkCenterDetailListCtrl
	ProductionVersionList             *productionversionlist.ProductionVersionListCtrl
	SupplyChainRelationshipList       *supplychainrelationshiplist.SupplyChainRelationshipListCtrl
	SupplyChainRelationshipDetailList *supplychainrelationshipdetaillist.SupplyChainRelationshipDetailListCtrl
	SupplyChainRelationshipExconfList *supplychainrelationshipexconflist.SupplyChainRelationshipExconfListCtrl
	BusinessPartnerList               *businesspartnerlist.BusinessPartnerListCtrl
	PriceMasterList                   *pricemasterlist.PriceMasterListCtrl
	PriceMasterDetailList             *pricemasterdetaillist.PriceMasterDetailListCtrl
	Update                            *update.Update
}

func NewController(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *Controller {
	return &Controller{
		cache:                            c,
		rmq:                              rmq,
		ctx:                              ctx,
		log:                              log,
		OrdersDetail:                     ordersdetail.NewOrdersDetailCtrl(ctx, c, rmq, log),
		OrdersList:                       orderslist.NewOrdersListCtrl(ctx, c, rmq, log),
		OrdersDetailList:                 ordersdetaillist.NewOrdersDetailListCtrl(ctx, c, rmq, log),
		OrdersDetailPagination:           ordersdetailpagination.NewOrdersDetailPaginationCtrl(ctx, c, rmq, log),
		DeliveryDocumentList:             deliverydocumentlist.NewDeliveryDocumentListCtrl(ctx, c, rmq, log),
		DeliveryDocumentDetailList:       deliverydocumentdetaillist.NewDeliveryDocumentDetailListCtrl(ctx, c, rmq, log),
		DeliveryDocumentDetail:           deliverydocumentdetail.NewDeliveryDocumentDetailCtrl(ctx, c, rmq, log),
		DeliveryDocumentDetailPagination: deliverydocumentdetailpagination.NewDeliveryDocumentDetailPaginationCtrl(ctx, c, rmq, log),
		InvoiceList:                      invoicelist.NewInvoiceListCtrl(ctx, c, rmq, log),
		InvoiceDetailList:                invoicedetaillist.NewInvoiceDetailListCtrl(ctx, c, rmq, log),
		InvoiceDetail:                    invoicedetail.NewInvoiceDetailCtrl(ctx, c, rmq, log),
		ProductionOrderList:              productionorderlist.NewProductionOrderListCtrl(ctx, c, rmq, log),
		ProductionOrderDetailList:        productionorderdetaillist.NewProductionOrderDetailListCtrl(ctx, c, rmq, log),
		ProductionOrderDetail:            productionorderdetail.NewProductionOrderDetailCtrl(ctx, c, rmq, log),
		ProductionOrderDetailPagination:  productionorderdetailpagination.NewProductionOrderDetailPaginationCtrl(ctx, c, rmq, log),
		ProductList:                      productlist.NewProductListCtrl(ctx, c, rmq, log),
		ProductDetailList:                productdetaillist.NewProductDetailListCtrl(ctx, c, rmq, log),
		BillOfMaterialList:               billofmateriallist.NewBillOfMaterialListCtrl(ctx, c, rmq, log),
		BillOfMaterialDetailList:         billofmaterialdetaillist.NewBillOfMaterialDetailListCtrl(ctx, c, rmq, log),
		ProductionVersionList:            productionversionlist.NewProductionVersionListCtrl(ctx, c, rmq, log),
		ProductionVersionDetailList:      productionversiondetaillist.NewProductionVersionDetailListCtrl(ctx, c, rmq, log),
		OperationsList:                   operationslist.NewOperationsListCtrl(ctx, c, rmq, log),
		OperationsDetailList:             operationsdetaillist.NewOperationsDetailListCtrl(ctx, c, rmq, log),
		EquipmentList:                    equipmentlist.NewEquipmentListCtrl(ctx, c, rmq, log),
		// EquipmenttDetailList:             equipmentdetaillist.NewEquipmentDetailListCtrl(ctx, c, rmq, log),
		WorkCenterList: workcenterlist.NewWorkCenterListCtrl(ctx, c, rmq, log),
		// WorkCentertDetailList:            workcenterdetaillist.NewWorkCenterDetailListCtrl(ctx, c, rmq, log),
		SupplyChainRelationshipList:       supplychainrelationshiplist.NewSupplyChainRelationshipListCtrl(ctx, c, rmq, log),
		SupplyChainRelationshipDetailList: supplychainrelationshipdetaillist.NewSupplyChainRelationshipDetailListCtrl(ctx, c, rmq, log),
		SupplyChainRelationshipExconfList: supplychainrelationshipexconflist.NewSupplyChainRelationshipExconfListCtrl(ctx, c, rmq, log),
		PriceMasterList:                   pricemasterlist.NewPriceMasterListCtrl(ctx, c, rmq, log),
		PriceMasterDetailList:             pricemasterdetaillist.NewPriceMasterDetailListCtrl(ctx, c, rmq, log),
		BusinessPartnerList:               businesspartnerlist.NewBusinessPartnerListCtrl(ctx, c, rmq, log),
		Update:                            update.NewUpdateCtrl(ctx, c, rmq, log),
	}
}
