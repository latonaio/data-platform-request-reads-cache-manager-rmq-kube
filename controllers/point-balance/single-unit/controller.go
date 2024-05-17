package controllersPointBalanceSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-balance"
	apiModuleRuntimesResponsesPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-balance"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type PointBalanceSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *PointBalanceSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "point-balance"
	redisKeyCategory2 := "point-balance-single-unit"
	businessPartner, _ := controller.GetInt("businessPartner")
	pointSymbol := "POYPO"

	PointBalanceSingleUnit := apiInputReader.PointBalanceGlobal{}

	PointBalanceSingleUnit = apiInputReader.PointBalanceGlobal{
		PointBalance: &apiInputReader.PointBalance{
			BusinessPartner: businessPartner,
			PointSymbol:     pointSymbol,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(businessPartner),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.PointBalance

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
			controller.request(PointBalanceSingleUnit)
		}()
	} else {
		controller.request(PointBalanceSingleUnit)
	}
}

func (
	controller *PointBalanceSingleUnitController,
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
	controller *PointBalanceSingleUnitController,
) request(
	input apiInputReader.PointBalanceGlobal,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	pointBalanceRes := *controller.createPointBalanceRequestPointBalance(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&pointBalanceRes,
	)
}

func (
	controller *PointBalanceSingleUnitController,
) fin(
	pointBalanceRes *apiModuleRuntimesResponsesPointBalance.PointBalanceRes,
) {

	data := apiOutputFormatter.PointBalance{}

	for _, v := range *pointBalanceRes.Message.PointBalance {
		//qrcode := services.CreateQRCodePointBalanceDocImage(
		//	pointBalanceDocRes,
		//	v.PointBalance,
		//)

		data.PointBalancePointBalance = append(data.PointBalancePointBalance,
			apiOutputFormatter.PointBalancePointBalance{
				BusinessPartner: v.BusinessPartner,
				PointSymbol:     v.PointSymbol,
				CurrentBalance:  v.CurrentBalance,
				LimitBalance:    v.LimitBalance,
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
