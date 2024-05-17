package services

import (
	apiModuleRuntimesResponsesBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
)

func CreateQRCodeBusinessPartnerDocImage(
	businessPartnerDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartner int,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*businessPartnerDocRes.Message.GeneralDoc)[0].DocID,
		FileExtension: (*businessPartnerDocRes.Message.GeneralDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeEventDocImage(
	eventDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	event int,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*eventDocRes.Message.HeaderDoc)[0].DocID,
		FileExtension: (*eventDocRes.Message.HeaderDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeSiteDocImage(
	siteDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	site int,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*siteDocRes.Message.HeaderDoc)[0].DocID,
		FileExtension: (*siteDocRes.Message.HeaderDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeProductDocImage(
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*productDocRes.Message.GeneralDoc)[0].DocID,
		FileExtension: (*productDocRes.Message.GeneralDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeProductStockDocImage(
	productStockDocRes *apiModuleRuntimesResponsesProductStock.ProductStockDocRes,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*productStockDocRes.Message.ProductStockDoc)[0].DocID,
		FileExtension: (*productStockDocRes.Message.ProductStockDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeProductionOrderItemDocImage(
	itemDocRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes,
) *apiOutputFormatter.QRCodeImage {
	img := &apiOutputFormatter.QRCodeImage{}

	img = &apiOutputFormatter.QRCodeImage{
		DocID:         (*itemDocRes.Message.ItemDoc)[0].DocID,
		FileExtension: (*itemDocRes.Message.ItemDoc)[0].FileExtension,
	}

	return img
}

func CreateQRCodeOrdersItemDocImage(
	itemDocRes *apiModuleRuntimesResponsesOrders.OrdersDocRes,
	orderId int,
	orderItem int,
) *apiOutputFormatter.QRCodeImage {
	for i, itemDoc := range *itemDocRes.Message.ItemDoc {
		if itemDoc.OrderID == orderId && itemDoc.OrderItem == orderItem {
			if itemDoc.DocType == "QRCODE" {
				return &apiOutputFormatter.QRCodeImage{
					DocID:         (*itemDocRes.Message.ItemDoc)[i].DocID,
					FileExtension: (*itemDocRes.Message.ItemDoc)[i].FileExtension,
				}
			}
		}
	}

	return nil
}

func CreateQRCodeDeliveryDocumentItemDocImage(
	itemDocRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes,
	deliveryDocument int,
	deliveryDocumentItem int,
) *apiOutputFormatter.QRCodeImage {
	for i, itemDoc := range *itemDocRes.Message.ItemDoc {
		if itemDoc.DeliveryDocument == deliveryDocument && itemDoc.DeliveryDocumentItem == deliveryDocumentItem {
			if itemDoc.DocType == "QRCODE" {
				return &apiOutputFormatter.QRCodeImage{
					DocID:         (*itemDocRes.Message.ItemDoc)[i].DocID,
					FileExtension: (*itemDocRes.Message.ItemDoc)[i].FileExtension,
				}
			}
		}
	}

	return nil
}

func CreateQRCodeBillOfMaterialHeaderDocImage(
	billOfMaterialHeaderDocRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialDocRes,
	billOfMaterial int,
) *apiOutputFormatter.QRCodeImage {
	for i, headerDoc := range *billOfMaterialHeaderDocRes.Message.HeaderDoc {
		if headerDoc.BillOfMaterial == billOfMaterial {
			if headerDoc.DocType == "QRCODE" {
				return &apiOutputFormatter.QRCodeImage{
					DocID:         (*billOfMaterialHeaderDocRes.Message.HeaderDoc)[i].DocID,
					FileExtension: (*billOfMaterialHeaderDocRes.Message.HeaderDoc)[i].FileExtension,
				}
			}
		}
	}

	return nil
}

func CreateQRCodeInspectionLotHeaderDocImage(
	inspectionLotHeaderDocRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotDocRes,
	inspectionLot int,
) *apiOutputFormatter.QRCodeImage {
	for i, headerDoc := range *inspectionLotHeaderDocRes.Message.HeaderDoc {
		if headerDoc.InspectionLot == inspectionLot {
			if headerDoc.DocType == "QRCODE" {
				return &apiOutputFormatter.QRCodeImage{
					DocID:         (*inspectionLotHeaderDocRes.Message.HeaderDoc)[i].DocID,
					FileExtension: (*inspectionLotHeaderDocRes.Message.HeaderDoc)[i].FileExtension,
				}
			}
		}
	}

	return nil
}

func CreateQRCodeProductionOrderHeaderDocImage(
	productionOrderHeaderDocRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes,
	productionOrder int,
) *apiOutputFormatter.QRCodeImage {
	for i, headerDoc := range *productionOrderHeaderDocRes.Message.HeaderDoc {
		if headerDoc.ProductionOrder == productionOrder {
			if headerDoc.DocType == "QRCODE" {
				return &apiOutputFormatter.QRCodeImage{
					DocID:         (*productionOrderHeaderDocRes.Message.HeaderDoc)[i].DocID,
					FileExtension: (*productionOrderHeaderDocRes.Message.HeaderDoc)[i].FileExtension,
				}
			}
		}
	}

	return nil
}
