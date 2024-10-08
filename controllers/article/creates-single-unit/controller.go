package controllersArticleCreatesSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/distribution-profile"
	apiModuleRuntimesRequestsArticleType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/article-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesDistributionProfile "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/distribution-profile"
	apiModuleRuntimesResponsesArticleType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/article-type"
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

type ArticleCreatesSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ArticleCreatesSingleUnit struct {
	BusinessPartnerGeneral            []apiOutputFormatter.BusinessPartnerGeneral            `json:"BusinessPartnerGeneral"`
	BusinessPartnerPersonOrganization []apiOutputFormatter.BusinessPartnerPersonOrganization `json:"BusinessPartnerPersonOrganization"`
	BusinessPartnerRoleText           []apiOutputFormatter.BusinessPartnerRoleText           `json:"BusinessPartnerRoleText"`
	ArticleTypeText                   []apiOutputFormatter.ArticleTypeText                   `json:"ArticleTypeText"`
	DistributionProfileText           []apiOutputFormatter.DistributionProfileText           `json:"DistributionProfileText"`
	SiteAddress                       []apiOutputFormatter.SiteAddress                       `json:"SiteAddress"`
	SiteHeader                        []apiOutputFormatter.SiteHeader                        `json:"SiteHeader"`
	SiteAddressWithHeader             []apiOutputFormatter.SiteAddressWithHeader             `json:"SiteAddressWithHeader"`
	ShopHeader                        []apiOutputFormatter.ShopHeader                        `json:"ShopHeader"`
}

