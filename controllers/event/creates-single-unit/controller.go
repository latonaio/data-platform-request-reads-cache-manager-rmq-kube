package controllersEventCreatesSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/distribution-profile"
	apiModuleRuntimesRequestsEventType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event-type"
	apiModuleRuntimesRequestsPointConditionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-condition-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/distribution-profile"
	apiModuleRuntimesResponsesEventType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event-type"
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

type EventCreatesSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type EventCreatesSingleUnit struct {
	BusinessPartnerGeneral            []apiOutputFormatter.BusinessPartnerGeneral            `json:"BusinessPartnerGeneral"`
	BusinessPartnerPersonOrganization []apiOutputFormatter.BusinessPartnerPersonOrganization `json:"BusinessPartnerPersonOrganization"`
	BusinessPartnerRoleText           []apiOutputFormatter.BusinessPartnerRoleText           `json:"BusinessPartnerRoleText"`
	EventTypeText                     []apiOutputFormatter.EventTypeText                     `json:"EventTypeText"`
	DistributionProfileText           []apiOutputFormatter.DistributionProfileText           `json:"DistributionProfileText"`
	PointConditionTypeText            []apiOutputFormatter.PointConditionTypeText            `json:"PointConditionTypeText"`
	SiteAddress                       []apiOutputFormatter.SiteAddress                       `json:"SiteAddress"`
	SiteHeader                        []apiOutputFormatter.SiteHeader                        `json:"SiteHeader"`
	SiteAddressWithHeader             []apiOutputFormatter.SiteAddressWithHeader             `json:"SiteAddressWithHeader"`
	ShopHeader                        []apiOutputFormatter.ShopHeader                        `json:"ShopHeader"`
}

