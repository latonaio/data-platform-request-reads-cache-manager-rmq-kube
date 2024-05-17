package controllersPointTransactionSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPointTransaction "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction"
	apiModuleRuntimesRequestsPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop-type"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPointTransaction "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-transaction"
	apiModuleRuntimesResponsesPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-transaction-type"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiModuleRuntimesResponsesShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type PointTransactionSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *PointTransactionSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "point-transaction"
	redisKeyCategory2 := "point-transaction-single-unit"
	pointTransaction, _ := controller.GetInt("pointTransaction")

	PointTransactionSingleUnit := apiInputReader.PointTransaction{}

	PointTransactionSingleUnit = apiInputReader.PointTransaction{
		PointTransactionHeader: &apiInputReader.PointTransactionHeader{
			PointTransaction: pointTransaction,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(pointTransaction),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.PointTransaction

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
			controller.request(PointTransactionSingleUnit)
		}()
	} else {
		controller.request(PointTransactionSingleUnit)
	}
}

func (
	controller *PointTransactionSingleUnitController,
) createPointTransactionRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
) *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes {
	responseJsonData := apiModuleRuntimesResponsesPointTransaction.PointTransactionRes{}
	responseBody := apiModuleRuntimesRequestsPointTransaction.PointTransactionReadsHeader(
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
		controller.CustomLogger.Error("createPointTransactionRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	pointTransactionRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*pointTransactionRes.Message.Header))

	for _, v := range *pointTransactionRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Sender,
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
	controller *PointTransactionSingleUnitController,
) CreatePointTransactionTypeRequestText(
	requestPram *apiInputReader.Request,
	pointTransactionRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
) *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes {

	pointTransactionType := &(*pointTransactionRes.Message.Header)[0].PointTransactionType

	var inputPointTransactionType *string

	if pointTransactionType != nil {
		inputPointTransactionType = pointTransactionType
	}

	input := apiModuleRuntimesRequestsPointTransactionType.PointTransactionType{
		PointTransactionType: *inputPointTransactionType,
	}

	responseJsonData := apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes{}
	responseBody := apiModuleRuntimesRequestsPointTransactionType.PointTransactionTypeReadsText(
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
		controller.CustomLogger.Error("CreatePointTransactionTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

// if SenderObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) createShopRequestHeaderSenderObject(
	requestPram *apiInputReader.Request,
	pointTransactionRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	input := make([]apiModuleRuntimesRequestsShop.Header, len(*pointTransactionRes.Message.Header))

	for _, v := range *pointTransactionRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsShop.Header{
			Shop:	v.SenderObject,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsHeader(
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
		controller.CustomLogger.Error("createShopRequestHeaderSenderObject Unmarshal error")
	}

	return &responseJsonData
}

// if SenderObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) createShopRequestAddressSenderObject(
	requestPram *apiInputReader.Request,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	input := make([]apiModuleRuntimesRequestsShop.Address, len(*shopHeaderRes.Message.Header))

	for _, v := range *shopHeaderRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsShop.Address{
			Shop:	v.Shop,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsAddresses(
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
		controller.CustomLogger.Error("createShopRequestAddressSenderObject Unmarshal error")
	}

	return &responseJsonData
}

// if ReceiverObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) createShopRequestHeaderReceiverObject(
	requestPram *apiInputReader.Request,
	pointTransactionRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	input := make([]apiModuleRuntimesRequestsShop.Header, len(*pointTransactionRes.Message.Header))

	for _, v := range *pointTransactionRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsShop.Header{
			Shop:	v.ReceiverObject,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsHeader(
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
		controller.CustomLogger.Error("createShopRequestHeaderReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

// if ReceiverObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) createShopRequestAddressReceiverObject(
	requestPram *apiInputReader.Request,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	input := make([]apiModuleRuntimesRequestsShop.Address, len(*shopHeaderRes.Message.Header))

	for _, v := range *shopHeaderRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsShop.Address{
			Shop:	v.Shop,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsAddresses(
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
		controller.CustomLogger.Error("createShopRequestAddressReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) request(
	input apiInputReader.PointTransaction,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createPointTransactionRequestHeader(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&headerRes,
	)

	pointTransactionTypeTextRes := controller.CreatePointTransactionTypeRequestText(
		controller.UserInfo,
		&headerRes,
	)
	
	shopHeaderResSenderObject := *controller.createShopRequestHeaderSenderObject(
		controller.UserInfo,
		&headerRes,
	)

	shopAddressResSenderObject := *controller.createShopRequestAddressSenderObject(
		controller.UserInfo,
		&shopHeaderResSenderObject,
	)

	shopHeaderResReceiverObject := *controller.createShopRequestHeaderReceiverObject(
		controller.UserInfo,
		&headerRes,
	)
	
	shopAddressResReceiverObject := *controller.createShopRequestAddressReceiverObject(
		controller.UserInfo,
		&shopHeaderResReceiverObject,
	)

	controller.fin(
		&headerRes,
		&businessPartnerRes,
		pointTransactionTypeTextRes,
		shopHeaderResSenderObject,
		shopAddressResSenderObject,
		shopHeaderResReceiverObject,
		shopAddressResReceiverObject,
	)
}

func (
	controller *PointTransactionSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	pointTransactionTypeTextRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes,
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopAddressResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopAddressResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) {
	//businessPartnerMapper := services.BusinessPartnerNameMapper(
	//	businessPartnerRes,
	//)

	pointTransactionTypeTextMapper := services.PointTransactionTypeTextMapper(
		pointTransactionTypeTextRes.Message.Text,
	)

	data := apiOutputFormatter.PointTransaction{}

	for _, v := range *headerRes.Message.Header {
		//qrcode := services.CreateQRCodePointTransactionDocImage(
		//	headerDocRes,
		//	v.PointTransaction,
		//)

		data.PointTransactionHeader = append(data.PointTransactionHeader,
			apiOutputFormatter.PointTransactionHeader{
				PointTransaction:         				v.PointTransaction,
				PointTransactionType:     				v.PointTransactionType,
				PointTransactionTypeName: 				pointTransactionTypeTextMapper[v.PointTransactionType].PointTransactionTypeName,
				PointTransactionDate:     				v.PointTransactionDate,
				PointTransactionTime:     				v.PointTransactionTime,
				SenderObjectType:          				v.SenderObjectType,
				SenderObject:          					v.SenderObject,
				//SenderName:							v.SenderName,  // func, Mapper対応
				ReceiverObjectType:						v.ReceiverObjectType,
				ReceiverObject:							v.ReceiverObject,
				//ReceiverName:			  				v.ReceiverName,  // func, Mapper対応
				PlusMinus:                  			v.PlusMinus,
				PointTransactionAmount:					v.PointTransactionAmount,
				PointTransactionObjectType: 			v.PointTransactionObjectType,
				//PointTransactionObjectTypeName:		v.PointTransactionObjectTypeName, // func, Mapper対応
				PointTransactionObject:               	v.PointTransactionObject,
				SenderPointBalanceAfterTransaction:   	v.SenderPointBalanceAfterTransaction,
				ReceiverPointBalanceAfterTransaction: 	v.ReceiverPointBalanceAfterTransaction,
				Attendance:								v.Attendance,
				Participation:							v.Participation,
				IsCancelled:                          	v.IsCancelled,
				Images:		apiOutputFormatter.Images{
					//PointTransaction: img,
					//QRCode:           qrcode,
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
