package controllersMessageList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsMessage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/message"
	apiModuleRuntimesRequestsMessageType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/message-type"
	apiModuleRuntimesResponsesMessage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/message"
	apiModuleRuntimesResponsesMessageType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/message-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type MessageListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *MessageListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	sender, _ := controller.GetInt("sender")
	receiver, _ := controller.GetInt("receiver")
	messageType := controller.GetString("messageType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "sender"
	redisKeyCategory2 := "receiver"
	redisKeyCategory3 := "messageType"

	isMarkedForDeletion := false

	MessageHeader := apiInputReader.Message{
		MessageHeader: &apiInputReader.MessageHeader{
			Sender:              &sender,
			Receiver:            &receiver,
			MessageType:         &messageType,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
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
		var responseData apiOutputFormatter.Message

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
			controller.request(MessageHeader)
		}()
	} else {
		controller.request(MessageHeader)
	}
}

func (
	controller *MessageListController,
) createMessageRequestHeadersByReceiver(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
) *apiModuleRuntimesResponsesMessage.MessageRes {
	responseJsonData := apiModuleRuntimesResponsesMessage.MessageRes{}
	responseBody := apiModuleRuntimesRequestsMessage.MessageReadsHeader(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"レシーバに対してのメッセージヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createMessageRequestHeadersByReceiver Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageListController,
) createMessageRequestHeadersBySender(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
) *apiModuleRuntimesResponsesMessage.MessageRes {
	responseJsonData := apiModuleRuntimesResponsesMessage.MessageRes{}
	responseBody := apiModuleRuntimesRequestsMessage.MessageReadsHeadersBySender(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Header) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"センダに対してのメッセージヘッダが見つかりませんでした",
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
		controller.CustomLogger.Error("createMessageRequestHeadersBySender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageListController,
) CreateMessageTypeRequestText(
	requestPram *apiInputReader.Request,
	messageRes *apiModuleRuntimesResponsesMessage.MessageRes,
) *apiModuleRuntimesResponsesMessageType.MessageTypeRes {

	messageType := &(*messageRes.Message.Header)[0].MessageType

	var inputMessageType *string

	if messageType != nil {
		inputMessageType = messageType
	}

	input := apiModuleRuntimesRequestsMessageType.MessageType{
		MessageType: *inputMessageType,
	}

	responseJsonData := apiModuleRuntimesResponsesMessageType.MessageTypeRes{}
	responseBody := apiModuleRuntimesRequestsMessageType.MessageTypeReadsText(
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
		controller.CustomLogger.Error("CreateMessageTypeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageListController,
) request(
	input apiInputReader.Message,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headersByReceiverRes := *controller.createMessageRequestHeadersByReceiver(
		controller.UserInfo,
		input,
	)

	headersBySenderRes := *controller.createMessageRequestHeadersBySender(
		controller.UserInfo,
		input,
	)

	messageTypeTextByReceiverRes := controller.CreateMessageTypeRequestText(
		controller.UserInfo,
		&headersByReceiverRes,
	)

	messageTypeTextBySenderRes := controller.CreateMessageTypeRequestText(
		controller.UserInfo,
		&headersBySenderRes,
	)

	controller.fin(
		&headersByReceiverRes,
		&headersBySenderRes,
		messageTypeTextByReceiverRes,
		messageTypeTextBySenderRes,
	)
}

func (
	controller *MessageListController,
) fin(
	headersByReceiverRes *apiModuleRuntimesResponsesMessage.MessageRes,
	headersBySenderRes *apiModuleRuntimesResponsesMessage.MessageRes,
	messageTypeTextByReceiverRes *apiModuleRuntimesResponsesMessageType.MessageTypeRes,
	messageTypeTextBySenderRes *apiModuleRuntimesResponsesMessageType.MessageTypeRes,
) {

	//	messageHeadersMapper := services.MessageHeadersMapper(
	//		headersByReceiverRes,
	//		headersBySenderRes,
	//	)

	messageTypeTextMapperByReceiver := services.MessageTypeTextMapper(
		messageTypeTextByReceiverRes.Message.Text,
	)

	messageTypeTextMapperBySender := services.MessageTypeTextMapper(
		messageTypeTextBySenderRes.Message.Text,
	)

	data := apiOutputFormatter.Message{}

	for _, v := range *headersByReceiverRes.Message.Header {
		data.MessageHeader = append(data.MessageHeader,
			apiOutputFormatter.MessageHeader{
				Message:         v.Message,
				MessageType:     v.MessageType,
				MessageTypeName: messageTypeTextMapperByReceiver[v.MessageType].MessageTypeName,
				Title:           v.Title,
				LongText:        v.LongText,
				CreationDate:    v.CreationDate,
				CreationTime:    v.CreationTime,
				LastChangeDate:  v.LastChangeDate,
				LastChangeTime:  v.LastChangeTime,
			},
		)
	}

	for _, v := range *headersBySenderRes.Message.Header {
		data.MessageHeader = append(data.MessageHeader,
			apiOutputFormatter.MessageHeader{
				Message:         v.Message,
				MessageType:     v.MessageType,
				MessageTypeName: messageTypeTextMapperBySender[v.MessageType].MessageTypeName,
				Title:           v.Title,
				LongText:        v.LongText,
				CreationDate:    v.CreationDate,
				CreationTime:    v.CreationTime,
				LastChangeDate:  v.LastChangeDate,
				LastChangeTime:  v.LastChangeTime,
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
