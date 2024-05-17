package friend

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type FriendReq struct {
	General  General   `json:"Friend"`
	Generals []General `json:"Friends"`
	Accepter []string  `json:"accepter"`
}

type General struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	Friend                    int     `json:"Friend"`
	BPBusinessPartnerType     *string `json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType *string `json:"FriendBusinessPartnerType"`
	RankType                  *string `json:"RankType"`
	Rank                      *int    `json:"Rank"`
	FriendIsBlocked           *bool   `json:"FriendIsBlocked"`
	CreationDate              *string `json:"CreationDate"`
	CreationTime              *string `json:"CreationTime"`
	LastChangeDate            *string `json:"LastChangeDate"`
	LastChangeTime            *string `json:"LastChangeTime"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

func CreateFriendRequestGeneral(
	requestPram *apiInputReader.Request,
	friendGeneral *apiInputReader.FriendGeneral,
) FriendReq {
	req := FriendReq{
		General: General{
			BusinessPartner:     friendGeneral.BusinessPartner,
			Friend:              friendGeneral.Friend,
			FriendIsBlocked:     &friendGeneral.FriendIsBlocked,
			IsMarkedForDeletion: friendGeneral.IsMarkedForDeletion,
		},
		Accepter: []string{
			"General",
		},
	}
	return req
}

func CreateFriendRequestGenerals(
	requestPram *apiInputReader.Request,
	friendGenerals *apiInputReader.FriendGeneral,
) FriendReq {
	req := FriendReq{
		General: General{
			BusinessPartner:     friendGenerals.Friend,
			FriendIsBlocked:     &friendGenerals.FriendIsBlocked,
			IsMarkedForDeletion: friendGenerals.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func FriendReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Friend,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_FRIEND_SRV"
	aPIType := "reads"

	var request FriendReq

	if accepter == "General" {
		request = CreateFriendRequestGeneral(
			requestPram,
			&apiInputReader.FriendGeneral{
				BusinessPartner:     input.FriendGeneral.BusinessPartner,
				Friend:              input.FriendGeneral.Friend,
				FriendIsBlocked:     input.FriendGeneral.FriendIsBlocked,
				IsMarkedForDeletion: input.FriendGeneral.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Generals" {
		request = CreateFriendRequestGenerals(
			requestPram,
			&apiInputReader.FriendGeneral{
				BusinessPartner:     input.FriendGeneral.BusinessPartner,
				FriendIsBlocked:     input.FriendGeneral.FriendIsBlocked,
				IsMarkedForDeletion: input.FriendGeneral.IsMarkedForDeletion,
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
