package services

import (
	apiModuleRuntimesResponsesActPurpose "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/act-purpose"
	apiModuleRuntimesResponsesArticle "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article"
	apiModuleRuntimesResponsesArticleType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article-type"
	apiModuleRuntimesResponsesBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/batch-master-record"
	apiModuleRuntimesResponsesBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/country"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/distribution-profile"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesEventType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event-type"
	apiModuleRuntimesResponsesIncoterms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/incoterms"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesLanguage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/language"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesMessage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/message"
	apiModuleRuntimesResponsesMessageType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/message-type"
	apiModuleRuntimesResponsesObjectType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/object-type"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPaymentTerms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/payment-terms"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesPointConditionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-condition-type"
	apiModuleRuntimesResponsesPointConsumptionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-consumption-type"
	apiModuleRuntimesResponsesPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-transaction-type"
	apiModuleRuntimesResponsesPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/post"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/project"
	apiModuleRuntimesResponsesRank "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/rank"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiModuleRuntimesResponsesShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop-type"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiModuleRuntimesResponsesSiteType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"strconv"
)

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

func BusinessPartnerRoleTextMapper(
	businessPartnerRoleText *[]apiModuleRuntimesResponsesBusinessPartnerRole.Text,
) map[string]apiModuleRuntimesResponsesBusinessPartnerRole.Text {
	businessPartnerRoleTextMapper := map[string]apiModuleRuntimesResponsesBusinessPartnerRole.Text{}

	for _, v := range *businessPartnerRoleText {
		businessPartnerRoleTextMapper[v.BusinessPartnerRole] = apiModuleRuntimesResponsesBusinessPartnerRole.Text{
			BusinessPartnerRole:     v.BusinessPartnerRole,
			Language:                v.Language,
			BusinessPartnerRoleName: v.BusinessPartnerRoleName,
		}
	}

	return businessPartnerRoleTextMapper
}

func SiteMapper(
	sites *apiModuleRuntimesResponsesSite.SiteRes,
) map[int]apiModuleRuntimesResponsesSite.Header {
	siteMapper := map[int]apiModuleRuntimesResponsesSite.Header{}

	for _, v := range *sites.Message.Header {
		siteMapper[v.Site] = apiModuleRuntimesResponsesSite.Header{
			Site:        v.Site,
			Description: v.Description,
		}
	}

	return siteMapper
}

func ShopMapper(
	shops *apiModuleRuntimesResponsesShop.ShopRes,
) map[int]apiModuleRuntimesResponsesShop.Header {
	shopMapper := map[int]apiModuleRuntimesResponsesShop.Header{}

	for _, v := range *shops.Message.Header {
		shopMapper[v.Shop] = apiModuleRuntimesResponsesShop.Header{
			Shop:        v.Shop,
			Description: v.Description,
		}
	}

	return shopMapper
}

func EventHeadersMapper(
	eventHeaderRes *apiModuleRuntimesResponsesEvent.EventRes,
) map[string]apiModuleRuntimesResponsesEvent.Header {
	eventHeadersMapper := map[string]apiModuleRuntimesResponsesEvent.Header{}

	for _, v := range *eventHeaderRes.Message.Header {
		eventHeadersMapper[strconv.Itoa(v.Event)] = v
	}

	return eventHeadersMapper
}

func EventCountersMapper(
	eventCounterRes *apiModuleRuntimesResponsesEvent.EventRes,
) map[string]apiModuleRuntimesResponsesEvent.Counter {
	eventCountersMapper := map[string]apiModuleRuntimesResponsesEvent.Counter{}

	for _, v := range *eventCounterRes.Message.Counter {
		eventCountersMapper[strconv.Itoa(v.Event)] = v
	}

	return eventCountersMapper
}

func ArticleHeadersMapper(
	articleHeaderRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) map[string]apiModuleRuntimesResponsesArticle.Header {
	articleHeadersMapper := map[string]apiModuleRuntimesResponsesArticle.Header{}

	for _, v := range *articleHeaderRes.Message.Header {
		articleHeadersMapper[strconv.Itoa(v.Article)] = v
	}

	return articleHeadersMapper
}

