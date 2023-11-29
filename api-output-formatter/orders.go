package apiOutputFormatter

type Orders struct {
	OrdersHeader             []OrdersHeader             `json:"Header"`
	OrdersHeaderWithItem     []OrdersHeaderWithItem     `json:"HeaderWithItem"`
	OrdersItem               []OrdersItem               `json:"Item"`
	OrdersSingleUnit         []OrdersSingleUnit         `json:"SingleUnit"`
	OrdersItemScheduleLine   []OrdersItemScheduleLine   `json:"ItemScheduleLine"`
	OrdersItemPricingElement []OrdersItemPricingElement `json:"ItemPricingElement"`
}

type OrdersHeader struct {
	OrderID              int     `json:"OrderID"`
	Buyer                int     `json:"Buyer"`
	BuyerName            string  `json:"BuyerName"`
	Seller               int     `json:"Seller"`
	SellerName           string  `json:"SellerName"`
	HeaderDeliveryStatus *string `json:"HeaderDeliveryStatus"`
	OrderType            string  `json:"OrderType"`
	IsCancelled          *bool   `json:"IsCancelled"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}

type OrdersHeaderWithItem struct {
	OrderID             int    `json:"OrderID"`
	OrderDate           string `json:"OrderDate"`
	PaymentTerms        string `json:"PaymentTerms"`
	PaymentTermsName    string `json:"PaymentTermsName"`
	PaymentMethod       string `json:"PaymentMethod"`
	TransactionCurrency string `json:"TransactionCurrency"`
	OrderType           string `json:"OrderType"`
	Buyer               int    `json:"Buyer"`
	BuyerName           string `json:"BuyerName"`
	Seller              int    `json:"Seller"`
	SellerName          string `json:"SellerName"`
}

type OrdersItem struct {
	OrderItem                   int     `json:"OrderItem"`
	Product                     string  `json:"Product"`
	Buyer                       int     `json:"Buyer"`
	BuyerName                   string  `json:"BuyerName"`
	Seller                      int     `json:"Seller"`
	SellerName                  string  `json:"SellerName"`
	OrderItemTextByBuyer        string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller       string  `json:"OrderItemTextBySeller"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
	DeliveryUnit                string  `json:"DeliveryUnit"`
	RequestedDeliveryDate       string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       string  `json:"RequestedDeliveryTime"`
	NetAmount                   float32 `json:"NetAmount"`
	IsCancelled                 *bool   `json:"IsCancelled"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
	Images                      Images  `json:"Images"`
}

type OrdersSingleUnit struct {
	OrderID               int     `json:"OrderID"`
	OrderItem             int     `json:"OrderItem"`
	Product               string  `json:"Product"`
	RequestedDeliveryDate string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime string  `json:"RequestedDeliveryTime"`
	GrossAmount           float32 `json:"GrossAmount"`
	OrderType             *string `json:"OrderType"`
	Buyer                 int     `json:"Buyer"`
	BuyerName             string  `json:"BuyerName"`
	Seller                int     `json:"Seller"`
	SellerName            string  `json:"SellerName"`
	OrderItemTextByBuyer  string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller string  `json:"OrderItemTextBySeller"`
	ConditionCurrency     string  `json:"ConditionCurrency"`
	Images                Images  `json:"Images"`
}

type OrdersItemScheduleLine struct {
	OrderID                              int     `json:"OrderID"`
	OrderItem                            int     `json:"OrderItem"`
	ScheduleLine                         int     `json:"ScheduleLine"`
	Product                              string  `json:"Product"`
	RequestedDeliveryDate                string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                string  `json:"RequestedDeliveryTime"`
	Buyer                                *int    `json:"Buyer"`
	BuyerName                            string  `json:"BuyerName"`
	Seller                               *int    `json:"Seller"`
	SellerName                           string  `json:"SellerName"`
	StockConfirmationBusinessPartner     int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationBusinessPartnerName string  `json:"StockConfirmationBusinessPartnerName"`
	StockConfirmationPlant               string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantName           string  `json:"StockConfirmationPlantName"`
	DeliveredQuantityInBaseUnit          float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit        float32 `json:"UndeliveredQuantityInBaseUnit"`
}

type OrdersItemPricingElement struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	PricingProcedureCounter int     `json:"PricingProcedureCounter"`
	ConditionRateValue      float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit  int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity  int     `json:"ConditionScaleQuantity"`
	ConditionCurrency       string  `json:"ConditionCurrency"`
	ConditionQuantity       float32 `json:"ConditionQuantity"`
	ConditionAmount         float32 `json:"ConditionAmount"`
	ConditionType           string  `json:"ConditionType"`
}
