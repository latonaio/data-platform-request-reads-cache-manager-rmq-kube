package services

import (
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
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

func CreateProductImage(
	pdRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartner int,
	product string,
) *apiOutputFormatter.ProductImage {
	img := &apiOutputFormatter.ProductImage{}

	for _, pmdResHeaderV := range *pdRes.Message.HeaderDoc {
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
