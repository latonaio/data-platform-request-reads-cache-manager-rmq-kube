package controllersPostListByFriend

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	apiModuleRuntimesRequestsFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/post/post"
	apiModuleRuntimesRequestsPostDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/post/post-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesFriend "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/friend"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/post"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type PostListByFriendController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PostListByFriend struct {
	BusinessPartnerPerson        []apiOutputFormatter.BusinessPartnerPerson        `json:"BusinessPartnerPerson"`
	BusinessPartnerAddress       []apiOutputFormatter.BusinessPartnerAddress       `json:"BusinessPartnerAddress"`
	BusinessPartnerSNS           []apiOutputFormatter.BusinessPartnerSNS           `json:"BusinessPartnerSNS"`
	FriendGeneral                []apiOutputFormatter.FriendGeneral                `json:"FriendGeneral"`
	PostHeader                   []apiOutputFormatter.PostHeader                   `json:"PostHeader"`
	PostHeaderWithInstagramMedia []apiOutputFormatter.PostHeaderWithInstagramMedia `json:"PostHeaderWithInstagramMedia"`
}

func (controller *PostListByFriendController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	friend, _ := controller.GetInt("friend")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "post"
	redisKeyCategory2 := "listByFriend"
	redisKeyCategory3 := friend

	PostListByFriendPosts := apiInputReader.Post{}
	BusinessPartnerPerson := apiInputReader.BusinessPartner{}
	BusinessPartnerAddress := apiInputReader.BusinessPartner{}
	BusinessPartnerDocGeneralDoc := apiInputReader.BusinessPartner{}
	BusinessPartnerSNS := apiInputReader.BusinessPartner{}
	FriendGeneral := apiInputReader.Friend{}

	isMarkedForDeletion := false
	isPublished := true
	friendIsBlocked := false

	PostListByFriendPosts = apiInputReader.Post{
		PostHeader: &apiInputReader.PostHeader{
			PostOwner:           &friend,
			IsPublished:         &isPublished,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerDocGeneralDoc = apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: friend,
		},
	}

	BusinessPartnerPerson = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     friend,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerAddress = apiInputReader.BusinessPartner{
		BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
			BusinessPartner: friend,
		},
	}

	BusinessPartnerSNS = apiInputReader.BusinessPartner{
		BusinessPartnerSNS: &apiInputReader.BusinessPartnerSNS{
			BusinessPartner:     friend,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
			strconv.Itoa(redisKeyCategory3),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData PostListByFriend

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
			controller.request(PostListByFriendPosts, BusinessPartnerDocGeneralDoc, BusinessPartnerPerson, BusinessPartnerAddress, BusinessPartnerSNS, FriendGeneral, isMarkedForDeletion)
		}()
	} else {
		controller.request(PostListByFriendPosts, BusinessPartnerDocGeneralDoc, BusinessPartnerPerson, BusinessPartnerAddress, BusinessPartnerSNS, FriendGeneral, isMarkedForDeletion)
	}
}

