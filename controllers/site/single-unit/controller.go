package controllersSiteSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsSiteType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site-type"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiModuleRuntimesResponsesSiteType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type SiteSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type SiteSingleUnit struct {
	SiteHeader                          []apiOutputFormatter.SiteHeader             `json:"SiteHeader"`
	SiteAddress                         []apiOutputFormatter.SiteAddress            `json:"SiteAddress"`
	BusinessPartnerGeneralSiteOwner     []apiOutputFormatter.BusinessPartnerGeneral `json:"BusinessPartnerGeneralSiteOwner"`
	BusinessPartnerPersonCreateUser     []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPersonCreateUser"`
	BusinessPartnerPersonLastChangeUser []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPersonLastChangeUser"`
}

func (controller *SiteSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "site"
	redisKeyCategory2 := "site-single-unit"
	site, _ := controller.GetInt("site")

	isReleased := true
	isMarkedForDeletion := false

	//docType := "QRCODE"

	SiteSingleUnitSiteHeader := apiInputReader.Site{}

	//	BusinessPartnerPerson := apiInputReader.BusinessPartner{}

	SiteSingleUnitSiteHeader = apiInputReader.Site{
		SiteHeader: &apiInputReader.SiteHeader{
			Site:                site,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		SiteAddress: &apiInputReader.SiteAddress{
			Site: site,
		},
		SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{
			Site: site,
			//DocType:				    &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	//	BusinessPartnerPerson = apiInputReader.BusinessPartner{
	//		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
	//			BusinessPartner:     businessPartner,
	//			IsMarkedForDeletion: &isMarkedForDeletion,
	//		},
	//	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(site),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData SiteSingleUnit

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
				SiteSingleUnitSiteHeader,
			)
		}()
	} else {
		controller.request(
			SiteSingleUnitSiteHeader,
		)
	}
}

