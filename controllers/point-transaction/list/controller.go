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
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "businessPartner"
	//	redisKeyCategory2 := "pointTransactionType"

	//	isCancelled := false

	PointTransactionHeader := apiInputReader.PointTransaction{
		PointTransactionHeader: &apiInputReader.PointTransactionHeader{
			SenderObjectType:   	"BUSINESS_PARTNER",
			SenderObject:   		controller.UserInfo.BusinessPartner,
			ReceiverObjectType:   	"BUSINESS_PARTNER",
			ReceiverObject: 		controller.UserInfo.BusinessPartner,
			//PointTransactionType: &pointTransactionType,
			//IsCancelled:          &isCancelled,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			//			redisKeyCategory2,
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

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"レシーバに対してのポイント取引ヘッダが見つかりませんでした",
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

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"センダに対してのポイント取引ヘッダが見つかりませんでした",
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

	headersByReceiverRes := *controller.createPointTransactionRequestHeadersByReceiver(
		controller.UserInfo,
		input,
	)

	headersBySenderRes := *controller.createPointTransactionRequestHeadersBySender(
		controller.UserInfo,
		input,
	)

	pointTransactionTypeTextByReceiverRes := controller.CreatePointTransactionTypeRequestText(
		controller.UserInfo,
		&headersByReceiverRes,
	)

	pointTransactionTypeTextBySenderRes := controller.CreatePointTransactionTypeRequestText(
		controller.UserInfo,
		&headersBySenderRes,
	)

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

	pointConditionTypeTextMapperByReceiver := services.PointTransactionTypeTextMapper(
		pointTransactionTypeTextByReceiverRes.Message.Text,
	)

	pointConditionTypeTextMapperBySender := services.PointTransactionTypeTextMapper(
		pointTransactionTypeTextBySenderRes.Message.Text,
	)

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
