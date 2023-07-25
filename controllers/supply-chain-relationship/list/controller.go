package controllersSupplyChainRelationshipList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsSupplyChainRelationship "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/supply-chain-relationship"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesSupplyChainRelationship "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/supply-chain-relationship"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SupplyChainRelationshipListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	buyer  = "buyer"
	seller = "seller"
)

func (controller *SupplyChainRelationshipListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "supply-chain-relationship"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	supplyChainRelationshipGeneral := apiInputReader.SupplyChainRelationship{}

	if userType == buyer {
		supplyChainRelationshipGeneral = apiInputReader.SupplyChainRelationship{
			SupplyChainRelationshipGeneral: &apiInputReader.SupplyChainRelationshipGeneral{
				Buyer:               controller.UserInfo.BusinessPartner,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == seller {
		supplyChainRelationshipGeneral = apiInputReader.SupplyChainRelationship{
			SupplyChainRelationshipGeneral: &apiInputReader.SupplyChainRelationshipGeneral{
				Seller:              controller.UserInfo.BusinessPartner,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
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
		var responseData apiOutputFormatter.SupplyChainRelationship

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
			controller.request(supplyChainRelationshipGeneral)
		}()
	} else {
		controller.request(supplyChainRelationshipGeneral)
	}
}

func (
	controller *SupplyChainRelationshipListController,
) createSupplyChainRelationshipRequestGeneralByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.SupplyChainRelationship,
) *apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes {
	responseJsonData := apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes{}
	responseBody := apiModuleRuntimesRequestsSupplyChainRelationship.SupplyChainRelationshipReads(
		requestPram,
		input,
		&controller.Controller,
		"GeneralsByBuyer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("SupplyChainRelationshipReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SupplyChainRelationshipListController,
) createSupplyChainRelationshipRequestGeneralBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.SupplyChainRelationship,
) *apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes {
	responseJsonData := apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes{}
	responseBody := apiModuleRuntimesRequestsSupplyChainRelationship.SupplyChainRelationshipReads(
		requestPram,
		input,
		&controller.Controller,
		"GeneralsBySeller",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("SupplyChainRelationshipReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SupplyChainRelationshipListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	supplyChainRelationshipRes *apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*supplyChainRelationshipRes.Message.General))

	for _, v := range *supplyChainRelationshipRes.Message.General {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		generals,
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
	controller *SupplyChainRelationshipListController,
) request(
	input apiInputReader.SupplyChainRelationship,
) {
	defer services.Recover(controller.CustomLogger)

	supplyChainRelationshipsRes := apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.SupplyChainRelationshipGeneral.Buyer != nil {
		supplyChainRelationshipsRes = *controller.createSupplyChainRelationshipRequestGeneralByBuyer(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&supplyChainRelationshipsRes,
		)
	}

	if input.SupplyChainRelationshipGeneral.Seller != nil {
		supplyChainRelationshipsRes = *controller.createSupplyChainRelationshipRequestGeneralBySeller(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&supplyChainRelationshipsRes,
		)
	}

	controller.fin(
		&supplyChainRelationshipsRes,
		&businessPartnerRes,
	)
}

func (
	controller *SupplyChainRelationshipListController,
) fin(
	supplyChainRelationshipRes *apiModuleRuntimesResponsesSupplyChainRelationship.SupplyChainRelationshipRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.SupplyChainRelationship{}

	for _, v := range *supplyChainRelationshipRes.Message.General {
		data.SupplyChainRelationshipGeneral = append(data.SupplyChainRelationshipGeneral,
			apiOutputFormatter.SupplyChainRelationshipGeneral{
				SupplyChainRelationshipID: v.SupplyChainRelationshipID,
				Buyer:                     v.Buyer,
				BuyerName:                 businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                    v.Seller,
				SellerName:                businessPartnerMapper[v.Seller].BusinessPartnerName,
				IsMarkedForDeletion:       v.IsMarkedForDeletion,
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
