package event

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type EventReq struct {
	Header   Header   `json:"Event"`
	Headers  []Header `json:"Events"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Event                         int                     `json:"Event"`
	EventType                     *string                 `json:"EventType"`
	EventOwner                    *int                    `json:"EventOwner"`
	EventOwnerBusinessPartnerRole *string                 `json:"EventOwnerBusinessPartnerRole"`
	PersonResponsible             *string                 `json:"PersonResponsible"`
	URL							  *string				  `json:"URL"`
	ValidityStartDate             *string                 `json:"ValidityStartDate"`
	ValidityStartTime             *string                 `json:"ValidityStartTime"`
	ValidityEndDate               *string                 `json:"ValidityEndDate"`
	ValidityEndTime               *string                 `json:"ValidityEndTime"`
	OperationStartDate            *string                 `json:"OperationStartDate"`
	OperationStartTime            *string                 `json:"OperationStartTime"`
	OperationEndDate              *string                 `json:"OperationEndDate"`
	OperationEndTime              *string                 `json:"OperationEndTime"`
	Description                   *string                 `json:"Description"`
	LongText                      *string                 `json:"LongText"`
	Introduction                  *string                 `json:"Introduction"`
	Site                          *int                    `json:"Site"`
	Capacity					  *int			          `json:"Capacity"`
	Shop                          *int                    `json:"Shop"`
	Project                       *int                    `json:"Project"`
	WBSElement                    *int                    `json:"WBSElement"`
	Tag1                          *string                 `json:"Tag1"`
	Tag2                          *string                 `json:"Tag2"`
	Tag3                          *string                 `json:"Tag3"`
	Tag4                          *string                 `json:"Tag4"`
	DistributionProfile           *string                 `json:"DistributionProfile"`
	PointConditionType            *string                 `json:"PointConditionType"`
	QuestionnaireType             *string                 `json:"QuestionnaireType"`
	QuestionnaireTemplate         *string                 `json:"QuestionnaireTemplate"`
	CreationDate                  *string                 `json:"CreationDate"`
	CreationTime                  *string                 `json:"CreationTime"`
	LastChangeDate                *string                 `json:"LastChangeDate"`
	LastChangeTime                *string                 `json:"LastChangeTime"`
	CreateUser                    *int                    `json:"CreateUser"`
	LastChangeUser                *int                    `json:"LastChangeUser"`
	IsReleased                    *bool                   `json:"IsReleased"`
	IsCancelled                   *bool                   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool                   `json:"IsMarkedForDeletion"`
	Partner                       []Partner               `json:"Partner"`
	Address                       []Address               `json:"Address"`
	Campaign                      []Campaign              `json:"Campaign"`
	Game                          []Game                  `json:"Game"`
	Participation	              []Participation		  `json:"Participation"`
	Attendance                    []Attendance            `json:"Attendance"`
	PointTransaction              []PointTransaction      `json:"PointTransaction"`
	PointConditionElement         []PointConditionElement `json:"PointConditionElement"`
	Counter						  []Counter				  `json:"Counter"`
}

type Partner struct {
	Event                   int     `json:"Event"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type Address struct {
	Event          int      `json:"Event"`
	AddressID      int      `json:"AddressID"`
	PostalCode     *string  `json:"PostalCode"`
	LocalSubRegion *string  `json:"LocalSubRegion"`
	LocalRegion    *string  `json:"LocalRegion"`
	Country        *string  `json:"Country"`
	GlobalRegion   *string  `json:"GlobalRegion"`
	TimeZone       *string  `json:"TimeZone"`
	District       *string  `json:"District"`
	StreetName     *string  `json:"StreetName"`
	CityName       *string  `json:"CityName"`
	Building       *string  `json:"Building"`
	Floor          *int     `json:"Floor"`
	Room           *int     `json:"Room"`
	XCoordinate    *float32 `json:"XCoordinate"`
	YCoordinate    *float32 `json:"YCoordinate"`
	ZCoordinate    *float32 `json:"ZCoordinate"`
	Site           *int     `json:"Site"`
}

type Campaign struct {
	Event               int     `json:"Event"`
	Campaign            int     `json:"Campaign"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsReleased          *bool   `json:"IsReleased"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Game struct {
	Event               int     `json:"Event"`
	Game                int     `json:"Game"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsReleased          *bool   `json:"IsReleased"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Participation struct {
	Event               int     `json:"Event"`
	Participator        int     `json:"Participator"`
	Participation       *int    `json:"Participation"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsReleased          *bool   `json:"IsReleased"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Attendance struct {
	Event               int     `json:"Event"`
	Attender            int     `json:"Attender"`
	Attendance          *int    `json:"Attendance"`
	Participation       *int    `json:"Participation"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsReleased          *bool   `json:"IsReleased"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type PointTransaction struct {
	Event                          int      `json:"Event"`
	Sender                         int      `json:"Sender"`
	Receiver                       int      `json:"Receiver"`
	PointConditionRecord           int      `json:"PointConditionRecord"`
	PointConditionSequentialNumber int      `json:"PointConditionSequentialNumber"`
	PointTransaction               *int     `json:"PointTransaction"`
	PointSymbol                    *string  `json:"PointSymbol"`
	PointTransactionType           *string  `json:"PointTransactionType"`
	PointConditionType             *string  `json:"PointConditionType"`
	PointConditionRateValue        *float32 `json:"PointConditionRateValue"`
	PointConditionRatio            *float32 `json:"PointConditionRatio"`
	PlusMinus                      *string  `json:"PlusMinus"`
	CreationDate                   *string  `json:"CreationDate"`
	CreationTime                   *string  `json:"CreationTime"`
	LastChangeDate                 *string  `json:"LastChangeDate"`
	LastChangeTime                 *string  `json:"LastChangeTime"`
	IsCancelled                    *bool    `json:"IsCancelled"`
}

type PointConditionElement struct {
	Event                          int      `json:"Event"`
	PointConditionRecord           int      `json:"PointConditionRecord"`
	PointConditionSequentialNumber int      `json:"PointConditionSequentialNumber"`
	PointSymbol                    *string  `json:"PointSymbol"`
	Sender                         *int     `json:"Sender"`
	PointTransactionType           *string  `json:"PointTransactionType"`
	PointConditionType             *string  `json:"PointConditionType"`
	PointConditionRateValue        *float32 `json:"PointConditionRateValue"`
	PointConditionRatio            *float32 `json:"PointConditionRatio"`
	PlusMinus                      *string  `json:"PlusMinus"`
	CreationDate                   *string  `json:"CreationDate"`
	LastChangeDate                 *string  `json:"LastChangeDate"`
	IsReleased                     *bool    `json:"IsReleased"`
	IsCancelled                    *bool    `json:"IsCancelled"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}

type Counter struct {
	Event                  int     `json:"Event"`
	NumberOfLikes          *int    `json:"NumberOfLikes"`
	NumberOfParticipations *int    `json:"NumberOfParticipations"`
	NumberOfAttendances    *int    `json:"NumberOfAttendances"`
	CreationDate           *string `json:"CreationDate"`
	CreationTime           *string `json:"CreationTime"`
	LastChangeDate         *string `json:"LastChangeDate"`
	LastChangeTime         *string `json:"LastChangeTime"`
}

func CreateEventRequestHeader(
	requestPram *apiInputReader.Request,
	eventHeader *apiInputReader.EventHeader,
) EventReq {
	req := EventReq{
		Header: Header{
			Event:               eventHeader.Event,
			IsReleased:          eventHeader.IsReleased,
			IsCancelled:         eventHeader.IsCancelled,
			IsMarkedForDeletion: eventHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateEventRequestHeaders(
	requestPram *apiInputReader.Request,
	eventHeaders *apiInputReader.EventHeader,
) EventReq {
	req := EventReq{
		Header: Header{
			IsReleased:          eventHeaders.IsReleased,
			IsCancelled:         eventHeaders.IsCancelled,
			IsMarkedForDeletion: eventHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreateEventRequestHeadersByEvents(
	requestPram *apiInputReader.Request,
	eventHeaders []Header,
) EventReq {
	req := EventReq{
		Headers: eventHeaders,
		Accepter: []string{
			"HeadersByEvents",
		},
	}
	return req
}

func CreateEventRequestHeadersByEventOwner(
	requestPram *apiInputReader.Request,
	eventHeaders *apiInputReader.EventHeader,
) EventReq {
	req := EventReq{
		Header: Header{
			EventOwner:			 eventHeaders.EventOwner,
			IsReleased:          eventHeaders.IsReleased,
			IsCancelled:         eventHeaders.IsCancelled,
			IsMarkedForDeletion: eventHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByEventOwner",
		},
	}
	return req
}

func CreateEventRequestPartner(
	requestPram *apiInputReader.Request,
	eventPartner *apiInputReader.EventPartner,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPartner.Event,
			Partner: []Partner{
				{
					PartnerFunction: eventPartner.PartnerFunction,
					BusinessPartner: eventPartner.BusinessPartner,
				},
			},
		},
		Accepter: []string{
			"Partner",
		},
	}
	return req
}

func CreateEventRequestPartners(
	requestPram *apiInputReader.Request,
	eventPartners *apiInputReader.EventPartner,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPartners.Event,
			Partner: []Partner{
				{
					//					IsMarkedForDeletion:           eventPartners.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Partners",
		},
	}
	return req
}

func CreateEventRequestAddress(
	requestPram *apiInputReader.Request,
	eventAddress *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventAddress.Event,
			Address: []Address{
				{
					AddressID: eventAddress.AddressID,
				},
			},
		},
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func CreateEventRequestAddresses(
	requestPram *apiInputReader.Request,
	eventAddresses *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventAddresses.Event,
			Address: []Address{
				{
					//					IsMarkedForDeletion:           eventAddresses.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Addresses",
		},
	}
	return req
}

func CreateEventRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	eventAddresses *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: eventAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegion",
		},
	}
	return req
}

func CreateEventRequestAddressesByLocalSubRegions(
	requestPram *apiInputReader.Request,
	eventAddresses *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: eventAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegions",
		},
	}
	return req
}

func CreateEventRequestAddressesByLocalRegion(
	requestPram *apiInputReader.Request,
	eventAddresses *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Address: []Address{
				{
					LocalRegion: eventAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegion",
		},
	}
	return req
}

func CreateEventRequestAddressesByLocalRegions(
	requestPram *apiInputReader.Request,
	eventAddresses *apiInputReader.EventAddress,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventAddresses.Event,
			Address: []Address{
				{
					LocalRegion: eventAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegions",
		},
	}
	return req
}

func CreateEventRequestCampaign(
	requestPram *apiInputReader.Request,
	eventCampaign *apiInputReader.EventCampaign,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventCampaign.Event,
			Campaign: []Campaign{
				{
					Campaign: eventCampaign.Campaign,
				},
			},
		},
		Accepter: []string{
			"Campaign",
		},
	}
	return req
}

func CreateEventRequestCampaigns(
	requestPram *apiInputReader.Request,
	eventCampaigns *apiInputReader.EventCampaign,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventCampaigns.Event,
			Campaign: []Campaign{
				{
					//					IsMarkedForDeletion:           eventCampaigns.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Campaigns",
		},
	}
	return req
}

func CreateEventRequestGame(
	requestPram *apiInputReader.Request,
	eventGame *apiInputReader.EventGame,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventGame.Event,
			Game: []Game{
				{
					Game: eventGame.Game,
				},
			},
		},
		Accepter: []string{
			"Game",
		},
	}
	return req
}

func CreateEventRequestGames(
	requestPram *apiInputReader.Request,
	eventGames *apiInputReader.EventGame,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventGames.Event,
			Game: []Game{
				{
					//					IsMarkedForDeletion:           eventGames.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Games",
		},
	}
	return req
}

func CreateEventRequestParticipation(
	requestPram *apiInputReader.Request,
	eventParticipation *apiInputReader.EventParticipation,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventParticipation.Event,
			Participation: []Participation{
				{
					Participator:		eventParticipation.Participator,
					IsCancelled:		eventParticipation.IsCancelled,
				},
			},
		},
		Accepter: []string{
			"Participation",
		},
	}
	return req
}

func CreateEventRequestAttendance(
	requestPram *apiInputReader.Request,
	eventAttendance *apiInputReader.EventAttendance,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventAttendance.Event,
			Attendance: []Attendance{
				{
					Attender:			eventAttendance.Attender,
					IsCancelled:		eventAttendance.IsCancelled,
				},
			},
		},
		Accepter: []string{
			"Attendance",
		},
	}
	return req
}

func CreateEventRequestPointTransaction(
	requestPram *apiInputReader.Request,
	eventPointTransaction *apiInputReader.EventPointTransaction,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPointTransaction.Event,
			PointTransaction: []PointTransaction{
				{
					Sender:                         eventPointTransaction.Sender,
					Receiver:                       eventPointTransaction.Receiver,
					PointConditionRecord:           eventPointTransaction.PointConditionRecord,
					PointConditionSequentialNumber: eventPointTransaction.PointConditionSequentialNumber,
				},
			},
		},
		Accepter: []string{
			"PointTransaction",
		},
	}
	return req
}

func CreateEventRequestPointTransactions(
	requestPram *apiInputReader.Request,
	eventPointTransactions *apiInputReader.EventPointTransaction,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPointTransactions.Event,
			PointTransaction: []PointTransaction{
				{
					//					IsCancelled:			eventPointTransactions.IsCancelled,
				},
			},
		},
		Accepter: []string{
			"PointTransactions",
		},
	}
	return req
}

func CreateEventRequestPointConditionElement(
	requestPram *apiInputReader.Request,
	eventPointConditionElement *apiInputReader.EventPointConditionElement,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPointConditionElement.Event,
			PointConditionElement: []PointConditionElement{
				{
					PointConditionRecord:           eventPointConditionElement.PointConditionRecord,
					PointConditionSequentialNumber: eventPointConditionElement.PointConditionSequentialNumber,
				},
			},
		},
		Accepter: []string{
			"PointConditionElement",
		},
	}
	return req
}

func CreateEventRequestPointConditionElements(
	requestPram *apiInputReader.Request,
	eventPointConditionElements *apiInputReader.EventPointConditionElement,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventPointConditionElements.Event,
			PointConditionElement: []PointConditionElement{
				{
					//					IsReleased:				eventPointConditionElements.IsReleased,
					//					IsCancelled:			eventPointConditionElements.IsCancelled,
					//					IsMarkedForDeletion:	eventPointConditionElements.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"PointConditionElements",
		},
	}
	return req
}

func CreateEventRequestCounter(
	requestPram *apiInputReader.Request,
	eventCounter *apiInputReader.EventCounter,
) EventReq {
	req := EventReq{
		Header: Header{
			Event: eventCounter.Event,
			Counter: []Counter{
				{
				},
			},
		},
		Accepter: []string{
			"Counter",
		},
	}
	return req
}

func CreateEventRequestCountersByEvents(
	requestPram *apiInputReader.Request,
	eventHeaders []Header,
) EventReq {
	req := EventReq{
		Headers: eventHeaders,
		Accepter: []string{
			"CountersByEvents",
		},
	}
	return req
}

func EventReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "reads"

	var request EventReq

	if accepter == "Header" {
		request = CreateEventRequestHeader(
			requestPram,
			&apiInputReader.EventHeader{
				Event:               input.EventHeader.Event,
				IsReleased:          input.EventHeader.IsReleased,
				IsCancelled:         input.EventHeader.IsCancelled,
				IsMarkedForDeletion: input.EventHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreateEventRequestHeaders(
			requestPram,
			&apiInputReader.EventHeader{
				IsReleased:          input.EventHeader.IsReleased,
				IsCancelled:         input.EventHeader.IsCancelled,
				IsMarkedForDeletion: input.EventHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByEventOwner" {
		request = CreateEventRequestHeadersByEventOwner(
			requestPram,
			&apiInputReader.EventHeader{
				EventOwner:			 input.EventHeader.EventOwner,
				IsReleased:          input.EventHeader.IsReleased,
				IsCancelled:         input.EventHeader.IsCancelled,
				IsMarkedForDeletion: input.EventHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Partner" {
		request = CreateEventRequestPartner(
			requestPram,
			&apiInputReader.EventPartner{
				Event:           input.EventPartner.Event,
				PartnerFunction: input.EventPartner.PartnerFunction,
				BusinessPartner: input.EventPartner.BusinessPartner,
			},
		)
	}

	if accepter == "Partners" {
		request = CreateEventRequestPartners(
			requestPram,
			&apiInputReader.EventPartner{
				Event: input.EventPartner.Event,
			},
		)
	}

	if accepter == "Address" {
		request = CreateEventRequestAddress(
			requestPram,
			&apiInputReader.EventAddress{
				Event:     input.EventAddress.Event,
				AddressID: input.EventAddress.AddressID,
			},
		)
	}

	if accepter == "Addresses" {
		request = CreateEventRequestAddresses(
			requestPram,
			&apiInputReader.EventAddress{
				Event: input.EventAddress.Event,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegion" {
		request = CreateEventRequestAddressesByLocalSubRegion(
			requestPram,
			&apiInputReader.EventAddress{
				LocalSubRegion: input.EventAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegions" {
		request = CreateEventRequestAddressesByLocalSubRegions(
			requestPram,
			&apiInputReader.EventAddress{
				LocalSubRegion: input.EventAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegion" {
		request = CreateEventRequestAddressesByLocalRegion(
			requestPram,
			&apiInputReader.EventAddress{
				LocalRegion: input.EventAddress.LocalRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegions" {
		request = CreateEventRequestAddressesByLocalRegions(
			requestPram,
			&apiInputReader.EventAddress{
				LocalRegion: input.EventAddress.LocalRegion,
			},
		)
	}

	if accepter == "Campaign" {
		request = CreateEventRequestCampaign(
			requestPram,
			&apiInputReader.EventCampaign{
				Event:    input.EventCampaign.Event,
				Campaign: input.EventCampaign.Campaign,
			},
		)
	}

	if accepter == "Campaigns" {
		request = CreateEventRequestCampaigns(
			requestPram,
			&apiInputReader.EventCampaign{
				Event: input.EventCampaign.Event,
			},
		)
	}

	if accepter == "Game" {
		request = CreateEventRequestGame(
			requestPram,
			&apiInputReader.EventGame{
				Event: input.EventGame.Event,
				Game:  input.EventGame.Game,
			},
		)
	}

	if accepter == "Games" {
		request = CreateEventRequestGames(
			requestPram,
			&apiInputReader.EventGame{
				Event: input.EventGame.Event,
			},
		)
	}

	if accepter == "Participation" {
		request = CreateEventRequestParticipation(
			requestPram,
			&apiInputReader.EventParticipation{
				Event:                          input.EventParticipation.Event,
				Participator:                   input.EventParticipation.Participator,
				IsCancelled:					input.EventParticipation.IsCancelled,
			},
		)
	}
	
	if accepter == "Attendance" {
		request = CreateEventRequestAttendance(
			requestPram,
			&apiInputReader.EventAttendance{
				Event:                          input.EventAttendance.Event,
				Attender:       	            input.EventAttendance.Attender,
				IsCancelled:					input.EventAttendance.IsCancelled,
			},
		)
	}

	if accepter == "PointTransaction" {
		request = CreateEventRequestPointTransaction(
			requestPram,
			&apiInputReader.EventPointTransaction{
				Event:                          input.EventPointTransaction.Event,
				Sender:                         input.EventPointTransaction.Sender,
				Receiver:                       input.EventPointTransaction.Receiver,
				PointConditionRecord:           input.EventPointTransaction.PointConditionRecord,
				PointConditionSequentialNumber: input.EventPointTransaction.PointConditionSequentialNumber,
			},
		)
	}

	if accepter == "PointTransactions" {
		request = CreateEventRequestPointTransactions(
			requestPram,
			&apiInputReader.EventPointTransaction{
				Event: input.EventPointTransaction.Event,
			},
		)
	}

	if accepter == "PointConditionElement" {
		request = CreateEventRequestPointConditionElement(
			requestPram,
			&apiInputReader.EventPointConditionElement{
				Event:                          input.EventPointConditionElement.Event,
				PointConditionRecord:           input.EventPointConditionElement.PointConditionRecord,
				PointConditionSequentialNumber: input.EventPointConditionElement.PointConditionSequentialNumber,
			},
		)
	}

	if accepter == "PointConditionElements" {
		request = CreateEventRequestPointConditionElements(
			requestPram,
			&apiInputReader.EventPointConditionElement{
				Event: input.EventPointConditionElement.Event,
			},
		)
	}

	if accepter == "Counter" {
		request = CreateEventRequestCounter(
			requestPram,
			&apiInputReader.EventCounter{
				Event:                          input.EventCounter.Event,
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

func EventReadsHeadersByEvents(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	//if accepter == "HeadersByEvents" {
	//	request = CreateEventRequestHeadersByEvents(
	//		requestPram,
	//		&apiInputReader.EventHeader{
	//			Event:               input.EventHeader.Event,
	//			IsReleased:          input.EventHeader.IsReleased,
	//			IsCancelled:         input.EventHeader.IsCancelled,
	//			IsMarkedForDeletion: input.EventHeader.IsMarkedForDeletion,
	//		},
	//	)
	//}

	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "reads"

	request := CreateEventRequestHeadersByEvents(
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

func EventReadsCountersByEvents(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	//if accepter == "CountersByEvents" {
	//	request = CreateEventRequestCountersByEvents(
	//		requestPram,
	//		&apiInputReader.EventCounter{
	//			Event:               input.EventCounter.Event,
	//		},
	//	)
	//}

	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "reads"

	request := CreateEventRequestCountersByEvents(
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
