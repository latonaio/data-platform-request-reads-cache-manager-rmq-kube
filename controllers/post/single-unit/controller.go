package controllersPostSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	//	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	//	apiModuleRuntimesRequestsFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiModuleRuntimesRequestsPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/post/post"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	//	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	//	apiModuleRuntimesResponsesFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/friend"
	apiModuleRuntimesResponsesPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/post"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type PostSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PostSingleUnit struct {
	PostHeader         []apiOutputFormatter.PostHeader         `json:"PostHeader"`
	PostInstagramMedia []apiOutputFormatter.PostInstagramMedia `json:"PostInstagramMedia"`
	SiteHeader         []apiOutputFormatter.SiteHeader         `json:"SiteHeader"`
	//	BusinessPartnerPerson	[]apiOutputFormatter.BusinessPartnerPerson	`json:"BusinessPartnerPerson"`
}

func (controller *PostSingleUnitController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	post, _ := controller.GetInt("post")
	friend, _ := controller.GetInt("friend")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "post"
	redisKeyCategory2 := "friend"
	redisKeyCategory3 := friend
	redisKeyCategory4 := post

	PostSingleUnitPost := apiInputReader.Post{}
	BusinessPartnerDocGeneralDoc := apiInputReader.BusinessPartner{}

	//	BusinessPartnerPerson := apiInputReader.BusinessPartner{}
	//	FriendGeneral := apiInputReader.Friend{}

	isMarkedForDeletion := false
	//	friendIsBlocked := false

	PostSingleUnitPost = apiInputReader.Post{
		PostHeader: &apiInputReader.PostHeader{
			Post: post,
			//			IsPublished:         nil,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		PostInstagramMedia: &apiInputReader.PostInstagramMedia{
			Post:                post,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerDocGeneralDoc = apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: friend,
		},
	}

	//	BusinessPartnerPerson = apiInputReader.BusinessPartner{
	//		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
	//			BusinessPartner:     friend,
	//			IsMarkedForDeletion: &isMarkedForDeletion,
	//		},
	//	}

	//	FriendGeneral = apiInputReader.Friend{
	//		FriendGeneral: &apiInputReader.FriendGeneral{
	//			BusinessPartner:     *controller.UserInfo.BusinessPartner,
	//			Friend:              friend,
	//			FriendIsBlocked:     friendIsBlocked,
	//			IsMarkedForDeletion: &isMarkedForDeletion,
	//		},
	//	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
			strconv.Itoa(redisKeyCategory4),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData PostSingleUnit

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
			controller.request(PostSingleUnitPost, BusinessPartnerDocGeneralDoc)
		}()
	} else {
		controller.request(PostSingleUnitPost, BusinessPartnerDocGeneralDoc)
	}
}

//func (
//	controller *PostSingleUnitController,
//) createBusinessPartnerRequestPerson(
//	requestPram *apiInputReader.Request,
//	inputBusinessPartnerPerson apiInputReader.BusinessPartner,
//) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
//	var input apiModuleRuntimesRequestsBusinessPartner.Person
//
//	input = apiModuleRuntimesRequestsBusinessPartner.Person{
//		BusinessPartner: inputBusinessPartnerPerson.BusinessPartnerPerson.BusinessPartner,
//	}
//
//	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
//	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPerson(
//		requestPram,
//		input,
//		&controller.Controller,
//	)
//
//	err := json.Unmarshal(responseBody, &responseJsonData)
//	if err != nil {
//		services.HandleError(
//			&controller.Controller,
//			err,
//			nil,
//		)
//		controller.CustomLogger.Error("createBusinessPartnerRequestPerson Unmarshal error")
//	}
//
//	return &responseJsonData
//}

//func (
//	controller *PostSingleUnitController,
//) createFriendRequestGeneral(
//	requestPram *apiInputReader.Request,
//	input apiInputReader.Friend,
//) *apiModuleRuntimesResponsesFriend.FriendRes {
//	responseJsonData := apiModuleRuntimesResponsesFriend.FriendRes{}
//	responseBody := apiModuleRuntimesRequestsFriend.FriendReads(
//		requestPram,
//		input,
//		&controller.Controller,
//		"General",
//	)
//
//	err := json.Unmarshal(responseBody, &responseJsonData)
//	if err != nil {
//		services.HandleError(
//			&controller.Controller,
//			err,
//			nil,
//		)
//		controller.CustomLogger.Error("createFriendRequestGeneral Unmarshal error")
//	}
//
//	return &responseJsonData
//}

