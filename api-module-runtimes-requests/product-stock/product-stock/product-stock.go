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
	ProductStock ProductStock `json:"ProductStock"`
	Accepter     []string     `json:"accepter"`
}

type ProductStock struct {
	Product                                *string                           `json:"Product"`
	BusinessPartner                        *int                              `json:"BusinessPartner"`
	Plant                                  *string                           `json:"Plant"`
	SupplyChainRelationshipID              *int                              `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int                              `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int                              `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  *int                              `json:"Buyer"`
	Seller                                 *int                              `json:"Seller"`
	DeliverToParty                         *int                              `json:"DeliverToParty"`
	DeliverFromParty                       *int                              `json:"DeliverFromParty"`
	DeliverToPlant                         *string                           `json:"DeliverToPlant"`
	DeliverFromPlant                       *string                           `json:"DeliverFromPlant"`
	InventoryStockType                     *string                           `json:"InventoryStockType"`
	ProductStock                           *float32                          `json:"ProductStock"`
	CreationDate                           *string                           `json:"CreationDate"`
	CreationTime                           *string                           `json:"CreationTime"`
	LastChangeDate                         *string                           `json:"LastChangeDate"`
	LastChangeTime                         *string                           `json:"LastChangeTime"`
	ProductStockAvailability               []ProductStockAvailability        `json:"ProductStockAvailability"`
	ProductStockByStorageBinByBatch        []ProductStockByStorageBinByBatch `json:"ProductStockByStorageBinByBatch"`
}

type ProductStockAvailability struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockByStorageBinByBatch struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	StorageLocation                        string   `json:"StorageLocation"`
	StorageBin                             string   `json:"StorageBin"`
	Batch                                  string   `json:"Batch"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	InventoryStockType                     string   `json:"InventoryStockType"`
	ProductStock                           *float32 `json:"ProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

func CreateProductStockRequestProductStock(
	requestPram *apiInputReader.Request,
	productStockHeader *apiInputReader.ProductStockHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			Product:         &productStockHeader.Product,
			BusinessPartner: &productStockHeader.BusinessPartner,
			Plant:           &productStockHeader.Plant,
		},
		Accepter: []string{
			"ProductStock",
		},
	}
	return req
}

func CreateProductStockRequestProductStocksByStorageBinByBatch(
	requestPram *apiInputReader.Request,
	productStockByStorageBinByBatchHeader *apiInputReader.ProductStockByStorageBinByBatchHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			ProductStockByStorageBinByBatch: []ProductStockByStorageBinByBatch{
				{
					Product:         productStockByStorageBinByBatchHeader.Product,
					BusinessPartner: productStockByStorageBinByBatchHeader.BusinessPartner,
					Plant:           productStockByStorageBinByBatchHeader.Plant,
				},
			},
		},
		Accepter: []string{
			"ProductStocksByStorageBinByBatch",
		},
	}
	return req
}

func CreateProductStockRequestProductStockByBuyer(
	requestPram *apiInputReader.Request,
	productStockHeader *apiInputReader.ProductStockHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			Buyer: productStockHeader.Buyer,
		},
		Accepter: []string{
			"ProductStocksByBuyer",
		},
	}
	return req
}

func CreateProductStockRequestProductStockBySeller(
	requestPram *apiInputReader.Request,
	productStockHeader *apiInputReader.ProductStockHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			Seller: productStockHeader.Seller,
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

	if accepter == "ProductStock" {
		request = CreateProductStockRequestProductStock(
			requestPram,
			&apiInputReader.ProductStockHeader{
				Product:         input.ProductStockHeader.Product,
				BusinessPartner: input.ProductStockHeader.BusinessPartner,
				Plant:           input.ProductStockHeader.Plant,
			},
		)
	}

	if accepter == "ProductStocksByStorageBinByBatch" {
		request = CreateProductStockRequestProductStocksByStorageBinByBatch(
			requestPram,
			&apiInputReader.ProductStockByStorageBinByBatchHeader{
				Product:         input.ProductStockByStorageBinByBatchHeader.Product,
				BusinessPartner: input.ProductStockByStorageBinByBatchHeader.BusinessPartner,
				Plant:           input.ProductStockByStorageBinByBatchHeader.Plant,
			},
		)
	}

	if accepter == "ProductStocksByBuyer" {
		request = CreateProductStockRequestProductStockByBuyer(
			requestPram,
			&apiInputReader.ProductStockHeader{
				Buyer: input.ProductStockHeader.Buyer,
			},
		)
	}

	if accepter == "ProductStocksBySeller" {
		request = CreateProductStockRequestProductStockBySeller(
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

func CreateProductStockRequestProductStockAvailability(
	requestPram *apiInputReader.Request,
	productStockAvailabilityHeader *apiInputReader.ProductStockAvailabilityHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			ProductStockAvailability: []ProductStockAvailability{
				{
					Product:         productStockAvailabilityHeader.Product,
					BusinessPartner: productStockAvailabilityHeader.BusinessPartner,
					Plant:           productStockAvailabilityHeader.Plant,
				},
			},
		},

		Accepter: []string{
			"ProductStockAvailability",
		},
	}
	return req
}

func CreateProductStockAvailabilityRequestProductStockAvailabilityByBuyer(
	requestPram *apiInputReader.Request,
	productStockAvailabilityHeader *apiInputReader.ProductStockAvailabilityHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			ProductStockAvailability: []ProductStockAvailability{
				{
					Buyer: *productStockAvailabilityHeader.Buyer,
				},
			},
		},
		Accepter: []string{
			"ProductStockAvailabilitiesByBuyer",
		},
	}
	return req
}

func CreateProductStockAvailabilityRequestProductStockAvailabilityBySeller(
	requestPram *apiInputReader.Request,
	productStockAvailabilityHeader *apiInputReader.ProductStockAvailabilityHeader,
) ProductStockReq {
	req := ProductStockReq{
		ProductStock: ProductStock{
			ProductStockAvailability: []ProductStockAvailability{
				{
					Seller: *productStockAvailabilityHeader.Seller,
				},
			},
		},
		Accepter: []string{
			"ProductStockAvailabilitiesBySeller",
		},
	}
	return req
}

func ProductStockAvailabilityReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_STOCK_SRV"
	aPIType := "reads"

	var request ProductStockReq

	if accepter == "ProductStockAvailability" {
		request = CreateProductStockRequestProductStockAvailability(
			requestPram,
			&apiInputReader.ProductStockAvailabilityHeader{
				Product:         input.ProductStockAvailabilityHeader.Product,
				BusinessPartner: input.ProductStockAvailabilityHeader.BusinessPartner,
				Plant:           input.ProductStockAvailabilityHeader.Plant,
			},
		)
	}

	if accepter == "ProductStockAvailabilitiesByBuyer" {
		request = CreateProductStockAvailabilityRequestProductStockAvailabilityByBuyer(
			requestPram,
			&apiInputReader.ProductStockAvailabilityHeader{
				Buyer: input.ProductStockAvailabilityHeader.Buyer,
			},
		)
	}

	if accepter == "ProductStockAvailabilitiesBySeller" {
		request = CreateProductStockAvailabilityRequestProductStockAvailabilityBySeller(
			requestPram,
			&apiInputReader.ProductStockAvailabilityHeader{
				Seller: input.ProductStockAvailabilityHeader.Seller,
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
