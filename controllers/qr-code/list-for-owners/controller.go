package controllersQRCodeListForOwners

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/event/event-doc"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesResponsesEvent "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/event"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"sync"
)

type QRCodeListForOwnersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type QRCodeListForOwners struct {
	EventHeader []apiOutputFormatter.EventHeader `json:"EventHeader"`
	ShopHeader  []apiOutputFormatter.ShopHeader  `json:"ShopHeader"`
}

func (controller *QRCodeListForOwnersController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "organizationBusinessPartner"
	redisKeyCategory2 := "qr-code-list-for-owners"
	organizationBusinessPartner, _ := controller.GetInt("organizationBusinessPartner")

	QRCodeListForOwnersEvent := apiInputReader.Event{}
	QRCodeListForOwnersShop := apiInputReader.Shop{}

	//	docType := "QRCODE"

	isReleased := true
	isCancelled := false
	isMarkedForDeletion := false

	QRCodeListForOwnersEvent = apiInputReader.Event{
		EventHeader: &apiInputReader.EventHeader{
			EventOwner:          &organizationBusinessPartner,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	QRCodeListForOwnersShop = apiInputReader.Shop{
		ShopHeader: &apiInputReader.ShopHeader{
			ShopOwner:           organizationBusinessPartner,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
		var responseData QRCodeListForOwners

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
			controller.request(QRCodeListForOwnersEvent, QRCodeListForOwnersShop)
		}()
	} else {
		controller.request(QRCodeListForOwnersEvent, QRCodeListForOwnersShop)
	}
}

func (
	controller *QRCodeListForOwnersController,
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
	controller *QRCodeListForOwnersController,
) createEventDocRequest(
	requestPram *apiInputReader.Request,
	shopHeaderRes apiModuleRuntimesResponsesEvent.EventRes,
) *apiModuleRuntimesResponsesEvent.EventDocRes {
	var input = apiInputReader.Event{}

	for _, v := range *shopHeaderRes.Message.Header {
		input = apiInputReader.Event{
			EventDocHeaderDoc: &apiInputReader.EventDocHeaderDoc{
				Event: v.Event,
			},
		}
	}

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
	controller *QRCodeListForOwnersController,
) createShopRequestHeadersByShopOwner(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByShopOwner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createShopRequestHeadersByShopOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QRCodeListForOwnersController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	var input = apiInputReader.Shop{}

	for _, v := range *shopHeaderRes.Message.Header {
		input = apiInputReader.Shop{
			ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{
				Shop: v.Shop,
			},
		}
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopDocRes{}
	responseBody := apiModuleRuntimesRequestsShopDoc.ShopDocReads(
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
		controller.CustomLogger.Error("createShopDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QRCodeListForOwnersController,
) request(
	inputEvent apiInputReader.Event,
	inputShop apiInputReader.Shop,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(2)

	var eventHeaderRes apiModuleRuntimesResponsesEvent.EventRes
	var shopHeaderRes apiModuleRuntimesResponsesShop.ShopRes

	var eventHeaderDocRes *apiModuleRuntimesResponsesEvent.EventDocRes
	var shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	go func() {
		defer wg.Done()
		eventHeaderRes = *controller.createEventRequestHeadersByEventOwner(
			controller.UserInfo,
			inputEvent,
		)
		eventHeaderDocRes = controller.createEventDocRequest(
			controller.UserInfo,
			eventHeaderRes,
		)
	}()

	go func() {
		defer wg.Done()
		shopHeaderRes = *controller.createShopRequestHeadersByShopOwner(
			controller.UserInfo,
			inputShop,
		)
		shopHeaderDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			shopHeaderRes,
		)
	}()

	wg.Wait()

	controller.fin(
		&eventHeaderRes,
		eventHeaderDocRes,
		&shopHeaderRes,
		shopHeaderDocRes,
	)
}

func (
	controller *QRCodeListForOwnersController,
) fin(
	eventHeaderRes *apiModuleRuntimesResponsesEvent.EventRes,
	eventHeaderDocRes *apiModuleRuntimesResponsesEvent.EventDocRes,
	shopHeaderRes *apiModuleRuntimesResponsesShop.ShopRes,
	shopHeaderDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
) {

	data := QRCodeListForOwners{}

	for _, v := range *eventHeaderRes.Message.Header {
		//		img := services.ReadEventImage(
		//			headerDocRes,
		//			v.Event,
		//		)

		qrcode := services.CreateQRCodeEventDocImage(
			eventHeaderDocRes,
			v.Event,
		)

		//		documentImage := services.ReadDocumentImageEvent(
		//			headerDocRes,
		//			v.Event,
		//		)

		data.EventHeader = append(data.EventHeader,
			apiOutputFormatter.EventHeader{
				Event:       v.Event,
				EventOwner:  v.EventOwner,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					//					Event:              img,
					QRCode: qrcode,
					//					DocumentImageEvent: documentImage,
				},
			},
		)
	}

	for _, v := range *shopHeaderRes.Message.Header {
		//		img := services.ReadShopImage(
		//			headerDocRes,
		//			v.Shop,
		//		)

		qrcode := services.CreateQRCodeShopDocImage(
			shopHeaderDocRes,
			v.Shop,
		)

		//		documentImage := services.ReadDocumentImageShop(
		//			headerDocRes,
		//			v.Shop,
		//		)

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				ShopOwner:   v.ShopOwner,
				Description: v.Description,
				Images: apiOutputFormatter.Images{
					//					Shop:              img,
					QRCode: qrcode,
					//					DocumentImageShop: documentImage,
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
