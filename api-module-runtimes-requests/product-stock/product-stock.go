package apiModuleRuntimesRequestsProductStock

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type ProductStockReq struct {
	Header   Header   `json:"ProductStock"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Product                                string                     `json:"Product"`
	BusinessPartner                        int                        `json:"BusinessPartner"`
	Plant                                  string                     `json:"Plant"`
	SupplyChainRelationshipID              int                        `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int                        `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int                        `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int                        `json:"Buyer"`
	Seller                                 int                        `json:"Seller"`
	DeliverToParty                         int                        `json:"DeliverToParty"`
	DeliverFromParty                       int                        `json:"DeliverFromParty"`
	DeliverToPlant                         string                     `json:"DeliverToPlant"`
	DeliverFromPlant                       string                     `json:"DeliverFromPlant"`
	InventoryStockType                     string                     `json:"InventoryStockType"`
	ProductStock                           *float32                   `json:"ProductStock"`
	CreationDate                           *string                    `json:"CreationDate"`
	CreationTime                           *string                    `json:"CreationTime"`
	LastChangeDate                         *string                    `json:"LastChangeDate"`
	LastChangeTime                         *string                    `json:"LastChangeTime"`
}

func CreateProductStockRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	productStockHeader *apiInputReader.ProductStockHeader,
) ProductStockReq {
	req := ProductStockReq{
		Header: Header{
			Buyer: *productStockHeader.Buyer,
		},
		Accepter: []string{
			"ProductStocksByBuyer",
		},
	}
	return req
}

func CreateProductStockRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	productStockHeader *apiInputReader.ProductStockHeader,
) ProductStockReq {
	req := ProductStockReq{
		Header: Header{
			Seller: *productStockHeader.Seller,
		},
		Accepter: []string{
			"ProductStocksBySeller",
		},
	}
	return req
}

func ProductStockReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_STOCK_SRV"
	aPIType := "reads"

	var request ProductStockReq

	if accepter == "ProductStocksByBuyer" {
		request = CreateProductStockRequestHeaderByBuyer(
			requestPram,
			&apiInputReader.ProductStockHeader{
				Buyer: input.ProductStockHeader.Buyer,
			},
		)
	}

	if accepter == "ProductStocksBySeller" {
		request = CreateProductStockRequestHeaderBySeller(
			requestPram,
			&apiInputReader.ProductStockHeader{
				Seller: input.ProductStockHeader.Seller,
			},
		)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
