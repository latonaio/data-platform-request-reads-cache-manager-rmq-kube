package controllersAttendanceListForOwnersByEvent

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsAttendance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/attendance/attendance"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesResponsesAttendance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/attendance"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type AttendanceListForOwnersByEventController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type AttendanceListForOwnersByEvent struct {
	EventHeader      []apiOutputFormatter.EventHeader      `json:"EventHeader"`
	AttendanceHeader []apiOutputFormatter.AttendanceHeader `json:"AttendanceHeader"`
}

func (controller *AttendanceListForOwnersByEventController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	event, _ := controller.GetInt("event")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "event"
	redisKeyCategory2 := "attendance-list-for-owners"

	isCancelled := false

	Attendance := apiInputReader.Attendance{}
	Event := apiInputReader.Event{}

	attendanceObjectType := "EVENT"

	Attendance = apiInputReader.Attendance{
		AttendanceHeader: &apiInputReader.AttendanceHeader{
			AttendanceObjectType: &attendanceObjectType,
			AttendanceObject:     &event,
			IsCancelled:          &isCancelled,
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
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Attendance

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
			controller.request(Attendance, Event)
		}()
	} else {
		controller.request(Attendance, Event)
	}
}

func (
	controller *AttendanceListForOwnersByEventController,
) createAttendanceRequestHeadersByEvent(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
) *apiModuleRuntimesResponsesAttendance.AttendanceRes {
	responseJsonData := apiModuleRuntimesResponsesAttendance.AttendanceRes{}
	responseBody := apiModuleRuntimesRequestsAttendance.AttendanceReadsHeadersByEvent(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//	if len(*responseJsonData.Message.Address) == 0 {
	//		status := 500
	//		services.HandleError(
	//			&controller.Controller,
	//			"有効な現地参加が見つかりませんでした",
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
		controller.CustomLogger.Error("createAttendanceRequestHeadersByEvent Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceListForOwnersByEventController,
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
	controller *AttendanceListForOwnersByEventController,
) request(
	inputAttendance apiInputReader.Attendance,
	inputEvent apiInputReader.Event,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	attendanceHeaderRes := *controller.createAttendanceRequestHeadersByEvent(
		controller.UserInfo,
		inputAttendance,
	)

	eventHeaderRes := *controller.createEventRequestHeader(
		controller.UserInfo,
		inputEvent,
	)

	controller.fin(
		&attendanceHeaderRes,
		&eventHeaderRes,
	)
}

func (
	controller *AttendanceListForOwnersByEventController,
) fin(
	attendanceHeaderRes *apiModuleRuntimesResponsesAttendance.AttendanceRes,
	eventHeaderRes *apiModuleRuntimesResponsesEvent.EventRes,
) {

	data := AttendanceListForOwnersByEvent{}

	for _, v := range *eventHeaderRes.Message.Header {

		data.EventHeader = append(data.EventHeader,
			apiOutputFormatter.EventHeader{
				Event:       v.Event,
				Description: v.Description,
			},
		)
	}

	for _, v := range *attendanceHeaderRes.Message.Header {

		data.AttendanceHeader = append(data.AttendanceHeader,
			apiOutputFormatter.AttendanceHeader{
				Attendance:           v.Attendance,
				Attender:             v.Attender,
				AttendanceObjectType: v.AttendanceObjectType,
				AttendanceObject:     v.AttendanceObject,
				Participation:        v.Participation,
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
