package apiModuleRuntimesRequestsOrdersDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type OrdersDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Orders"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	OrderID                  *int    `json:"OrderID"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
	ItemDoc                  ItemDoc `json:"ItemDoc"`
}

type ItemDoc struct {
	OrderID                  *int    `json:"OrderID"`
	OrderItem                *int    `json:"OrderItem"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateOrdersDocRequestItemDoc(
	requestPram *apiInputReader.Request,
	itemDoc apiInputReader.OrdersDocItemDoc,
) OrdersDocReq {
	docIssuerBusinessPartner := 201 // TODO 暫定対応

	req := OrdersDocReq{
		HeaderDoc: HeaderDoc{
			OrderID: &itemDoc.OrderID,
			ItemDoc: ItemDoc{
				OrderID:                  &itemDoc.OrderID,
				OrderItem:                itemDoc.OrderItem,
				DocType:                  itemDoc.DocType,
				DocIssuerBusinessPartner: &docIssuerBusinessPartner,
				//DocIssuerBusinessPartner: itemDoc.DocIssuerBusinessPartner,
			},
		},
		Accepter: []string{
			"ItemDoc",
		},
	}
	return req
}

func OrdersDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_ORDERS_DOC_SRV"
	aPIType := "reads"

	var request OrdersDocReq

	if accepter == "ItemDoc" {
		request = CreateOrdersDocRequestItemDoc(
			requestPram,
			apiInputReader.OrdersDocItemDoc{
				OrderID:                  input.OrdersDocItemDoc.OrderID,
				OrderItem:                input.OrdersDocItemDoc.OrderItem,
				DocType:                  input.OrdersDocItemDoc.DocType,
				DocIssuerBusinessPartner: input.OrdersDocItemDoc.DocIssuerBusinessPartner,
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
