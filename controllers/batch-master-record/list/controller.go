package controllersBatchMasterRecordList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesResponsesBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/batch-master-record"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BatchMasterRecordListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *BatchMasterRecordListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "batch-master-record"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	batchMasterRecordHeader := apiInputReader.BatchMasterRecord{}

	batchMasterRecordHeader = apiInputReader.BatchMasterRecord{
		BatchMasterRecordHeader: &apiInputReader.BatchMasterRecordHeader{
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			userType,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.BatchMasterRecord

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
			controller.request(batchMasterRecordHeader)
		}()
	} else {
		controller.request(batchMasterRecordHeader)
	}
}

func (
	controller *BatchMasterRecordListController,
) createBatchMasterRecordRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.BatchMasterRecord,
) *apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes {
	responseJsonData := apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes{}
	//responseBody := apiModuleRuntimesRequestsBatchMasterRecord.BatchMasterRecordReads(
	//	requestPram,
	//	//apiModuleRuntimesRequestsBatchMasterRecord.Batch{},
	//	&controller.Controller,
	//	"Batches",
	//)

	err := json.Unmarshal(nil, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BatchMasterRecordReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BatchMasterRecordListController,
) request(
	input apiInputReader.BatchMasterRecord,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes{}

	headerRes = *controller.createBatchMasterRecordRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&headerRes,
	)
}

func (
	controller *BatchMasterRecordListController,
) fin(
	headerRes *apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes,
) {

	data := apiOutputFormatter.BatchMasterRecord{}

	for _, v := range *headerRes.Message.Batch {

		data.BatchMasterRecordHeader = append(data.BatchMasterRecordHeader,
			apiOutputFormatter.BatchMasterRecordHeader{
				Product:           v.Product,
				BusinessPartner:   v.BusinessPartner,
				Plant:             v.Plant,
				Batch:             v.Batch,
				ValidityStartDate: v.ValidityStartDate,
				ValidityStartTime: v.ValidityStartTime,
				ValidityEndDate:   v.ValidityEndDate,
				ValidityEndTime:   v.ValidityEndTime,
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
