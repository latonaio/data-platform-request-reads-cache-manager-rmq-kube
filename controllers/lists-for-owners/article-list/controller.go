package controllersArticleListForOwners

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

type ArticleListForOwnersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ArticleListForOwnersController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	organizationBusinessPartner, _ := controller.GetInt("organizationBusinessPartner")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "organizationBusinessPartner"
	redisKeyCategory2 := "list-for-owners-article"

	isMarkedForDeletion := false

	Article := apiInputReader.Article{
		ArticleHeader: &apiInputReader.ArticleHeader{
			ArticleOwner:        &organizationBusinessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
			controller.request(Article)
		}()
	} else {
		controller.request(Article)
	}
}

func (
	controller *ArticleListForOwnersController,
) createArticleRequestHeadersByArticleOwner(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByArticleOwner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//	if len(*responseJsonData.Message.Address) == 0 {
	//		status := 500
	//		services.HandleError(
	//			&controller.Controller,
	//			"記事オーナーに対しての記事が見つかりませんでした",
	//			&status,
	//		)
	//		return nil
	//	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createArticleRequestHeadersByArticleOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleListForOwnersController,
) createArticleRequestCountersByArticles(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	var input []apiModuleRuntimesRequestsArticle.Header

	for _, v := range *articleRes.Message.Header {
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
			"記事に対して有効な記事カウンタデータが見つかりませんでした",
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
	controller *ArticleListForOwnersController,
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
	controller *ArticleListForOwnersController,
) request(
	input apiInputReader.Article,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createArticleRequestHeadersByArticleOwner(
		controller.UserInfo,
		input,
	)

	counterRes := *controller.createArticleRequestCountersByArticles(
		controller.UserInfo,
		&headerRes,
	)

	headerDocRes := controller.createArticleDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&headerRes,
		&counterRes,
		headerDocRes,
	)
}

func (
	controller *ArticleListForOwnersController,
) fin(
	headerRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	counterRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	headerDocRes *apiModuleRuntimesResponsesArticle.ArticleDocRes,
) {

	articleCountersMapper := services.ArticleCountersMapper(
		counterRes,
	)

	data := apiOutputFormatter.Article{}

	for _, v := range *headerRes.Message.Header {

		numberOfLikes := articleCountersMapper[strconv.Itoa(v.Article)].NumberOfLikes

		img := services.ReadArticleImage(
			headerDocRes,
			v.Article,
		)

		data.ArticleHeader = append(data.ArticleHeader,
			apiOutputFormatter.ArticleHeader{
				Article:           v.Article,
				Description:       v.Description,
				Introduction:      v.Introduction,
				ValidityStartDate: v.ValidityStartDate,
				ValidityStartTime: v.ValidityStartTime,
				ValidityEndDate:   v.ValidityEndDate,
				ValidityEndTime:   v.ValidityEndTime,
				Tag1:              v.Tag1,
				Tag2:              v.Tag2,
				Tag3:              v.Tag3,
				Tag4:              v.Tag4,
				LastChangeDate:	   v.LastChangeDate,
				LastChangeTime:	   v.LastChangeTime,
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
