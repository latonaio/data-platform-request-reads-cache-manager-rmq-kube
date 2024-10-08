package controllersParticipationListForOwnersByEvent

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/participation/participation"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/participation"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ParticipationListForOwnersByEventController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ParticipationListForOwnersByEvent struct {
	EventHeader         []apiOutputFormatter.EventHeader         `json:"EventHeader"`
	ParticipationHeader []apiOutputFormatter.ParticipationHeader `json:"ParticipationHeader"`
}

func (controller *ParticipationListForOwnersByEventController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	event, _ := controller.GetInt("event")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "participation"
	redisKeyCategory2 := "list-for-owners-by-event"
	redisKeyCategory3 := event

	isCancelled := false

	Participation := apiInputReader.Participation{}
	Event := apiInputReader.Event{}

	participationObjectType := "EVENT"

	Participation = apiInputReader.Participation{
		ParticipationHeader: &apiInputReader.ParticipationHeader{
			ParticipationObjectType: &participationObjectType,
			ParticipationObject:     &event,
			IsCancelled:             &isCancelled,
		},
	}

	Event = apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			Event: event,
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
		var responseData apiOutputFormatter.Participation

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
			controller.request(Participation, Event)
		}()
	} else {
		controller.request(Participation, Event)
	}
}

func (
	controller *ParticipationListForOwnersByEventController,
) createParticipationRequestHeadersByEvent(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
) *apiModuleRuntimesResponsesParticipation.ParticipationRes {
	responseJsonData := apiModuleRuntimesResponsesParticipation.ParticipationRes{}
	responseBody := apiModuleRuntimesRequestsParticipation.ParticipationReadsHeadersByEvent(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//	if len(*responseJsonData.Message.Address) == 0 {
	//		status := 500
	//		services.HandleError(
	//			&controller.Controller,
	//			"有効な現地参加権利が見つかりませんでした",
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
		controller.CustomLogger.Error("createParticipationRequestHeadersByEvent Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationListForOwnersByEventController,
) createEventRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationListForOwnersByEventController,
) request(
	inputParticipation apiInputReader.Participation,
	inputEvent apiInputReader.Event,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	participationHeaderRes := *controller.createParticipationRequestHeadersByEvent(
		controller.UserInfo,
		inputParticipation,
	)

	eventHeaderRes := *controller.createEventRequestHeader(
		controller.UserInfo,
		inputEvent,
	)

	controller.fin(
		&participationHeaderRes,
		&eventHeaderRes,
	)
}

func (
	controller *ParticipationListForOwnersByEventController,
) fin(
	participationHeaderRes *apiModuleRuntimesResponsesParticipation.ParticipationRes,
	eventHeaderRes *apiModuleRuntimesResponsesEvent.EventRes,
) {

	data := ParticipationListForOwnersByEvent{}

	for _, v := range *eventHeaderRes.Message.Header {

		data.EventHeader = append(data.EventHeader,
			apiOutputFormatter.EventHeader{
				Event:       v.Event,
				Description: v.Description,
			},
		)
	}

	for _, v := range *participationHeaderRes.Message.Header {

		data.ParticipationHeader = append(data.ParticipationHeader,
			apiOutputFormatter.ParticipationHeader{
				Participation:           v.Participation,
				Participator:            v.Participator,
				ParticipationObjectType: v.ParticipationObjectType,
				ParticipationObject:     v.ParticipationObject,
				Attendance:              v.Attendance,
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
