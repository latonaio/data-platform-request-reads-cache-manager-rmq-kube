package services

import (
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
