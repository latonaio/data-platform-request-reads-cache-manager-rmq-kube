package apiInputReader

type DeliveryDocument struct {
	DeliveryDocumentHeader     *DeliveryDocumentHeader
	DeliveryDocumentItems      *DeliveryDocumentItems
	DeliveryDocumentItem       *DeliveryDocumentItem
	DeliveryDocumentDocItemDoc *DeliveryDocumentDocItemDoc
}

type DeliveryDocumentHeader struct {
	DeliveryDocument                int     `json:"DeliveryDocument"`
	DeliverToParty                  *int    `json:"DeliverToParty"`
	DeliverFromParty                *int    `json:"DeliverFromParty"`
	HeaderCompleteDeliveryIsDefined *bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryBlockStatus       *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderDeliveryStatus            *string `json:"HeaderDeliveryStatus"`
	IsCancelled                     *bool   `json:"IsCancelled"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentItems struct {
	DeliveryDocument              int   `json:"DeliveryDocument"`
	DeliveryDocumentItem          *int  `json:"DeliveryDocumentItem"`
	ItemCompleteDeliveryIsDefined *bool `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryBlockStatus       *bool `json:"ItemDeliveryBlockStatus"`
	IsCancelled                   *bool `json:"IsCancelled"`
	IsMarkedForDeletion           *bool `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument              int   `json:"DeliveryDocument"`
	DeliveryDocumentItem          int   `json:"DeliveryDocumentItem"`
	ItemCompleteDeliveryIsDefined *bool `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryBlockStatus       *bool `json:"ItemDeliveryBlockStatus"`
	IsCancelled                   *bool `json:"IsCancelled"`
	IsMarkedForDeletion           *bool `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentDocItemDoc struct {
	DeliveryDocument         int    `json:"DeliveryDocument"`
	DeliveryDocumentItem     int    `json:"DeliveryDocumentItem"`
	DocType                  string `json:"DocType"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
