package controllersAfterPointAcquisition

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesRequestsPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-balance"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-balance"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type AfterPointAcquisitionController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *AfterPointAcquisitionController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "point-acquisition"
	redisKeyCategory2 := "after-point-acquisition"
	businessPartner, _ := controller.GetInt("businessPartner")
	event, _ := controller.GetInt("event")
	pointSymbol := "POYPO"

	PointBalancePointBalance := apiInputReader.PointBalanceGlobal{}

	PointBalancePointBalance = apiInputReader.PointBalanceGlobal{
		PointBalance: &apiInputReader.PointBalance{
			BusinessPartner: businessPartner,
			PointSymbol:     pointSymbol,
		},
	}

	EventSingleUnit := apiInputReader.Event{}

	isReleased := true
	isCancelled := false
	isMarkedForDeletion := false

	//docType := "QRCODE"

	EventSingleUnit = apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			Event:               event,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventPointConditionElement: &apiInputReader.EventPointConditionElement{
			Event:               event,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{
			Event: event,
			//DocType:				     &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(businessPartner),
			strconv.Itoa(event),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.AfterPointAcquisitionGlobal

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
			controller.request(PointBalancePointBalance, EventSingleUnit)
		}()
	} else {
		controller.request(PointBalancePointBalance, EventSingleUnit)
	}
}

func (
	controller *AfterPointAcquisitionController,
) createPointBalanceRequestPointBalance(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
) *apiModuleRuntimesResponsesPointBalance.PointBalanceRes {
	responseJsonData := apiModuleRuntimesResponsesPointBalance.PointBalanceRes{}
	responseBody := apiModuleRuntimesRequestsPointBalance.PointBalanceReadsPointBalance(
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
		controller.CustomLogger.Error("createPointBalanceRequestPointBalance Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterPointAcquisitionController,
) createEventRequestPointConditionElement(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
) *apiModuleRuntimesResponsesEvent.EventRes {
	responseJsonData := apiModuleRuntimesResponsesEvent.EventRes{}
	responseBody := apiModuleRuntimesRequestsEvent.EventReads(
		requestPram,
		input,
		&controller.Controller,
		"PointConditionElements",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEventRequestPointConditionElement Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterPointAcquisitionController,
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
	controller *AfterPointAcquisitionController,
) request(
	pointBalanceGlobal apiInputReader.PointBalanceGlobal,
	eventSingleUnit apiInputReader.Event,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	pointBalanceRes := *controller.createPointBalanceRequestPointBalance(
		controller.UserInfo,
		pointBalanceGlobal,
	)

	eventPointConditionElementRes := *controller.createEventRequestPointConditionElement(
		controller.UserInfo,
		eventSingleUnit,
	)

	eventHeaderDocRes := controller.createEventDocRequest(
		controller.UserInfo,
		eventSingleUnit,
	)

	controller.fin(
		&pointBalanceRes,
		&eventPointConditionElementRes,
		eventHeaderDocRes,
	)
}

func (
	controller *AfterPointAcquisitionController,
) fin(
	pointBalanceRes *apiModuleRuntimesResponsesPointBalance.PointBalanceRes,
	eventPointConditionElementRes *apiModuleRuntimesResponsesEvent.EventRes,
	eventHeaderDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
) {

	data := apiOutputFormatter.AfterPointAcquisitionGlobal{}

	for _, v := range *pointBalanceRes.Message.PointBalance {

		data.AfterPointAcquisition = append(data.AfterPointAcquisition,
			apiOutputFormatter.AfterPointAcquisition{
				BusinessPartner: &v.BusinessPartner,
				PointSymbol:     &v.PointSymbol,
				CurrentBalance:  &v.CurrentBalance,
				LimitBalance:    v.LimitBalance,
			},
		)
	}

	for _, v := range *eventPointConditionElementRes.Message.PointConditionElement {
		data.AfterPointAcquisition = append(data.AfterPointAcquisition,
			apiOutputFormatter.AfterPointAcquisition{
				Event:                          &v.Event,
				PointConditionRecord:           &v.PointConditionRecord,
				PointConditionSequentialNumber: &v.PointConditionSequentialNumber,
				PointSymbol:                    &v.PointSymbol,
				Sender:                         &v.Sender,
				PointTransactionType:           &v.PointTransactionType,
				PointConditionType:             &v.PointConditionType,
				PointConditionRateValue:        &v.PointConditionRateValue,
				PointConditionRatio:            &v.PointConditionRatio,
				PlusMinus:                      &v.PlusMinus,
			},
		)
	}

	for _, v := range *eventHeaderDocRes.Message.HeaderDoc {
		img := services.ReadEventImage(
			eventHeaderDocRes,
			v.Event,
		)

		data.AfterPointAcquisition = append(data.AfterPointAcquisition,
			apiOutputFormatter.AfterPointAcquisition{
				Event: &v.Event,
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
