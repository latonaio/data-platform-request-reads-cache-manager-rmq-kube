package controllersPostListMe

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsPost "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/post/post"
	apiModuleRuntimesRequestsPostDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/post/post-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
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

type PostListMeController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PostListMe struct {
	BusinessPartnerPerson        []apiOutputFormatter.BusinessPartnerPerson        `json:"BusinessPartnerPerson"`
	BusinessPartnerAddress       []apiOutputFormatter.BusinessPartnerAddress       `json:"BusinessPartnerAddress"`
	BusinessPartnerSNS           []apiOutputFormatter.BusinessPartnerSNS           `json:"BusinessPartnerSNS"`
	PostHeader                   []apiOutputFormatter.PostHeader                   `json:"PostHeader"`
	PostHeaderWithInstagramMedia []apiOutputFormatter.PostHeaderWithInstagramMedia `json:"PostHeaderWithInstagramMedia"`
}

func (controller *PostListMeController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	businessPartner, _ := controller.GetInt("businessPartner")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "post"
	redisKeyCategory2 := "listMe"
	redisKeyCategory3 := businessPartner

	PostListMePosts := apiInputReader.Post{}
	BusinessPartnerPerson := apiInputReader.BusinessPartner{}
	BusinessPartnerAddress := apiInputReader.BusinessPartner{}
	BusinessPartnerDocGeneralDoc := apiInputReader.BusinessPartner{}
	BusinessPartnerSNS := apiInputReader.BusinessPartner{}

	isMarkedForDeletion := false

	PostListMePosts = apiInputReader.Post{
		PostHeader: &apiInputReader.PostHeader{
			PostOwner: &businessPartner,
			//			IsPublished:         nil,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerDocGeneralDoc = apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: businessPartner,
		},
	}

	BusinessPartnerPerson = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	BusinessPartnerAddress = apiInputReader.BusinessPartner{
		BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
			BusinessPartner: businessPartner,
		},
	}

	BusinessPartnerSNS = apiInputReader.BusinessPartner{
		BusinessPartnerSNS: &apiInputReader.BusinessPartnerSNS{
			BusinessPartner:     businessPartner,
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
		var responseData PostListMe

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
			controller.request(PostListMePosts, BusinessPartnerDocGeneralDoc, BusinessPartnerPerson, BusinessPartnerAddress, BusinessPartnerSNS, isMarkedForDeletion)
		}()
	} else {
		controller.request(PostListMePosts, BusinessPartnerDocGeneralDoc, BusinessPartnerPerson, BusinessPartnerAddress, BusinessPartnerSNS, isMarkedForDeletion)
	}
}

func (
	controller *PostListMeController,
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
	controller *PostListMeController,
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
	controller *PostListMeController,
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
	controller *PostListMeController,
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
			"あなたのビジネスパートナヘッダに画像が見つかりませんでした",
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
	controller *PostListMeController,
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
	//		"あなたの投稿が見つかりませんでした",
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
	controller *PostListMeController,
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
	controller *PostListMeController,
) createPostDocRequest(
	requestPram *apiInputReader.Request,
	headerRes apiModuleRuntimesResponsesPost.PostRes,
) *apiModuleRuntimesResponsesPost.PostDocRes {
	var input apiInputReader.Post = apiInputReader.Post{
		PostDocHeaderDoc: &apiInputReader.PostDocHeaderDoc{},
	}

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
	controller *PostListMeController,
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
	controller *PostListMeController,
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
	controller *PostListMeController,
) request(
	postListMe apiInputReader.Post,
	businessPartnerDocGeneralDoc apiInputReader.BusinessPartner,
	businessPartnerPerson apiInputReader.BusinessPartner,
	businessPartnerAddress apiInputReader.BusinessPartner,
	businessPartnerSNS apiInputReader.BusinessPartner,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(6)

	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerSNSRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerGeneralDocRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes

	var headerDocRes apiModuleRuntimesResponsesPost.PostDocRes
	var instagramMediasRes apiModuleRuntimesResponsesPost.PostRes

	var businessPartnerAddressRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	headerRes := *controller.createPostRequestHeadersByPostOwner(
		controller.UserInfo,
		postListMe,
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
		businessPartnerGeneralDocRes = *controller.createBusinessPartnerDocRequest(
			businessPartnerDocGeneralDoc,
		)
	}()

	go func() {
		defer wg.Done()
		if headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {
			headerDocRes = *controller.createPostDocRequest(
				controller.UserInfo,
				headerRes,
			)
		}
	}()

	go func() {
		defer wg.Done()
		if headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {
			instagramMediasRes = *controller.createPostRequestInstagramMediasByPosts(
				controller.UserInfo,
				&headerRes,
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
		&headerRes,
		&businessPartnerPersonRes,
		&businessPartnerSNSRes,
		&businessPartnerGeneralDocRes,
		&headerDocRes,
		&instagramMediasRes,
		&businessPartnerAddressRes,
		localSubRegionTextRes,
		localRegionTextRes,
	)
}

func (
	controller *PostListMeController,
) fin(
	headerRes *apiModuleRuntimesResponsesPost.PostRes,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerSNSRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
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

	data := PostListMe{}

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
				XURL:            v.XURL,
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
