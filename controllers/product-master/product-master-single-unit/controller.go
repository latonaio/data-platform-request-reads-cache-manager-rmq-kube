package controllersProductSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductMasterSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductMasterSingleUnitController) Get() {
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "single-unit"
	product := controller.GetString("product")

	isMarkedForDeletion := false
	qrCodeDocType := "QRCODE"

	productMasterSingleUnit := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{
			Product:             product,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductMasterGeneralDoc: &apiInputReader.ProductMasterGeneralDoc{
			Product:                  product,
			BusinessPartner:          controller.UserInfo.BusinessPartner,
			DocType:                  &qrCodeDocType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
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
			controller.request(productMasterSingleUnit)
		}()
	} else {
		controller.request(productMasterSingleUnit)
	}
}

func (
	controller *ProductMasterSingleUnitController,
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGeneral(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product:             input.ProductMasterGeneral.Product,
			IsMarkedForDeletion: input.ProductMasterGeneral.IsMarkedForDeletion,
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
		controller.CustomLogger.Error("createProductMasterRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterSingleUnitController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,

) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterSingleUnitController,
) createProductMasterQrCodeDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterQrCodeDocReads(
		requestPram,
		input,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterQrCodeDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterSingleUnitController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	productMasterRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 1)
	isMarkedForDeletion := false

	for _, v := range *productMasterRes.Message.General {
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
	controller *ProductMasterSingleUnitController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	productRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
		input,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		productRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	productQrCodeDocRes := controller.createProductMasterQrCodeDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		productRes,
		productDescByBPRes,
		productDocRes,
		productQrCodeDocRes,
	)
}

func (
	controller *ProductMasterSingleUnitController,
) fin(
	productRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	productQrCodeDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *productRes.Message.General {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		qrcode := services.CreateQRCodeProductDocImage(
			productQrCodeDocRes,
		)

		data.ProductMasterSingleUnit = append(data.ProductMasterSingleUnit,
			apiOutputFormatter.ProductMasterSingleUnit{
				Product:                       v.Product,
				ProductName:                   &productDescription,
				ProductType:                   &v.ProductType,
				GrossWeight:                   v.GrossWeight,
				NetWeight:                     v.NetWeight,
				WeightUnit:                    v.WeightUnit,
				InternalCapacityQuantity:      v.InternalCapacityQuantity,
				InternalCapacityQuantityUnit:  v.InternalCapacityQuantityUnit,
				SizeOrDimensionText:           v.SizeOrDimensionText,
				ProductStandardID:             v.ProductStandardID,
				IndustryStandardName:          v.IndustryStandardName,
				ItemCategory:                  &v.ItemCategory,
				CountryOfOrigin:               v.CountryOfOrigin,
				CountryOfOriginLanguage:       v.CountryOfOriginLanguage,
				LocalRegionOfOrigin:           v.LocalRegionOfOrigin,
				LocalSubRegionOfOrigin:        v.LocalSubRegionOfOrigin,
				BarcodeType:                   v.BarcodeType,
				MarkingOfMaterial:             v.MarkingOfMaterial,
				ProductAccountAssignmentGroup: v.ProductAccountAssignmentGroup,
				ValidityEndDate:               v.ValidityEndDate,
				CreationDate:                  v.CreationDate,
				LastChangeDate:                v.LastChangeDate,
				IsMarkedForDeletion:           v.IsMarkedForDeletion,
				Images: &apiOutputFormatter.Images{
					Product: img,
					QRCode:  qrcode,
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
