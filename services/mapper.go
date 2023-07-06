package services

import (
	apiModuleRuntimesResponses "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
)

func DescriptionMapper(
	productDescByBP *[]apiModuleRuntimesResponses.ProductDescByBP,
) map[string]apiModuleRuntimesResponses.ProductDescByBP {
	descriptionMapper := map[string]apiModuleRuntimesResponses.ProductDescByBP{}

	for _, v := range *productDescByBP {
		descriptionMapper[v.Product] = v
	}

	return descriptionMapper
}

func PlantMapper(
	plantGeneral *[]apiModuleRuntimesResponses.PlantGeneral,
) map[string]apiModuleRuntimesResponses.PlantGeneral {
	plantMapper := map[string]apiModuleRuntimesResponses.PlantGeneral{}

	for _, v := range *plantGeneral {
		plantMapper[v.Plant] = v
	}

	return plantMapper
}

func CreateProductImage(
	pdRes *apiModuleRuntimesResponses.ProductMasterDocRes,
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