func ArticleCountersMapper(
	articleCounterRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) map[string]apiModuleRuntimesResponsesArticle.Counter {
	articleCountersMapper := map[string]apiModuleRuntimesResponsesArticle.Counter{}

	for _, v := range *articleCounterRes.Message.Counter {
		articleCountersMapper[strconv.Itoa(v.Article)] = v
	}

	return articleCountersMapper
}

func SiteHeadersMapper(
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
) map[string]apiModuleRuntimesResponsesSite.Header {
	siteHeadersMapper := map[string]apiModuleRuntimesResponsesSite.Header{}

	for _, v := range *siteHeaderRes.Message.Header {
		siteHeadersMapper[strconv.Itoa(v.Site)] = v
	}

	return siteHeadersMapper
}

func SiteCountersMapper(
	siteCounterRes *apiModuleRuntimesResponsesSite.SiteRes,
) map[string]apiModuleRuntimesResponsesSite.Counter {
	siteCountersMapper := map[string]apiModuleRuntimesResponsesSite.Counter{}

	for _, v := range *siteCounterRes.Message.Counter {
		siteCountersMapper[strconv.Itoa(v.Site)] = v
	}

	return siteCountersMapper
}

func ShopHeadersMapper(
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
) map[string]apiModuleRuntimesResponsesShop.Header {
	shopHeadersMapper := map[string]apiModuleRuntimesResponsesShop.Header{}

	for _, v := range *shopHeaderRes.Message.Header {
		shopHeadersMapper[strconv.Itoa(v.Shop)] = v
	}

	return shopHeadersMapper
}

func PostHeadersMapper(
	postHeaderRes *apiModuleRuntimesResponsesPost.PostRes,
) map[string]apiModuleRuntimesResponsesPost.Header {
	postHeadersMapper := map[string]apiModuleRuntimesResponsesPost.Header{}

	for _, v := range *postHeaderRes.Message.Header {
		postHeadersMapper[strconv.Itoa(v.Post)] = v
	}

	return postHeadersMapper
}

func ActPurposeTextMapper(
	actPurposeText *[]apiModuleRuntimesResponsesActPurpose.Text,
) map[string]apiModuleRuntimesResponsesActPurpose.Text {
	actPurposeTextMapper := map[string]apiModuleRuntimesResponsesActPurpose.Text{}

	for _, v := range *actPurposeText {
		actPurposeTextMapper[v.ActPurpose] = apiModuleRuntimesResponsesActPurpose.Text{
			ActPurpose:     v.ActPurpose,
			Language:       v.Language,
			ActPurposeName: v.ActPurposeName,
		}
	}

	return actPurposeTextMapper
}

func EventTypeTextMapper(
	eventTypeText *[]apiModuleRuntimesResponsesEventType.Text,
) map[string]apiModuleRuntimesResponsesEventType.Text {
	eventTypeTextMapper := map[string]apiModuleRuntimesResponsesEventType.Text{}

	for _, v := range *eventTypeText {
		eventTypeTextMapper[v.EventType] = apiModuleRuntimesResponsesEventType.Text{
			EventType:     v.EventType,
			Language:      v.Language,
			EventTypeName: v.EventTypeName,
		}
	}

	return eventTypeTextMapper
}

func ArticleTypeTextMapper(
	articleTypeText *[]apiModuleRuntimesResponsesArticleType.Text,
) map[string]apiModuleRuntimesResponsesArticleType.Text {
	articleTypeTextMapper := map[string]apiModuleRuntimesResponsesArticleType.Text{}

	for _, v := range *articleTypeText {
		articleTypeTextMapper[v.ArticleType] = apiModuleRuntimesResponsesArticleType.Text{
			ArticleType:     v.ArticleType,
			Language:        v.Language,
			ArticleTypeName: v.ArticleTypeName,
		}
	}

	return articleTypeTextMapper
}

func SiteTypeTextMapper(
	siteTypeText *[]apiModuleRuntimesResponsesSiteType.Text,
) map[string]apiModuleRuntimesResponsesSiteType.Text {
	siteTypeTextMapper := map[string]apiModuleRuntimesResponsesSiteType.Text{}

	for _, v := range *siteTypeText {
		siteTypeTextMapper[v.SiteType] = apiModuleRuntimesResponsesSiteType.Text{
			SiteType:     v.SiteType,
			Language:     v.Language,
			SiteTypeName: v.SiteTypeName,
		}
	}

	return siteTypeTextMapper
}