func (
	controller *PostListByFriendController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerPerson apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Person

	input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputBusinessPartnerPerson.BusinessPartnerPerson.BusinessPartner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPerson Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) createBusinessPartnerRequestAddresses(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerAddress apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Address

	input = apiModuleRuntimesRequestsBusinessPartner.Address{
		BusinessPartner: inputBusinessPartnerAddress.BusinessPartnerAddress.BusinessPartner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) createBusinessPartnerRequestSNS(
	requestPram *apiInputReader.Request,
	inputBusinessPartnerSNS apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.SNS

	input = apiModuleRuntimesRequestsBusinessPartner.SNS{
		BusinessPartner: inputBusinessPartnerSNS.BusinessPartnerSNS.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsSNS(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestSNS Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) createFriendRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.Friend,
) *apiModuleRuntimesResponsesFriend.FriendRes {
	responseJsonData := apiModuleRuntimesResponsesFriend.FriendRes{}
	responseBody := apiModuleRuntimesRequestsFriend.FriendReads(
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
		controller.CustomLogger.Error("createFriendRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
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
	controller *PostListByFriendController,
) createPostRequestHeadersByPostOwner(
	requestPram *apiInputReader.Request,
	input apiInputReader.Post,
) *apiModuleRuntimesResponsesPost.PostRes {
	responseJsonData := apiModuleRuntimesResponsesPost.PostRes{}
	responseBody := apiModuleRuntimesRequestsPost.PostReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByPostOwner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if len(*responseJsonData.Message.Header) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"フレンドの投稿が見つかりませんでした",
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
		controller.CustomLogger.Error("createPostRequestHeadersByPostOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) createPostRequestInstagramMediasByPosts(
	requestPram *apiInputReader.Request,
	postRes *apiModuleRuntimesResponsesPost.PostRes,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesPost.PostRes {
	var input []apiModuleRuntimesRequestsPost.Header

	for _, v := range *postRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsPost.Header{
			Post:                v.Post,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesPost.PostRes{}
	responseBody := apiModuleRuntimesRequestsPost.PostReadsInstagramMediasByPosts(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.InstagramMedia == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"投稿に対してInstagramメディアデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createPostRequestInstagramMediasByPosts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) createPostDocRequest(
	requestPram *apiInputReader.Request,
	headerRes apiModuleRuntimesResponsesPost.PostRes,
) *apiModuleRuntimesResponsesPost.PostDocRes {
	var input apiInputReader.Post = apiInputReader.Post{
		PostDocHeaderDoc: &apiInputReader.PostDocHeaderDoc{},
	}

	//for _, v := range *headerRes.Message.Header {
	//	input = apiInputReader.Post{
	//		PostDocHeaderDoc: &apiInputReader.PostDocHeaderDoc{
	//			Post: v.Post,
	//		},
	//	}
	//}

	responseJsonData := apiModuleRuntimesResponsesPost.PostDocRes{}
	responseBody := apiModuleRuntimesRequestsPostDoc.PostDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.HeaderDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"投稿ヘッダに画像または動画が見つかりませんでした",
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
		controller.CustomLogger.Error("createPostDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressRes.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressRes.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostListByFriendController,
) request(
	postListByFriendPosts apiInputReader.Post,
	businessPartnerDocGeneralDoc apiInputReader.BusinessPartner,
	businessPartnerPerson apiInputReader.BusinessPartner,
	businessPartnerAddress apiInputReader.BusinessPartner,
	businessPartnerSNS apiInputReader.BusinessPartner,
	friendGeneral apiInputReader.Friend,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(7)

	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerSNSRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var friendGeneralRes apiModuleRuntimesResponsesFriend.FriendRes
	var businessPartnerGeneralDocRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes

	var headerDocRes apiModuleRuntimesResponsesPost.PostDocRes
	var instagramMediasRes apiModuleRuntimesResponsesPost.PostRes

	var businessPartnerAddressRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	headersRes := *controller.createPostRequestHeadersByPostOwner(
		controller.UserInfo,
		postListByFriendPosts,
	)

	go func() {
		defer wg.Done()
		businessPartnerPersonRes = *controller.createBusinessPartnerRequestPerson(
			controller.UserInfo,
			businessPartnerPerson,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerSNSRes = *controller.createBusinessPartnerRequestSNS(
			controller.UserInfo,
			businessPartnerSNS,
		)
	}()

	go func() {
		defer wg.Done()
		friendGeneralRes = *controller.createFriendRequestGeneral(
			controller.UserInfo,
			friendGeneral,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerGeneralDocRes = *controller.createBusinessPartnerDocRequest(
			businessPartnerDocGeneralDoc,
		)
	}()

	go func() {
		defer wg.Done()
		if headersRes.Message.Header != nil && len(*headersRes.Message.Header) != 0 {
			headerDocRes = *controller.createPostDocRequest(
				controller.UserInfo,
				headersRes,
			)
		}
	}()

	go func() {
		defer wg.Done()
		if headersRes.Message.Header != nil && len(*headersRes.Message.Header) != 0 {
			instagramMediasRes = *controller.createPostRequestInstagramMediasByPosts(
				controller.UserInfo,
				&headersRes,
				isMarkedForDeletion,
			)
		}
	}()

	go func() {
		defer wg.Done()
		businessPartnerAddressRes = *controller.createBusinessPartnerRequestAddresses(
			controller.UserInfo,
			businessPartnerAddress,
		)
		localSubRegionTextRes = controller.CreateLocalSubRegionRequestText(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)
		localRegionTextRes = controller.CreateLocalRegionRequestText(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)
	}()

	wg.Wait()

	controller.fin(
		&headersRes,
		&businessPartnerPersonRes,
		&businessPartnerSNSRes,
		&friendGeneralRes,
		&businessPartnerGeneralDocRes,
		&headerDocRes,
		&instagramMediasRes,
		&businessPartnerAddressRes,
		localSubRegionTextRes,
		localRegionTextRes,
	)
}

func (
	controller *PostListByFriendController,
) fin(
	headerRes *apiModuleRuntimesResponsesPost.PostRes,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerSNSRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	friendGeneralRes *apiModuleRuntimesResponsesFriend.FriendRes,
	businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	headerDocRes *apiModuleRuntimesResponsesPost.PostDocRes,
	instagramMediasRes *apiModuleRuntimesResponsesPost.PostRes,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {
	var postHeadersMapper map[string]apiModuleRuntimesResponsesPost.Header

	if headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {
		postHeadersMapper = services.PostHeadersMapper(
			headerRes,
		)
	}

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	data := PostListByFriend{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadPostImage(
			headerDocRes,
			v.Post,
		)
		data.PostHeader = append(data.PostHeader,
			apiOutputFormatter.PostHeader{
				Post:         v.Post,
				PostOwner:    v.PostOwner,
				Description:  v.Description,
				LongText:     v.LongText,
				Tag1:         v.Tag1,
				Tag2:         v.Tag2,
				Tag3:         v.Tag3,
				Tag4:         v.Tag4,
				IsPublished:  v.IsPublished,
				CreationDate: v.CreationDate,
				CreationTime: v.CreationTime,
				Images: apiOutputFormatter.Images{
					Post: img,
				},
			},
		)
	}

	if *headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {
		for _, v := range *instagramMediasRes.Message.InstagramMedia {
			postOwner := postHeadersMapper[strconv.Itoa(v.Post)].PostOwner
			description := postHeadersMapper[strconv.Itoa(v.Post)].Description
			longText := postHeadersMapper[strconv.Itoa(v.Post)].LongText
			tag1 := postHeadersMapper[strconv.Itoa(v.Post)].Tag1
			tag2 := postHeadersMapper[strconv.Itoa(v.Post)].Tag2
			tag3 := postHeadersMapper[strconv.Itoa(v.Post)].Tag3
			tag4 := postHeadersMapper[strconv.Itoa(v.Post)].Tag4
			isPublished := postHeadersMapper[strconv.Itoa(v.Post)].IsPublished
			creationDate := postHeadersMapper[strconv.Itoa(v.Post)].CreationDate
			creationTime := postHeadersMapper[strconv.Itoa(v.Post)].CreationTime
			img := services.ReadPostImage(
				headerDocRes,
				v.Post,
			)

			data.PostHeaderWithInstagramMedia = append(data.PostHeaderWithInstagramMedia,
				apiOutputFormatter.PostHeaderWithInstagramMedia{
					Post:               v.Post,
					PostOwner:          postOwner,
					Description:        description,
					LongText:           longText,
					Tag1:               tag1,
					Tag2:               tag2,
					Tag3:               tag3,
					Tag4:               tag4,
					IsPublished:        isPublished,
					CreationDate:       creationDate,
					CreationTime:       creationTime,
					InstagramMediaID:   v.InstagramMediaID,
					InstagramMediaType: v.InstagramMediaType,
					InstagramMediaURL:  v.InstagramMediaURL,
					InstagramMediaDate: v.InstagramMediaDate,
					InstagramMediaTime: v.InstagramMediaTime,
					InstagramUserName:  v.InstagramUserName,
					Images: apiOutputFormatter.Images{
						Post: img,
					},
				},
			)
		}
	}

	for _, v := range *businessPartnerPersonRes.Message.Person {

		img := services.ReadBusinessPartnerImage(
			businessPartnerGeneralDocRes,
			v.BusinessPartner,
		)

		qrcode := services.CreateQRCodeBusinessPartnerDocImage(
			businessPartnerGeneralDocRes,
			v.BusinessPartner,
		)

		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
				ProfileComment:  v.ProfileComment,
				Images: apiOutputFormatter.Images{
					BusinessPartner: img,
					QRCode:          qrcode,
				},
			},
		)
	}

	for _, v := range *businessPartnerSNSRes.Message.SNS {
		data.BusinessPartnerSNS = append(data.BusinessPartnerSNS,
			apiOutputFormatter.BusinessPartnerSNS{
				BusinessPartner: v.BusinessPartner,
				InstagramURL:    v.InstagramURL,
				TikTokURL:       v.TikTokURL,
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

	for _, v := range *businessPartnerAddressRes.Message.Address {
		data.BusinessPartnerAddress = append(data.BusinessPartnerAddress,
			apiOutputFormatter.BusinessPartnerAddress{
				BusinessPartner:    v.BusinessPartner,
				AddressID:          v.AddressID,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
				Country:            &v.Country,
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
