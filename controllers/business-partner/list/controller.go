package controllersBusinessPartnerList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BusinessPartnerListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

//const (
//	buyer	= "buyer"
//	seller	= "seller"
//)

func (controller *BusinessPartnerListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "business-partner"
	redisKeyCategory2 := "list"
	//userType :=
	businessPartner, _ := controller.GetInt("businessPartner")
	userType := controller.GetString(":userType")

	businessPartnerGeneral := apiInputReader.BusinessPartner{
		BusinessPartnerGeneral: &apiInputReader.BusinessPartnerGeneral{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			userType,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.BusinessPartner

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
			controller.request(businessPartnerGeneral)
		}()
	} else {
		controller.request(businessPartnerGeneral)
	}
}

func (
	controller *BusinessPartnerListController,
) createBusinessPartnerRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGenerals(
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerListController,
) request(
	input apiInputReader.BusinessPartner,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	generalRes := controller.createBusinessPartnerRequestGenerals(
		controller.UserInfo,
		input,
	)

	controller.fin(
		generalRes,
	)
}

func (
	controller *BusinessPartnerListController,
) fin(
	generalRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	data := apiOutputFormatter.BusinessPartner{}

	for _, v := range *generalRes.Message.General {

		data.BusinessPartnerGeneral = append(data.BusinessPartnerGeneral,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: v.BusinessPartnerName,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
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