func ShopTypeTextMapper(
	shopTypeText *[]apiModuleRuntimesResponsesShopType.Text,
) map[string]apiModuleRuntimesResponsesShopType.Text {
	shopTypeTextMapper := map[string]apiModuleRuntimesResponsesShopType.Text{}

	for _, v := range *shopTypeText {
		shopTypeTextMapper[v.ShopType] = apiModuleRuntimesResponsesShopType.Text{
			ShopType:     v.ShopType,
			Language:     v.Language,
			ShopTypeName: v.ShopTypeName,
		}
	}

	return shopTypeTextMapper
}

func DistributionProfileTextMapper(
	distributionProfileText *[]apiModuleRuntimesResponsesDistributionProfile.Text,
) map[string]apiModuleRuntimesResponsesDistributionProfile.Text {
	distributionProfileTextMapper := map[string]apiModuleRuntimesResponsesDistributionProfile.Text{}

	for _, v := range *distributionProfileText {
		distributionProfileTextMapper[v.DistributionProfile] = apiModuleRuntimesResponsesDistributionProfile.Text{
			DistributionProfile:     v.DistributionProfile,
			Language:                v.Language,
			DistributionProfileName: v.DistributionProfileName,
		}
	}

	return distributionProfileTextMapper
}

func PointConditionTypeTextMapper(
	pointConditionTypeText *[]apiModuleRuntimesResponsesPointConditionType.Text,
) map[string]apiModuleRuntimesResponsesPointConditionType.Text {
	pointConditionTypeTextMapper := map[string]apiModuleRuntimesResponsesPointConditionType.Text{}

	for _, v := range *pointConditionTypeText {
		pointConditionTypeTextMapper[v.PointConditionType] = apiModuleRuntimesResponsesPointConditionType.Text{
			PointConditionType:     v.PointConditionType,
			Language:               v.Language,
			PointConditionTypeName: v.PointConditionTypeName,
		}
	}

	return pointConditionTypeTextMapper
}

func PointConsumptionTypeTextMapper(
	pointConsumptionTypeText *[]apiModuleRuntimesResponsesPointConsumptionType.Text,
) map[string]apiModuleRuntimesResponsesPointConsumptionType.Text {
	pointConsumptionTypeTextMapper := map[string]apiModuleRuntimesResponsesPointConsumptionType.Text{}

	for _, v := range *pointConsumptionTypeText {
		pointConsumptionTypeTextMapper[v.PointConsumptionType] = apiModuleRuntimesResponsesPointConsumptionType.Text{
			PointConsumptionType:     v.PointConsumptionType,
			Language:                 v.Language,
			PointConsumptionTypeName: v.PointConsumptionTypeName,
		}
	}

	return pointConsumptionTypeTextMapper
}

func PointTransactionTypeTextMapper(
	pointTransactionTypeText *[]apiModuleRuntimesResponsesPointTransactionType.Text,
) map[string]apiModuleRuntimesResponsesPointTransactionType.Text {
	pointTransactionTypeTextMapper := map[string]apiModuleRuntimesResponsesPointTransactionType.Text{}

	for _, v := range *pointTransactionTypeText {
		pointTransactionTypeTextMapper[v.PointTransactionType] = apiModuleRuntimesResponsesPointTransactionType.Text{
			PointTransactionType:     v.PointTransactionType,
			Language:                 v.Language,
			PointTransactionTypeName: v.PointTransactionTypeName,
		}
	}

	return pointTransactionTypeTextMapper
}

func ObjectTypeTextMapper(
	objectTypeText *[]apiModuleRuntimesResponsesObjectType.Text,
) map[string]apiModuleRuntimesResponsesObjectType.Text {
	objectTypeTextMapper := map[string]apiModuleRuntimesResponsesObjectType.Text{}

	for _, v := range *objectTypeText {
		objectTypeTextMapper[v.ObjectType] = apiModuleRuntimesResponsesObjectType.Text{
			ObjectType:     v.ObjectType,
			Language:       v.Language,
			ObjectTypeName: v.ObjectTypeName,
		}
	}

	return objectTypeTextMapper
}

