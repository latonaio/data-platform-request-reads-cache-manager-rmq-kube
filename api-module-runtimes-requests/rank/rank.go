package apiModuleRuntimesRequestsRank

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type RankReq struct {
	Rank         Rank       `json:"Rank"`
	Ranks        []Rank     `json:"Ranks"`
	Accepter     []string   `json:"accepter"`
}

type Rank struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	CreationDate		*string	`json:"CreationDate"`
	LastChangeDate		*string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	RankType			string	 `json:"RankType"`
	Rank				int		 `json:"Rank"`
	Language          	string   `json:"Language"`
	RankName			*string	 `json:"RankdName"`
	CreationDate		*string  `json:"CreationDate"`
	LastChangeDate		*string	 `json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	 `json:"IsMarkedForDeletion"`
}

func CreateRankRequestRanksByRanks(
	requestPram *apiInputReader.Request,
	input []Rank,
) RankReq {
	req := RankReq{
		Ranks: input,
		Accepter: []string{
			"RanksByRanks",
		},
	}
	return req
}

func CreateRankRequestRanks(
	requestPram *apiInputReader.Request,
	input apiInputReader.Rank,
) RankReq {
	isMarkedForDeletion := false

	req := RankReq{
		Ranks: []Rank{
			{
			    RankType:            input.RankType,
			    Rank:                input.Rank,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"Ranks",
		},
	}
	return req
}

func CreateRankRequestText(
	requestPram *apiInputReader.Request,
	input Rank,
) RankReq {
	isMarkedForDeletion := false

	req := RankReq{
		Rank: Rank{
			RankType:            input.RankType,
			Rank:                input.Rank,
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

func CreateRankRequestTexts(
	requestPram *apiInputReader.Request,
	input Rank,
) RankReq {
	isMarkedForDeletion := false

	req := RankReq{
		Rank: Rank{
			RankType:    input.RankType,
			Rank:        input.Rank,
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

func RankReadsRanks(
	requestPram *apiInputReader.Request,
	input apiInputReader.Rank,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_RANK_SRV"
	aPIType := "reads"

	var request RankReq

	request = CreateRankRequestRanks(
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

func RankReadsRanksByRanks(
	requestPram *apiInputReader.Request,
	input []Rank,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_RANK_SRV"
	aPIType := "reads"

	var request RankReq

	request = CreateRankRequestRanksByRanks(
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

func RankReadsText(
	requestPram *apiInputReader.Request,
	input Rank,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_RANK_SRV"
	aPIType := "reads"

	var request RankReq

	request = CreateRankRequestText(
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

func RankReadsTexts(
	requestPram *apiInputReader.Request,
	input Rank,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_RANK_SRV"
	aPIType := "reads"

	var request RankReq

	request = CreateRankRequestTexts(
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
