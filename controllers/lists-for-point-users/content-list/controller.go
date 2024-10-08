package controllersContentListForPointUsers

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ContentListForPointUsersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ContentListForPointUsersController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	localSubRegion := controller.GetString("localSubRegion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "point-transaction"
	redisKeyCategory2 := "content-list"
	redisKeyCategory3 := localSubRegion

	isReleased := true
	isCancelled := false
	isMarkedForDeletion := false

	EventAddress := apiInputReader.Event{
		EventAddress: &apiInputReader.EventAddress{
			LocalSubRegion: &localSubRegion,
		},
		EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{},
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
		var responseData apiOutputFormatter.Event

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
			controller.request(EventAddress, isReleased, isCancelled, isMarkedForDeletion)
		}()
	} else {
		controller.request(EventAddress, isReleased, isCancelled, isMarkedForDeletion)
	}
}

func (
	controller *ContentListForPointUsersController,
) createEventRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"AddressesByLocalSubRegion",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if len(*responseJsonData.Message.Address) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対してのイベントアドレスが見つかりませんでした",
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
		controller.CustomLogger.Error("createEventRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForPointUsersController,
) createEventRequestHeadersByEvents(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
	isReleased bool,
	isCancelled bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesEvent.EventRes {
	var input []apiModuleRuntimesRequestsEvent.Header

	for _, v := range *eventRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsEvent.Header{
			Event:               v.Event,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReadsHeadersByEvents(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if responseJsonData.Message.Header == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対して有効なイベントヘッダデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createEventRequestHeadersByEvents Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForPointUsersController,
) createEventRequestCountersByEvents(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesEvent.EventRes {
	var input []apiModuleRuntimesRequestsEvent.Header

	for _, v := range *eventRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsEvent.Header{
			Event: v.Event,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReadsCountersByEvents(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if responseJsonData.Message.Counter == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対して有効なイベントカウンタデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createEventRequestCountersByEvents Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForPointUsersController,
) createEventDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventDocRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventDocRes{}
	responseBody := apiModuleRuntimesRequestsEventDoc.EventDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if responseJsonData.Message.HeaderDoc == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"イベントヘッダに画像が見つかりませんでした",
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
		controller.CustomLogger.Error("createEventDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForPointUsersController,
) request(
	input apiInputReader.Event,
	isReleased bool,
	isCancelled bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var headerRes *apiModuleRuntimesResponsesEvent.EventRes
	var counterRes *apiModuleRuntimesResponsesEvent.EventRes
	var headerDocRes *apiModuleRuntimesResponsesEvent.EventDocRes

	addressRes := *controller.createEventRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		input,
	)

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
		headerRes = controller.createEventRequestHeadersByEvents(
			controller.UserInfo,
			&addressRes,
			isReleased,
			isCancelled,
			isMarkedForDeletion,
		)

		counterRes = controller.createEventRequestCountersByEvents(
			controller.UserInfo,
			&addressRes,
		)
		headerDocRes = controller.createEventDocRequest(
			controller.UserInfo,
			input,
		)
	}

	controller.fin(
		&addressRes,
		headerRes,
		counterRes,
		headerDocRes,
	)
}

func (
	controller *ContentListForPointUsersController,
) fin(
	addressRes *apiModuleRuntimesResponsesEvent.EventRes,
	headerRes *apiModuleRuntimesResponsesEvent.EventRes,
	counterRes *apiModuleRuntimesResponsesEvent.EventRes,
	headerDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
) {
	data := apiOutputFormatter.Event{}

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
		eventHeadersMapper := services.EventHeadersMapper(
			headerRes,
		)
		eventCountersMapper := services.EventCountersMapper(
			counterRes,
		)

		for _, v := range *addressRes.Message.Address {
			eventType := eventHeadersMapper[strconv.Itoa(v.Event)].EventType
			validityStartDate := eventHeadersMapper[strconv.Itoa(v.Event)].ValidityStartDate
			validityStartTime := eventHeadersMapper[strconv.Itoa(v.Event)].ValidityStartTime
			validityEndDate := eventHeadersMapper[strconv.Itoa(v.Event)].ValidityEndDate
			validityEndTime := eventHeadersMapper[strconv.Itoa(v.Event)].ValidityEndTime
			introduction := eventHeadersMapper[strconv.Itoa(v.Event)].Introduction
			tag1 := eventHeadersMapper[strconv.Itoa(v.Event)].Tag1
			tag2 := eventHeadersMapper[strconv.Itoa(v.Event)].Tag2
			tag3 := eventHeadersMapper[strconv.Itoa(v.Event)].Tag3
			tag4 := eventHeadersMapper[strconv.Itoa(v.Event)].Tag4
			lastChangeDate := eventHeadersMapper[strconv.Itoa(v.Event)].LastChangeDate
			lastChangeTime := eventHeadersMapper[strconv.Itoa(v.Event)].LastChangeTime

			numberOfLikes := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfLikes
			numberOfParticipations := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfParticipations
			numberOfAttendances := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfAttendances

			img := services.ReadEventImage(
				headerDocRes,
				v.Event,
			)

			data.EventAddressWithHeader = append(data.EventAddressWithHeader,
				apiOutputFormatter.EventAddressWithHeader{
					Event:                  v.Event,
					AddressID:              v.AddressID,
					LocalSubRegion:         v.LocalSubRegion,
					LocalRegion:            v.LocalRegion,
					EventType:              eventType,
					ValidityStartDate:      validityStartDate,
					ValidityStartTime:      validityStartTime,
					ValidityEndDate:        validityEndDate,
					ValidityEndTime:        validityEndTime,
					Introduction:           introduction,
					Tag1:                   tag1,
					Tag2:                   tag2,
					Tag3:                   tag3,
					Tag4:                   tag4,
					LastChangeDate:			lastChangeDate,
					LastChangeTime:			lastChangeTime,
					NumberOfLikes:          &numberOfLikes,
					NumberOfParticipations: &numberOfParticipations,
					NumberOfAttendances:    &numberOfAttendances,

					Images: apiOutputFormatter.Images{
						Event: img,
					},
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