func MessageTypeTextMapper(
	messageTypeText *[]apiModuleRuntimesResponsesMessageType.Text,
) map[string]apiModuleRuntimesResponsesMessageType.Text {
	messageTypeTextMapper := map[string]apiModuleRuntimesResponsesMessageType.Text{}

	for _, v := range *messageTypeText {
		messageTypeTextMapper[v.MessageType] = apiModuleRuntimesResponsesMessageType.Text{
			MessageType:     v.MessageType,
			Language:        v.Language,
			MessageTypeName: v.MessageTypeName,
		}
	}

	return messageTypeTextMapper
}

func LocalSubRegionTextMapper(
	localSubRegionText *[]apiModuleRuntimesResponsesLocalSubRegion.Text,
) map[string]apiModuleRuntimesResponsesLocalSubRegion.Text {
	localSubRegionTextMapper := map[string]apiModuleRuntimesResponsesLocalSubRegion.Text{}

	for _, v := range *localSubRegionText {
		localSubRegionTextMapper[v.LocalSubRegion] = apiModuleRuntimesResponsesLocalSubRegion.Text{
			LocalSubRegion:     v.LocalSubRegion,
			LocalRegion:        v.LocalRegion,
			Country:            v.Country,
			Language:           v.Language,
			LocalSubRegionName: v.LocalSubRegionName,
		}
	}

	return localSubRegionTextMapper
}

func LocalRegionTextMapper(
	localRegionText *[]apiModuleRuntimesResponsesLocalRegion.Text,
) map[string]apiModuleRuntimesResponsesLocalRegion.Text {
	localRegionTextMapper := map[string]apiModuleRuntimesResponsesLocalRegion.Text{}

	for _, v := range *localRegionText {
		localRegionTextMapper[v.LocalRegion] = apiModuleRuntimesResponsesLocalRegion.Text{
			LocalRegion:     v.LocalRegion,
			Country:         v.Country,
			Language:        v.Language,
			LocalRegionName: v.LocalRegionName,
		}
	}

	return localRegionTextMapper
}

func CountryTextMapper(
	countryText *[]apiModuleRuntimesResponsesCountry.Text,
) map[string]apiModuleRuntimesResponsesCountry.Text {
	countryTextMapper := map[string]apiModuleRuntimesResponsesCountry.Text{}

	for _, v := range *countryText {
		countryTextMapper[v.Country] = apiModuleRuntimesResponsesCountry.Text{
			Country:     v.Country,
			Language:    v.Language,
			CountryName: v.CountryName,
		}
	}

	return countryTextMapper
}

func LanguageTextMapper(
	languageText *[]apiModuleRuntimesResponsesLanguage.Text,
) map[string]apiModuleRuntimesResponsesLanguage.Text {
	languageTextMapper := map[string]apiModuleRuntimesResponsesLanguage.Text{}

	for _, v := range *languageText {
		languageTextMapper[v.Language] = apiModuleRuntimesResponsesLanguage.Text{
			Language:               v.Language,
			CorrespondenceLanguage: v.CorrespondenceLanguage,
			LanguageName:           v.LanguageName,
		}
	}

	return languageTextMapper
}

func RankTextMapper(
	rankText *[]apiModuleRuntimesResponsesRank.Text,
) map[int]apiModuleRuntimesResponsesRank.Text {
	rankTextMapper := map[int]apiModuleRuntimesResponsesRank.Text{}

	for _, v := range *rankText {
		rankTextMapper[v.Rank] = apiModuleRuntimesResponsesRank.Text{
			RankType: v.RankType,
			Rank:     v.Rank,
			Language: v.Language,
			RankName: v.RankName,
		}
	}

	return rankTextMapper
}

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

