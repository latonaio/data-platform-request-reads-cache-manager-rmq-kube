package apiInputReader

type Orders struct {
	OrdersHeader            *OrdersHeader
	OrdersItems             *OrdersItems
	OrdersItem              *OrdersItem
	OrdersItemScheduleLines *OrdersItemScheduleLines
	OrdersDocItemDoc        *OrdersDocItemDoc
}

type OrdersHeader struct {
	OrderID                         int     `json:"OrderID"`
	Buyer                           *int    `json:"Buyer"`
	Seller                          *int    `json:"Seller"`
	HeaderCompleteDeliveryIsDefined *bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryBlockStatus       *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderDeliveryStatus            *string `json:"HeaderDeliveryStatus"`
	IsCancelled                     *bool   `json:"IsCancelled"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type OrdersItems struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     *int    `json:"OrderItem"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryBlockStatus       *bool   `json:"ItemDeliveryBlockStatus"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	IsCancelled                   *bool   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type OrdersItem struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     int     `json:"OrderItem"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryBlockStatus       *bool   `json:"ItemDeliveryBlockStatus"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	IsCancelled                   *bool   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type OrdersItemScheduleLines struct {
	OrderID   int `json:"OrderID"`
	OrderItem int `json:"OrderItem"`
}

type OrdersDocItemDoc struct {
	OrderID                  int    `json:"OrderID"`
	OrderItem                int    `json:"OrderItem"`
	DocType                  string `json:"DocType"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
