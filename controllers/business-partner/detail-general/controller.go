package controllersBusinessPartnerDetailGeneral

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
	"strconv"
)

type BusinessPartnerDetailGeneralController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *BusinessPartnerDetailGeneralController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	businessPartner, _ := controller.GetInt("businessPartner")
	redisKeyCategory1 := "business-partner"
	redisKeyCategory2 := "detail-general"
	redisKeyCategory3 := businessPartner
	userType := controller.GetString(":userType")

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	businessPartnerGeneralDetails := apiInputReader.BusinessPartner{
		BusinessPartnerGeneral: &apiInputReader.BusinessPartnerGeneral{
			BusinessPartner: businessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
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
			controller.request(businessPartnerGeneralDetails)
		}()
	} else {
		controller.request(businessPartnerGeneralDetails)
	}
}

func (
	controller *BusinessPartnerDetailGeneralController,
) createBusinessPartnerRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	general := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	general = append(general, apiModuleRuntimesRequestsBusinessPartner.General{
		BusinessPartner: input.BusinessPartnerGeneral.BusinessPartner,
	})

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		general,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerDetailGeneralController,
) request(
	input apiInputReader.BusinessPartner,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	generalRes := controller.createBusinessPartnerRequestGeneral(
		controller.UserInfo,
		input,
	)

	controller.fin(
		generalRes,
	)
}

func (
	controller *BusinessPartnerDetailGeneralController,
) fin(
	generalRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	data := apiOutputFormatter.BusinessPartner{}

	for _, v := range *generalRes.Message.General {

		data.BusinessPartnerGeneral = append(data.BusinessPartnerGeneral,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: v.BusinessPartnerName,
			},
		)

		data.BusinessPartnerDetailGeneral = append(data.BusinessPartnerDetailGeneral,
			apiOutputFormatter.BusinessPartnerDetailGeneral{
				BusinessPartnerFullName:  v.BusinessPartnerFullName,
				Industry:                 v.Industry,
				LegalEntityRegistration:  v.LegalEntityRegistration,
				Country:                  v.Country,
				Language:                 &v.Language,
				Currency:                 v.Currency,
				AddressID:                v.AddressID,
				BusinessPartnerIsBlocked: v.BusinessPartnerIsBlocked,
				CreationDate:             v.CreationDate,
				LastChangeDate:           v.LastChangeDate,
				IsMarkedForDeletion:      v.IsMarkedForDeletion,
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