func BusinessPartnerPersonMapper(
	businessPartnerPerson *[]apiModuleRuntimesResponsesBusinessPartner.Person,
) map[int]apiModuleRuntimesResponsesBusinessPartner.Person {
	businessPartnerPersonMapper := map[int]apiModuleRuntimesResponsesBusinessPartner.Person{}

	for _, v := range *businessPartnerPerson {
		businessPartnerPersonMapper[v.BusinessPartner] = v
	}

	return businessPartnerPersonMapper
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

func OrdersAddressesMapper(
	ordersAddressesRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) map[int]apiModuleRuntimesResponsesOrders.Address {
	ordersAddressMapper := map[int]apiModuleRuntimesResponsesOrders.Address{}

	for _, v := range *ordersAddressesRes.Message.Address {
		ordersAddressMapper[v.AddressID] = apiModuleRuntimesResponsesOrders.Address{
			OrderID:     v.OrderID,
			AddressID:   v.AddressID,
			PostalCode:  v.PostalCode,
			LocalRegion: v.LocalRegion,
			Country:     v.Country,
			District:    v.District,
			StreetName:  v.StreetName,
			CityName:    v.CityName,
			Building:    v.Building,
			Floor:       v.Floor,
			Room:        v.Room,
		}
	}

	return ordersAddressMapper
}

func ReadEventImage(
	eventRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	event int,
) *apiOutputFormatter.EventImage {
	img := &apiOutputFormatter.EventImage{}

	for _, eventResHeaderV := range *eventRes.Message.HeaderDoc {

		if &event != nil &&
			eventResHeaderV.Event == event {
			img = &apiOutputFormatter.EventImage{
				BusinessPartnerID: eventResHeaderV.DocIssuerBusinessPartner,
				DocID:             eventResHeaderV.DocID,
				FileExtension:     eventResHeaderV.FileExtension,
			}
		}

	}

	return img
}

func ReadArticleImage(
	articleRes *apiModuleRuntimesResponsesArticle.ArticleDocRes,
	article int,
) *apiOutputFormatter.ArticleImage {
	img := &apiOutputFormatter.ArticleImage{}

	for _, articleResHeaderV := range *articleRes.Message.HeaderDoc {

		if &article != nil &&
			articleResHeaderV.Article == article {
			img = &apiOutputFormatter.ArticleImage{
				BusinessPartnerID: articleResHeaderV.DocIssuerBusinessPartner,
				DocID:             articleResHeaderV.DocID,
				FileExtension:     articleResHeaderV.FileExtension,
			}
		}

	}

	return img
}

func ReadSiteImage(
	siteRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	site int,
) *apiOutputFormatter.SiteImage {
	img := &apiOutputFormatter.SiteImage{}

	for _, siteResHeaderV := range *siteRes.Message.HeaderDoc {

		if &site != nil &&
			siteResHeaderV.Site == site {
			img = &apiOutputFormatter.SiteImage{
				BusinessPartnerID: siteResHeaderV.DocIssuerBusinessPartner,
				DocID:             siteResHeaderV.DocID,
				FileExtension:     siteResHeaderV.FileExtension,
			}
		}

	}

	return img
}

func ReadShopImage(
	shopRes *apiModuleRuntimesResponsesShop.ShopDocRes,
	shop int,
) *apiOutputFormatter.ShopImage {
	img := &apiOutputFormatter.ShopImage{}

	for _, shopResHeaderV := range *shopRes.Message.HeaderDoc {

		if &shop != nil &&
			shopResHeaderV.Shop == shop {
			img = &apiOutputFormatter.ShopImage{
				BusinessPartnerID: shopResHeaderV.DocIssuerBusinessPartner,
				DocID:             shopResHeaderV.DocID,
				FileExtension:     shopResHeaderV.FileExtension,
			}
		}

	}

	return img
}

func ReadBusinessPartnerImage(
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartner int,
) *apiOutputFormatter.BusinessPartnerImage {
	img := &apiOutputFormatter.BusinessPartnerImage{}

	for _, businessPartnerResGeneralV := range *businessPartnerRes.Message.GeneralDoc {

		if &businessPartner != nil &&
			businessPartnerResGeneralV.BusinessPartner == businessPartner {
			img = &apiOutputFormatter.BusinessPartnerImage{
				BusinessPartnerID: businessPartnerResGeneralV.DocIssuerBusinessPartner,
				DocID:             businessPartnerResGeneralV.DocID,
				FileExtension:     businessPartnerResGeneralV.FileExtension,
			}
		}

	}

	return img
}

func ReadMessageImage(
	messageRes *apiModuleRuntimesResponsesMessage.MessageDocRes,
	message int,
) *apiOutputFormatter.MessageImage {
	img := &apiOutputFormatter.MessageImage{}

	for _, messageResHeaderV := range *messageRes.Message.HeaderDoc {

		if &message != nil &&
			messageResHeaderV.Message == message {
			img = &apiOutputFormatter.MessageImage{
				BusinessPartnerID: messageResHeaderV.DocIssuerBusinessPartner,
				DocID:             messageResHeaderV.DocID,
				FileExtension:     messageResHeaderV.FileExtension,
			}
		}

	}

	return img
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

func ReadDocumentImageEvent(
	eventHeaderDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	event int,
) *apiOutputFormatter.DocumentImageEvent {
	for _, headerDoc := range *eventHeaderDocRes.Message.HeaderDoc {
		if headerDoc.Event == event {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageEvent{
					Event:         headerDoc.Event,
					DocID:         headerDoc.DocID,
					FileExtension: headerDoc.FileExtension,
				}
			}
		}
	}

	return nil
}

func ReadDocumentImageArticle(
	articleHeaderDocRes *apiModuleRuntimesResponsesArticle.ArticleDocRes,
	article int,
) *apiOutputFormatter.DocumentImageArticle {
	for _, headerDoc := range *articleHeaderDocRes.Message.HeaderDoc {
		if headerDoc.Article == article {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageArticle{
					Article:       headerDoc.Article,
					DocID:         headerDoc.DocID,
					FileExtension: headerDoc.FileExtension,
				}
			}
		}
	}

	return nil
}

