package services

import (
	apiModuleRuntimesResponsesBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/batch-master-record"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
)

//func AddressMapper(
//	address *[]apiModuleRuntimesResponsesAddress.Address,
//) map[string]apiModuleRuntimesResponsesAddress.Address {
//	addressMapper := map[string]apiModuleRuntimesResponsesAddress.Address{}
//
//	for _, v := range *address {
//		addressMapper[v.Address] = v
//	}
//
//	return addressMapper
//}

func ProductDescriptionMapper(
	productDescription *[]apiModuleRuntimesResponsesProductMaster.ProductDescription,
) map[string]apiModuleRuntimesResponsesProductMaster.ProductDescription {
	descriptionMapper := map[string]apiModuleRuntimesResponsesProductMaster.ProductDescription{}

	for _, v := range *productDescription {
		descriptionMapper[v.Product] = v
	}

	return descriptionMapper
}

func ProductDescByBPMapper(
	productDescByBP *[]apiModuleRuntimesResponsesProductMaster.ProductDescByBP,
) map[string]apiModuleRuntimesResponsesProductMaster.ProductDescByBP {
	productDescByBPMapper := map[string]apiModuleRuntimesResponsesProductMaster.ProductDescByBP{}

	for _, v := range *productDescByBP {
		productDescByBPMapper[v.Product] = v
	}

	return productDescByBPMapper
}

func BusinessPartnerNameMapper(
	businessPartners *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) map[int]apiModuleRuntimesResponsesBusinessPartner.General {
	businessPartnerMapper := map[int]apiModuleRuntimesResponsesBusinessPartner.General{}

	for _, v := range *businessPartners.Message.General {
		//businessPartnerMapper[v.BusinessPartner] = v.BusinessPartnerName
		businessPartnerMapper[v.BusinessPartner] = apiModuleRuntimesResponsesBusinessPartner.General{
			BusinessPartner:     v.BusinessPartner,
			BusinessPartnerName: v.BusinessPartnerName,
		}
	}

	return businessPartnerMapper
}

func PlantMapper(
	plantGeneral *[]apiModuleRuntimesResponsesPlant.General,
) map[string]apiModuleRuntimesResponsesPlant.General {
	plantMapper := map[string]apiModuleRuntimesResponsesPlant.General{}

	for _, v := range *plantGeneral {
		plantMapper[v.Plant] = apiModuleRuntimesResponsesPlant.General{
			Plant:     v.Plant,
			PlantName: v.PlantName,
		}
	}

	return plantMapper
}

func GeneralsMapper(
	productMasterGeneral *[]apiModuleRuntimesResponsesProductMaster.General,
) map[string]apiModuleRuntimesResponsesProductMaster.General {
	productMasterGeneralMapper := map[string]apiModuleRuntimesResponsesProductMaster.General{}

	for _, v := range *productMasterGeneral {
		productMasterGeneralMapper[v.Product] = apiModuleRuntimesResponsesProductMaster.General{
			Product:                  v.Product,
			InternalCapacityQuantity: v.InternalCapacityQuantity,
			SizeOrDimensionText:      v.SizeOrDimensionText,
		}
	}

	return productMasterGeneralMapper
}

func BatchMapper(
	batch *[]apiModuleRuntimesResponsesBatchMasterRecord.Batch,
) map[string]apiModuleRuntimesResponsesBatchMasterRecord.Batch {
	batchMapper := map[string]apiModuleRuntimesResponsesBatchMasterRecord.Batch{}

	for _, v := range *batch {
		batchMapper[v.Batch] = apiModuleRuntimesResponsesBatchMasterRecord.Batch{
			Batch:             v.Batch,
			ValidityStartDate: v.ValidityStartDate,
			ValidityStartTime: v.ValidityStartTime,
			ValidityEndDate:   v.ValidityEndDate,
			ValidityEndTime:   v.ValidityEndTime,
		}
	}

	return batchMapper
}

func ReadProductImage(
	pdRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartner int,
	product string,
) *apiOutputFormatter.ProductImage {
	img := &apiOutputFormatter.ProductImage{}

	for _, pmdResHeaderV := range *pdRes.Message.GeneralDoc {
		if &pmdResHeaderV.DocIssuerBusinessPartner != nil &&
			pmdResHeaderV.DocIssuerBusinessPartner == businessPartner &&
			&product != nil &&
			pmdResHeaderV.Product == product {
			img = &apiOutputFormatter.ProductImage{
				BusinessPartnerID: pmdResHeaderV.DocIssuerBusinessPartner,
				DocID:             pmdResHeaderV.DocID,
				FileExtension:     pmdResHeaderV.FileExtension,
			}
		}
	}

	return img
}

func ReadDocumentImageOrders(
	itemDocRes *apiModuleRuntimesResponsesOrders.OrdersDocRes,
	ordersId int,
	ordersItem int,
) *apiOutputFormatter.DocumentImageOrders {
	for _, itemDoc := range *itemDocRes.Message.ItemDoc {
		if itemDoc.OrderID == ordersId && itemDoc.OrderItem == ordersItem {
			if itemDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageOrders{
					OrdersID:      itemDoc.OrderID,
					OrdersItem:    itemDoc.OrderItem,
					DocID:         itemDoc.DocID,
					FileExtension: itemDoc.FileExtension,
				}
			}
		}
	}

	return nil
}

func ReadDocumentImageDeliveryDocument(
	itemDocRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes,
	deliveryDocument int,
	deliveryDocumentItem int,
) *apiOutputFormatter.DocumentImageDeliveryDocument {
	for _, itemDoc := range *itemDocRes.Message.ItemDoc {
		if itemDoc.DeliveryDocument == deliveryDocument && itemDoc.DeliveryDocumentItem == deliveryDocumentItem {
			if itemDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageDeliveryDocument{
					DeliveryDocument:     itemDoc.DeliveryDocument,
					DeliveryDocumentItem: itemDoc.DeliveryDocumentItem,
					DocID:                itemDoc.DocID,
					FileExtension:        itemDoc.FileExtension,
				}
			}
		}
	}

	return nil
}
