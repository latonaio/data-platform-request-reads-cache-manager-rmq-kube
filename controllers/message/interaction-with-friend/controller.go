package controllersMessageInteractionWithFriend

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	apiModuleRuntimesRequestsFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsMessage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/message/message"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/friend"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesMessage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/message"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type MessageInteractionWithFriendController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type MessageInteractionWithFriend struct {
	MessageHeaderFromMeToFriend  []apiOutputFormatter.MessageHeader             `json:"MessageHeaderFromMeToFriend"`
	MessageHeaderFromFriendToMe  []apiOutputFormatter.MessageHeader             `json:"MessageHeaderFromFriendToMe"`
	BusinessPartnerPersonMe      []apiOutputFormatter.BusinessPartnerPerson     `json:"BusinessPartnerPersonMe"`
	BusinessPartnerPersonFriend  []apiOutputFormatter.BusinessPartnerPerson     `json:"BusinessPartnerPersonFriend"`
	BusinessPartnerAddressFriend []apiOutputFormatter.BusinessPartnerAddress    `json:"BusinessPartnerAddressFriend"`
	BusinessPartnerGeneralDoc    []apiOutputFormatter.BusinessPartnerGeneralDoc `json:"BusinessPartnerGeneralDoc"`
	FriendGeneral                []apiOutputFormatter.FriendGeneral             `json:"FriendGeneral"`
}

func (controller *MessageInteractionWithFriendController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	friend, _ := controller.GetInt("friend")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	messageType := controller.GetString("messageType")
	redisKeyCategory1 := "sender"
	redisKeyCategory2 := "friend"
	redisKeyCategory3 := "messageType"

	isMarkedForDeletion := false
	friendIsBlocked := false

	MessageHeaderFromMeToFriend := apiInputReader.Message{}
	MessageHeaderFromFriendToMe := apiInputReader.Message{}
	BusinessPartnerPersonMe := apiInputReader.BusinessPartner{}
	BusinessPartnerPersonFriend := apiInputReader.BusinessPartner{}
	BusinessPartnerAddressFriend := apiInputReader.BusinessPartner{}
	FriendGeneral := apiInputReader.Friend{}

	MessageHeaderFromMeToFriend = apiInputReader.Message{
		MessageHeader: &apiInputReader.MessageHeader{
			Sender:              controller.UserInfo.BusinessPartner,
			Receiver:            &friend,
			MessageType:         &messageType,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	MessageHeaderFromFriendToMe = apiInputReader.Message{
		MessageHeader: &apiInputReader.MessageHeader{
			Sender:              &friend,
			Receiver:            controller.UserInfo.BusinessPartner,
			MessageType:         &messageType,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	docType := "IMAGE"

	BusinessPartnerPersonMe = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     *controller.UserInfo.BusinessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: *controller.UserInfo.BusinessPartner,
			DocType:         &docType,
		},
	}

	BusinessPartnerPersonFriend = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     friend,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerAddressFriend = apiInputReader.BusinessPartner{
		BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
			BusinessPartner: friend,
		},
	}

	FriendGeneral = apiInputReader.Friend{
		FriendGeneral: &apiInputReader.FriendGeneral{
			BusinessPartner:     *controller.UserInfo.BusinessPartner,
			Friend:              friend,
			FriendIsBlocked:     friendIsBlocked,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData MessageInteractionWithFriend

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
				MessageHeaderFromMeToFriend,
				MessageHeaderFromFriendToMe,
				BusinessPartnerPersonMe,
				BusinessPartnerPersonFriend,
				BusinessPartnerAddressFriend,
				FriendGeneral,
			)
		}()
	} else {
		controller.request(
			MessageHeaderFromMeToFriend,
			MessageHeaderFromFriendToMe,
			BusinessPartnerPersonMe,
			BusinessPartnerPersonFriend,
			BusinessPartnerAddressFriend,
			FriendGeneral,
		)
	}
}

