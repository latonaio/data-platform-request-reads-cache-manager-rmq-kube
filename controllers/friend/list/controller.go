package controllersFriendList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	apiModuleRuntimesRequestsFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/friend"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type FriendListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *FriendListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	businessPartner, _ := controller.GetInt("businessPartner")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "businessPartner"
	redisKeyCategory2 := "friend-list"
	redisKeyCategory3 := businessPartner

	friendIsBlocked := false
	isMarkedForDeletion := false

	FriendGeneral := apiInputReader.Friend{
		FriendGeneral: &apiInputReader.FriendGeneral{
			BusinessPartner:     businessPartner,
			FriendIsBlocked:     friendIsBlocked,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartner := apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{},
	}

	//FriendBusinessPartnerDocGeneralDoc := apiInputReader.BusinessPartner{
	//	BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
	//		BusinessPartner: Friend, // 取得できたFriendをBusinessPartnerに置き換える
	//		//DocType:					&docType,
	//		DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
	//	},
	//}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Friend

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
			controller.request(FriendGeneral, BusinessPartner)
		}()
	} else {
		controller.request(FriendGeneral, BusinessPartner)
	}
}

func (
	controller *FriendListController,
) createFriendRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.Friend,
) *apiModuleRuntimesResponsesFriend.FriendRes {
	responseJsonData := apiModuleRuntimesResponsesFriend.FriendRes{}
	responseBody := apiModuleRuntimesRequestsFriend.FriendReads(
		requestPram,
		input,
		&controller.Controller,
		"Generals",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if len(*responseJsonData.Message.General) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"フレンドが見つかりませんでした",
	//		&status,
	//	)
	//	return nil
	//}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createFriendRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *FriendListController,
) createBusinessPartnerDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerDoc.BusinessPartnerDocReads(
		requestPram,
		input,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.GeneralDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ビジネスパートナヘッダに画像が見つかりませんでした",
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
		controller.CustomLogger.Error("createBusinessPartnerDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *FriendListController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	friendGeneralRes apiModuleRuntimesResponsesFriend.FriendRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*friendGeneralRes.Message.General))

	for _, v := range *friendGeneralRes.Message.General {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Friend,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPersonsByBusinessPartners(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPerson Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *FriendListController,
) request(
	inputFriend apiInputReader.Friend,
	inputBusinessPartner apiInputReader.BusinessPartner,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes
	var businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes

	generalRes := *controller.createFriendRequestGenerals(
		controller.UserInfo,
		inputFriend,
	)

	if generalRes.Message.General != nil && len(*generalRes.Message.General) != 0 {
		businessPartnerGeneralDocRes = controller.createBusinessPartnerDocRequest(
			controller.UserInfo,
			inputBusinessPartner,
		)

		businessPartnerPersonRes = controller.createBusinessPartnerRequestPerson(
			controller.UserInfo,
			generalRes,
		)
	}

	controller.fin(
		&generalRes,
		businessPartnerGeneralDocRes,
		businessPartnerPersonRes,
	)
}

func (
	controller *FriendListController,
) fin(
	friendRes *apiModuleRuntimesResponsesFriend.FriendRes,
	businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	var businessPartnerPersonMapper map[int]apiModuleRuntimesResponsesBusinessPartner.Person

	if friendRes.Message.General != nil && len(*friendRes.Message.General) != 0 {
		businessPartnerPersonMapper = services.BusinessPartnerPersonMapper(
			businessPartnerPersonRes.Message.Person,
		)
	}

	data := apiOutputFormatter.Friend{}

	for _, v := range *friendRes.Message.General {
		img := services.ReadBusinessPartnerImage(
			businessPartnerGeneralDocRes,
			v.Friend,
		)

		data.FriendGeneral = append(data.FriendGeneral,
			apiOutputFormatter.FriendGeneral{
				BusinessPartner: v.BusinessPartner,
				Friend:          v.Friend,
				//CommunityRank:	 v.CommunityRank,
				FriendNickName: businessPartnerPersonMapper[v.Friend].NickName, // Mapper対応
				Images: apiOutputFormatter.Images{
					BusinessPartner: img,
				},
			},
		)
	}

	//for _, v := range *businessPartnerPersonRes.Message.Person {
	//	data.FriendGeneral = append(data.FriendGeneral,
	//		apiOutputFormatter.FriendGeneral{
	//			BusinessPartner: v.BusinessPartner,
	//			FriendNickName:  v.NickName,
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