func (controller *EventCreatesSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)

	businessPartner, _ := controller.GetInt("businessPartner")

	businessPartnerRole := controller.GetString("businessPartnerRole") // UIでログイン時にロール選択で決まる、該当箇所開発時までは当面固定値にするか

	localSubRegion := controller.GetString("localSubRegion")

	redisKeyCategory1 := "event"
	redisKeyCategory2 := "creates-single-unit"
	redisKeyCategory3 := localSubRegion
	redisKeyCategory4 := businessPartner
	redisKeyCategory5 := businessPartnerRole
	//redisKeyCategory1 := "businessPartner"
	//redisKeyCategory2 := "businessPartnerRole"
	//redisKeyCategory3 := "event-creates-single-unit"

	EventCreatesSingleUnitBP := apiInputReader.BusinessPartner{}
	EventCreatesSingleUnitEventType := apiInputReader.EventTypeGlobal{}
	EventCreatesSingleUnitDistributionProfile := apiInputReader.DistributionProfileGlobal{}
	EventCreatesSingleUnitPointConditionType := apiInputReader.PointConditionTypeGlobal{}
	EventCreatesSingleUnitSiteAddress := apiInputReader.Site{}

	isReleased := true
	isMarkedForDeletion := false

	EventCreatesSingleUnitBP = apiInputReader.BusinessPartner{
		BusinessPartnerGeneral: &apiInputReader.BusinessPartnerGeneral{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerPersonOrganization: &apiInputReader.BusinessPartnerPersonOrganization{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	EventCreatesSingleUnitEventType = apiInputReader.EventTypeGlobal{
		EventTypeText: &apiInputReader.EventTypeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	EventCreatesSingleUnitDistributionProfile = apiInputReader.DistributionProfileGlobal{
		DistributionProfileText: &apiInputReader.DistributionProfileText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	EventCreatesSingleUnitPointConditionType = apiInputReader.PointConditionTypeGlobal{
		PointConditionTypeText: &apiInputReader.PointConditionTypeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	EventCreatesSingleUnitSiteAddress = apiInputReader.Site{
		SiteAddress: &apiInputReader.SiteAddress{
			LocalSubRegion: &localSubRegion,
		},
		SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
			strconv.Itoa(redisKeyCategory4),
			redisKeyCategory5,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData EventCreatesSingleUnit

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
			controller.request(
				EventCreatesSingleUnitBP,
				EventCreatesSingleUnitEventType,
				EventCreatesSingleUnitDistributionProfile,
				EventCreatesSingleUnitPointConditionType,
				EventCreatesSingleUnitSiteAddress,
				businessPartnerRole,
				isReleased,
				isMarkedForDeletion,
			)
		}()
	} else {
		controller.request(
			EventCreatesSingleUnitBP,
			EventCreatesSingleUnitEventType,
			EventCreatesSingleUnitDistributionProfile,
			EventCreatesSingleUnitPointConditionType,
			EventCreatesSingleUnitSiteAddress,
			businessPartnerRole,
			isReleased,
			isMarkedForDeletion,
		)
	}
}

func (
	controller *EventCreatesSingleUnitController,
) createBusinessPartnerRequestGeneral(
	requestPram *apiInputReader.Request,
	inputEventCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
		BusinessPartner:     inputEventCreatesSingleUnitBP.BusinessPartnerGeneral.BusinessPartner,
		IsMarkedForDeletion: inputEventCreatesSingleUnitBP.BusinessPartnerPersonOrganization.IsMarkedForDeletion,
	})

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
	controller *EventCreatesSingleUnitController,
) createBusinessPartnerRequestPersonOrganization(
	requestPram *apiInputReader.Request,
	inputEventCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{}

	input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{
		BusinessPartner: inputEventCreatesSingleUnitBP.BusinessPartnerPersonOrganization.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPersonOrganization(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonOrganization Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) createBusinessPartnerRequestGeneralOrganizationBP(
	requestPram *apiInputReader.Request,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner:     v.OrganizationBusinessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneralOrganizationBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerRole string,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {
	inputBusinessPartnerRole := businessPartnerRole

	input := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRole{
		BusinessPartnerRole: inputBusinessPartnerRole,
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
	controller *EventCreatesSingleUnitController,
) CreateEventTypeRequestTexts(
	requestPram *apiInputReader.Request,
	inputEventCreatesSingleUnitEventType apiInputReader.EventTypeGlobal,
) *apiModuleRuntimesResponsesEventType.EventTypeRes {
	input := apiModuleRuntimesRequestsEventType.EventType{
		Text: []apiModuleRuntimesRequestsEventType.Text{
			{
				Language:            inputEventCreatesSingleUnitEventType.EventTypeText.Language,
				IsMarkedForDeletion: inputEventCreatesSingleUnitEventType.EventTypeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesEventType.EventTypeRes{}
	responseBody := apiModuleRuntimesRequestsEventType.EventTypeReadsTexts(
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
		controller.CustomLogger.Error("CreateEventTypeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) CreateDistributionProfileRequestTexts(
	requestPram *apiInputReader.Request,
	inputEventCreatesSingleUnitDistributionProfile apiInputReader.DistributionProfileGlobal,
) *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes {
	input := apiModuleRuntimesRequestsDistributionProfile.DistributionProfile{
		Text: []apiModuleRuntimesRequestsDistributionProfile.Text{
			{
				Language:            inputEventCreatesSingleUnitDistributionProfile.DistributionProfileText.Language,
				IsMarkedForDeletion: inputEventCreatesSingleUnitDistributionProfile.DistributionProfileText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes{}
	responseBody := apiModuleRuntimesRequestsDistributionProfile.DistributionProfileReadsTexts(
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
		controller.CustomLogger.Error("CreateDistributionProfileRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) CreatePointConditionTypeRequestTexts(
	requestPram *apiInputReader.Request,
	inputEventCreatesSingleUnitPointConditionType apiInputReader.PointConditionTypeGlobal,
) *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes {
	input := apiModuleRuntimesRequestsPointConditionType.PointConditionType{
		Text: []apiModuleRuntimesRequestsPointConditionType.Text{
			{
				Language:            inputEventCreatesSingleUnitPointConditionType.PointConditionTypeText.Language,
				IsMarkedForDeletion: inputEventCreatesSingleUnitPointConditionType.PointConditionTypeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes{}
	responseBody := apiModuleRuntimesRequestsPointConditionType.PointConditionTypeReadsTexts(
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
		controller.CustomLogger.Error("CreatePointConditionTypeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) createSiteRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteRes {
	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
		requestPram,
		input,
		&controller.Controller,
		"AddressesByLocalSubRegion",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対してのサイトが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) createSiteRequestHeadersBySites(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	for _, v := range *siteRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site:                v.Site,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReadsHeadersBySites(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.Header == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対して有効なサイトヘッダデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestHeadersBySites Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
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
	controller *EventCreatesSingleUnitController,
) createShopRequestHeadersByShopOwner(
	requestPram *apiInputReader.Request,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesShop.ShopRes {
	var input apiModuleRuntimesRequestsShop.Header

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		shopOwner := v.OrganizationBusinessPartner

		input = apiModuleRuntimesRequestsShop.Header{
			ShopOwner:           &shopOwner,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsHeadersByShopOwner(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"OrganizationBusinessPartnerに対して有効な店舗が見つかりませんでした",
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
		controller.CustomLogger.Error("createShopRequestHeadersByShopOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesSingleUnitController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	var input = apiInputReader.Shop{}

	for _, v := range *shopHeaderRes.Message.Header {
		input = apiInputReader.Shop{
			ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
				Shop: v.Shop,
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
	controller *EventCreatesSingleUnitController,
) request(
	inputEventCreatesSingleUnitBP apiInputReader.BusinessPartner,
	inputEventCreatesSingleUnitEventType apiInputReader.EventTypeGlobal,
	inputEventCreatesSingleUnitDistributionProfile apiInputReader.DistributionProfileGlobal,
	inputEventCreatesSingleUnitPointConditionType apiInputReader.PointConditionTypeGlobal,
	inputEventCreatesSingleUnitSiteAddress apiInputReader.Site,
	businessPartnerRole string,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(7)

	var businessPartnerGeneralRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerPersonOrganizationRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerGeneralOrganizationBPRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	var eventTypeTextRes *apiModuleRuntimesResponsesEventType.EventTypeRes

	var distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes

	var pointConditionTypeTextRes *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes

	var siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes

	var siteHeaderRes apiModuleRuntimesResponsesSite.SiteRes
	var shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes

	var siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes
	var shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	go func() {
		defer wg.Done()
		businessPartnerGeneralRes = *controller.createBusinessPartnerRequestGeneral(
			controller.UserInfo,
			inputEventCreatesSingleUnitBP,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerPersonOrganizationRes = *controller.createBusinessPartnerRequestPersonOrganization(
			controller.UserInfo,
			inputEventCreatesSingleUnitBP,
		)
		businessPartnerGeneralOrganizationBPRes = *controller.createBusinessPartnerRequestGeneralOrganizationBP(
			controller.UserInfo,
			&businessPartnerPersonOrganizationRes,
			isMarkedForDeletion,
		)
		shopHeaderRes = *controller.createShopRequestHeadersByShopOwner(
			controller.UserInfo,
			&businessPartnerPersonOrganizationRes,
			isReleased,
			isMarkedForDeletion,
		)
		shopHeaderDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			shopHeaderRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerRoleTextRes = controller.CreateBusinessPartnerRoleRequestText(
			controller.UserInfo,
			businessPartnerRole,
		)
	}()

	go func() {
		defer wg.Done()
		eventTypeTextRes = controller.CreateEventTypeRequestTexts(
			controller.UserInfo,
			inputEventCreatesSingleUnitEventType,
		)
	}()

	go func() {
		defer wg.Done()
		distributionProfileTextRes = controller.CreateDistributionProfileRequestTexts(
			controller.UserInfo,
			inputEventCreatesSingleUnitDistributionProfile,
		)
	}()

	go func() {
		defer wg.Done()
		pointConditionTypeTextRes = controller.CreatePointConditionTypeRequestTexts(
			controller.UserInfo,
			inputEventCreatesSingleUnitPointConditionType,
		)
	}()

	go func() {
		defer wg.Done()
		siteAddressRes = controller.createSiteRequestAddressesByLocalSubRegion(
			controller.UserInfo,
			inputEventCreatesSingleUnitSiteAddress,
		)
		siteHeaderRes = *controller.createSiteRequestHeadersBySites(
			controller.UserInfo,
			siteAddressRes,
			isReleased,
			isMarkedForDeletion,
		)
		siteHeaderDocRes = controller.createSiteDocRequest(
			controller.UserInfo,
			inputEventCreatesSingleUnitSiteAddress,
		)
	}()

	wg.Wait()

	controller.fin(
		&businessPartnerGeneralRes,
		&businessPartnerPersonOrganizationRes,
		&businessPartnerGeneralOrganizationBPRes,
		businessPartnerRoleTextRes,
		eventTypeTextRes,
		distributionProfileTextRes,
		pointConditionTypeTextRes,
		siteAddressRes,
		&siteHeaderRes,
		siteHeaderDocRes,
		&shopHeaderRes,
		shopHeaderDocRes,
	)
}

func (
	controller *EventCreatesSingleUnitController,
) fin(
	businessPartnerGeneralRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerGeneralOrganizationBPRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	eventTypeTextRes *apiModuleRuntimesResponsesEventType.EventTypeRes,
	distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes,
	pointConditionTypeTextRes *apiModuleRuntimesResponsesPointConditionType.PointConditionTypeRes,
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
) {
	businessPartnerNameMapper := services.BusinessPartnerNameMapper(
		businessPartnerGeneralRes,
	)

	businessPartnerNameMapperOrganizationBP := services.BusinessPartnerNameMapper(
		businessPartnerGeneralOrganizationBPRes,
	)

	businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
		businessPartnerRoleTextRes.Message.Text,
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

	siteHeadersMapper := services.SiteHeadersMapper(
		siteHeaderRes,
	)

	data := EventCreatesSingleUnit{}

	for _, v := range *businessPartnerGeneralRes.Message.General {
		data.BusinessPartnerGeneral = append(data.BusinessPartnerGeneral,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: businessPartnerNameMapper[v.BusinessPartner].BusinessPartnerName,
			},
		)
	}

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		data.BusinessPartnerPersonOrganization = append(data.BusinessPartnerPersonOrganization,
			apiOutputFormatter.BusinessPartnerPersonOrganization{
				BusinessPartner:                 v.BusinessPartner,
				OrganizationBusinessPartner:     v.OrganizationBusinessPartner,
				OrganizationBusinessPartnerName: businessPartnerNameMapperOrganizationBP[v.OrganizationBusinessPartner].BusinessPartnerName,
			},
		)
	}

	for _, v := range *businessPartnerRoleTextRes.Message.Text {
		data.BusinessPartnerRoleText = append(data.BusinessPartnerRoleText,
			apiOutputFormatter.BusinessPartnerRoleText{
				BusinessPartnerRole:     v.BusinessPartnerRole,
				Language:                v.Language,
				BusinessPartnerRoleName: businessPartnerRoleTextMapper[v.BusinessPartnerRole].BusinessPartnerRoleName,
			},
		)
	}

	for _, v := range *eventTypeTextRes.Message.Text {
		data.EventTypeText = append(data.EventTypeText,
			apiOutputFormatter.EventTypeText{
				EventType:     v.EventType,
				Language:      v.Language,
				EventTypeName: eventTypeTextMapper[v.EventType].EventTypeName,
			},
		)
	}

	for _, v := range *distributionProfileTextRes.Message.Text {
		data.DistributionProfileText = append(data.DistributionProfileText,
			apiOutputFormatter.DistributionProfileText{
				DistributionProfile:     v.DistributionProfile,
				Language:                v.Language,
				DistributionProfileName: distributionProfileTextMapper[v.DistributionProfile].DistributionProfileName,
			},
		)
	}

	for _, v := range *pointConditionTypeTextRes.Message.Text {
		data.PointConditionTypeText = append(data.PointConditionTypeText,
			apiOutputFormatter.PointConditionTypeText{
				PointConditionType:     v.PointConditionType,
				Language:               v.Language,
				PointConditionTypeName: pointConditionTypeTextMapper[v.PointConditionType].PointConditionTypeName,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {
		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				Description: v.Description,
			},
		)
	}

	for _, v := range *siteAddressRes.Message.Address {
		siteType := siteHeadersMapper[strconv.Itoa(v.Site)].SiteType
		validityStartDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartDate
		validityStartTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartTime
		validityEndDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndDate
		validityEndTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndTime
		description := siteHeadersMapper[strconv.Itoa(v.Site)].Description
		introduction := siteHeadersMapper[strconv.Itoa(v.Site)].Introduction
		tag1 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag1
		tag2 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag2
		tag3 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag3
		tag4 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag4
		img := services.ReadSiteImage(
			siteHeaderDocRes,
			v.Site,
		)

		data.SiteAddressWithHeader = append(data.SiteAddressWithHeader,
			apiOutputFormatter.SiteAddressWithHeader{
				Site:              v.Site,
				AddressID:         v.AddressID,
				LocalSubRegion:    v.LocalSubRegion,
				LocalRegion:       v.LocalRegion,
				SiteType:          siteType,
				ValidityStartDate: validityStartDate,
				ValidityStartTime: validityStartTime,
				ValidityEndDate:   validityEndDate,
				ValidityEndTime:   validityEndTime,
				Description:       description,
				Introduction:      introduction,
				Tag1:              tag1,
				Tag2:              tag2,
				Tag3:              tag3,
				Tag4:              tag4,
				PostalCode:        v.PostalCode,
				Country:           v.Country,
				GlobalRegion:      v.GlobalRegion,
				TimeZone:          v.TimeZone,
				District:          v.District,
				StreetName:        v.StreetName,
				CityName:          v.CityName,
				Building:          v.Building,
				Floor:             v.Floor,
				Room:              v.Room,
				XCoordinate:       v.XCoordinate,
				YCoordinate:       v.YCoordinate,
				ZCoordinate:       v.ZCoordinate,
				Images: apiOutputFormatter.Images{
					Site: img,
				},
			},
		)
	}

	for _, v := range *shopHeaderRes.Message.Header {
		img := services.ReadShopImage(
			shopHeaderDocRes,
			v.Shop,
		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Shop: img,
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
