package controllersPointTransactionList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsPointTransaction "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction"
	apiModuleRuntimesRequestsPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-transaction-type"
	apiModuleRuntimesResponsesPointTransaction "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-transaction"
	apiModuleRuntimesResponsesPointTransactionType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-transaction-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type PointTransactionListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PointTransactionList struct {
	PointTransactionHeaderByReceiver []apiOutputFormatter.PointTransactionHeader `json:"PointTransactionHeaderByReceiver"`
	PointTransactionHeaderBySender   []apiOutputFormatter.PointTransactionHeader `json:"PointTransactionHeaderBySender"`
}

func (controller *PointTransactionListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	//	sender, _ := controller.GetInt("sender")
	//	receiver, _ := controller.GetInt("receiver")
	//	pointTransactionType := controller.GetString("pointTransactionType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "point-transaction"
	redisKeyCategory2 := "list"
	redisKeyCategory3 := *controller.UserInfo.BusinessPartner
	//	isCancelled := false

	objectType := "BUSINESS_PARTNER"

	PointTransactionHeader := apiInputReader.PointTransaction{
		PointTransactionHeader: &apiInputReader.PointTransactionHeader{
			SenderObjectType:   &objectType,
			SenderObject:       controller.UserInfo.BusinessPartner,
			ReceiverObjectType: &objectType,
			ReceiverObject:     controller.UserInfo.BusinessPartner,
			//PointTransactionType: &pointTransactionType,
			//IsCancelled:          &isCancelled,
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
		var responseData PointTransactionList

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
			controller.request(PointTransactionHeader)
		}()
	} else {
		controller.request(PointTransactionHeader)
	}
}

func (
	controller *PointTransactionListController,
) createPointTransactionRequestHeadersByReceiver(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
) *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes {
	responseJsonData := apiModuleRuntimesResponsesPointTransaction.PointTransactionRes{}
	responseBody := apiModuleRuntimesRequestsPointTransaction.PointTransactionReadsHeadersByReceiver(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if len(*responseJsonData.Message.Header) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"レシーバに対してのポイント取引ヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createPointTransactionRequestHeadersByReceiver Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionListController,
) createPointTransactionRequestHeadersBySender(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
) *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes {
	responseJsonData := apiModuleRuntimesResponsesPointTransaction.PointTransactionRes{}
	responseBody := apiModuleRuntimesRequestsPointTransaction.PointTransactionReadsHeadersBySender(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if len(*responseJsonData.Message.Header) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"センダに対してのポイント取引ヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createPointTransactionRequestHeadersBySender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionListController,
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
	controller *PointTransactionListController,
) request(
	input apiInputReader.PointTransaction,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var pointTransactionTypeTextByReceiverRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes
	var pointTransactionTypeTextBySenderRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes

	headersByReceiverRes := *controller.createPointTransactionRequestHeadersByReceiver(
		controller.UserInfo,
		input,
	)

	headersBySenderRes := *controller.createPointTransactionRequestHeadersBySender(
		controller.UserInfo,
		input,
	)

	if headersByReceiverRes.Message.Header != nil && len(*headersByReceiverRes.Message.Header) != 0 {
		pointTransactionTypeTextByReceiverRes = controller.CreatePointTransactionTypeRequestText(
			controller.UserInfo,
			&headersByReceiverRes,
		)
	}

	if headersBySenderRes.Message.Header != nil && len(*headersBySenderRes.Message.Header) != 0 {
		pointTransactionTypeTextBySenderRes = controller.CreatePointTransactionTypeRequestText(
			controller.UserInfo,
			&headersBySenderRes,
		)
	}

	controller.fin(
		&headersByReceiverRes,
		&headersBySenderRes,
		pointTransactionTypeTextByReceiverRes,
		pointTransactionTypeTextBySenderRes,
	)
}

func (
	controller *PointTransactionListController,
) fin(
	headersByReceiverRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
	headersBySenderRes *apiModuleRuntimesResponsesPointTransaction.PointTransactionRes,
	pointTransactionTypeTextByReceiverRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes,
	pointTransactionTypeTextBySenderRes *apiModuleRuntimesResponsesPointTransactionType.PointTransactionTypeRes,
) {

	//	pointTransactionHeadersMapper := services.PointTransactionHeadersMapper(
	//		headersByReceiverRes,
	//		headersBySenderRes,
	//	)

	var pointConditionTypeTextMapperByReceiver map[string]apiModuleRuntimesResponsesPointTransactionType.Text
	var pointConditionTypeTextMapperBySender map[string]apiModuleRuntimesResponsesPointTransactionType.Text

	if headersByReceiverRes.Message.Header != nil && len(*headersByReceiverRes.Message.Header) != 0 {
		pointConditionTypeTextMapperByReceiver = services.PointTransactionTypeTextMapper(
			pointTransactionTypeTextByReceiverRes.Message.Text,
		)
	}

	if headersBySenderRes.Message.Header != nil && len(*headersBySenderRes.Message.Header) != 0 {
		pointConditionTypeTextMapperBySender = services.PointTransactionTypeTextMapper(
			pointTransactionTypeTextBySenderRes.Message.Text,
		)
	}

	data := PointTransactionList{}

	for _, v := range *headersByReceiverRes.Message.Header {
		data.PointTransactionHeaderByReceiver = append(data.PointTransactionHeaderByReceiver,
			apiOutputFormatter.PointTransactionHeader{
				PointTransaction:                     v.PointTransaction,
				PointTransactionType:                 v.PointTransactionType,
				PointTransactionTypeName:             pointConditionTypeTextMapperByReceiver[v.PointTransactionType].PointTransactionTypeName,
				PointTransactionDate:                 v.PointTransactionDate,
				PointTransactionTime:                 v.PointTransactionTime,
				PointSymbol:                          v.PointSymbol,
				PlusMinus:                            v.PlusMinus,
				PointTransactionAmount:               v.PointTransactionAmount,
				SenderPointBalanceAfterTransaction:   v.SenderPointBalanceAfterTransaction,
				ReceiverPointBalanceAfterTransaction: v.ReceiverPointBalanceAfterTransaction,
				ValidityStartDate:                    v.ValidityStartDate,
				ValidityEndDate:                      v.ValidityEndDate,
				IsCancelled:                          v.IsCancelled,
			},
		)
	}

	for _, v := range *headersBySenderRes.Message.Header {
		data.PointTransactionHeaderBySender = append(data.PointTransactionHeaderBySender,
			apiOutputFormatter.PointTransactionHeader{
				PointTransaction:                     v.PointTransaction,
				PointTransactionType:                 v.PointTransactionType,
				PointTransactionTypeName:             pointConditionTypeTextMapperBySender[v.PointTransactionType].PointTransactionTypeName,
				PointTransactionDate:                 v.PointTransactionDate,
				PointTransactionTime:                 v.PointTransactionTime,
				PointSymbol:                          v.PointSymbol,
				PlusMinus:                            v.PlusMinus,
				PointTransactionAmount:               v.PointTransactionAmount,
				SenderPointBalanceAfterTransaction:   v.SenderPointBalanceAfterTransaction,
				ReceiverPointBalanceAfterTransaction: v.ReceiverPointBalanceAfterTransaction,
				ValidityStartDate:                    v.ValidityStartDate,
				ValidityEndDate:                      v.ValidityEndDate,
				IsCancelled:                          v.IsCancelled,
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
