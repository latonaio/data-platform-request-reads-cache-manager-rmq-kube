package main

import (
	"context"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-api-request-reads-cache-manager-rmq-kube/config"
	"data-platform-api-request-reads-cache-manager-rmq-kube/controller"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

func main() {
	ctx := context.Background()
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), "", nil, 0)
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	rmqCtrl, err := rmqsessioncontroller.NewRMQSessionCtrl(conf, l)
	if err != nil {
		l.Fatal(err.Error())
	}
	cache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l)
	cacheCtrl := controller.NewController(ctx, cache, rmqCtrl, l)

	l.Info("READY!")
	for msg := range iter {
		reqMsg := msg
		go func(msg rabbitmq.RabbitmqMessage) {
			err = callProcess(reqMsg, cacheCtrl, conf)
			if err != nil {
				l.Error(err)
			}
			reqMsg.Success()
		}(reqMsg)
	}
}

func recovery(l *logger.Logger, err *error) {
	if e := recover(); e != nil {
		*err = xerrors.Errorf("error occurred: %w", e)
		l.Error("%+v", *err)
		*err = nil
		return
	}
}

func callProcess(msg rabbitmq.RabbitmqMessage, cacheCtrl *controller.Controller, conf *config.Conf) (err error) {
	l := logger.NewLogger()
	defer recovery(l, &err)
	input := map[string]interface{}{}
	err = json.Unmarshal(msg.Raw(), &input)

	l.JsonParseOut(input)
	if err != nil {
		return err
	}
	switch input["ui_function"] {
	case "OrdersDetail":
		err = cacheCtrl.OrdersDetail.OrdersDetail(msg)
	case "OrdersListBuyer":
		err = cacheCtrl.OrdersList.OrdersList(msg, l)
	case "OrdersListSeller":
		err = cacheCtrl.OrdersList.OrdersList(msg, l)
	case "OrdersDetailPagination":
		err = cacheCtrl.OrdersDetailPagination.Pagination(msg)
	case "OrdersDetailList":
		err = cacheCtrl.OrdersDetailList.OrdersDetailList(msg)
	case "DeliveryDocumentList":
		err = cacheCtrl.DeliveryDocumentList.DeliveryDocumentList(msg)
	case "DeliveryDocumentDetailList":
		err = cacheCtrl.DeliveryDocumentDetailList.DeliveryDocumentDetailList(msg)
	case "DeliveryDocumentDetail":
		err = cacheCtrl.DeliveryDocumentDetail.DeliveryDocumentDetail(msg)
	case "DeliveryDocumentDetailPagination":
		err = cacheCtrl.DeliveryDocumentDetailPagination.Pagination(msg)
	case "InvoiceDocumentList":
		err = cacheCtrl.InvoiceList.InvoiceList(msg)
	case "InvoiceDocumentDetailList":
		err = cacheCtrl.InvoiceDetailList.InvoiceDetailList(msg)
	case "InvoiceDocumentDetail":
		err = cacheCtrl.InvoiceDetail.InvoiceDetail(msg)
	case "ProductionOrderList":
		err = cacheCtrl.ProductionOrderList.ProductionOrderList(msg, l)
	case "ProductionOrderDetailList":
		err = cacheCtrl.ProductionOrderDetailList.ProductionOrderDetailList(msg)
	case "ProductionOrderDetail":
		err = cacheCtrl.ProductionOrderDetail.ProductionOrderDetail(msg, l)
	case "ProductionOrderDetailPagination":
		err = cacheCtrl.ProductionOrderDetailPagination.Pagination(msg)
	case "ProductList":
		err = cacheCtrl.ProductList.ProductList(msg, l)
	case "ProductDetailExconfList":
		err = cacheCtrl.ProductDetailList.ProductDetailList(msg, l)
	case "BillOfMaterialList":
		err = cacheCtrl.BillOfMaterialList.BillOfMaterialList(msg)
	case "BillOfMaterialDetailList":
		err = cacheCtrl.BillOfMaterialDetailList.BillOfMaterialDetailList(msg)
	case "ProductionVersionList":
		err = cacheCtrl.ProductionVersionList.ProductionVersionList(msg)
	// case "ProductionVersionDetailList":
	// 	err = cacheCtrl.ProductionVersionDetailList.ProductionVersionDetailList(msg)
	case "OperationsList":
		err = cacheCtrl.OperationsList.OperationsList(msg, l)
	// case "OperationsDetailList":
	// 	err = cacheCtrl.OperationsDetailList.OperationsDetailList(msg)
	case "EquipmentList":
		err = cacheCtrl.EquipmentList.EquipmentList(msg)
	// case "EquipmentDetailList":
	// 	err = cacheCtrl.EquipmentDetailList.EquipmentDetailList(msg)
	case "WorkCenterList":
		err = cacheCtrl.WorkCenterList.WorkCenterList(msg, l)
	// case "WorkCenterDetailList":
	// 	err = cacheCtrl.WorkCenterDetailList.WorkCenterDetailList(msg)
	case "SupplyChainRelationshipList":
		err = cacheCtrl.SupplyChainRelationshipList.SupplyChainRelationshipList(msg, l)
	case "BusinessPartnerList":
		err = cacheCtrl.BusinessPartnerList.BusinessPartnerList(msg, l)

		// case "StorageStockList":
		// 	err = cacheCtrl.WorkCenterDetailList.WorkCenterDetailList(msg)
		// case "StorageStockDetailList":
		// 	err = cacheCtrl.WorkCenterDetailList.WorkCenterDetailList(msg)
	case "PriceMasterList":
		err = cacheCtrl.PriceMasterList.PriceMasterList(msg, l)
	case "update":
		err = cacheCtrl.Update.Update(msg)
	default:
		l.Info("unknow ui-function %v", input["ui_function"])
	}

	return err
}

func getSessionID(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawSID, ok := m["runtime_session_id"]
	if !ok {
		return "", xerrors.Errorf("runtime_session_id not included")
	}

	return fmt.Sprintf("%v", rawSID), nil
}