func (
	controller *PostSingleUnitController,
) createBusinessPartnerDocRequest(
	businessPartnerDocGeneralDoc apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerDoc.BusinessPartnerDocReads(
		nil,
		businessPartnerDocGeneralDoc,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.GeneralDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"フレンドのビジネスパートナヘッダに画像が見つかりませんでした",
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
	controller *PostSingleUnitController,
) createPostRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Post,
) *apiModuleRuntimesResponsesPost.PostRes {
	responseJsonData := apiModuleRuntimesResponsesPost.PostRes{}
	responseBody := apiModuleRuntimesRequestsPost.PostReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"投稿が見つかりませんでした",
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
		controller.CustomLogger.Error("createPostRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostSingleUnitController,
) createPostRequestInstagramMedia(
	requestPram *apiInputReader.Request,
	input apiInputReader.Post,
) *apiModuleRuntimesResponsesPost.PostRes {
	responseJsonData := apiModuleRuntimesResponsesPost.PostRes{}
	responseBody := apiModuleRuntimesRequestsPost.PostReads(
		requestPram,
		input,
		&controller.Controller,
		"InstagramMedia",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createPostRequestInstagramMedia Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostSingleUnitController,
) createSiteRequestHeader(
	requestPram *apiInputReader.Request,
	postRes *apiModuleRuntimesResponsesPost.PostRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	input := apiInputReader.Site{}

	for _, v := range *postRes.Message.Header {
		input = apiInputReader.Site{
			SiteHeader: &apiInputReader.SiteHeader{
				Site: *v.Site,
			},
		}
	}

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
	controller *PostSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	postRes *apiModuleRuntimesResponsesPost.PostRes,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
	input := apiInputReader.Site{}

	for _, v := range *postRes.Message.Header {
		input = apiInputReader.Site{
			SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{
				Site: *v.Site,
			},
		}
	}

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
	controller *PostSingleUnitController,
) request(
	postSingleUnit apiInputReader.Post,
	businessPartnerDocGeneralDoc apiInputReader.BusinessPartner,
	// businessPartnerPerson apiInputReader.BusinessPartner,
	// friendGeneral apiInputReader.Friend,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(3)

	var headerRes apiModuleRuntimesResponsesPost.PostRes
	var instagramMediaRes apiModuleRuntimesResponsesPost.PostRes
	var businessPartnerGeneralDocRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes

	var siteHeaderRes apiModuleRuntimesResponsesSite.SiteRes
	var siteDocRes apiModuleRuntimesResponsesSite.SiteDocRes

	//	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	//	var friendGeneralRes apiModuleRuntimesResponsesFriend.FriendRes

	go func() {
		defer wg.Done()
		headerRes = *controller.createPostRequestHeader(
			controller.UserInfo,
			postSingleUnit,
		)
		siteHeaderRes = *controller.createSiteRequestHeader(
			controller.UserInfo,
			&headerRes,
		)
		siteDocRes = *controller.createSiteDocRequest(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		instagramMediaRes = *controller.createPostRequestInstagramMedia(
			controller.UserInfo,
			postSingleUnit,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerGeneralDocRes = *controller.createBusinessPartnerDocRequest(
			businessPartnerDocGeneralDoc,
		)
	}()

	//	go func() {
	//		defer wg.Done()
	//		businessPartnerPersonRes = *controller.createBusinessPartnerRequestPerson(
	//			controller.UserInfo,
	//			businessPartnerPerson,
	//		)
	//	}()

	//	go func() {
	//		defer wg.Done()
	//		friendGeneralRes = *controller.createFriendRequestGeneral(
	//			controller.UserInfo,
	//			friendGeneral,
	//		)
	//	}()

	wg.Wait()

	controller.fin(
		&headerRes,
		&siteHeaderRes,
		&siteDocRes,
		&instagramMediaRes,
		&businessPartnerGeneralDocRes,
		//		&businessPartnerPersonRes,
		//		&friendGeneralRes,
	)
}

func (
	controller *PostSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesPost.PostRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	instagramMediaRes *apiModuleRuntimesResponsesPost.PostRes,
	businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	// businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	// friendGeneralRes *apiModuleRuntimesResponsesFriend.FriendRes,
) {

	data := PostSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		data.PostHeader = append(data.PostHeader,
			apiOutputFormatter.PostHeader{
				Post:         v.Post,
				PostOwner:    v.PostOwner,
				Description:  v.Description,
				LongText:     v.LongText,
				Site:         v.Site,
				Tag1:         v.Tag1,
				Tag2:         v.Tag2,
				Tag3:         v.Tag3,
				Tag4:         v.Tag4,
				IsPublished:  v.IsPublished,
				CreationDate: v.CreationDate,
				CreationTime: v.CreationTime,
			},
		)
	}

	for _, v := range *instagramMediaRes.Message.InstagramMedia {
		data.PostInstagramMedia = append(data.PostInstagramMedia,
			apiOutputFormatter.PostInstagramMedia{
				Post:                    v.Post,
				InstagramMediaID:        v.InstagramMediaID,
				InstagramMediaPermaLink: v.InstagramMediaPermaLink,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {
		img := services.ReadSiteImage(
			siteDocRes,
			v.Site,
		)

		documentImage := services.ReadDocumentImageSite(
			siteDocRes,
			v.Site,
		)

		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				URL:         v.URL,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Site:              img,
					DocumentImageSite: documentImage,
				},
			},
		)
	}

	//	for _, v := range *businessPartnerPersonRes.Message.Person {
	//
	//		img := services.ReadBusinessPartnerImage(
	//			businessPartnerGeneralDocRes,
	//			v.BusinessPartner,
	//		)
	//
	//		qrcode := services.CreateQRCodeBusinessPartnerDocImage(
	//			businessPartnerGeneralDocRes,
	//			v.BusinessPartner,
	//		)
	//
	//		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
	//			apiOutputFormatter.BusinessPartnerPerson{
	//				BusinessPartner: v.BusinessPartner,
	//				NickName:        v.NickName,
	//				ProfileComment:  v.ProfileComment,
	//				Images: apiOutputFormatter.Images{
	//					BusinessPartner: img,
	//					QRCode:          qrcode,
	//				},
	//			},
	//		)
	//	}

	//	for _, v := range *friendGeneralRes.Message.General {
	//		data.FriendGeneral = append(data.FriendGeneral,
	//			apiOutputFormatter.FriendGeneral{
	//				BusinessPartner: v.BusinessPartner,
	//				Friend:          v.Friend,
	//				RankType:        v.RankType,
	//				Rank:            v.Rank,
	//			},
	//		)
	//	}

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
