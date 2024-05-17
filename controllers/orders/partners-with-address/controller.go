package controllersOrdersPartnersWithAddress

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type OrdersPartnersWithAddressController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *OrdersPartnersWithAddressController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "orders-partners-with-address-controller"
	orderId, _ := controller.GetInt("orderId")

	OrdersPartnersWithAddress := apiInputReader.Orders{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	OrdersPartnersWithAddress = apiInputReader.Orders{
		OrdersHeader: &apiInputReader.OrdersHeader{
			OrderID:                         orderId,
			HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
			HeaderDeliveryStatus:            &headerDeliveryStatus,
			IsCancelled:                     &isCancelled,
			IsMarkedForDeletion:             &isMarkedForDeletion,
		},
		OrdersPartner: &apiInputReader.OrdersPartner{
			OrderID: orderId,
		},
		OrdersAddress: &apiInputReader.OrdersAddress{
			OrderID: orderId,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(orderId),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Orders

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
			controller.request(OrdersPartnersWithAddress)
		}()
	} else {
		controller.request(OrdersPartnersWithAddress)
	}
}

func (
	controller *OrdersPartnersWithAddressController,
) createOrdersRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
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
		controller.CustomLogger.Error("OrdersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersPartnersWithAddressController,
) createOrdersRequestPartners(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"Partners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("OrdersPartnersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersPartnersWithAddressController,
) createOrdersRequestAddresses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"Addresses",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("OrdersAddressesReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersPartnersWithAddressController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*ordersRes.Message.Address)[0].LocalRegion
	country := &(*ordersRes.Message.Address)[0].Country

	var inputLocalRegion *string
	var inputCountry *string

	if localRegion != nil {
		inputLocalRegion = *localRegion
		inputCountry = *country
	}

	input := apiModuleRuntimesRequestsLocalRegion.LocalRegion{
		LocalRegion: *inputLocalRegion,
		Country:     *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesLocalRegion.LocalRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalRegion.LocalRegionReadsText(
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
		controller.CustomLogger.Error("LocalRegionReadsText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersPartnersWithAddressController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	ordersPartnersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*ordersPartnersRes.Message.Partner))

	for _, v := range *ordersPartnersRes.Message.Partner {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
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
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersPartnersWithAddressController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	ordersHeaderRes := *controller.createOrdersRequestHeader(
		controller.UserInfo,
		input,
	)

	ordersPartnersRes := controller.createOrdersRequestPartners(
		controller.UserInfo,
		input,
	)

	ordersAddressesRes := controller.createOrdersRequestAddresses(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		ordersPartnersRes,
	)

	localRegionTextRes := controller.CreateLocalRegionRequestText(
		controller.UserInfo,
		ordersAddressesRes,
	)

	controller.fin(
		&ordersHeaderRes,
		ordersPartnersRes,
		ordersAddressesRes,
		&businessPartnerRes,
		localRegionTextRes,
	)
}

func (
	controller *OrdersPartnersWithAddressController,
) fin(
	ordersHeaderRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersPartnersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersAddressesRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

	ordersAddressesMapper := services.OrdersAddressesMapper(
		ordersAddressesRes,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *ordersHeaderRes.Message.Header {
		data.OrdersHeaderWithItem = append(data.OrdersHeaderWithItem,
			apiOutputFormatter.OrdersHeaderWithItem{
				OrderID:               v.OrderID,
				Buyer:                 v.Buyer,
				BuyerName:             businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                v.Seller,
				SellerName:            businessPartnerMapper[v.Seller].BusinessPartnerName,
				RequestedDeliveryDate: v.RequestedDeliveryDate,
				RequestedDeliveryTime: v.RequestedDeliveryTime,
				TotalGrossAmount:      v.TotalGrossAmount,
			},
		)
	}

	for _, v := range *ordersPartnersRes.Message.Partner {
		cityName := ordersAddressesMapper[*v.AddressID].CityName
		streetName := ordersAddressesMapper[*v.AddressID].StreetName
		building := ordersAddressesMapper[*v.AddressID].Building
		//var floor *int = ordersAddressesMapper[*v.AddressID].Floor
		//var room *int = ordersAddressesMapper[*v.AddressID].Room

		var cityNameStr string
		if cityName != nil {
			cityNameStr = *cityName
		} else {
			cityNameStr = ""
		}

		var streetNameStr string
		if streetName != nil {
			streetNameStr = *streetName
		} else {
			streetNameStr = ""
		}

		var buildingStr string
		if building != nil {
			buildingStr = *building
		} else {
			buildingStr = ""
		}

		//var floorStr string
		//if floor != nil {
		//	floorStr = fmt.Sprintf("%dF", *floor)
		//} else {
		//	floorStr = ""
		//}
		//
		//var roomStr string
		//if room != nil {
		//	roomStr = fmt.Sprintf("%d", *room)
		//} else {
		//	roomStr = ""
		//}

		addressIdentifier := fmt.Sprintf(
			//"%s%s%s%s %s",
			"%s%s%s",
			cityNameStr,
			streetNameStr,
			buildingStr,
			//floorStr,
			//roomStr,
		)

		var localRegionName *string

		localRegionTextMapperName := localRegionTextMapper[*ordersAddressesMapper[*v.AddressID].LocalRegion].LocalRegionName
		if &localRegionTextMapperName != nil {
			localRegionName = &localRegionTextMapperName
		}

		var businessPartnerName *string
		businessPartnerMapperName := businessPartnerMapper[v.BusinessPartner].BusinessPartnerName
		if &businessPartnerMapperName != nil {
			businessPartnerName = &businessPartnerMapperName
		}

		data.OrdersPartner = append(data.OrdersPartner,
			apiOutputFormatter.OrdersPartner{
				OrderID:             v.OrderID,
				PartnerFunction:     v.PartnerFunction,
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: businessPartnerName,
				Country:             v.Country,
				PostalCode:          ordersAddressesMapper[*v.AddressID].PostalCode,
				AddressIdentifier:   &addressIdentifier,
				LocalRegionName:     localRegionName,
			},
		)
	}

	for _, v := range *ordersAddressesRes.Message.Address {
		var localRegionName *string

		localRegionTextMapperName := localRegionTextMapper[*v.LocalRegion].LocalRegionName
		if &localRegionTextMapperName != nil {
			localRegionName = &localRegionTextMapperName
		}

		data.OrdersAddress = append(data.OrdersAddress,
			apiOutputFormatter.OrdersAddress{
				AddressID:       &v.AddressID,
				OrderID:         v.OrderID,
				PostalCode:      v.PostalCode,
				LocalRegion:     v.LocalRegion,
				LocalRegionName: localRegionName,
				Country:         v.Country,
				StreetName:      v.StreetName,
				CityName:        v.CityName,
				Building:        v.Building,
				Floor:           v.Floor,
				Room:            v.Room,
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
