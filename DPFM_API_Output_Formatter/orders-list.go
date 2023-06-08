package dpfm_api_output_formatter

type OrdersList struct {
	Orders []Orders `json:"Orders"`
}
type Orders struct {
	OrderID        int     `json:"OrderID"`
	SellerName     string  `json:"SellerName"`
	Seller         *int    `json:"Seller"`
	BuyerName      string  `json:"BuyerName"`
	Buyer          *int    `json:"Buyer"`
	DeliveryStatus *string `json:"DeliveryStatus"`

	OrderDate           *string `json:"OrderDate"`
	PaymentTerms        *string `json:"PaymentTerms"`
	PaymentTermsName    *string `json:"PaymentTermsName"`
	PaymentMethod       *string `json:"PaymentMethod"`
	PaymentMethodName   *string `json:"PaymentMethodName"`
	TransactionCurrency *string `json:"TransactionCurrency"`
	OrderType           *string `json:"OrderType"`
	IsCancelled         bool    `json:"IsCancelled"`
	IsMarkedForDeletion bool    `json:"IsMarkedForDeletion"`
}
