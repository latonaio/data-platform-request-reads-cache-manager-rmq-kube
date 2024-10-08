package controllersContentListForOwners

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

type ContentListForOwnersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ContentListForOwnersController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	organizationBusinessPartner, _ := controller.GetInt("organizationBusinessPartner")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "lists-for-owners"
	redisKeyCategory2 := "content-list"
	redisKeyCategory3 := organizationBusinessPartner

	//	isReleased := true
	//	isCancelled := false
	isMarkedForDeletion := false

	Event := apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			EventOwner:          &organizationBusinessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{},
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
			controller.request(Event)
		}()
	} else {
		controller.request(Event)
	}
}

func (
	controller *ContentListForOwnersController,
) createEventRequestHeadersByEventOwner(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByEventOwner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//	if len(*responseJsonData.Message.Address) == 0 {
	//		status := 500
	//		services.HandleError(
	//			&controller.Controller,
	//			"イベントオーナーに対してのイベントが見つかりませんでした",
	//			&status,
	//		)
	//		return nil
	//	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestHeadersByEventOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForOwnersController,
) createEventRequestCountersByEvents(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesEvent.EventRes {
	var input []apiModuleRuntimesRequestsEvent.Header

	for _, v := range *eventRes.Message.Header {
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

	if responseJsonData.Message.Counter == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"イベントに対して有効なイベントカウンタデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createEventRequestCountersByEvents Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForOwnersController,
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

	if responseJsonData.Message.HeaderDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"イベントヘッダに画像が見つかりませんでした",
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
		controller.CustomLogger.Error("createEventDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ContentListForOwnersController,
) request(
	input apiInputReader.Event,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createEventRequestHeadersByEventOwner(
		controller.UserInfo,
		input,
	)

	counterRes := *controller.createEventRequestCountersByEvents(
		controller.UserInfo,
		&headerRes,
	)

	headerDocRes := controller.createEventDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&headerRes,
		&counterRes,
		headerDocRes,
	)
}

func (
	controller *ContentListForOwnersController,
) fin(
	headerRes *apiModuleRuntimesResponsesEvent.EventRes,
	counterRes *apiModuleRuntimesResponsesEvent.EventRes,
	headerDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
) {

	eventCountersMapper := services.EventCountersMapper(
		counterRes,
	)

	data := apiOutputFormatter.Event{}

	for _, v := range *headerRes.Message.Header {

		numberOfLikes := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfLikes
		numberOfParticipations := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfParticipations
		numberOfAttendances := eventCountersMapper[strconv.Itoa(v.Event)].NumberOfAttendances

		img := services.ReadEventImage(
			headerDocRes,
			v.Event,
		)

		data.EventHeader = append(data.EventHeader,
			apiOutputFormatter.EventHeader{
				Event:                  v.Event,
				Description:            v.Description,
				Introduction:           v.Introduction,
				ValidityStartDate:      v.ValidityStartDate,
				ValidityStartTime:      v.ValidityStartTime,
				ValidityEndDate:        v.ValidityEndDate,
				ValidityEndTime:        v.ValidityEndTime,
				OperationStartDate:     v.OperationStartDate,
				OperationStartTime:     v.OperationStartTime,
				OperationEndDate:       v.OperationEndDate,
				OperationEndTime:       v.OperationEndTime,
				Tag1:                   v.Tag1,
				Tag2:                   v.Tag2,
				Tag3:                   v.Tag3,
				Tag4:                   v.Tag4,
				LastChangeDate:			v.LastChangeDate,
				LastChangeTime:			v.LastChangeTime,
				NumberOfLikes:          &numberOfLikes,
				NumberOfParticipations: &numberOfParticipations,
				NumberOfAttendances:    &numberOfAttendances,

				Images: apiOutputFormatter.Images{
					Event: img,
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
