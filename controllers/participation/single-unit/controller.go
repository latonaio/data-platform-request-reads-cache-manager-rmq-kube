package controllersParticipationSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/distribution-profile"
	apiModuleRuntimesRequestsEventType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event-type"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/participation/participation"
	apiModuleRuntimesRequestsParticipationDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/participation/participation-doc"
	apiModuleRuntimesRequestsPointConditionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-condition-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/distribution-profile"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesEventType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event-type"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/participation"
	apiModuleRuntimesResponsesPointConditionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-condition-type"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type ParticipationSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ParticipationSingleUnit struct {
	EventHeader                []apiOutputFormatter.EventHeader                `json:"EventHeader"`
	EventAddress               []apiOutputFormatter.EventAddress               `json:"EventAddress"`
	EventPointConditionElement []apiOutputFormatter.EventPointConditionElement `json:"EventPointConditionElement"`
	SiteHeader                 []apiOutputFormatter.SiteHeader                 `json:"SiteHeader"`
	SiteAddress                []apiOutputFormatter.SiteAddress                `json:"SiteAddress"`
	ShopHeader                 []apiOutputFormatter.ShopHeader                 `json:"ShopHeader"`
	ParticipationHeader        []apiOutputFormatter.ParticipationHeader        `json:"ParticipationHeader"`
}

