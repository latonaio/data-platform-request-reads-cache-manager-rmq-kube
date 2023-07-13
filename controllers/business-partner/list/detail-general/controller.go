package controllersBusinessPartnerDetailGeneral

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
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
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	businessPartner, _ := controller.GetInt("businessPartner")
	redisKeyCategory1 := "business-partner"
	redisKeyCategory2 := "detail-general"
	redisKeyCategory3 := businessPartner
	userType := controller.GetString("userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

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
	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		input,
		&controller.Controller,
		"General",
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
) createBusinessPartnerDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesBusinessPartnerDoc.BusinessPartnerDocRes {
	responseJsonData := apiModuleRuntimesResponsesBusinessPartnerDoc.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequests.BusinessPartnerDocReads(
		requestPram,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerDocReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerDetailGeneralController,
) request(
	input apiInputReader.BusinessPartner,
) {
	defer services.Recover(controller.CustomLogger)

	bGeneralRes := controller.createBusinessPartnerRequestGeneral(
		controller.UserInfo,
		input,
	)

	controller.fin(
		bGeneralRes,
	)
}

func (
	controller *BusinessPartnerDetailGeneralController,
) fin(
	bGeneralRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	data := apiOutputFormatter.BusinessPartner{}

	for _, v := range *bGeneralRes.Message.General {

		data.BusinessPartnerGeneral = append(data.BusinessPartnerGeneral,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:			v.BusinessPartner,
				BusinessPartnerName:		&BusinessPartnerName,
			},
		)

		data.BusinessPartnerDetailGeneral = append(data.BusinessPartnerDetailGeneral,
			apiOutputFormatter.BusinessPartnerDetailGeneral{
				BusinessPartnerFullName:		v.BusinessPartnerFullName,
				Industry:						v.Industry,
				LegalEntityRegistration:		v.LegalEntityRegistration,
				Country:						v.Country,
				Language:						v.Language,
				Currency:						v.Currency,
				AddressID:						v.AddressID,
				BusinessPartnerIsBlocked:		v.BusinessPartnerIsBlocked,
				CreationDate:					v.CreationDate,
				LastChangeDate:					v.LastChangeDate,
				IsMarkedForDeletion:			v.IsMarkedForDeletion,
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
