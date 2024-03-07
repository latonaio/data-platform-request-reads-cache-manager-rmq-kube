package controllersInspectionLotList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
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

//const (
//	buyer  = "buyer"
//	seller = "seller"
//)

func (controller *InspectionLotListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
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
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	partnersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*partnersRes.Message.Partner))

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
) request(
	input apiInputReader.InspectionLot,
) {
	defer services.Recover(controller.CustomLogger)

	headersRes := *controller.createInspectionLotRequestHeaders(
		controller.UserInfo,
		input,
	)

	partnersRes := *controller.createInspectionLotRequestPartners(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&partnersRes,
	)

	controller.fin(
		&headersRes,
		&partnersRes,
		&businessPartnerRes,
	)
}

func (
	controller *InspectionLotListController,
) fin(
	headersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	partnersRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
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
		//partnerFunction := inspectionLotListMapper[v.InspectionLot].PartnerFunction
		//businessPartner := inspectionLotListMapper[v.InspectionLot].BusinessPartner
		//businessPartnerName := inspectionLotListMapper[v.InspectionLot].BusinessPartnerName

		data.InspectionLotHeader = append(data.InspectionLotHeader,
			apiOutputFormatter.InspectionLotHeader{
				InspectionLot:     v.InspectionLot,
				InspectionLotDate: v.InspectionLotDate,
				Product:           v.Product,
				//PartnerFunction:     &partnerFunction,
				//BusinessPartner:     &businessPartner,
				//BusinessPartnerName: businessPartnerName,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
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
