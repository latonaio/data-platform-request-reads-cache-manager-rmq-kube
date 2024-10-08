package controllersAttendanceList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsAttendance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/attendance/attendance"
	apiModuleRuntimesRequestsAttendanceDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/attendance/attendance-doc"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesResponsesAttendance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/attendance"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type AttendanceListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *AttendanceListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	attender, _ := controller.GetInt("attender")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "attender"
	redisKeyCategory2 := "attendance-list"

	isReleased := true
	isCancelled := false
	isMarkedForDeletion := false

	Attendance := apiInputReader.Attendance{
		AttendanceHeader: &apiInputReader.AttendanceHeader{
			Attender: &attender,
		},
		AttendanceDocHeaderDoc: &apiInputReader.AttendanceDocHeaderDoc{},
	}

	Event := apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{},
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
	controller *AttendanceListController,
) createAttendanceRequestHeadersByAttender(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
) *apiModuleRuntimesResponsesAttendance.AttendanceRes {
	responseJsonData := apiModuleRuntimesResponsesAttendance.AttendanceRes{}
	responseBody := apiModuleRuntimesRequestsAttendance.AttendanceReadsHeadersByAttender(
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
		controller.CustomLogger.Error("createAttendanceRequestHeadersByAttender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceListController,
) createAttendanceDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Attendance,
) *apiModuleRuntimesResponsesAttendance.AttendanceDocRes {
	responseJsonData := apiModuleRuntimesResponsesAttendance.AttendanceDocRes{}
	responseBody := apiModuleRuntimesRequestsAttendanceDoc.AttendanceDocReads(
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
			"現地参加権利ヘッダに画像が見つかりませんでした",
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
		controller.CustomLogger.Error("createAttendanceDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceListController,
) CreateEventRequestHeader(
	requestPram *apiInputReader.Request,
	attendanceRes *apiModuleRuntimesResponsesAttendance.AttendanceRes,
	inputEvent apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	var input = apiInputReader.Event{}

	for _, v := range *attendanceRes.Message.Header {
		if v.AttendanceObjectType == "EVENT" {
			input = apiInputReader.Event{
				EventHeader: &apiInputReader.EventHeader{
					IsReleased:          inputEvent.EventHeader.IsReleased,
					IsCancelled:         inputEvent.EventHeader.IsCancelled,
					IsMarkedForDeletion: inputEvent.EventHeader.IsMarkedForDeletion,
				},
			}
		}
	}

	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"Headers",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CreateEventRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceListController,
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
	controller *AttendanceListController,
) CreateSiteRequestHeader(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	isReleased := true
	isMarkedForDeletion := false

	for _, v := range *eventRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site:                v.Site,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReadsHeadersBySites(
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
		controller.CustomLogger.Error("CreateSiteRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceListController,
) request(
	inputAttendance apiInputReader.Attendance,
	inputEvent apiInputReader.Event,
	// isCancelled bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createAttendanceRequestHeadersByAttender(
		controller.UserInfo,
		inputAttendance,
	)

	headerDocRes := controller.createAttendanceDocRequest(
		controller.UserInfo,
		inputAttendance,
	)

	eventHeaderRes := *controller.CreateEventRequestHeader(
		controller.UserInfo,
		&headerRes,
		inputEvent,
	)

	eventHeaderDocRes := controller.createEventDocRequest(
		controller.UserInfo,
		inputEvent,
	)

	siteHeaderRes := *controller.CreateSiteRequestHeader(
		controller.UserInfo,
		&eventHeaderRes,
	)

	controller.fin(
		&headerRes,
		headerDocRes,
		&eventHeaderRes,
		eventHeaderDocRes,
		&siteHeaderRes,
	)
}

func (
	controller *AttendanceListController,
) fin(
	headerRes *apiModuleRuntimesResponsesAttendance.AttendanceRes,
	headerDocRes *apiModuleRuntimesResponsesAttendance.AttendanceDocRes,
	eventHeaderRes *apiModuleRuntimesResponsesEvent.EventRes,
	eventHeaderDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
) {

	eventHeadersMapper := services.EventHeadersMapper(
		eventHeaderRes,
	)

	siteHeadersMapper := services.SiteHeadersMapper(
		siteHeaderRes,
	)

	data := apiOutputFormatter.Attendance{}

	for _, v := range *headerRes.Message.Header {
		event := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].Event
		eventDescription := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].Description
		eventOperationStartDate := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].OperationStartDate
		eventOperationStartTime := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].OperationStartTime
		eventOperationEndDate := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].OperationEndDate
		eventOperationEndTime := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].OperationEndTime
		eventSite := eventHeadersMapper[strconv.Itoa(v.AttendanceObject)].Site

		site := siteHeadersMapper[strconv.Itoa(eventSite)].Site
		siteDescription := siteHeadersMapper[strconv.Itoa(eventSite)].Description

		//img := services.ReadAttendanceImage(
		//	headerDocRes,
		//	v.Attendance,
		//)

		img := services.ReadEventImage(
			eventHeaderDocRes,
			v.AttendanceObject,
		)

		data.AttendanceHeaderWithEvent = append(data.AttendanceHeaderWithEvent,
			apiOutputFormatter.AttendanceHeaderWithEvent{
				Attendance:              v.Attendance,
				Attender:                v.Attender,
				AttendanceObjectType:    v.AttendanceObjectType,
				AttendanceObject:        v.AttendanceObject,
				Participation:           v.Participation,
				Event:                   event,
				EventDescription:        eventDescription,
				EventOperationStartDate: eventOperationStartDate,
				EventOperationStartTime: eventOperationStartTime,
				EventOperationEndDate:   eventOperationEndDate,
				EventOperationEndTime:   eventOperationEndTime,
				EventSite:               eventSite,
				Site:                    site,
				SiteDescription:         siteDescription,

				Images: apiOutputFormatter.Images{
					//Attendance: img,
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
