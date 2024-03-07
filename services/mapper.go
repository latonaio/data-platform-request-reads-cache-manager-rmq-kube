package services

import (
	apiModuleRuntimesResponsesBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/batch-master-record"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesIncoterms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/incoterms"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPaymentTerms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/payment-terms"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/project"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"strconv"
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
		//plantMapper[v.Plant] = apiModuleRuntimesResponsesPlant.General{
		plantMapper[strconv.Itoa(v.BusinessPartner)] = apiModuleRuntimesResponsesPlant.General{
			Plant:     v.Plant,
			PlantName: v.PlantName,
		}
	}

	return plantMapper
}

func ProjectMapper(
	projectProject *[]apiModuleRuntimesResponsesProject.Project,
) map[int]apiModuleRuntimesResponsesProject.Project {
	projectMapper := map[int]apiModuleRuntimesResponsesProject.Project{}

	for _, v := range *projectProject {
		projectMapper[v.Project] = apiModuleRuntimesResponsesProject.Project{
			Project:            v.Project,
			ProjectDescription: v.ProjectDescription,
		}
	}

	return projectMapper
}

func WBSElementMapper(
	projectWBSElement *[]apiModuleRuntimesResponsesProject.WBSElement,
) map[int]apiModuleRuntimesResponsesProject.WBSElement {
	wBSElementMapper := map[int]apiModuleRuntimesResponsesProject.WBSElement{}

	for _, v := range *projectWBSElement {
		wBSElementMapper[v.Project] = apiModuleRuntimesResponsesProject.WBSElement{
			Project:               v.Project,
			WBSElement:            v.WBSElement,
			ResponsiblePersonName: v.ResponsiblePersonName,
			WBSElementDescription: v.WBSElementDescription,
		}
	}

	return wBSElementMapper
}

func IncotermsTextMapper(
	incotermsText *[]apiModuleRuntimesResponsesIncoterms.Text,
) map[string]apiModuleRuntimesResponsesIncoterms.Text {
	incotermsTextMapper := map[string]apiModuleRuntimesResponsesIncoterms.Text{}

	for _, v := range *incotermsText {
		incotermsTextMapper[v.Incoterms] = apiModuleRuntimesResponsesIncoterms.Text{
			Incoterms:     v.Incoterms,
			Language:      v.Language,
			IncotermsName: v.IncotermsName,
		}
	}

	return incotermsTextMapper
}

func PaymentTermsTextMapper(
	paymentTermsText *[]apiModuleRuntimesResponsesPaymentTerms.Text,
) map[string]apiModuleRuntimesResponsesPaymentTerms.Text {
	paymentTermsTextMapper := map[string]apiModuleRuntimesResponsesPaymentTerms.Text{}

	for _, v := range *paymentTermsText {
		paymentTermsTextMapper[v.PaymentTerms] = apiModuleRuntimesResponsesPaymentTerms.Text{
			PaymentTerms:     v.PaymentTerms,
			Language:         v.Language,
			PaymentTermsName: v.PaymentTermsName,
		}
	}

	return paymentTermsTextMapper
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

func InspectionLotListMapper(
	header *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	partner *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) map[int]apiModuleRuntimesResponsesInspectionLot.Header {
	inspectionListMapper := map[int]apiModuleRuntimesResponsesInspectionLot.Header{}

	for _, h := range *partner.Message.Partner {
		for _, p := range *header.Message.Header {
			if h.InspectionLot == p.InspectionLot {
				inspectionListMapper[h.InspectionLot] = apiModuleRuntimesResponsesInspectionLot.Header{
					InspectionLot:     p.InspectionLot,
					Product:           p.Product,
					InspectionLotDate: p.InspectionLotDate,
				}
			}
		}
	}

	return inspectionListMapper
}

func ReadProductImage(
	pdRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartner int,
	product string,
) *apiOutputFormatter.ProductImage {
	img := &apiOutputFormatter.ProductImage{}

	for _, pmdResHeaderV := range *pdRes.Message.GeneralDoc {
		//if &pmdResHeaderV.DocIssuerBusinessPartner != nil &&
		//	pmdResHeaderV.DocIssuerBusinessPartner == businessPartner &&
		//	&product != nil &&
		//	pmdResHeaderV.Product == product {
		//	img = &apiOutputFormatter.ProductImage{
		//		//BusinessPartnerID: pmdResHeaderV.DocIssuerBusinessPartner,
		//		DocID:         pmdResHeaderV.DocID,
		//		FileExtension: pmdResHeaderV.FileExtension,
		//	}
		//}

		if &product != nil &&
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

func ReadDocumentImageInspectionLot(
	inspectionLotHeaderDocRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotDocRes,
	inspectionLot int,
) *apiOutputFormatter.DocumentImageInspectionLot {
	for _, headerDoc := range *inspectionLotHeaderDocRes.Message.HeaderDoc {
		if headerDoc.InspectionLot == inspectionLot {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageInspectionLot{
					InspectionLot: headerDoc.InspectionLot,
					DocID:         headerDoc.DocID,
					FileExtension: headerDoc.FileExtension,
				}
			}
		}
	}

	return nil
}
