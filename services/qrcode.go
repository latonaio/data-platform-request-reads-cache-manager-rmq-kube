package services

import (
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
)

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
