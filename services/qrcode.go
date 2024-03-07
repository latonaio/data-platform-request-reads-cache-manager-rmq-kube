package services

import (
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
)

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