func (
	controller *MessageInteractionWithFriendController,
) createMessageRequestHeadersBySenderReceiverFromMeToFriend(
	requestPram *apiInputReader.Request,
	inputMessageHeaderFromMeToFriend apiInputReader.Message,
) *apiModuleRuntimesResponsesMessage.MessageRes {
	responseJsonData := apiModuleRuntimesResponsesMessage.MessageRes{}
	responseBody := apiModuleRuntimesRequestsMessage.MessageReadsHeader(
		requestPram,
		inputMessageHeaderFromMeToFriend,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"センダレシーバ(MeToFriend)に対してのメッセージヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createMessageRequestHeadersBySenderReceiverFromMeToFriend Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createMessageRequestHeadersBySenderReceiverFromFriendToMe(
	requestPram *apiInputReader.Request,
	inputMessageHeaderFromFriendToMe apiInputReader.Message,
) *apiModuleRuntimesResponsesMessage.MessageRes {
	responseJsonData := apiModuleRuntimesResponsesMessage.MessageRes{}
	responseBody := apiModuleRuntimesRequestsMessage.MessageReadsHeader(
		requestPram,
		inputMessageHeaderFromFriendToMe,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"センダレシーバ(FriendToMe)に対してのメッセージヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createMessageRequestHeadersBySenderReceiverFromFriendToMe Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createBusinessPartnerDocRequestMe(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	var input apiInputReader.BusinessPartner

	docType := "IMAGE"

	input = apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: *controller.UserInfo.BusinessPartner,
			DocType:         &docType,
		},
	}

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
		controller.CustomLogger.Error("createBusinessPartnerDocRequestMe Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createBusinessPartnerRequestPersonMe(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerPersonMe apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputBusinessPartnerPersonMe.BusinessPartnerPerson.BusinessPartner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonMe Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createBusinessPartnerDocRequestFriend(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerPersonFriend apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputBusinessPartnerPersonFriend.BusinessPartnerPerson.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPerson(
		requestPram,
		input,
		&controller.Controller,
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
		controller.CustomLogger.Error("createBusinessPartnerDocRequestFriend Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createBusinessPartnerRequestPersonFriend(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerPersonFriend apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputBusinessPartnerPersonFriend.BusinessPartnerPerson.BusinessPartner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonFriend Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createBusinessPartnerRequestAddressFriend(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerAddressFriend apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Address{
		BusinessPartner: inputBusinessPartnerAddressFriend.BusinessPartnerAddress.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsAddresses(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestAddressFriend Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) createFriendRequestGeneral(
	requestPram *apiInputReader.Request,
	inputFriendGeneral apiInputReader.Friend,
) *apiModuleRuntimesResponsesFriend.FriendRes {
	responseJsonData := apiModuleRuntimesResponsesFriend.FriendRes{}
	responseBody := apiModuleRuntimesRequestsFriend.FriendReads(
		requestPram,
		inputFriendGeneral,
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
		controller.CustomLogger.Error("createFriendRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageInteractionWithFriendController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressFriendRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*businessPartnerAddressFriendRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*businessPartnerAddressFriendRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressFriendRes.Message.Address)[0].Country

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
	controller *MessageInteractionWithFriendController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressFriendRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*businessPartnerAddressFriendRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressFriendRes.Message.Address)[0].Country

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
	controller *MessageInteractionWithFriendController,
) request(
	inputMessageHeaderFromMeToFriend apiInputReader.Message,
	inputMessageHeaderFromFriendToMe apiInputReader.Message,
	inputBusinessPartnerPersonMe apiInputReader.BusinessPartner,
	inputBusinessPartnerPersonFriend apiInputReader.BusinessPartner,
	inputBusinessPartnerAddressFriend apiInputReader.BusinessPartner,
	inputFriendGeneral apiInputReader.Friend,

) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headersBySenderReceiverResFromMeToFriend := *controller.createMessageRequestHeadersBySenderReceiverFromMeToFriend(
		controller.UserInfo,
		inputMessageHeaderFromMeToFriend,
	)

	headersBySenderReceiverResFromFriendToMe := *controller.createMessageRequestHeadersBySenderReceiverFromFriendToMe(
		controller.UserInfo,
		inputMessageHeaderFromFriendToMe,
	)

	businessPartnerGeneralDocResMe := controller.createBusinessPartnerDocRequestMe(
		controller.UserInfo,
	)

	businessPartnerPersonResMe := controller.createBusinessPartnerRequestPersonMe(
		controller.UserInfo,
		inputBusinessPartnerPersonMe,
	)

	businessPartnerGeneralDocResFriend := controller.createBusinessPartnerDocRequestFriend(
		controller.UserInfo,
		inputBusinessPartnerPersonFriend,
	)

	businessPartnerPersonResFriend := controller.createBusinessPartnerRequestPersonFriend(
		controller.UserInfo,
		inputBusinessPartnerPersonFriend,
	)

	businessPartnerAddressResFriend := controller.createBusinessPartnerRequestAddressFriend(
		controller.UserInfo,
		inputBusinessPartnerAddressFriend,
	)

	friendGeneralRes := controller.createFriendRequestGeneral(
		controller.UserInfo,
		inputFriendGeneral,
	)

	friendLocalSubRegionTextRes := controller.CreateLocalSubRegionRequestText(
		controller.UserInfo,
		businessPartnerAddressResFriend,
	)

	friendLocalRegionTextRes := controller.CreateLocalRegionRequestText(
		controller.UserInfo,
		businessPartnerAddressResFriend,
	)

	controller.fin(
		&headersBySenderReceiverResFromMeToFriend,
		&headersBySenderReceiverResFromFriendToMe,
		businessPartnerGeneralDocResMe,
		businessPartnerPersonResMe,
		businessPartnerGeneralDocResFriend,
		businessPartnerPersonResFriend,
		businessPartnerAddressResFriend,
		friendGeneralRes,
		friendLocalSubRegionTextRes,
		friendLocalRegionTextRes,
	)
}

func (
	controller *MessageInteractionWithFriendController,
) fin(
	headersBySenderReceiverResFromMeToFriend *apiModuleRuntimesResponsesMessage.MessageRes,
	headersBySenderReceiverResFromFriendToMe *apiModuleRuntimesResponsesMessage.MessageRes,
	businessPartnerGeneralDocResMe *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartnerPersonResMe *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerGeneralDocResFriend *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartnerPersonResFriend *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerAddressResFriend *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	friendGeneralRes *apiModuleRuntimesResponsesFriend.FriendRes,
	friendLocalSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	friendLocalRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		friendLocalSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		friendLocalRegionTextRes.Message.Text,
	)

	data := MessageInteractionWithFriend{}

	for _, v := range *headersBySenderReceiverResFromMeToFriend.Message.Header {
		data.MessageHeaderFromMeToFriend = append(data.MessageHeaderFromMeToFriend,
			apiOutputFormatter.MessageHeader{
				Message:        v.Message,
				LongText:       v.LongText,
				MessageIsRead:  v.MessageIsRead,
				CreationDate:   v.CreationDate,
				CreationTime:   v.CreationTime,
				LastChangeDate: v.LastChangeDate,
				LastChangeTime: v.LastChangeTime,
				IsCancelled:    v.IsCancelled,
			},
		)
	}

	for _, v := range *headersBySenderReceiverResFromFriendToMe.Message.Header {
		data.MessageHeaderFromFriendToMe = append(data.MessageHeaderFromFriendToMe,
			apiOutputFormatter.MessageHeader{
				Message:        v.Message,
				LongText:       v.LongText,
				MessageIsRead:  v.MessageIsRead,
				CreationDate:   v.CreationDate,
				CreationTime:   v.CreationTime,
				LastChangeDate: v.LastChangeDate,
				LastChangeTime: v.LastChangeTime,
				IsCancelled:    v.IsCancelled,
			},
		)
	}

	for _, v := range *businessPartnerGeneralDocResMe.Message.GeneralDoc {

		img := services.ReadBusinessPartnerImage(
			businessPartnerGeneralDocResMe,
			v.BusinessPartner,
		)

		data.BusinessPartnerGeneralDoc = append(data.BusinessPartnerGeneralDoc,
			apiOutputFormatter.BusinessPartnerGeneralDoc{
				BusinessPartner: v.BusinessPartner,
				Images: apiOutputFormatter.Images{
					BusinessPartner: img,
				},
			},
		)
	}

	for _, v := range *businessPartnerPersonResMe.Message.Person {
		data.BusinessPartnerPersonMe = append(data.BusinessPartnerPersonMe,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
				ProfileComment:  v.ProfileComment,
			},
		)
	}

	for _, v := range *businessPartnerGeneralDocResFriend.Message.GeneralDoc {

		img := services.ReadBusinessPartnerImage(
			businessPartnerGeneralDocResFriend,
			v.BusinessPartner,
		)

		data.BusinessPartnerGeneralDoc = append(data.BusinessPartnerGeneralDoc,
			apiOutputFormatter.BusinessPartnerGeneralDoc{
				BusinessPartner: v.BusinessPartner,
				Images: apiOutputFormatter.Images{
					BusinessPartner: img,
				},
			},
		)
	}

	for _, v := range *businessPartnerPersonResFriend.Message.Person {
		data.BusinessPartnerPersonFriend = append(data.BusinessPartnerPersonFriend,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
				ProfileComment:  v.ProfileComment,
			},
		)
	}

	for _, v := range *businessPartnerAddressResFriend.Message.Address {
		data.BusinessPartnerAddressFriend = append(data.BusinessPartnerAddressFriend,
			apiOutputFormatter.BusinessPartnerAddress{
				BusinessPartner:    v.BusinessPartner,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
			},
		)
	}

	for _, v := range *friendGeneralRes.Message.General {
		data.FriendGeneral = append(data.FriendGeneral,
			apiOutputFormatter.FriendGeneral{
				BusinessPartner: v.BusinessPartner,
				Friend:          v.Friend,
				RankType:        v.RankType,
				Rank:            v.Rank,
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
