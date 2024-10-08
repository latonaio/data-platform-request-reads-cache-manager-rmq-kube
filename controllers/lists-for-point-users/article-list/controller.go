package controllersArticleListForPointUsers

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsArticle "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article/article"
	apiModuleRuntimesRequestsArticleDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article/article-doc"
	apiModuleRuntimesResponsesArticle "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ArticleListForPointUsersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ArticleListForPointUsersController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	localSubRegion := controller.GetString("localSubRegion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "localSubRegion"
	redisKeyCategory2 := "listObjectArticle"

	isReleased := true
	isMarkedForDeletion := false

	ArticleAddress := apiInputReader.Article{
		ArticleAddress: &apiInputReader.ArticleAddress{
			LocalSubRegion: &localSubRegion,
		},
		ArticleDocHeaderDoc: &apiInputReader.ArticleDocHeaderDoc{},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Article

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
			controller.request(ArticleAddress, isReleased, isMarkedForDeletion)
		}()
	} else {
		controller.request(ArticleAddress, isReleased, isMarkedForDeletion)
	}
}

func (
	controller *ArticleListForPointUsersController,
) createArticleRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReads(
		requestPram,
		input,
		&controller.Controller,
		"AddressesByLocalSubRegion",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対しての記事が見つかりませんでした",
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
		controller.CustomLogger.Error("createArticleRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleListForPointUsersController,
) createArticleRequestHeadersByArticles(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	var input []apiModuleRuntimesRequestsArticle.Header

	for _, v := range *articleRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsArticle.Header{
			Article:             v.Article,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReadsHeadersByArticles(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.Header == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対して有効な記事ヘッダデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createArticleRequestHeadersByArticles Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleListForPointUsersController,
) createArticleRequestCountersByArticles(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	var input []apiModuleRuntimesRequestsArticle.Header

	for _, v := range *articleRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsArticle.Header{
			Article: v.Article,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReadsCountersByArticles(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.Counter == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対して有効な記事カウンタデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createArticleRequestCountersByArticles Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleListForPointUsersController,
) createArticleDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleDocRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleDocRes{}
	responseBody := apiModuleRuntimesRequestsArticleDoc.ArticleDocReads(
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
			"記事ヘッダに画像が見つかりませんでした",
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
		controller.CustomLogger.Error("createArticleDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleListForPointUsersController,
) request(
	input apiInputReader.Article,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	addressRes := *controller.createArticleRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		input,
	)

	headerRes := *controller.createArticleRequestHeadersByArticles(
		controller.UserInfo,
		&addressRes,
		isReleased,
		isMarkedForDeletion,
	)

	counterRes := *controller.createArticleRequestCountersByArticles(
		controller.UserInfo,
		&addressRes,
	)

	headerDocRes := controller.createArticleDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&addressRes,
		&headerRes,
		&counterRes,
		headerDocRes,
	)
}

func (
	controller *ArticleListForPointUsersController,
) fin(
	addressRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	headerRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	counterRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	headerDocRes *apiModuleRuntimesResponsesArticle.ArticleDocRes,
) {

	articleHeadersMapper := services.ArticleHeadersMapper(
		headerRes,
	)

	articleCountersMapper := services.ArticleCountersMapper(
		counterRes,
	)

	data := apiOutputFormatter.Article{}

	for _, v := range *addressRes.Message.Address {
		articleType := articleHeadersMapper[strconv.Itoa(v.Article)].ArticleType
		validityStartDate := articleHeadersMapper[strconv.Itoa(v.Article)].ValidityStartDate
		validityStartTime := articleHeadersMapper[strconv.Itoa(v.Article)].ValidityStartTime
		validityEndDate := articleHeadersMapper[strconv.Itoa(v.Article)].ValidityEndDate
		validityEndTime := articleHeadersMapper[strconv.Itoa(v.Article)].ValidityEndTime
		introduction := articleHeadersMapper[strconv.Itoa(v.Article)].Introduction
		tag1 := articleHeadersMapper[strconv.Itoa(v.Article)].Tag1
		tag2 := articleHeadersMapper[strconv.Itoa(v.Article)].Tag2
		tag3 := articleHeadersMapper[strconv.Itoa(v.Article)].Tag3
		tag4 := articleHeadersMapper[strconv.Itoa(v.Article)].Tag4
		lastChangeDate := articleHeadersMapper[strconv.Itoa(v.Article)].LastChangeDate
		lastChangeTime := articleHeadersMapper[strconv.Itoa(v.Article)].LastChangeTime

		numberOfLikes := articleCountersMapper[strconv.Itoa(v.Article)].NumberOfLikes

		img := services.ReadArticleImage(
			headerDocRes,
			v.Article,
		)

		data.ArticleAddressWithHeader = append(data.ArticleAddressWithHeader,
			apiOutputFormatter.ArticleAddressWithHeader{
				Article:           v.Article,
				AddressID:         v.AddressID,
				LocalSubRegion:    v.LocalSubRegion,
				LocalRegion:       v.LocalRegion,
				ArticleType:       articleType,
				ValidityStartDate: validityStartDate,
				ValidityStartTime: validityStartTime,
				ValidityEndDate:   validityEndDate,
				ValidityEndTime:   validityEndTime,
				Introduction:      introduction,
				Tag1:              tag1,
				Tag2:              tag2,
				Tag3:              tag3,
				Tag4:              tag4,
				LastChangeDate:    lastChangeDate,
				LastChangeTime:    lastChangeTime,
				NumberOfLikes:     &numberOfLikes,

				Images: apiOutputFormatter.Images{
					Article: img,
				},
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