func (
	controller *SiteSingleUnitController,
) createSiteRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteRes {
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
	controller *SiteSingleUnitController,
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
	controller *SiteSingleUnitController,
) createSiteRequestAddresses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteRes {
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
	controller *SiteSingleUnitController,
) createBusinessPartnerRequestGeneralSiteOwner(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*siteRes.Message.Header))

	for _, v := range *siteRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: *v.SiteOwner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneralSiteOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) createBusinessPartnerRequestPersonCreateUser(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := apiModuleRuntimesRequestsBusinessPartner.Person{}

	for _, v := range *siteRes.Message.Header {
		input = apiModuleRuntimesRequestsBusinessPartner.Person{
			BusinessPartner: v.CreateUser,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPerson(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonCreateUser Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) createBusinessPartnerRequestPersonLastChangeUser(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := apiModuleRuntimesRequestsBusinessPartner.Person{}

	for _, v := range *siteRes.Message.Header {
		input = apiModuleRuntimesRequestsBusinessPartner.Person{
			BusinessPartner: v.LastChangeUser,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPerson(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonLastChangeUser Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) CreateSiteTypeRequestText(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesSiteType.SiteTypeRes {

	siteType := &(*siteRes.Message.Header)[0].SiteType

	var inputSiteType *string

	if siteType != nil {
		inputSiteType = siteType
	}

	input := apiModuleRuntimesRequestsSiteType.SiteType{
		SiteType: *inputSiteType,
	}

	responseJsonData := apiModuleRuntimesResponsesSiteType.SiteTypeRes{}
	responseBody := apiModuleRuntimesRequestsSiteType.SiteTypeReadsText(
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
		controller.CustomLogger.Error("CreateSiteTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*siteRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*siteRes.Message.Address)[0].LocalRegion
	country := &(*siteRes.Message.Address)[0].Country

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
		controller.CustomLogger.Error("LocalSubRegionReadsText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*siteRes.Message.Address)[0].LocalRegion
	country := &(*siteRes.Message.Address)[0].Country

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
		controller.CustomLogger.Error("LocalRegionReadsText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteSingleUnitController,
) request(
	inputSiteHeader apiInputReader.Site,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createSiteRequestHeader(
		controller.UserInfo,
		inputSiteHeader,
	)

	addressRes := *controller.createSiteRequestAddresses(
		controller.UserInfo,
		inputSiteHeader,
	)

	headerDocRes := controller.createSiteDocRequest(
		controller.UserInfo,
		inputSiteHeader,
	)

	businessPartnerGeneralResSiteOwner := *controller.createBusinessPartnerRequestGeneralSiteOwner(
		controller.UserInfo,
		&headerRes,
	)

	businessPartnerPersonResCreateUser := *controller.createBusinessPartnerRequestPersonCreateUser(
		controller.UserInfo,
		&headerRes,
	)

	businessPartnerPersonResLastChangeUser := *controller.createBusinessPartnerRequestPersonLastChangeUser(
		controller.UserInfo,
		&headerRes,
	)

	siteTypeTextRes := controller.CreateSiteTypeRequestText(
		controller.UserInfo,
		&headerRes,
	)

	localSubRegionTextRes := controller.CreateLocalSubRegionRequestText(
		controller.UserInfo,
		&addressRes,
	)

	localRegionTextRes := controller.CreateLocalRegionRequestText(
		controller.UserInfo,
		&addressRes,
	)

	controller.fin(
		&headerRes,
		&addressRes,
		headerDocRes,
		&businessPartnerGeneralResSiteOwner,
		&businessPartnerPersonResCreateUser,
		&businessPartnerPersonResLastChangeUser,
		siteTypeTextRes,
		localSubRegionTextRes,
		localRegionTextRes,
	)
}

func (
	controller *SiteSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesSite.SiteRes,
	addressRes *apiModuleRuntimesResponsesSite.SiteRes,
	headerDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	businessPartnerGeneralResSiteOwner *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonResCreateUser *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonResLastChangeUser *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	siteTypeTextRes *apiModuleRuntimesResponsesSiteType.SiteTypeRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {

	siteTypeTextMapper := services.SiteTypeTextMapper(
		siteTypeTextRes.Message.Text,
	)

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	data := SiteSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadSiteImage(
			headerDocRes,
			v.Site,
		)

		qrcode := services.CreateQRCodeSiteDocImage(
			headerDocRes,
			v.Site,
		)

		documentImage := services.ReadDocumentImageSite(
			headerDocRes,
			v.Site,
		)

		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:         v.Site,
				SiteType:     v.SiteType,
				SiteTypeName: siteTypeTextMapper[v.SiteType].SiteTypeName,
				Brand:        v.Brand,
				//BrandDescription:				  	v.BrandDescription,
				PersonResponsible:       *v.PersonResponsible,
				URL:                     v.URL,
				DailyOperationStartTime: v.DailyOperationStartTime,
				DailyOperationEndTime:   v.DailyOperationEndTime,
				Description:             v.Description,
				LongText:                v.LongText,
				Introduction:            v.Introduction,
				OperationRemarks:        v.OperationRemarks,
				AvailabilityOfParking:   v.AvailabilityOfParking,
				NumberOfParkingSpaces:   v.NumberOfParkingSpaces,
				SuperiorSite:            v.SuperiorSite,
				Tag1:                    v.Tag1,
				Tag2:                    v.Tag2,
				Tag3:                    v.Tag3,
				Tag4:                    v.Tag4,
				//				PointConsumptionType:     		  	v.PointConsumptionType,
				//				PointConsumptionTypeName: 		  	pointConsumptionTypeTextMapper[v.PointConsumptionType].PointConsumptionTypeName,
				CreateUser: v.CreateUser,
				//				CreateUserFullName:					v.CreateUserFullName,
				//				CreateUserNickName:					v.CreateUserNickName,
				LastChangeUser: v.LastChangeUser,
				//				LastChangeUserFullName:				v.LastChangeUserFullName,
				//				LastChangeUserNickName:           	v.LastChangeUserNickName,
				Images: apiOutputFormatter.Images{
					Site:              img,
					QRCode:            qrcode,
					DocumentImageSite: documentImage,
				},
			},
		)
	}

	for _, v := range *addressRes.Message.Address {
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

	for _, v := range *businessPartnerGeneralResSiteOwner.Message.General {
		data.BusinessPartnerGeneralSiteOwner = append(data.BusinessPartnerGeneralSiteOwner,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: v.BusinessPartnerName,
			},
		)
	}

	for _, v := range *businessPartnerPersonResCreateUser.Message.Person {
		data.BusinessPartnerPersonCreateUser = append(data.BusinessPartnerPersonCreateUser,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
			},
		)
	}

	for _, v := range *businessPartnerPersonResLastChangeUser.Message.Person {
		data.BusinessPartnerPersonLastChangeUser = append(data.BusinessPartnerPersonLastChangeUser,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
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
