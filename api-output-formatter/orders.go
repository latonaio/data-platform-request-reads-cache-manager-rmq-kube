package apiOutputFormatter

type Orders struct {
	OrdersHeader  []OrdersHeader  `json:"Header"`
	OrdersItem    []OrdersItem    `json:"Item"`
}

type OrdersHeader struct {
	OrderID                   int     `json:"OrderID"`
	Buyer                     int     `json:"Buyer"`
	BuyerName                 string  `json:"BuyerName"`
	Seller                    int     `json:"Seller"`
	SellerName                string  `json:"SellerName"`
    HeaderDeliveryStatus      *string `json:"HeaderDeliveryStatus"`
	OrderType                 string  `json:"OrderType"`
    IsCancelled               *bool   `json:"IsCancelled"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type OrdersHeaderWithItem struct {
	OrderID                   int     `json:"OrderID"`
	OrderDate                 string  `json:"OrderDate"`
	PaymentTerms			  string  `json:"PaymentTerms"`
	PaymentTermsName		  string  `json:"PaymentTermsName"`
	PaymentMethod			  string  `json:"PaymentMethod"`
	TransactionCurrency		  string  `json:"TransactionCurrency"`
	OrderType                 string  `json:"OrderType"`
	Buyer                     int     `json:"Buyer"`
	BuyerName                 string  `json:"BuyerName"`
	Seller                    int     `json:"Seller"`
	SellerName                string  `json:"SellerName"`
    IsCancelled               *bool   `json:"IsCancelled"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type OrdersItem struct {
	OrderItem                     int     `json:"OrderItem"`
	Product                       string  `json:"Product"`
    OrderItemTextByBuyer          string  `json:"OrderItemTextByBuyer"`
    OrderItemTextBySeller         string  `json:"OrderItemTextBySeller"`
    OrderQuantityInDeliveryUnit   float32 `json:"OrderQuantityInDeliveryUnit"`
    DeliveryUnit                  string  `json:"DeliveryUnit"`
    RequestedDeliveryDate         string  `json:"RequestedDeliveryDate"`
    NetAmount                     float32 `json:"NetAmount"`
    IsCancelled                   *bool   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
	Images                        Images  `json:"Images"`
}
