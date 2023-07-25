package controllersProductMasterDetailBusinessPartner

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductMasterDetailBusinessPartnerController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductMasterDetailBusinessPartnerController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	product := controller.GetString("product")
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "detail-business-partner"
	redisKeyCategory3 := product
	userType := controller.GetString(":userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	productMasterBusinessPartner := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{
			Product: product,
		},
		ProductMasterBusinessPartner: &apiInputReader.ProductMasterBusinessPartner{
			Product:             product,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
			userType,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.ProductMaster

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
			controller.request(productMasterBusinessPartner)
		}()
	} else {
		controller.request(productMasterBusinessPartner)
	}
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGenerals(
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
		controller.CustomLogger.Error("createProductMasterRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) createProductMasterRequestBusinessPartners(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsBusinessPartners(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product: input.ProductMasterBusinessPartner.Product,
		},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterDocReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*productDescByBPRes.Message.ProductDescByBP))
	isMarkedForDeletion := false

	for _, v := range *productDescByBPRes.Message.ProductDescByBP {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: *requestPram.BusinessPartner,
					ProductDescByBP: []apiModuleRuntimesRequestsProductMaster.ProductDescByBP{
						{
							Language:            *requestPram.Language,
							IsMarkedForDeletion: &isMarkedForDeletion,
						},
					},
				},
			},
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProductDescsByBP(
		requestPram,
		productDescsByBP,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	businessPartnerRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*businessPartnerRes.Message.BusinessPartner))

	for _, v := range *businessPartnerRes.Message.BusinessPartner {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
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
	controller *ProductMasterDetailBusinessPartnerController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger)

	generalRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
		input,
	)

	productMasterBusinessPartnerRes := controller.createProductMasterRequestBusinessPartners(
		controller.UserInfo,
		input,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		productMasterBusinessPartnerRes,
	)

	businessPartnerRes := controller.createBusinessPartnerRequest(
		controller.UserInfo,
		productMasterBusinessPartnerRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		generalRes,
		productMasterBusinessPartnerRes,
		productDescByBPRes,
		businessPartnerRes,
		productDocRes,
	)
}

func (
	controller *ProductMasterDetailBusinessPartnerController,
) fin(
	generalRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterbusinessPartnerRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	//businessPartnerMapper := services.BusinessPartnerNameMapper(
	//	businessPartnerRes,
	//)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *generalRes.Message.General {
		//img := services.CreateProductImage(
		//	productDocRes,
		//	v.BusinessPartner,
		//	v.Product,
		//)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.ProductMasterGeneralWithOthers = append(data.ProductMasterGeneralWithOthers,
			apiOutputFormatter.ProductMasterGeneralWithOthers{
				Product:           v.Product,
				ProductName:       &productDescription,
				ProductGroup:      v.ProductGroup,
				BaseUnit:          v.BaseUnit,
				ValidityStartDate: v.ValidityStartDate,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
			},
		)
	}

	for _, v := range *generalRes.Message.BusinessPartner {
		data.ProductMasterDetailBusinessPartner = append(data.ProductMasterDetailBusinessPartner,
			apiOutputFormatter.ProductMasterDetailBusinessPartner{
				BusinessPartner:        v.BusinessPartner,
				ValidityStartDate:      v.ValidityStartDate,
				ValidityEndDate:        v.ValidityEndDate,
				BusinessPartnerProduct: v.BusinessPartnerProduct,
				CreationDate:           v.CreationDate,
				LastChangeDate:         v.LastChangeDate,
				IsMarkedForDeletion:    v.IsMarkedForDeletion,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
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
