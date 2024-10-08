package controllersInspectionLotList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type InspectionLotListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *InspectionLotListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "inspection-lot"
	redisKeyCategory2 := "list"

	inspectionLotList := apiInputReader.InspectionLot{}

	inspectionLot, _ := controller.GetInt("inspectionLot")

	isReleased := true
	isMarkedForDeletion := false

	inspectionLotList = apiInputReader.InspectionLot{
		InspectionLotHeader: &apiInputReader.InspectionLotHeader{
			InspectionLot:       inspectionLot,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotPartner: &apiInputReader.InspectionLotPartner{
			InspectionLot: inspectionLot,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(inspectionLot),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.InspectionLot

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
			controller.request(inspectionLotList)
		}()
	} else {
		controller.request(inspectionLotList)
	}
}

func (
	controller *InspectionLotListController,
) createInspectionLotRequestHeaders(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"Headers",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InspectionLotHeaderReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotListController,
) createInspectionLotRequestHeadersExcludeRoleManufacturer(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersExcludeRoleManufacturer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InspectionLotHeaderReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotListController,
) createInspectionLotRequestPartners(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"Partners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InspectionLotPartnersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotListController,
) createInspectionLotRequestPartnersHeadersExcludeRoleManufacturer(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"PartnersExcludeRoleManufacturer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InspectionLotPartnersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	partnersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *partnersRes.Message.Partner {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
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
	controller *InspectionLotListController,
) createBusinessPartnerRequestRole(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Role

	input = apiModuleRuntimesRequestsBusinessPartner.Role{
		BusinessPartner: *requestPram.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsRole(
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
	controller *InspectionLotListController,
) request(
	input apiInputReader.InspectionLot,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var headersRes apiModuleRuntimesResponsesInspectionLot.InspectionLotRes
	var partnersRes apiModuleRuntimesResponsesInspectionLot.InspectionLotRes

	businessPartnerRoleRes := *controller.createBusinessPartnerRequestRole(
		controller.UserInfo,
	)

	var isIncludedUserRoleManufacture bool

	for _, v := range *businessPartnerRoleRes.Message.Role {
		if v.BusinessPartnerRole == "MANUFACTURER" {
			isIncludedUserRoleManufacture = true
			break
		} else {
			isIncludedUserRoleManufacture = false
		}
	}

	// ログインユーザのBPロールがMANUFACTURERを含むなら全件表示され、含まないならMANUFACTURERの行が表示されない
	if isIncludedUserRoleManufacture {
		headersRes = *controller.createInspectionLotRequestHeaders(
			controller.UserInfo,
			input,
		)
		partnersRes = *controller.createInspectionLotRequestPartners(
			controller.UserInfo,
			input,
		)
	} else {
		headersRes = *controller.createInspectionLotRequestHeadersExcludeRoleManufacturer(
			controller.UserInfo,
			input,
		)
		partnersRes = *controller.createInspectionLotRequestPartnersHeadersExcludeRoleManufacturer(
			controller.UserInfo,
			input,
		)
	}

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&partnersRes,
	)

	controller.fin(
		&headersRes,
		&partnersRes,
		&businessPartnerRes,
		&businessPartnerRoleRes,
	)
}

func (
	controller *InspectionLotListController,
) fin(
	headersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	partnersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	inspectionLotListMapper := services.InspectionLotListMapper(
		headersRes,
		partnersRes,
	)

	data := apiOutputFormatter.InspectionLot{}

	for _, v := range *headersRes.Message.Header {

		data.InspectionLotHeader = append(data.InspectionLotHeader,
			apiOutputFormatter.InspectionLotHeader{
				InspectionLot:     v.InspectionLot,
				InspectionLotDate: v.InspectionLotDate,
				Product:           v.Product,
			},
		)
	}

	for _, v := range *partnersRes.Message.Partner {
		product := inspectionLotListMapper[v.InspectionLot].Product
		inspectionLotDate := inspectionLotListMapper[v.InspectionLot].InspectionLotDate

		data.InspectionLotPartner = append(data.InspectionLotPartner,
			apiOutputFormatter.InspectionLotPartner{
				InspectionLot:       v.InspectionLot,
				PartnerFunction:     v.PartnerFunction,
				BusinessPartner:     strconv.Itoa(v.BusinessPartner),
				BusinessPartnerName: businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
				Product:             product,
				InspectionLotDate:   inspectionLotDate,
			},
		)
	}

	//for _, v := range *businessPartnerRes.Message.General {
	//	data.InspectionLotPartner = append(data.InspectionLotPartner,
	//		apiOutputFormatter.InspectionLotPartner{
	//			InspectionLot:       v.InspectionLot,
	//			PartnerFunction:     v.PartnerFunction,
	//			BusinessPartner:     v.BusinessPartner,
	//			BusinessPartnerName: businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
	//		},
	//	)
	//}

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