func (controller *ParticipationSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "participation"
	redisKeyCategory2 := "single-unit"

	participationObject, _ := controller.GetInt("participationObject")
	participation, _ := controller.GetInt("participation")

	ParticipationSingleUnitEvent := apiInputReader.Event{}
	ParticipationSingleUnitParticipation := apiInputReader.Participation{}

	isReleased := true
	isCancelled := false
	isMarkedForDeletion := false

	docType := "QRCODE"

	ParticipationSingleUnitEvent = apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			Event:               participationObject,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventPointConditionElement: &apiInputReader.EventPointConditionElement{
			Event:               participationObject,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventAddress: &apiInputReader.EventAddress{
			Event: participationObject,
		},
		EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{
			Event: participationObject,
			//DocType:                &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	ParticipationSingleUnitParticipation = apiInputReader.Participation{
		ParticipationHeader: &apiInputReader.ParticipationHeader{
			Participation: participation,
//			IsCancelled:   &isCancelled,
		},
		ParticipationDocHeaderDoc: &apiInputReader.ParticipationDocHeaderDoc{
			Participation:            participation,
			DocType:                  &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(participationObject),
			strconv.Itoa(participation),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData ParticipationSingleUnit

		err := json.Unmarshal(cacheData, &responseData)

		if err != nil {
			services.HandleError(
				&controller.Controller,
				err,
				nil,
			)
		}

		services.Respond(
			&controller.Controller,
			&responseData,
		)
	}

	if cacheData != nil {
		go func() {
			controller.request(ParticipationSingleUnitEvent, ParticipationSingleUnitParticipation)
		}()
	} else {
		controller.request(ParticipationSingleUnitEvent, ParticipationSingleUnitParticipation)
	}
}

func (
	controller *ParticipationSingleUnitController,
) createEventRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createEventRequestPointConditionElement(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"PointConditionElements",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestPointConditionElement Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createEventDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventDocRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventDocRes{}
	responseBody := apiModuleRuntimesRequestsEventDoc.EventDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createEventRequestAddresses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"Addresses",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"イベントのアドレスが見つかりませんでした",
			&status,
		)
		return nil
	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createBusinessPartnerRequestGeneral(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*eventRes.Message.Header))

	for _, v := range *eventRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.EventOwner,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {

	businessPartnerRole := &(*eventRes.Message.Header)[0].EventOwnerBusinessPartnerRole

	var inputBusinessPartnerRole *string

	if businessPartnerRole != nil {
		inputBusinessPartnerRole = businessPartnerRole
	}

	input := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRole{
		BusinessPartnerRole: *inputBusinessPartnerRole,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRoleReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateBusinessPartnerRoleRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createSiteRequestHeader(
	requestPram *apiInputReader.Request,
	eventHeaderRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	header := eventHeaderRes.Message.Header

	var input = apiInputReader.Site{}

	input.SiteHeader = &apiInputReader.SiteHeader{
		Site:                (*header)[0].Site,
		IsReleased:          (*header)[0].IsReleased,
		IsMarkedForDeletion: (*header)[0].IsMarkedForDeletion,
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createSiteRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	eventHeaderRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
	var input = apiInputReader.Site{}

	for _, v := range *eventHeaderRes.Message.Header {
		input = apiInputReader.Site{
			SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{
				Site: v.Site,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteDocRes{}
	responseBody := apiModuleRuntimesRequestsSiteDoc.SiteDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createSiteDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createShopRequestHeader(
	requestPram *apiInputReader.Request,
	eventHeaderRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	header := eventHeaderRes.Message.Header

	var input = apiInputReader.Shop{}

	input.ShopHeader = &apiInputReader.ShopHeader{
		Shop:                *(*header)[0].Shop,
		IsReleased:          (*header)[0].IsReleased,
		IsMarkedForDeletion: (*header)[0].IsMarkedForDeletion,
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createShopRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	eventHeaderRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	var input = apiInputReader.Shop{}

	for _, v := range *eventHeaderRes.Message.Header {
		input = apiInputReader.Shop{
			ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
				Shop: *v.Shop,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopDocRes{}
	responseBody := apiModuleRuntimesRequestsShopDoc.ShopDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createShopDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreateEventTypeRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesEventType.EventTypeRes {

	eventType := &(*eventRes.Message.Header)[0].EventType

	var inputEventType *string

	if eventType != nil {
		inputEventType = eventType
	}

	input := apiModuleRuntimesRequestsEventType.EventType{
		EventType: *inputEventType,
	}

	responseJsonData := apiModuleRuntimesResponsesEventType.EventTypeRes{}
	responseBody := apiModuleRuntimesRequestsEventType.EventTypeReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateEventTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreateDistributionProfileRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes {

	distributionProfile := &(*eventRes.Message.Header)[0].DistributionProfile

	var inputDistributionProfile *string

	if distributionProfile != nil {
		inputDistributionProfile = distributionProfile
	}

	input := apiModuleRuntimesRequestsDistributionProfile.DistributionProfile{
		DistributionProfile: *inputDistributionProfile,
	}

	responseJsonData := apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes{}
	responseBody := apiModuleRuntimesRequestsDistributionProfile.DistributionProfileReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateDistributionProfileRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreatePointConditionTypeRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes {

	pointConditionType := &(*eventRes.Message.Header)[0].PointConditionType

	var inputPointConditionType *string

	if pointConditionType != nil {
		inputPointConditionType = pointConditionType
	}

	input := apiModuleRuntimesRequestsPointConditionType.PointConditionType{
		PointConditionType: *inputPointConditionType,
	}

	responseJsonData := apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes{}
	responseBody := apiModuleRuntimesRequestsPointConditionType.PointConditionTypeReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreatePointConditionTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*eventRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*eventRes.Message.Address)[0].LocalRegion
	country := &(*eventRes.Message.Address)[0].Country

	var inputLocalSubRegion *string
	var inputLocalRegion *string
	var inputCountry *string

	if localRegion != nil {
		inputLocalSubRegion = localSubRegion
		inputLocalRegion = localRegion
		inputCountry = country
	}

	input := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegion{
		LocalSubRegion: *inputLocalSubRegion,
		LocalRegion:    *inputLocalRegion,
		Country:        *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegionReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateLocalSubRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*eventRes.Message.Address)[0].LocalRegion
	country := &(*eventRes.Message.Address)[0].Country

	var inputLocalRegion *string
	var inputCountry *string

	if localRegion != nil {
		inputLocalRegion = localRegion
		inputCountry = country
	}

	input := apiModuleRuntimesRequestsLocalRegion.LocalRegion{
		LocalRegion: *inputLocalRegion,
		Country:     *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesLocalRegion.LocalRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalRegion.LocalRegionReadsText(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateLocalRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createSiteRequestAddresses(
	requestPram *apiInputReader.Request,
	headerRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input = apiInputReader.Site{}

	for _, v := range *headerRes.Message.Header {
		input = apiInputReader.Site{
			SiteAddress: &apiInputReader.SiteAddress{
				Site: v.Site,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
		requestPram,
		input,
		&controller.Controller,
		"Addresses",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"サイトのアドレスが見つかりませんでした",
			&status,
		)
		return nil
	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createSiteRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createParticipationRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
) *apiModuleRuntimesResponsesParticipation.ParticipationRes {
	responseJsonData := apiModuleRuntimesResponsesParticipation.ParticipationRes{}
	responseBody := apiModuleRuntimesRequestsParticipation.ParticipationReadsHeader(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createParticipationRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) createParticipationDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
) *apiModuleRuntimesResponsesParticipation.ParticipationDocRes {
	responseJsonData := apiModuleRuntimesResponsesParticipation.ParticipationDocRes{}
	responseBody := apiModuleRuntimesRequestsParticipationDoc.ParticipationDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createParticipationDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationSingleUnitController,
) request(
	inputEvent apiInputReader.Event,
	inputParticipation apiInputReader.Participation,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(15)

	var pointConditionElementRes apiModuleRuntimesResponsesEvent.EventRes
	var addressRes apiModuleRuntimesResponsesEvent.EventRes

	var businessPartnerGeneralRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	var siteHeaderRes apiModuleRuntimesResponsesSite.SiteRes

	var shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes

	var eventTypeTextRes *apiModuleRuntimesResponsesEventType.EventTypeRes

	var distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes

	var pointConditionTypeTextRes *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes

	var siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes

	var shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	var headerDocRes *apiModuleRuntimesResponsesEvent.EventDocRes

	var siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes

	var localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var participationHeaderRes *apiModuleRuntimesResponsesParticipation.ParticipationRes
	var participationHeaderDocRes *apiModuleRuntimesResponsesParticipation.ParticipationDocRes

	headerRes := *controller.createEventRequestHeader(
		controller.UserInfo,
		inputEvent,
	)

	go func() {
		defer wg.Done()
		pointConditionElementRes = *controller.createEventRequestPointConditionElement(
			controller.UserInfo,
			inputEvent,
		)
		controller.CustomLogger.Debug("complete pointConditionElementRes")
	}()

	go func() {
		defer wg.Done()
		addressRes = *controller.createEventRequestAddresses(
			controller.UserInfo,
			inputEvent,
		)

		localSubRegionTextRes = controller.CreateLocalSubRegionRequestText(
			controller.UserInfo,
			&addressRes,
		)

		localRegionTextRes = controller.CreateLocalRegionRequestText(
			controller.UserInfo,
			&addressRes,
		)
	}()

	go func() {
		defer wg.Done()
		headerDocRes = controller.createEventDocRequest(
			controller.UserInfo,
			inputEvent,
		)
		controller.CustomLogger.Debug("complete headerDocRes")
	}()

	go func() {
		defer wg.Done()
		businessPartnerGeneralRes = *controller.createBusinessPartnerRequestGeneral(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerRoleTextRes = controller.CreateBusinessPartnerRoleRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteHeaderRes = *controller.createSiteRequestHeader(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		shopHeaderRes = *controller.createShopRequestHeader(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		eventTypeTextRes = controller.CreateEventTypeRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		distributionProfileTextRes = controller.CreateDistributionProfileRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		pointConditionTypeTextRes = controller.CreatePointConditionTypeRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteHeaderDocRes = controller.createSiteDocRequest(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		shopHeaderDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteAddressRes = controller.createSiteRequestAddresses(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		participationHeaderRes = controller.createParticipationRequestHeader(
			controller.UserInfo,
			inputParticipation,
		)
	}()

	go func() {
		defer wg.Done()
		participationHeaderDocRes = controller.createParticipationDocRequest(
			controller.UserInfo,
			inputParticipation,
		)
	}()

	wg.Wait()

	controller.fin(
		&headerRes,
		&pointConditionElementRes,
		&addressRes,
		&businessPartnerGeneralRes,
		businessPartnerRoleTextRes,
		&siteHeaderRes,
		&shopHeaderRes,
		eventTypeTextRes,
		distributionProfileTextRes,
		pointConditionTypeTextRes,
		localSubRegionTextRes,
		localRegionTextRes,
		headerDocRes,
		siteHeaderDocRes,
		shopHeaderDocRes,
		siteAddressRes,
		participationHeaderRes,
		participationHeaderDocRes,
	)
}

func (
	controller *ParticipationSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesEvent.EventRes,
	pointConditionElementRes *apiModuleRuntimesResponsesEvent.EventRes,
	addressRes *apiModuleRuntimesResponsesEvent.EventRes,
	businessPartnerGeneralRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
	eventTypeTextRes *apiModuleRuntimesResponsesEventType.EventTypeRes,
	distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes,
	pointConditionTypeTextRes *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	headerDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
	participationHeaderRes *apiModuleRuntimesResponsesParticipation.ParticipationRes,
	participationHeaderDocRes *apiModuleRuntimesResponsesParticipation.ParticipationDocRes,
) {
	businessPartnerNameMapper := services.BusinessPartnerNameMapper(
		businessPartnerGeneralRes,
	)

	businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
		businessPartnerRoleTextRes.Message.Text,
	)

	siteMapper := services.SiteMapper(
		siteHeaderRes,
	)

	shopMapper := services.ShopMapper(
		shopHeaderRes,
	)

	eventTypeTextMapper := services.EventTypeTextMapper(
		eventTypeTextRes.Message.Text,
	)

	distributionProfileTextMapper := services.DistributionProfileTextMapper(
		distributionProfileTextRes.Message.Text,
	)

	pointConditionTypeTextMapper := services.PointConditionTypeTextMapper(
		pointConditionTypeTextRes.Message.Text,
	)

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	data := ParticipationSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadEventImage(
			headerDocRes,
			v.Event,
		)

		qrcode := services.CreateQRCodeEventDocImage(
			headerDocRes,
			v.Event,
		)

		documentImage := services.ReadDocumentImageEvent(
			headerDocRes,
			v.Event,
		)

		shopDescription := shopMapper[*v.Shop].Description

		data.EventHeader = append(data.EventHeader,
			apiOutputFormatter.EventHeader{
				Event:                             v.Event,
				EventType:                         v.EventType,
				EventTypeName:                     eventTypeTextMapper[v.EventType].EventTypeName,
				EventOwner:                        v.EventOwner,
				EventOwnerName:                    businessPartnerNameMapper[v.EventOwner].BusinessPartnerName,
				EventOwnerBusinessPartnerRoleName: businessPartnerRoleTextMapper[v.EventOwnerBusinessPartnerRole].BusinessPartnerRoleName,
				PersonResponsible:                 v.PersonResponsible,
				ValidityStartDate:                 v.ValidityStartDate,
				ValidityStartTime:                 v.ValidityStartTime,
				ValidityEndDate:                   v.ValidityEndDate,
				ValidityEndTime:                   v.ValidityEndTime,
				OperationStartDate:                v.OperationStartDate,
				OperationStartTime:                v.OperationStartTime,
				OperationEndDate:                  v.OperationEndDate,
				OperationEndTime:                  v.OperationEndTime,
				Description:                       v.Description,
				LongText:                          v.LongText,
				Introduction:                      v.Introduction,
				Site:                              v.Site,
				SiteDescription:                   siteMapper[v.Site].Description,
				Capacity:                          v.Capacity,
				Shop:                              v.Shop,
				ShopDescription:                   &shopDescription,
				Tag1:                              v.Tag1,
				Tag2:                              v.Tag2,
				Tag3:                              v.Tag3,
				Tag4:                              v.Tag4,
				DistributionProfile:               v.DistributionProfile,
				DistributionProfileName:           distributionProfileTextMapper[v.DistributionProfile].DistributionProfileName,
				PointConditionType:                v.PointConditionType,
				PointConditionTypeName:            pointConditionTypeTextMapper[v.PointConditionType].PointConditionTypeName,
				QuestionnaireType:                 v.QuestionnaireType,
				//				QuestionnaireTypeName:             questionnaireTypeTextMapper[v.QuestionnaireType].QuestionnaireTypeName,
				QuestionnaireTemplate: v.QuestionnaireTemplate,
				//				QuestionnaireTemplateName:         questionnaireTemplateTextMapper[v.QuestionnaireTemplate].QuestionnaireTemplateName,
				CreateUser: v.CreateUser,
				//				CreateUserFullName:         	   v.CreateUserFullName,
				//				CreateUserNickName:         	   v.CreateUserNickName,
				LastChangeUser: v.LastChangeUser,
				//				LastChangeUserFullName:            v.LastChangeUserFullName,
				//				LastChangeUserNickName:            v.LastChangeUserNickName,
				Images: apiOutputFormatter.Images{
					Event:              img,
					QRCode:             qrcode,
					DocumentImageEvent: documentImage,
				},
			},
		)
	}

	for _, v := range *pointConditionElementRes.Message.PointConditionElement {
		data.EventPointConditionElement = append(data.EventPointConditionElement,
			apiOutputFormatter.EventPointConditionElement{
				Event:                          v.Event,
				PointConditionRecord:           v.PointConditionRecord,
				PointConditionSequentialNumber: v.PointConditionSequentialNumber,
				PointSymbol:                    v.PointSymbol,
				Sender:                         v.Sender,
				PointTransactionType:           v.PointTransactionType,
				PointConditionType:             v.PointConditionType,
				PointConditionRateValue:        v.PointConditionRateValue,
				PointConditionRatio:            v.PointConditionRatio,
				PlusMinus:                      v.PlusMinus,
			},
		)
	}

	for _, v := range *addressRes.Message.Address {
		data.EventAddress = append(data.EventAddress,
			apiOutputFormatter.EventAddress{
				Event:              v.Event,
				AddressID:          v.AddressID,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
				PostalCode:         &v.PostalCode,
				StreetName:         v.StreetName,
				CityName:           v.CityName,
				Building:           v.Building,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {

		img := services.ReadSiteImage(
			siteHeaderDocRes,
			v.Site,
		)

		qrcode := services.CreateQRCodeSiteDocImage(
			siteHeaderDocRes,
			v.Site,
		)

		documentImage := services.ReadDocumentImageSite(
			siteHeaderDocRes,
			v.Site,
		)

		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Site:              img,
					QRCode:            qrcode,
					DocumentImageSite: documentImage,
				},
			},
		)
	}

	for _, v := range *siteAddressRes.Message.Address {
		data.SiteAddress = append(data.SiteAddress,
			apiOutputFormatter.SiteAddress{
				Site:               v.Site,
				AddressID:          v.AddressID,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
				PostalCode:         v.PostalCode,
				StreetName:         v.StreetName,
				CityName:           v.CityName,
				Building:           v.Building,
			},
		)
	}

	for _, v := range *shopHeaderRes.Message.Header {

		img := services.ReadShopImage(
			shopHeaderDocRes,
			v.Shop,
		)

		qrcode := services.CreateQRCodeShopDocImage(
			shopHeaderDocRes,
			v.Shop,
		)

		documentImage := services.ReadDocumentImageShop(
			shopHeaderDocRes,
			v.Shop,
		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Shop:              img,
					QRCode:            qrcode,
					DocumentImageShop: documentImage,
				},
			},
		)
	}

	for _, v := range *participationHeaderRes.Message.Header {

		qrcode := services.CreateQRCodeParticipationDocImage(
			participationHeaderDocRes,
			v.Participation,
		)

		data.ParticipationHeader = append(data.ParticipationHeader,
			apiOutputFormatter.ParticipationHeader{
				Participation:           v.Participation,
				ParticipationDate:       v.ParticipationDate,
				ParticipationTime:       v.ParticipationTime,
				Participator:            v.Participator,
				ParticipationObjectType: v.ParticipationObjectType,
				ParticipationObject:     v.ParticipationObject,
				IsCancelled:             v.IsCancelled,
				Images: apiOutputFormatter.Images{
					QRCode: qrcode,
				},
			},
		)
	}

	err := controller.RedisCache.SetCache(
		controller.RedisKey,
		data,
	)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