func (controller *ArticleCreatesSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "businessPartner"
	redisKeyCategory2 := "businessPartnerRole"
	redisKeyCategory3 := "article-creates-single-unit"

	businessPartner, _ := controller.GetInt("businessPartner")

	businessPartnerRole := controller.GetString("businessPartnerRole") // UIでログイン時にロール選択で決まる、該当箇所開発時までは当面固定値にするか

	localSubRegion := controller.GetString("localSubRegion")

	ArticleCreatesSingleUnitBP := apiInputReader.BusinessPartner{}
	ArticleCreatesSingleUnitArticleType := apiInputReader.ArticleTypeGlobal{}
	ArticleCreatesSingleUnitDistributionProfile := apiInputReader.DistributionProfileGlobal{}
	ArticleCreatesSingleUnitSiteAddress := apiInputReader.Site{}

	isReleased := true
	isMarkedForDeletion := false

	ArticleCreatesSingleUnitBP = apiInputReader.BusinessPartner{
		BusinessPartnerGeneral: &apiInputReader.BusinessPartnerGeneral{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerPersonOrganization: &apiInputReader.BusinessPartnerPersonOrganization{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	ArticleCreatesSingleUnitArticleType = apiInputReader.ArticleTypeGlobal{
		ArticleTypeText: &apiInputReader.ArticleTypeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	ArticleCreatesSingleUnitDistributionProfile = apiInputReader.DistributionProfileGlobal{
		DistributionProfileText: &apiInputReader.DistributionProfileText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	ArticleCreatesSingleUnitSiteAddress = apiInputReader.Site{
		SiteAddress: &apiInputReader.SiteAddress{
			LocalSubRegion: &localSubRegion,
		},
		SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{},
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
		var responseData ArticleCreatesSingleUnit

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
				ArticleCreatesSingleUnitBP,
				ArticleCreatesSingleUnitArticleType,
				ArticleCreatesSingleUnitDistributionProfile,
				ArticleCreatesSingleUnitSiteAddress,
				businessPartnerRole,
				isReleased,
				isMarkedForDeletion,
			)
		}()
	} else {
		controller.request(
			ArticleCreatesSingleUnitBP,
			ArticleCreatesSingleUnitArticleType,
			ArticleCreatesSingleUnitDistributionProfile,
			ArticleCreatesSingleUnitSiteAddress,
			businessPartnerRole,
			isReleased,
			isMarkedForDeletion,
		)
	}
}

func (
	controller *ArticleCreatesSingleUnitController,
) createBusinessPartnerRequestGeneral(
	requestPram *apiInputReader.Request,
	inputArticleCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
		BusinessPartner:     inputArticleCreatesSingleUnitBP.BusinessPartnerGeneral.BusinessPartner,
		IsMarkedForDeletion: inputArticleCreatesSingleUnitBP.BusinessPartnerPersonOrganization.IsMarkedForDeletion,
	})

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
	controller *ArticleCreatesSingleUnitController,
) createBusinessPartnerRequestPersonOrganization(
	requestPram *apiInputReader.Request,
	inputArticleCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{}

	input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{
		BusinessPartner: inputArticleCreatesSingleUnitBP.BusinessPartnerPersonOrganization.BusinessPartner,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPersonOrganization(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonOrganization Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) createBusinessPartnerRequestGeneralOrganizationBP(
	requestPram *apiInputReader.Request,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner:     v.OrganizationBusinessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneralOrganizationBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerRole string,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {
	inputBusinessPartnerRole := businessPartnerRole

	input := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRole{
		BusinessPartnerRole: inputBusinessPartnerRole,
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
	controller *ArticleCreatesSingleUnitController,
) CreateArticleTypeRequestTexts(
	requestPram *apiInputReader.Request,
	inputArticleCreatesSingleUnitArticleType apiInputReader.ArticleTypeGlobal,
) *apiModuleRuntimesResponsesArticleType.ArticleTypeRes {
	input := apiModuleRuntimesRequestsArticleType.ArticleType{
		Text: []apiModuleRuntimesRequestsArticleType.Text{
			{
				Language:            inputArticleCreatesSingleUnitArticleType.ArticleTypeText.Language,
				IsMarkedForDeletion: inputArticleCreatesSingleUnitArticleType.ArticleTypeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesArticleType.ArticleTypeRes{}
	responseBody := apiModuleRuntimesRequestsArticleType.ArticleTypeReadsTexts(
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
		controller.CustomLogger.Error("CreateArticleTypeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) CreateDistributionProfileRequestTexts(
	requestPram *apiInputReader.Request,
	inputArticleCreatesSingleUnitDistributionProfile apiInputReader.DistributionProfileGlobal,
) *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes {
	input := apiModuleRuntimesRequestsDistributionProfile.DistributionProfile{
		Text: []apiModuleRuntimesRequestsDistributionProfile.Text{
			{
				Language:            inputArticleCreatesSingleUnitDistributionProfile.DistributionProfileText.Language,
				IsMarkedForDeletion: inputArticleCreatesSingleUnitDistributionProfile.DistributionProfileText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes{}
	responseBody := apiModuleRuntimesRequestsDistributionProfile.DistributionProfileReadsTexts(
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
		controller.CustomLogger.Error("CreateDistributionProfileRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) createSiteRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteRes {
	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
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
			"ローカルサブ地域に対してのサイトが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) createSiteRequestHeadersBySites(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	for _, v := range *siteRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site:                v.Site,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReadsHeadersBySites(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.Header == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対して有効なサイトヘッダデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestHeadersBySites Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
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
	controller *ArticleCreatesSingleUnitController,
) createShopRequestHeadersByShopOwner(
	requestPram *apiInputReader.Request,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesShop.ShopRes {
	var input apiModuleRuntimesRequestsShop.Header

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		shopOwner := v.OrganizationBusinessPartner

		input = apiModuleRuntimesRequestsShop.Header{
			ShopOwner:           &shopOwner,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsHeadersByShopOwner(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"OrganizationBusinessPartnerに対して有効な店舗が見つかりませんでした",
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
		controller.CustomLogger.Error("createShopRequestHeadersByShopOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleCreatesSingleUnitController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	var input = apiInputReader.Shop{}

	for _, v := range *shopHeaderRes.Message.Header {
		input = apiInputReader.Shop{
			ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
				Shop: v.Shop,
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
	controller *ArticleCreatesSingleUnitController,
) request(
	inputArticleCreatesSingleUnitBP apiInputReader.BusinessPartner,
	inputArticleCreatesSingleUnitArticleType apiInputReader.ArticleTypeGlobal,
	inputArticleCreatesSingleUnitDistributionProfile apiInputReader.DistributionProfileGlobal,
	inputArticleCreatesSingleUnitSiteAddress apiInputReader.Site,
	businessPartnerRole string,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(6)

	var businessPartnerGeneralRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerPersonOrganizationRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerGeneralOrganizationBPRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	var articleTypeTextRes *apiModuleRuntimesResponsesArticleType.ArticleTypeRes

	var distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes

	var siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes

	var siteHeaderRes apiModuleRuntimesResponsesSite.SiteRes
	var shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes

	var siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes
	var shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	go func() {
		defer wg.Done()
		businessPartnerGeneralRes = *controller.createBusinessPartnerRequestGeneral(
			controller.UserInfo,
			inputArticleCreatesSingleUnitBP,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerPersonOrganizationRes = *controller.createBusinessPartnerRequestPersonOrganization(
			controller.UserInfo,
			inputArticleCreatesSingleUnitBP,
		)
		businessPartnerGeneralOrganizationBPRes = *controller.createBusinessPartnerRequestGeneralOrganizationBP(
			controller.UserInfo,
			&businessPartnerPersonOrganizationRes,
			isMarkedForDeletion,
		)
		shopHeaderRes = *controller.createShopRequestHeadersByShopOwner(
			controller.UserInfo,
			&businessPartnerPersonOrganizationRes,
			isReleased,
			isMarkedForDeletion,
		)
		shopHeaderDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			shopHeaderRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerRoleTextRes = controller.CreateBusinessPartnerRoleRequestText(
			controller.UserInfo,
			businessPartnerRole,
		)
	}()

	go func() {
		defer wg.Done()
		articleTypeTextRes = controller.CreateArticleTypeRequestTexts(
			controller.UserInfo,
			inputArticleCreatesSingleUnitArticleType,
		)
	}()

	go func() {
		defer wg.Done()
		distributionProfileTextRes = controller.CreateDistributionProfileRequestTexts(
			controller.UserInfo,
			inputArticleCreatesSingleUnitDistributionProfile,
		)
	}()

	go func() {
		defer wg.Done()
		siteAddressRes = controller.createSiteRequestAddressesByLocalSubRegion(
			controller.UserInfo,
			inputArticleCreatesSingleUnitSiteAddress,
		)
		siteHeaderRes = *controller.createSiteRequestHeadersBySites(
			controller.UserInfo,
			siteAddressRes,
			isReleased,
			isMarkedForDeletion,
		)
		siteHeaderDocRes = controller.createSiteDocRequest(
			controller.UserInfo,
			inputArticleCreatesSingleUnitSiteAddress,
		)
	}()

	wg.Wait()

	controller.fin(
		&businessPartnerGeneralRes,
		&businessPartnerPersonOrganizationRes,
		&businessPartnerGeneralOrganizationBPRes,
		businessPartnerRoleTextRes,
		articleTypeTextRes,
		distributionProfileTextRes,
		siteAddressRes,
		&siteHeaderRes,
		siteHeaderDocRes,
		&shopHeaderRes,
		shopHeaderDocRes,
	)
}

func (
	controller *ArticleCreatesSingleUnitController,
) fin(
	businessPartnerGeneralRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerGeneralOrganizationBPRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	articleTypeTextRes *apiModuleRuntimesResponsesArticleType.ArticleTypeRes,
	distributionProfileTextRes *apiModuleRuntimesResponsesDistributionProfile.DistributionProfileRes,
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
) {
	businessPartnerNameMapper := services.BusinessPartnerNameMapper(
		businessPartnerGeneralRes,
	)

	businessPartnerNameMapperOrganizationBP := services.BusinessPartnerNameMapper(
		businessPartnerGeneralOrganizationBPRes,
	)

	businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
		businessPartnerRoleTextRes.Message.Text,
	)

	articleTypeTextMapper := services.ArticleTypeTextMapper(
		articleTypeTextRes.Message.Text,
	)

	distributionProfileTextMapper := services.DistributionProfileTextMapper(
		distributionProfileTextRes.Message.Text,
	)

	siteHeadersMapper := services.SiteHeadersMapper(
		siteHeaderRes,
	)

	data := ArticleCreatesSingleUnit{}

	for _, v := range *businessPartnerGeneralRes.Message.General {
		data.BusinessPartnerGeneral = append(data.BusinessPartnerGeneral,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: businessPartnerNameMapper[v.BusinessPartner].BusinessPartnerName,
			},
		)
	}

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		data.BusinessPartnerPersonOrganization = append(data.BusinessPartnerPersonOrganization,
			apiOutputFormatter.BusinessPartnerPersonOrganization{
				BusinessPartner:                 v.BusinessPartner,
				OrganizationBusinessPartner:     v.OrganizationBusinessPartner,
				OrganizationBusinessPartnerName: businessPartnerNameMapperOrganizationBP[v.OrganizationBusinessPartner].BusinessPartnerName,
			},
		)
	}

	for _, v := range *businessPartnerRoleTextRes.Message.Text {
		data.BusinessPartnerRoleText = append(data.BusinessPartnerRoleText,
			apiOutputFormatter.BusinessPartnerRoleText{
				BusinessPartnerRole:     v.BusinessPartnerRole,
				Language:                v.Language,
				BusinessPartnerRoleName: businessPartnerRoleTextMapper[v.BusinessPartnerRole].BusinessPartnerRoleName,
			},
		)
	}

	for _, v := range *articleTypeTextRes.Message.Text {
		data.ArticleTypeText = append(data.ArticleTypeText,
			apiOutputFormatter.ArticleTypeText{
				ArticleType:     v.ArticleType,
				Language:      v.Language,
				ArticleTypeName: articleTypeTextMapper[v.ArticleType].ArticleTypeName,
			},
		)
	}

	for _, v := range *distributionProfileTextRes.Message.Text {
		data.DistributionProfileText = append(data.DistributionProfileText,
			apiOutputFormatter.DistributionProfileText{
				DistributionProfile:     v.DistributionProfile,
				Language:                v.Language,
				DistributionProfileName: distributionProfileTextMapper[v.DistributionProfile].DistributionProfileName,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {
		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				Description: v.Description,
			},
		)
	}

	for _, v := range *siteAddressRes.Message.Address {
		siteType := siteHeadersMapper[strconv.Itoa(v.Site)].SiteType
		validityStartDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartDate
		validityStartTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartTime
		validityEndDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndDate
		validityEndTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndTime
		description := siteHeadersMapper[strconv.Itoa(v.Site)].Description
		introduction := siteHeadersMapper[strconv.Itoa(v.Site)].Introduction
		tag1 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag1
		tag2 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag2
		tag3 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag3
		tag4 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag4
		img := services.ReadSiteImage(
			siteHeaderDocRes,
			v.Site,
		)

		data.SiteAddressWithHeader = append(data.SiteAddressWithHeader,
			apiOutputFormatter.SiteAddressWithHeader{
				Site:              v.Site,
				AddressID:         v.AddressID,
				LocalSubRegion:    v.LocalSubRegion,
				LocalRegion:       v.LocalRegion,
				SiteType:          siteType,
				ValidityStartDate: validityStartDate,
				ValidityStartTime: validityStartTime,
				ValidityEndDate:   validityEndDate,
				ValidityEndTime:   validityEndTime,
				Description:       description,
				Introduction:      introduction,
				Tag1:              tag1,
				Tag2:              tag2,
				Tag3:              tag3,
				Tag4:              tag4,
				PostalCode:        v.PostalCode,
				Country:           v.Country,
				GlobalRegion:      v.GlobalRegion,
				TimeZone:          v.TimeZone,
				District:          v.District,
				StreetName:        v.StreetName,
				CityName:          v.CityName,
				Building:          v.Building,
				Floor:             v.Floor,
				Room:              v.Room,
				XCoordinate:       v.XCoordinate,
				YCoordinate:       v.YCoordinate,
				ZCoordinate:       v.ZCoordinate,
				Images: apiOutputFormatter.Images{
					Site: img,
				},
			},
		)
	}

	for _, v := range *shopHeaderRes.Message.Header {
		img := services.ReadShopImage(
			shopHeaderDocRes,
			v.Shop,
		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					Shop: img,
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
