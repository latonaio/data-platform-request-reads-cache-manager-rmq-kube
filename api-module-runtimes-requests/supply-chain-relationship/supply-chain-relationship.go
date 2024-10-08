package apiModuleRuntimesRequestsSupplyChainRelationship

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type SupplyChainRelationshipReq struct {
	General  General  `json:"SupplyChainRelationship"`
	Accepter []string `json:"accepter"`
}

type General struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

func CreateSupplyChainRelationshipRequestGeneralByBuyer(
	requestPram *apiInputReader.Request,
	supplyChainRelationshipGeneral *apiInputReader.SupplyChainRelationshipGeneral,
) SupplyChainRelationshipReq {
	req := SupplyChainRelationshipReq{
		General: General{
			Buyer:               *supplyChainRelationshipGeneral.Buyer,
			IsMarkedForDeletion: supplyChainRelationshipGeneral.IsMarkedForDeletion,
		},
		Accepter: []string{
			"GeneralsByBuyer",
		},
	}
	return req
}

func CreateSupplyChainRelationshipRequestGeneralBySeller(
	requestPram *apiInputReader.Request,
	supplyChainRelationshipGeneral *apiInputReader.SupplyChainRelationshipGeneral,
) SupplyChainRelationshipReq {
	req := SupplyChainRelationshipReq{
		General: General{
			Seller:              *supplyChainRelationshipGeneral.Seller,
			IsMarkedForDeletion: supplyChainRelationshipGeneral.IsMarkedForDeletion,
		},
		Accepter: []string{
			"GeneralsBySeller",
		},
	}
	return req
}

func SupplyChainRelationshipReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.SupplyChainRelationship,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_SUPPLY_CHAIN_RELATIONSHIP_SRV"
	aPIType := "reads"

	var request SupplyChainRelationshipReq

	if accepter == "GeneralsByBuyer" {
		request = CreateSupplyChainRelationshipRequestGeneralByBuyer(
			requestPram,
			&apiInputReader.SupplyChainRelationshipGeneral{
				Buyer:               input.SupplyChainRelationshipGeneral.Buyer,
				IsMarkedForDeletion: input.SupplyChainRelationshipGeneral.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "GeneralsBySeller" {
		request = CreateSupplyChainRelationshipRequestGeneralBySeller(
			requestPram,
			&apiInputReader.SupplyChainRelationshipGeneral{
				Seller:              input.SupplyChainRelationshipGeneral.Seller,
				IsMarkedForDeletion: input.SupplyChainRelationshipGeneral.IsMarkedForDeletion,
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
		requestPram,
	)

	return responseBody
}