func ReadDocumentImageSite(
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	site int,
) *apiOutputFormatter.DocumentImageSite {
	for _, headerDoc := range *siteHeaderDocRes.Message.HeaderDoc {
		if headerDoc.Site == site {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageSite{
					Site:          headerDoc.Site,
					DocID:         headerDoc.DocID,
					FileExtension: headerDoc.FileExtension,
				}
			}
		}
	}

	return nil
}

func ReadDocumentImageShop(
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
	shop int,
) *apiOutputFormatter.DocumentImageShop {
	for _, headerDoc := range *shopHeaderDocRes.Message.HeaderDoc {
		if headerDoc.Shop == shop {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageShop{
					Shop:          headerDoc.Shop,
					DocID:         headerDoc.DocID,
					FileExtension: headerDoc.FileExtension,
				}
			}
		}
	}

	return nil
}

func ReadDocumentImageBusinessPartner(
	businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartner int,
) *apiOutputFormatter.DocumentImageBusinessPartner {
	for _, generalDoc := range *businessPartnerGeneralDocRes.Message.GeneralDoc {
		if generalDoc.BusinessPartner == businessPartner {
			if generalDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageBusinessPartner{
					BusinessPartner: generalDoc.BusinessPartner,
					DocID:           generalDoc.DocID,
					FileExtension:   generalDoc.FileExtension,
				}
			}
		}
	}

	return nil
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
					OrderID:       itemDoc.OrderID,
					OrderItem:     itemDoc.OrderItem,
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

func ReadDocumentImageBillOfMaterial(
	billOfMaterialHeaderDocRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialDocRes,
	billOfMaterial int,
) *apiOutputFormatter.DocumentImageBillOfMaterial {
	for _, headerDoc := range *billOfMaterialHeaderDocRes.Message.HeaderDoc {
		if headerDoc.BillOfMaterial == billOfMaterial {
			if headerDoc.DocType == "IMAGE" {
				return &apiOutputFormatter.DocumentImageBillOfMaterial{
					BillOfMaterial: headerDoc.BillOfMaterial,
					DocID:          headerDoc.DocID,
					FileExtension:  headerDoc.FileExtension,
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

func ReadPostImage(
	postDocRes *apiModuleRuntimesResponsesPost.PostDocRes,
	post int,
) *apiOutputFormatter.PostImage {
	for _, headerDoc := range *postDocRes.Message.HeaderDoc {
		if headerDoc.Post == post {
			return &apiOutputFormatter.PostImage{
				Post:          headerDoc.Post,
				DocID:         headerDoc.DocID,
				FileExtension: headerDoc.FileExtension,
			}
		}
	}

	return nil
}
