package controllersPointTransactionSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsObjectType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/object-type"
	apiModuleRuntimesRequestsPointTransaction "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction"
	apiModuleRuntimesRequestsPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction-type"
	apiModuleRuntimesRequestsShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop-type"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesObjectType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/object-type"
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
	"sync"
)

type PointTransactionSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PointTransactionSingleUnit struct {
	PointTransactionHeader    []apiOutputFormatter.PointTransactionHeader `json:"PointTransactionHeader"`
	ShopHeaderSenderObject    []apiOutputFormatter.ShopHeader             `json:"ShopHeaderSenderObject"`
	ShopAddressSenderObject   []apiOutputFormatter.ShopAddress            `json:"ShopAddressSenderObject"`
	ShopHeaderReceiverObject  []apiOutputFormatter.ShopHeader             `json:"ShopHeaderReceiverObject"`
	ShopAddressReceiverObject []apiOutputFormatter.ShopAddress            `json:"ShopAddressReceiverObject"`
}

func (controller *PointTransactionSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "point-transaction"
	redisKeyCategory2 := "single-unit"
	redisKeyCategory3 := *controller.UserInfo.BusinessPartner
	pointTransaction, _ := controller.GetInt("pointTransaction")

	PointTransactionSingleUnitPointTransaction := apiInputReader.PointTransaction{}

	PointTransactionSingleUnitPointTransaction = apiInputReader.PointTransaction{
		PointTransactionHeader: &apiInputReader.PointTransactionHeader{
			PointTransaction: pointTransaction,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
			strconv.Itoa(pointTransaction),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData PointTransactionSingleUnit

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
			controller.request(PointTransactionSingleUnitPointTransaction)
		}()
	} else {
		controller.request(PointTransactionSingleUnitPointTransaction)
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

func (
	controller *PointTransactionSingleUnitController,
) CreatePointTransactionObjectTypeRequestText(
	requestPram *apiInputReader.Request,
	pointTransactionRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
) *apiModuleRuntimesResponsesObjectType.ObjectTypeRes {

	objectType := &(*pointTransactionRes.Message.Header)[0].PointTransactionObjectType

	var inputObjectType *string

	if objectType != nil {
		inputObjectType = objectType
	}

	input := apiModuleRuntimesRequestsObjectType.ObjectType{
		ObjectType: *inputObjectType,
	}

	responseJsonData := apiModuleRuntimesResponsesObjectType.ObjectTypeRes{}
	responseBody := apiModuleRuntimesRequestsObjectType.ObjectTypeReadsText(
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
		controller.CustomLogger.Error("CreatePointTransactionObjectTypeRequestText Unmarshal error")
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
	var input = apiModuleRuntimesRequestsShop.Header{}

	for _, v := range *pointTransactionRes.Message.Header {
		input = apiModuleRuntimesRequestsShop.Header{
			Shop: v.SenderObject,
		}
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
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	var input = apiModuleRuntimesRequestsShop.Header{}

	for _, v := range *shopHeaderResSenderObject.Message.Header {
		input = apiModuleRuntimesRequestsShop.Header{
			Shop: v.Shop,
		}
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

// if SenderObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) CreateShopTypeRequestTextSenderObject(
	requestPram *apiInputReader.Request,
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShopType.ShopTypeRes {
	if shopHeaderResSenderObject == nil ||
		shopHeaderResSenderObject.Message.Header == nil ||
		len(*shopHeaderResSenderObject.Message.Header) == 0 {
		var shopTypeRes = apiModuleRuntimesResponsesShopType.ShopTypeRes{}
		return &shopTypeRes
	}

	shopType := &(*shopHeaderResSenderObject.Message.Header)[0].ShopType

	var inputShopType *string

	if shopType != nil {
		inputShopType = shopType
	}

	input := apiModuleRuntimesRequestsShopType.ShopType{
		ShopType: *inputShopType,
	}

	responseJsonData := apiModuleRuntimesResponsesShopType.ShopTypeRes{}
	responseBody := apiModuleRuntimesRequestsShopType.ShopTypeReadsText(
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
		controller.CustomLogger.Error("CreateShopTypeRequestTextSenderObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateLocalSubRegionRequestTextSenderObject(
	requestPram *apiInputReader.Request,
	shopAddressResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*shopAddressResSenderObject.Message.Address)[0].LocalSubRegion
	localRegion := &(*shopAddressResSenderObject.Message.Address)[0].LocalRegion
	country := &(*shopAddressResSenderObject.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestTextSenderObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateLocalRegionRequestTextSenderObject(
	requestPram *apiInputReader.Request,
	shopAddressResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*shopAddressResSenderObject.Message.Address)[0].LocalRegion
	country := &(*shopAddressResSenderObject.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalRegionRequestTextSenderObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) createBusinessPartnerRequestSenderObject(
	requestPram *apiInputReader.Request,
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*shopHeaderResSenderObject.Message.Header))

	for _, v := range *shopHeaderResSenderObject.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ShopOwner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestSenderObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateBusinessPartnerRoleRequestTextSenderObject(
	requestPram *apiInputReader.Request,
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {

	businessPartnerRole := &(*shopHeaderResSenderObject.Message.Header)[0].ShopOwnerBusinessPartnerRole

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
		controller.CustomLogger.Error("CreateBusinessPartnerRoleRequestTextSenderObject Unmarshal error")
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
	var input = apiModuleRuntimesRequestsShop.Header{}

	for _, v := range *pointTransactionRes.Message.Header {
		input = apiModuleRuntimesRequestsShop.Header{
			Shop: v.ReceiverObject,
		}
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
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopRes {
	var input = apiModuleRuntimesRequestsShop.Header{}

	for _, v := range *shopHeaderResReceiverObject.Message.Header {
		input = apiModuleRuntimesRequestsShop.Header{
			Shop: v.Shop,
		}
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

// if ReceiverObjectType = "SHOP"

func (
	controller *PointTransactionSingleUnitController,
) CreateShopTypeRequestTextReceiverObject(
	requestPram *apiInputReader.Request,
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShopType.ShopTypeRes {
	if shopHeaderResReceiverObject == nil ||
		shopHeaderResReceiverObject.Message.Header == nil ||
		len(*shopHeaderResReceiverObject.Message.Header) == 0 {
		var shopTypeRes = apiModuleRuntimesResponsesShopType.ShopTypeRes{}
		return &shopTypeRes
	}

	shopType := &(*shopHeaderResReceiverObject.Message.Header)[0].ShopType

	var inputShopType *string

	if shopType != nil {
		inputShopType = shopType
	}

	input := apiModuleRuntimesRequestsShopType.ShopType{
		ShopType: *inputShopType,
	}

	responseJsonData := apiModuleRuntimesResponsesShopType.ShopTypeRes{}
	responseBody := apiModuleRuntimesRequestsShopType.ShopTypeReadsText(
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
		controller.CustomLogger.Error("CreateShopTypeRequestTextReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateLocalSubRegionRequestTextReceiverObject(
	requestPram *apiInputReader.Request,
	shopAddressResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*shopAddressResReceiverObject.Message.Address)[0].LocalSubRegion
	localRegion := &(*shopAddressResReceiverObject.Message.Address)[0].LocalRegion
	country := &(*shopAddressResReceiverObject.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestTextReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateLocalRegionRequestTextReceiverObject(
	requestPram *apiInputReader.Request,
	shopAddressResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*shopAddressResReceiverObject.Message.Address)[0].LocalRegion
	country := &(*shopAddressResReceiverObject.Message.Address)[0].Country

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
		controller.CustomLogger.Error("CreateLocalRegionRequestTextReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) createBusinessPartnerRequestReceiverObject(
	requestPram *apiInputReader.Request,
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*shopHeaderResReceiverObject.Message.Header))

	for _, v := range *shopHeaderResReceiverObject.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ShopOwner,
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
		controller.CustomLogger.Error("createBusinessPartnerRequestReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) CreateBusinessPartnerRoleRequestTextReceiverObject(
	requestPram *apiInputReader.Request,
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {

	businessPartnerRole := &(*shopHeaderResReceiverObject.Message.Header)[0].ShopOwnerBusinessPartnerRole

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
		controller.CustomLogger.Error("CreateBusinessPartnerRoleRequestTextReceiverObject Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionSingleUnitController,
) request(
	input apiInputReader.PointTransaction,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(4)

	var pointTransactionTypeTextRes apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes
	var pointTransactionObjectTypeTextRes apiModuleRuntimesResponsesObjectType.ObjectTypeRes

	var shopHeaderResSenderObject apiModuleRuntimesResponsesShop.ShopRes
	var shopAddressResSenderObject apiModuleRuntimesResponsesShop.ShopRes
	var shopTypeTextResSenderObject apiModuleRuntimesResponsesShopType.ShopTypeRes

	var shopHeaderResReceiverObject apiModuleRuntimesResponsesShop.ShopRes
	var shopAddressResReceiverObject apiModuleRuntimesResponsesShop.ShopRes
	var shopTypeTextResReceiverObject apiModuleRuntimesResponsesShopType.ShopTypeRes

	var localSubRegionTextResSenderObject *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextResSenderObject *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var localSubRegionTextResReceiverObject *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextResReceiverObject *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var businessPartnerResSenderObject apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerResReceiverObject apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes

	var businessPartnerRoleTextResSenderObject apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes
	var businessPartnerRoleTextResReceiverObject apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	headerRes := *controller.createPointTransactionRequestHeader(
		controller.UserInfo,
		input,
	)

	go func() {
		defer wg.Done()
		pointTransactionTypeTextRes = *controller.CreatePointTransactionTypeRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		pointTransactionObjectTypeTextRes = *controller.CreatePointTransactionObjectTypeRequestText(
			controller.UserInfo,
			&headerRes,
		)
	}()

	go func() {
		defer wg.Done()
		if headerRes.Message.Header != nil && len(*headerRes.Message.Header) > 0 {
			for _, v := range *headerRes.Message.Header {
				if v.SenderObjectType == "SHOP" {
					shopHeaderResSenderObject = *controller.createShopRequestHeaderSenderObject(
						controller.UserInfo,
						&headerRes,
					)
					shopAddressResSenderObject = *controller.createShopRequestAddressSenderObject(
						controller.UserInfo,
						&shopHeaderResSenderObject,
					)
					shopTypeTextResSenderObject = *controller.CreateShopTypeRequestTextSenderObject(
						controller.UserInfo,
						&shopHeaderResSenderObject,
					)
					localSubRegionTextResSenderObject = controller.CreateLocalSubRegionRequestTextSenderObject(
						controller.UserInfo,
						&shopAddressResSenderObject,
					)
					localRegionTextResSenderObject = controller.CreateLocalRegionRequestTextSenderObject(
						controller.UserInfo,
						&shopAddressResSenderObject,
					)
					businessPartnerResSenderObject = *controller.createBusinessPartnerRequestSenderObject(
						controller.UserInfo,
						&shopHeaderResSenderObject,
					)
					businessPartnerRoleTextResSenderObject = *controller.CreateBusinessPartnerRoleRequestTextSenderObject(
						controller.UserInfo,
						&shopHeaderResSenderObject,
					)
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		if headerRes.Message.Header != nil && len(*headerRes.Message.Header) > 0 {
			for _, v := range *headerRes.Message.Header {
				if v.ReceiverObjectType == "SHOP" {
					shopHeaderResReceiverObject = *controller.createShopRequestHeaderReceiverObject(
						controller.UserInfo,
						&headerRes,
					)
					shopAddressResReceiverObject = *controller.createShopRequestAddressReceiverObject(
						controller.UserInfo,
						&shopHeaderResReceiverObject,
					)
					shopTypeTextResReceiverObject = *controller.CreateShopTypeRequestTextReceiverObject(
						controller.UserInfo,
						&shopHeaderResReceiverObject,
					)
					localSubRegionTextResReceiverObject = controller.CreateLocalSubRegionRequestTextReceiverObject(
						controller.UserInfo,
						&shopAddressResReceiverObject,
					)
					localRegionTextResReceiverObject = controller.CreateLocalRegionRequestTextReceiverObject(
						controller.UserInfo,
						&shopAddressResReceiverObject,
					)
					businessPartnerResReceiverObject = *controller.createBusinessPartnerRequestReceiverObject(
						controller.UserInfo,
						&shopHeaderResReceiverObject,
					)
					businessPartnerRoleTextResReceiverObject = *controller.CreateBusinessPartnerRoleRequestTextReceiverObject(
						controller.UserInfo,
						&shopHeaderResReceiverObject,
					)
				}
			}
		}
	}()

	wg.Wait()

	controller.fin(
		&headerRes,
		&pointTransactionTypeTextRes,
		&pointTransactionObjectTypeTextRes,
		&shopHeaderResSenderObject,
		&shopAddressResSenderObject,
		&shopTypeTextResSenderObject,
		localSubRegionTextResSenderObject,
		localRegionTextResSenderObject,
		&businessPartnerResSenderObject,
		&businessPartnerRoleTextResSenderObject,
		&shopHeaderResReceiverObject,
		&shopAddressResReceiverObject,
		&shopTypeTextResReceiverObject,
		localSubRegionTextResReceiverObject,
		localRegionTextResReceiverObject,
		&businessPartnerResReceiverObject,
		&businessPartnerRoleTextResReceiverObject,
	)
}

func (
	controller *PointTransactionSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
	pointTransactionTypeTextRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes,
	pointTransactionObjectTypeTextRes *apiModuleRuntimesResponsesObjectType.ObjectTypeRes,
	shopHeaderResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopAddressResSenderObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopTypeTextResSenderObject *apiModuleRuntimesResponsesShopType.ShopTypeRes,
	localSubRegionTextResSenderObject *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextResSenderObject *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	businessPartnerResSenderObject *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextResSenderObject *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	shopHeaderResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopAddressResReceiverObject *apiModuleRuntimesResponsesShop.ShopRes,
	shopTypeTextResReceiverObject *apiModuleRuntimesResponsesShopType.ShopTypeRes,
	localSubRegionTextResReceiverObject *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextResReceiverObject *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	businessPartnerResReceiverObject *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextResReceiverObject *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
) {

	pointTransactionTypeTextMapper := services.PointTransactionTypeTextMapper(
		pointTransactionTypeTextRes.Message.Text,
	)

	pointTransactionObjectTypeTextMapper := services.ObjectTypeTextMapper(
		pointTransactionObjectTypeTextRes.Message.Text,
	)

	var shopTypeTextMapperSenderObject = map[string]apiModuleRuntimesResponsesShopType.Text{}
	var localSubRegionTextMapperSenderObject = map[string]apiModuleRuntimesResponsesLocalSubRegion.Text{}
	var localRegionTextMapperSenderObject = map[string]apiModuleRuntimesResponsesLocalRegion.Text{}
	var businessPartnerMapperSenderObject = map[int]apiModuleRuntimesResponsesBusinessPartner.General{}
	var businessPartnerRoleTextMapperSenderObject = map[string]apiModuleRuntimesResponsesBusinessPartnerRole.Text{}

	var shopTypeTextMapperReceiverObject = map[string]apiModuleRuntimesResponsesShopType.Text{}
	var localSubRegionTextMapperReceiverObject = map[string]apiModuleRuntimesResponsesLocalSubRegion.Text{}
	var localRegionTextMapperReceiverObject = map[string]apiModuleRuntimesResponsesLocalRegion.Text{}
	var businessPartnerMapperReceiverObject = map[int]apiModuleRuntimesResponsesBusinessPartner.General{}
	var businessPartnerRoleTextMapperReceiverObject = map[string]apiModuleRuntimesResponsesBusinessPartnerRole.Text{}

	if shopTypeTextResSenderObject != nil && shopTypeTextResSenderObject.Message.Text != nil {
		shopTypeTextMapperSenderObject = services.ShopTypeTextMapper(
			shopTypeTextResSenderObject.Message.Text,
		)
	}

	if localSubRegionTextResSenderObject != nil && localSubRegionTextResSenderObject.Message.Text != nil {
		localSubRegionTextMapperSenderObject = services.LocalSubRegionTextMapper(
			localSubRegionTextResSenderObject.Message.Text,
		)
	}

	if localRegionTextResSenderObject != nil && localRegionTextResSenderObject.Message.Text != nil {
		localRegionTextMapperSenderObject = services.LocalRegionTextMapper(
			localRegionTextResSenderObject.Message.Text,
		)
	}

	if businessPartnerResSenderObject != nil && businessPartnerResSenderObject.Message.General != nil {
		businessPartnerMapperSenderObject = services.BusinessPartnerNameMapper(
			businessPartnerResSenderObject,
		)
	}

	if businessPartnerRoleTextResSenderObject != nil && businessPartnerRoleTextResSenderObject.Message.Text != nil {
		businessPartnerRoleTextMapperSenderObject = services.BusinessPartnerRoleTextMapper(
			businessPartnerRoleTextResSenderObject.Message.Text,
		)
	}

	if shopTypeTextResReceiverObject != nil && shopTypeTextResReceiverObject.Message.Text != nil {
		shopTypeTextMapperReceiverObject = services.ShopTypeTextMapper(
			shopTypeTextResReceiverObject.Message.Text,
		)
	}

	if localSubRegionTextResReceiverObject != nil && localSubRegionTextResReceiverObject.Message.Text != nil {
		localSubRegionTextMapperReceiverObject = services.LocalSubRegionTextMapper(
			localSubRegionTextResReceiverObject.Message.Text,
		)
	}

	if localRegionTextResReceiverObject != nil && localRegionTextResReceiverObject.Message.Text != nil {
		localRegionTextMapperReceiverObject = services.LocalRegionTextMapper(
			localRegionTextResReceiverObject.Message.Text,
		)
	}

	if businessPartnerResReceiverObject != nil && businessPartnerResReceiverObject.Message.General != nil {
		businessPartnerMapperReceiverObject = services.BusinessPartnerNameMapper(
			businessPartnerResReceiverObject,
		)
	}

	if businessPartnerRoleTextResReceiverObject != nil && businessPartnerRoleTextResReceiverObject.Message.Text != nil {
		businessPartnerRoleTextMapperReceiverObject = services.BusinessPartnerRoleTextMapper(
			businessPartnerRoleTextResReceiverObject.Message.Text,
		)
	}

	data := PointTransactionSingleUnit{}

	for _, v := range *headerRes.Message.Header {
		//qrcode := services.CreateQRCodePointTransactionDocImage(
		//	headerDocRes,
		//	v.PointTransaction,
		//)

		var senderObjectBusinessPartnerName string
		var receiverObjectBusinessPartnerName string

		var senderObjectBusinessPartnerRoleName string
		var receiverObjectBusinessPartnerRoleName string

		if businessPartnerMapperSenderObject != nil && shopHeaderResSenderObject.Message.Header != nil {
			for _, s := range *shopHeaderResSenderObject.Message.Header {
				senderObjectBusinessPartnerName = businessPartnerMapperSenderObject[s.ShopOwner].BusinessPartnerName
				senderObjectBusinessPartnerRoleName = businessPartnerRoleTextMapperSenderObject[s.ShopOwnerBusinessPartnerRole].BusinessPartnerRoleName
			}
		} else {
			senderObjectBusinessPartnerName = ""
			senderObjectBusinessPartnerRoleName = ""
		}

		if businessPartnerMapperReceiverObject != nil && shopHeaderResReceiverObject.Message.Header != nil {
			for _, s := range *shopHeaderResReceiverObject.Message.Header {
				receiverObjectBusinessPartnerName = businessPartnerMapperReceiverObject[s.ShopOwner].BusinessPartnerName
				receiverObjectBusinessPartnerRoleName = businessPartnerRoleTextMapperReceiverObject[s.ShopOwnerBusinessPartnerRole].BusinessPartnerRoleName
			}
		} else {
			receiverObjectBusinessPartnerName = ""
			receiverObjectBusinessPartnerRoleName = ""
		}

		data.PointTransactionHeader = append(data.PointTransactionHeader,
			apiOutputFormatter.PointTransactionHeader{
				PointTransaction:                      v.PointTransaction,
				PointTransactionType:                  v.PointTransactionType,
				PointTransactionTypeName:              pointTransactionTypeTextMapper[v.PointTransactionType].PointTransactionTypeName,
				PointTransactionDate:                  v.PointTransactionDate,
				PointTransactionTime:                  v.PointTransactionTime,
				SenderObjectType:                      v.SenderObjectType,
				SenderObject:                          v.SenderObject,
				SenderObjectBusinessPartnerName:       senderObjectBusinessPartnerName,
				SenderObjectBusinessPartnerRoleName:   senderObjectBusinessPartnerRoleName,
				ReceiverObjectType:                    v.ReceiverObjectType,
				ReceiverObject:                        v.ReceiverObject,
				ReceiverObjectBusinessPartnerName:     receiverObjectBusinessPartnerName,
				ReceiverObjectBusinessPartnerRoleName: receiverObjectBusinessPartnerRoleName,
				PlusMinus:                             v.PlusMinus,
				PointTransactionAmount:                v.PointTransactionAmount,
				PointTransactionObjectType:            v.PointTransactionObjectType,
				PointTransactionObjectTypeName:        pointTransactionObjectTypeTextMapper[v.PointTransactionObjectType].ObjectTypeName,
				PointTransactionObject:                v.PointTransactionObject,
				SenderPointBalanceAfterTransaction:    v.SenderPointBalanceAfterTransaction,
				ReceiverPointBalanceAfterTransaction:  v.ReceiverPointBalanceAfterTransaction,
				Attendance:                            v.Attendance,
				Participation:                         v.Participation,
				Invitation:                            v.Invitation,
				ValidityStartDate:                     v.ValidityStartDate,
				ValidityEndDate:                       v.ValidityEndDate,
				IsCancelled:                           v.IsCancelled,
				Images:                                apiOutputFormatter.Images{
					//PointTransaction: img,
					//QRCode:           qrcode,
				},
			},
		)
	}

	if shopHeaderResSenderObject != nil &&
		shopHeaderResSenderObject.Message.Header != nil {
		for _, v := range *shopHeaderResSenderObject.Message.Header {
			var shopTypeName string
			if shopTypeTextMapperSenderObject != nil {
				shopTypeName = shopTypeTextMapperSenderObject[v.ShopType].ShopTypeName
			} else {
				shopTypeName = ""
			}

			data.ShopHeaderSenderObject = append(data.ShopHeaderSenderObject,
				apiOutputFormatter.ShopHeader{
					Shop:                         v.Shop,
					ShopType:                     v.ShopType,
					ShopTypeName:                 shopTypeName,
					ShopOwner:                    v.ShopOwner,
					ShopOwnerBusinessPartnerRole: v.ShopOwnerBusinessPartnerRole,
					Description:                  v.Description,
				},
			)
		}
	}

	if shopAddressResSenderObject != nil &&
		shopAddressResSenderObject.Message.Address != nil {
		for _, v := range *shopAddressResSenderObject.Message.Address {

			var localSubRegionName string
			if localSubRegionTextMapperSenderObject != nil {
				localSubRegionName = localSubRegionTextMapperSenderObject[v.LocalSubRegion].LocalSubRegionName
			} else {
				localSubRegionName = ""
			}

			var localRegionName string
			if localRegionTextMapperSenderObject != nil {
				localRegionName = localRegionTextMapperSenderObject[v.LocalRegion].LocalRegionName
			} else {
				localRegionName = ""
			}

			data.ShopAddressSenderObject = append(data.ShopAddressSenderObject,
				apiOutputFormatter.ShopAddress{
					Shop:               v.Shop,
					AddressID:          v.AddressID,
					LocalSubRegion:     v.LocalSubRegion,
					LocalSubRegionName: localSubRegionName,
					LocalRegion:        v.LocalRegion,
					LocalRegionName:    localRegionName,
					PostalCode:         v.PostalCode,
					StreetName:         v.StreetName,
					CityName:           v.CityName,
					Building:           v.Building,
				},
			)
		}
	}

	if shopHeaderResReceiverObject != nil &&
		shopHeaderResReceiverObject.Message.Header != nil {
		for _, v := range *shopHeaderResReceiverObject.Message.Header {
			var shopTypeName string
			if shopTypeTextMapperReceiverObject != nil {
				shopTypeName = shopTypeTextMapperReceiverObject[v.ShopType].ShopTypeName
			} else {
				shopTypeName = ""
			}

			data.ShopHeaderReceiverObject = append(data.ShopHeaderReceiverObject,
				apiOutputFormatter.ShopHeader{
					Shop:                         v.Shop,
					ShopType:                     v.ShopType,
					ShopTypeName:                 shopTypeName,
					ShopOwner:                    v.ShopOwner,
					ShopOwnerBusinessPartnerRole: v.ShopOwnerBusinessPartnerRole,
					Description:                  v.Description,
				},
			)
		}
	}

	if shopAddressResReceiverObject != nil &&
		shopAddressResReceiverObject.Message.Address != nil {
		for _, v := range *shopAddressResReceiverObject.Message.Address {

			var localSubRegionName string
			if localSubRegionTextMapperReceiverObject != nil {
				localSubRegionName = localSubRegionTextMapperReceiverObject[v.LocalSubRegion].LocalSubRegionName
			} else {
				localSubRegionName = ""
			}

			var localRegionName string
			if localRegionTextMapperReceiverObject != nil {
				localRegionName = localRegionTextMapperReceiverObject[v.LocalRegion].LocalRegionName
			} else {
				localRegionName = ""
			}

			data.ShopAddressReceiverObject = append(data.ShopAddressReceiverObject,
				apiOutputFormatter.ShopAddress{
					Shop:               v.Shop,
					AddressID:          v.AddressID,
					LocalSubRegion:     v.LocalSubRegion,
					LocalSubRegionName: localSubRegionName,
					LocalRegion:        v.LocalRegion,
					LocalRegionName:    localRegionName,
					PostalCode:         v.PostalCode,
					StreetName:         v.StreetName,
					CityName:           v.CityName,
					Building:           v.Building,
				},
			)
		}
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
