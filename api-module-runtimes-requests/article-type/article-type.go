package apiModuleRuntimesRequestsArticleType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ArticleTypeReq struct {
	ArticleType  ArticleType   `json:"ArticleType"`
	ArticleTypes []ArticleType `json:"ArticleTypes"`
	Accepter     []string      `json:"accepter"`
}

type ArticleType struct {
	ArticleType         string  `json:"ArticleType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	ArticleType         string  `json:"ArticleType"`
	Language            string  `json:"Language"`
	ArticleTypeName     *string `json:"ArticleTypeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateArticleTypeRequestArticleTypesByArticleTypes(
	requestPram *apiInputReader.Request,
	input []ArticleType,
) ArticleTypeReq {
	req := ArticleTypeReq{
		ArticleTypes: input,
		Accepter: []string{
			"ArticleTypesByArticleTypes",
		},
	}
	return req
}

func CreateArticleTypeRequestArticleTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ArticleType,
) ArticleTypeReq {
	isMarkedForDeletion := false

	req := ArticleTypeReq{
		ArticleTypes: []ArticleType{
			{
				ArticleType:         input.ArticleType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"ArticleTypes",
		},
	}
	return req
}

func CreateArticleTypeRequestText(
	requestPram *apiInputReader.Request,
	input ArticleType,
) ArticleTypeReq {
	isMarkedForDeletion := false

	req := ArticleTypeReq{
		ArticleType: ArticleType{
			ArticleType:         input.ArticleType,
			IsMarkedForDeletion: &isMarkedForDeletion,
			Text: []Text{
				{
					Language:            *requestPram.Language,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Text",
		},
	}
	return req
}

func CreateArticleTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input ArticleType,
) ArticleTypeReq {
	isMarkedForDeletion := false

	req := ArticleTypeReq{
		ArticleType: ArticleType{
			ArticleType: input.ArticleType,
			Text: []Text{
				{
					Language:            *requestPram.Language,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Texts",
		},
	}
	return req
}

func ArticleTypeReadsArticleTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ArticleType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_TYPE_SRV"
	aPIType := "reads"

	var request ArticleTypeReq

	request = CreateArticleTypeRequestArticleTypes(
		requestPram,
		input,
	)

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

func ArticleTypeReadsArticleTypesByArticleTypes(
	requestPram *apiInputReader.Request,
	input []ArticleType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_TYPE_SRV"
	aPIType := "reads"

	var request ArticleTypeReq

	request = CreateArticleTypeRequestArticleTypesByArticleTypes(
		requestPram,
		input,
	)

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

func ArticleTypeReadsText(
	requestPram *apiInputReader.Request,
	input ArticleType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_TYPE_SRV"
	aPIType := "reads"

	var request ArticleTypeReq

	request = CreateArticleTypeRequestText(
		requestPram,
		input,
	)

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

func ArticleTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input ArticleType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_TYPE_SRV"
	aPIType := "reads"

	var request ArticleTypeReq

	request = CreateArticleTypeRequestTexts(
		requestPram,
		input,
	)

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
