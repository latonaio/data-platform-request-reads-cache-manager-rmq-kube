package controllersArticleSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsArticleType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article-type"
	apiModuleRuntimesRequestsArticle "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article/article"
	apiModuleRuntimesRequestsArticleDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article/article-doc"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/distribution-profile"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesArticle "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article"
	apiModuleRuntimesResponsesArticleType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article-type"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/distribution-profile"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
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

type ArticleSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ArticleSingleUnit struct {
	ArticleHeader	[]apiOutputFormatter.ArticleHeader	`json:"ArticleHeader"`
	ArticleAddress	[]apiOutputFormatter.ArticleAddress	`json:"ArticleAddress"`
	ArticleCounter	[]apiOutputFormatter.ArticleCounter	`json:"ArticleCounter"`
	SiteHeader		[]apiOutputFormatter.SiteHeader		`json:"SiteHeader"`
	SiteAddress		[]apiOutputFormatter.SiteAddress	`json:"SiteAddress"`
	ShopHeader		[]apiOutputFormatter.ShopHeader		`json:"ShopHeader"`
}

func (controller *ArticleSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "article"
	redisKeyCategory2 := "article-single-unit"
	article, _ := controller.GetInt("article")

	ArticleSingleUnitArticle := apiInputReader.Article{}

	isReleased := true
	isMarkedForDeletion := false

	//docType := "QRCODE"

	ArticleSingleUnitArticle = apiInputReader.Article{
		ArticleHeader: &apiInputReader.ArticleHeader{
			Article:             article,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ArticleAddress: &apiInputReader.ArticleAddress{
			Article:			article,
		},
		ArticleCounter: &apiInputReader.ArticleCounter{
			Article:			article,
		},
		ArticleDocHeaderDoc: &apiInputReader.ArticleDocHeaderDoc{
			Article:					article,
			//DocType:					&docType,
			DocIssuerBusinessPartner:	controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(article),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData ArticleSingleUnit

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
			controller.request(ArticleSingleUnitArticle)
		}()
	} else {
		controller.request(ArticleSingleUnitArticle)
	}
}

func (
	controller *ArticleSingleUnitController,
) createArticleRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReads(
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
		controller.CustomLogger.Error("createArticleRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
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
	controller *ArticleSingleUnitController,
) createArticleRequestAddresses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReads(
		requestPram,
		input,
		&controller.Controller,
		"Addresses",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"イベントのアドレスが見つかりませんでした",
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
		controller.CustomLogger.Error("createArticleRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) createArticleRequestCounter(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
) *apiModuleRuntimesResponsesArticle.ArticleRes {
	responseJsonData := apiModuleRuntimesResponsesArticle.ArticleRes{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleReads(
		requestPram,
		input,
		&controller.Controller,
		"Counter",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

		if len(*responseJsonData.Message.Counter) == 0 {
			status := 500
			services.HandleError(
				&controller.Controller,
				"記事にカウンタデータがありません",
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
		controller.CustomLogger.Error("createArticleRequestCounter Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) createBusinessPartnerRequestGeneral(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*articleRes.Message.Header))

	for _, v := range *articleRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ArticleOwner,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
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
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {

	businessPartnerRole := &(*articleRes.Message.Header)[0].ArticleOwnerBusinessPartnerRole

	var inputBusinessPartnerRole *string

	if businessPartnerRole != nil {
		inputBusinessPartnerRole = businessPartnerRole
	}

	input := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRole{
		BusinessPartnerRole: *inputBusinessPartnerRole,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRoleReadsText(
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
		controller.CustomLogger.Error("CreateBusinessPartnerRoleRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) createSiteRequestHeader(
	requestPram *apiInputReader.Request,
	articleHeaderRes apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	header := articleHeaderRes.Message.Header

	var input = apiInputReader.Site{}

	input.SiteHeader = &apiInputReader.SiteHeader{
		Site:                (*header)[0].Site,
		IsReleased:          (*header)[0].IsReleased,
		IsMarkedForDeletion: (*header)[0].IsMarkedForDeletion,
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
	controller *ArticleSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	articleHeaderRes apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
	var input = apiInputReader.Site{}

	for _, v := range *articleHeaderRes.Message.Header {
		input = apiInputReader.Site{
			SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{
				Site: v.Site,
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
	controller *ArticleSingleUnitController,
) createShopRequestHeader(
	requestPram *apiInputReader.Request,
	articleHeaderRes apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	header := articleHeaderRes.Message.Header

	var input = apiInputReader.Shop{}

	input.ShopHeader = &apiInputReader.ShopHeader{
		Shop:                *(*header)[0].Shop,
		IsReleased:          (*header)[0].IsReleased,
		IsMarkedForDeletion: (*header)[0].IsMarkedForDeletion,
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
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
		controller.CustomLogger.Error("createShopRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	articleHeaderRes apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	var input = apiInputReader.Shop{}

	for _, v := range *articleHeaderRes.Message.Header {
		input = apiInputReader.Shop{
			ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
				Shop: *v.Shop,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopDocRes{}
	responseBody := apiModuleRuntimesRequestsShopDoc.ShopDocReads(
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
		controller.CustomLogger.Error("createShopDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) CreateArticleTypeRequestText(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesArticleType.ArticleTypeRes {

	articleType := &(*articleRes.Message.Header)[0].ArticleType

	var inputArticleType *string

	if articleType != nil {
		inputArticleType = articleType
	}

	input := apiModuleRuntimesRequestsArticleType.ArticleType{
		ArticleType: *inputArticleType,
	}

	responseJsonData := apiModuleRuntimesResponsesArticleType.ArticleTypeRes{}
	responseBody := apiModuleRuntimesRequestsArticleType.ArticleTypeReadsText(
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
		controller.CustomLogger.Error("CreateArticleTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) CreateDistributionProfileRequestText(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes {

	distributionProfile := &(*articleRes.Message.Header)[0].DistributionProfile

	var inputDistributionProfile *string

	if distributionProfile != nil {
		inputDistributionProfile = distributionProfile
	}

	input := apiModuleRuntimesRequestsDistributionProfile.DistributionProfile{
		DistributionProfile: *inputDistributionProfile,
	}

	responseJsonData := apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes{}
	responseBody := apiModuleRuntimesRequestsDistributionProfile.DistributionProfileReadsText(
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
		controller.CustomLogger.Error("CreateDistributionProfileRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*articleRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*articleRes.Message.Address)[0].LocalRegion
	country := &(*articleRes.Message.Address)[0].Country

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
	controller *ArticleSingleUnitController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	articleRes *apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*articleRes.Message.Address)[0].LocalRegion
	country := &(*articleRes.Message.Address)[0].Country

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
	controller *ArticleSingleUnitController,
) createSiteRequestAddresses(
	requestPram *apiInputReader.Request,
	headerRes apiModuleRuntimesResponsesArticle.ArticleRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input = apiInputReader.Site{}

	for _, v := range *headerRes.Message.Header {
		input = apiInputReader.Site{
			SiteAddress: &apiInputReader.SiteAddress{
				Site: v.Site,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
		requestPram,
		input,
		&controller.Controller,
		"Addresses",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"サイトのアドレスが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestAddresses Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleSingleUnitController,
) request(
	inputArticle apiInputReader.Article,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(12)

	var addressRes apiModuleRuntimesResponsesArticle.ArticleRes
	var counterRes apiModuleRuntimesResponsesArticle.ArticleRes

	var businessPartnerGeneralRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	var siteHeaderRes apiModuleRuntimesResponsesSite.SiteRes

	var shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes

	var articleTypeTextRes *apiModuleRuntimesResponsesArticleType.ArticleTypeRes

	var distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes

	var siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes

	var shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	var headerDocRes *apiModuleRuntimesResponsesArticle.ArticleDocRes

	var siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes

	var localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	headerRes := *controller.createArticleRequestHeader(
		controller.UserInfo,
		inputArticle,
	)

	go func() {
		defer wg.Done()
		addressRes = *controller.createArticleRequestAddresses(
			controller.UserInfo,
			inputArticle,
		)

		localSubRegionTextRes = controller.CreateLocalSubRegionRequestText(
			controller.UserInfo,
			&addressRes,
		)

		localRegionTextRes = controller.CreateLocalRegionRequestText(
			controller.UserInfo,
			&addressRes,
		)
	}()

	go func() {
		defer wg.Done()
		counterRes = *controller.createArticleRequestCounter(
			controller.UserInfo,
			inputArticle,
		)
	}()

	go func() {
		defer wg.Done()
		headerDocRes = controller.createArticleDocRequest(
			controller.UserInfo,
			inputArticle,
		)
		controller.CustomLogger.Debug("complete headerDocRes")
	}()

	go func() {
		defer wg.Done()
		businessPartnerGeneralRes = *controller.createBusinessPartnerRequestGeneral(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerRoleTextRes = controller.CreateBusinessPartnerRoleRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteHeaderRes = *controller.createSiteRequestHeader(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		shopHeaderRes = *controller.createShopRequestHeader(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		articleTypeTextRes = controller.CreateArticleTypeRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		distributionProfileTextRes = controller.CreateDistributionProfileRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteHeaderDocRes = controller.createSiteDocRequest(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		shopHeaderDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		siteAddressRes = controller.createSiteRequestAddresses(
			controller.UserInfo,
			headerRes,
		)
	}()

	wg.Wait()

	controller.fin(
		&headerRes,
		&addressRes,
		&counterRes,
		&businessPartnerGeneralRes,
		businessPartnerRoleTextRes,
		&siteHeaderRes,
		&shopHeaderRes,
		articleTypeTextRes,
		distributionProfileTextRes,
		localSubRegionTextRes,
		localRegionTextRes,
		headerDocRes,
		siteHeaderDocRes,
		shopHeaderDocRes,
		siteAddressRes,
	)
}

func (
	controller *ArticleSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	addressRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	counterRes *apiModuleRuntimesResponsesArticle.ArticleRes,
	businessPartnerGeneralRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
	articleTypeTextRes *apiModuleRuntimesResponsesArticleType.ArticleTypeRes,
	distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	headerDocRes *apiModuleRuntimesResponsesArticle.ArticleDocRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
) {
	businessPartnerNameMapper := services.BusinessPartnerNameMapper(
		businessPartnerGeneralRes,
	)

	businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
		businessPartnerRoleTextRes.Message.Text,
	)

	siteMapper := services.SiteMapper(
		siteHeaderRes,
	)

	shopMapper := services.ShopMapper(
		shopHeaderRes,
	)

	articleTypeTextMapper := services.ArticleTypeTextMapper(
		articleTypeTextRes.Message.Text,
	)

	distributionProfileTextMapper := services.DistributionProfileTextMapper(
		distributionProfileTextRes.Message.Text,
	)

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	data := ArticleSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadArticleImage(
			headerDocRes,
			v.Article,
		)

		qrcode := services.CreateQRCodeArticleDocImage(
			headerDocRes,
			v.Article,
		)

		documentImage := services.ReadDocumentImageArticle(
			headerDocRes,
			v.Article,
		)

		shopDescription := shopMapper[*v.Shop].Description

		data.ArticleHeader = append(data.ArticleHeader,
			apiOutputFormatter.ArticleHeader{
				Article:                             v.Article,
				ArticleType:                         v.ArticleType,
				ArticleTypeName:                     articleTypeTextMapper[v.ArticleType].ArticleTypeName,
				ArticleOwner:                        v.ArticleOwner,
				ArticleOwnerName:                    businessPartnerNameMapper[v.ArticleOwner].BusinessPartnerName,
				ArticleOwnerBusinessPartnerRoleName: businessPartnerRoleTextMapper[v.ArticleOwnerBusinessPartnerRole].BusinessPartnerRoleName,
				PersonResponsible:                   v.PersonResponsible,
				ValidityStartDate:                   v.ValidityStartDate,
				ValidityStartTime:                   v.ValidityStartTime,
				ValidityEndDate:                     v.ValidityEndDate,
				ValidityEndTime:                     v.ValidityEndTime,
				Description:                         v.Description,
				LongText:                            v.LongText,
				Introduction:                        v.Introduction,
				Site:                                v.Site,
				SiteDescription:                     siteMapper[v.Site].Description,
				Shop:                                v.Shop,
				ShopDescription:                     &shopDescription,
				Tag1:                                v.Tag1,
				Tag2:                                v.Tag2,
				Tag3:                                v.Tag3,
				Tag4:                                v.Tag4,
				DistributionProfile:                 v.DistributionProfile,
				DistributionProfileName:             distributionProfileTextMapper[v.DistributionProfile].DistributionProfileName,
				QuestionnaireType:                   v.QuestionnaireType,
				//QuestionnaireTypeName:             questionnaireTypeTextMapper[v.QuestionnaireType].QuestionnaireTypeName,
				QuestionnaireTemplate: v.QuestionnaireTemplate,
				//QuestionnaireTemplateName:         questionnaireTemplateTextMapper[v.QuestionnaireTemplate].QuestionnaireTemplateName,
				CreateUser: v.CreateUser,
				//CreateUserFullName:         	     v.CreateUserFullName,
				//CreateUserNickName:         	     v.CreateUserNickName,
				LastChangeUser: v.LastChangeUser,
				//LastChangeUserFullName:            v.LastChangeUserFullName,
				//LastChangeUserNickName:            v.LastChangeUserNickName,
				Images: apiOutputFormatter.Images{
					Article:              img,
					QRCode:               qrcode,
					DocumentImageArticle: documentImage,
				},
			},
		)
	}

	for _, v := range *addressRes.Message.Address {
		data.ArticleAddress = append(data.ArticleAddress,
			apiOutputFormatter.ArticleAddress{
				Article:            v.Article,
				AddressID:          v.AddressID,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
				PostalCode:         &v.PostalCode,
				StreetName:         v.StreetName,
				CityName:           v.CityName,
				Building:           v.Building,
			},
		)
	}
	
	for _, v := range *counterRes.Message.Counter {
		data.ArticleCounter = append(data.ArticleCounter,
			apiOutputFormatter.ArticleCounter{
				Article:			v.Article,
				NumberOfLikes:		v.NumberOfLikes,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {

		img := services.ReadSiteImage(
			siteHeaderDocRes,
			v.Site,
		)

		qrcode := services.CreateQRCodeSiteDocImage(
			siteHeaderDocRes,
			v.Site,
		)

		documentImage := services.ReadDocumentImageSite(
			siteHeaderDocRes,
			v.Site,
		)

		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Site:              img,
					QRCode:            qrcode,
					DocumentImageSite: documentImage,
				},
			},
		)
	}

	for _, v := range *siteAddressRes.Message.Address {
		data.SiteAddress = append(data.SiteAddress,
			apiOutputFormatter.SiteAddress{
				Site:               v.Site,
				AddressID:          v.AddressID,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapper[v.LocalRegion].LocalRegionName,
				PostalCode:         v.PostalCode,
				StreetName:         v.StreetName,
				CityName:           v.CityName,
				Building:           v.Building,
			},
		)
	}

	for _, v := range *shopHeaderRes.Message.Header {

		img := services.ReadShopImage(
			shopHeaderDocRes,
			v.Shop,
		)

		qrcode := services.CreateQRCodeShopDocImage(
			shopHeaderDocRes,
			v.Shop,
		)

		documentImage := services.ReadDocumentImageShop(
			shopHeaderDocRes,
			v.Shop,
		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Shop:              img,
					QRCode:            qrcode,
					DocumentImageShop: documentImage,
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
