package controllersShopSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiModuleRuntimesResponsesShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ShopSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ShopSingleUnit struct {
	ShopHeader                          []apiOutputFormatter.ShopHeader             `json:"ShopHeader"`
	ShopAddress                         []apiOutputFormatter.ShopAddress            `json:"ShopAddress"`
	BusinessPartnerGeneralShopOwner     []apiOutputFormatter.BusinessPartnerGeneral `json:"BusinessPartnerGeneralShopOwner"`
	BusinessPartnerPersonCreateUser     []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPersonCreateUser"`
	BusinessPartnerPersonLastChangeUser []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPersonLastChangeUser"`
}

func (controller *ShopSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "shop"
	redisKeyCategory2 := "single-unit"
	shop, _ := controller.GetInt("shop")

	isReleased := true
	isMarkedForDeletion := false

	ShopSingleUnitShopHeader := apiInputReader.Shop{}

	ShopSingleUnitShopHeader = apiInputReader.Shop{
		ShopHeader: &apiInputReader.ShopHeader{
			Shop:                shop,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ShopAddress: &apiInputReader.ShopAddress{
			Shop: shop,
		},
		ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
			Shop: shop,
			//DocType:				    &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(shop),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData ShopSingleUnit

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
				ShopSingleUnitShopHeader,
			)
		}()
	} else {
		controller.request(
			ShopSingleUnitShopHeader,
		)
	}
}

func (
	controller *ShopSingleUnitController,
) createShopRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
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
	controller *ShopSingleUnitController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
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
	controller *ShopSingleUnitController,
) createShopRequestAddresses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
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
		controller.CustomLogger.Error("createShopRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopSingleUnitController,
) createBusinessPartnerRequestGeneralShopOwner(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*shopRes.Message.Header))

	for _, v := range *shopRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: *v.ShopOwner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneralShopOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopSingleUnitController,
) createBusinessPartnerRequestPersonCreateUser(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := apiModuleRuntimesRequestsBusinessPartner.Person{}

	for _, v := range *shopRes.Message.Header {
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
	controller *ShopSingleUnitController,
) createBusinessPartnerRequestPersonLastChangeUser(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := apiModuleRuntimesRequestsBusinessPartner.Person{}

	for _, v := range *shopRes.Message.Header {
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
	controller *ShopSingleUnitController,
) CreateShopTypeRequestText(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShopType.ShopTypeRes {

	shopType := &(*shopRes.Message.Header)[0].ShopType

	var inputShopType *string

	if shopType != nil {
		inputShopType = shopType
	}

	input := apiModuleRuntimesRequestsShopType.ShopType{
		ShopType: *inputShopType,
	}

	responseJsonData := apiModuleRuntimesResponsesShopType.ShopTypeRes{}
	responseBody := apiModuleRuntimesRequestsShopType.ShopTypeReadsText(
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
		controller.CustomLogger.Error("CreateShopTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopSingleUnitController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*shopRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*shopRes.Message.Address)[0].LocalRegion
	country := &(*shopRes.Message.Address)[0].Country

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
	controller *ShopSingleUnitController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*shopRes.Message.Address)[0].LocalRegion
	country := &(*shopRes.Message.Address)[0].Country

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
	controller *ShopSingleUnitController,
) request(
	inputShopHeader apiInputReader.Shop,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createShopRequestHeader(
		controller.UserInfo,
		inputShopHeader,
	)

	addressRes := *controller.createShopRequestAddresses(
		controller.UserInfo,
		inputShopHeader,
	)

	headerDocRes := controller.createShopDocRequest(
		controller.UserInfo,
		inputShopHeader,
	)

	businessPartnerGeneralResShopOwner := *controller.createBusinessPartnerRequestGeneralShopOwner(
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

	shopTypeTextRes := controller.CreateShopTypeRequestText(
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
		&businessPartnerGeneralResShopOwner,
		&businessPartnerPersonResCreateUser,
		&businessPartnerPersonResLastChangeUser,
		shopTypeTextRes,
		localSubRegionTextRes,
		localRegionTextRes,
	)
}

func (
	controller *ShopSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesShop.ShopRes,
	addressRes *apiModuleRuntimesResponsesShop.ShopRes,
	headerDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
	businessPartnerGeneralResShopOwner *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonResCreateUser *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonResLastChangeUser *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	shopTypeTextRes *apiModuleRuntimesResponsesShopType.ShopTypeRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {

	shopTypeTextMapper := services.ShopTypeTextMapper(
		shopTypeTextRes.Message.Text,
	)

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	data := ShopSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadShopImage(
			headerDocRes,
			v.Shop,
		)

		qrcode := services.CreateQRCodeShopDocImage(
			headerDocRes,
			v.Shop,
		)

		documentImage := services.ReadDocumentImageShop(
			headerDocRes,
			v.Shop,
		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:                         v.Shop,
				ShopType:                     v.ShopType,
				ShopTypeName:                 shopTypeTextMapper[v.ShopType].ShopTypeName,
				Brand:                        v.Brand,
				//BrandDescription:           v.BrandDescription,
				PersonResponsible:            *v.PersonResponsible,
				URL:                          v.URL,
				DailyOperationStartTime:      v.DailyOperationStartTime,
				DailyOperationEndTime:        v.DailyOperationEndTime,
				Description:                  v.Description,
				LongText:                     v.LongText,
				Introduction:                 v.Introduction,
				OperationRemarks:             v.OperationRemarks,
				AvailabilityOfParking:        v.AvailabilityOfParking,
				NumberOfParkingSpaces:        v.NumberOfParkingSpaces,
				SuperiorShop:                 v.SuperiorShop,
				Tag1:                         v.Tag1,
				Tag2:                         v.Tag2,
				Tag3:                         v.Tag3,
				Tag4:                         v.Tag4,
				//PointConsumptionType:       v.PointConsumptionType,
				//PointConsumptionTypeName:   pointConsumptionTypeTextMapper[v.PointConsumptionType].PointConsumptionTypeName,
				CreateUser:                   v.CreateUser,
				//CreateUserFullName:         v.CreateUserFullName,
				//CreateUserNickName:         v.CreateUserNickName,
				LastChangeUser: v.LastChangeUser,
				//LastChangeUserFullName:     v.LastChangeUserFullName,
				//LastChangeUserNickName:     v.LastChangeUserNickName,
				Images: apiOutputFormatter.Images{
					Shop:              img,
					QRCode:            qrcode,
					DocumentImageShop: documentImage,
				},
			},
		)
	}

	for _, v := range *addressRes.Message.Address {
		data.ShopAddress = append(data.ShopAddress,
			apiOutputFormatter.ShopAddress{
				Shop:               v.Shop,
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

	for _, v := range *businessPartnerGeneralResShopOwner.Message.General {
		data.BusinessPartnerGeneralShopOwner = append(data.BusinessPartnerGeneralShopOwner,
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
