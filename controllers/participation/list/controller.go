package controllersParticipationList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesRequestsParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/participation/participation"
	apiModuleRuntimesRequestsParticipationDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/participation/participation-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesParticipation "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/participation"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ParticipationListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ParticipationListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	participator, _ := controller.GetInt("participator")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "participator"
	redisKeyCategory2 := "list"
	redisKeyCategory3 := participator

	//	isCancelled := false

	Participation := apiInputReader.Participation{
		ParticipationHeader: &apiInputReader.ParticipationHeader{
			Participator: &participator,
		},
		ParticipationDocHeaderDoc: &apiInputReader.ParticipationDocHeaderDoc{},
	}

	Event := apiInputReader.Event{
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
	controller *ParticipationListController,
) createParticipationRequestHeadersByParticipator(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
) *apiModuleRuntimesResponsesParticipation.ParticipationRes {
	responseJsonData := apiModuleRuntimesResponsesParticipation.ParticipationRes{}
	responseBody := apiModuleRuntimesRequestsParticipation.ParticipationReadsHeadersByParticipator(
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
		controller.CustomLogger.Error("createParticipationRequestHeadersByParticipator Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationListController,
) createParticipationDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
) *apiModuleRuntimesResponsesParticipation.ParticipationDocRes {
	responseJsonData := apiModuleRuntimesResponsesParticipation.ParticipationDocRes{}
	responseBody := apiModuleRuntimesRequestsParticipationDoc.ParticipationDocReads(
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
		controller.CustomLogger.Error("createParticipationDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationListController,
) CreateEventRequestHeader(
	requestPram *apiInputReader.Request,
	participationRes *apiModuleRuntimesResponsesParticipation.ParticipationRes,
) *apiModuleRuntimesResponsesEvent.EventRes {
	var input = apiInputReader.Event{}

	input = apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{},
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
	controller *ParticipationListController,
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
	controller *ParticipationListController,
) CreateSiteRequestHeader(
	requestPram *apiInputReader.Request,
	eventRes *apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	for _, v := range *eventRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site: v.Site,
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
	controller *ParticipationListController,
) request(
	inputParticipation apiInputReader.Participation,
	inputEvent apiInputReader.Event,
	// isCancelled bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := *controller.createParticipationRequestHeadersByParticipator(
		controller.UserInfo,
		inputParticipation,
	)

	headerDocRes := controller.createParticipationDocRequest(
		controller.UserInfo,
		inputParticipation,
	)

	eventHeaderRes := *controller.CreateEventRequestHeader(
		controller.UserInfo,
		&headerRes,
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
	controller *ParticipationListController,
) fin(
	headerRes *apiModuleRuntimesResponsesParticipation.ParticipationRes,
	headerDocRes *apiModuleRuntimesResponsesParticipation.ParticipationDocRes,
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

	data := apiOutputFormatter.Participation{}

	for _, v := range *headerRes.Message.Header {
		event := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].Event
		eventDescription := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].Description
		eventOperationStartDate := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].OperationStartDate
		eventOperationStartTime := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].OperationStartTime
		eventOperationEndDate := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].OperationEndDate
		eventOperationEndTime := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].OperationEndTime
		eventSite := eventHeadersMapper[strconv.Itoa(v.ParticipationObject)].Site

		site := siteHeadersMapper[strconv.Itoa(eventSite)].Site
		siteDescription := siteHeadersMapper[strconv.Itoa(eventSite)].Description

		//img := services.ReadParticipationImage(
		//	headerDocRes,
		//	v.Participation,
		//)

		img := services.ReadEventImage(
			eventHeaderDocRes,
			v.ParticipationObject,
		)

		data.ParticipationHeaderWithEvent = append(data.ParticipationHeaderWithEvent,
			apiOutputFormatter.ParticipationHeaderWithEvent{
				Participation:           v.Participation,
				Participator:            v.Participator,
				ParticipationObjectType: v.ParticipationObjectType,
				ParticipationObject:     v.ParticipationObject,
				Attendance:              v.Attendance,
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
					//Participation: img,
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
