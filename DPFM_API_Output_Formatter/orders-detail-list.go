package dpfm_api_output_formatter

type OrdersDetailList struct {
	Header  OrdersDetailHeader  `json:"Header"`
	Details []OrdersItemSummary `json:"Details"`
}
type OrdersItemSummary struct {
	OrderItem                   int    `json:"OrderItem"`
	Product                     string `json:"Product"`
	OrderItemTextByBuyer        string `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller       string `json:"OrderItemTextBySeller"`
	OrderQuantityInDeliveryUnit string `json:"OrderQuantityInDeliveryUnit"`
	DeliveryUnit                string `json:"DeliveryUnit"`
	ConditionRateValue          string `json:"ConditionRateValue"`
	RequestedDeliveryDate       string `json:"RequestedDeliveryDate"`
	NetAmount                   string `json:"NetAmount"`
	IsCancelled                 *bool  `json:"IsCancelled"`
	IsMarkedForDeletion         *bool  `json:"IsMarkedForDeletion"`
	SupplyChainRelationshipID   int    `json:"SupplyChainRelationshipID"`
	PricingProcedureCounter     int    `json:"PricingProcedureCounter"`
}

type OrdersDetailHeader struct {
	Index         int             `json:"Index"`
	Key           string          `json:"Key"`
	PaymentTerms  []PaymentTerms  `json:"PaymentTerms"`
	PaymentMethod []PaymentMethod `json:"PaymentMethod"`
	Currency      []Currency      `json:"Currency"`
	QuantityUnit  []QuantityUnit  `json:"QuantityUnit"`
}

type Allergen struct {
	Name    string `json:"AllergenName"`
	Contain *bool  `json:"AllergenIsContained"`
}

type PaymentTerms struct {
	PaymentTerms     string `json:"PaymentTerms"`
	PaymentTermsName string `json:"PaymentTermsName"`
}
type PaymentMethod struct {
	PaymentMethod     string `json:"PaymentMethod"`
	PaymentMethodName string `json:"PaymentMethodName"`
}
type Currency struct {
	Currency     string `json:"Currency"`
	CurrencyName string `json:"CurrencyName"`
}
type QuantityUnit struct {
	QuantityUnit     string `json:"QuantityUnit"`
	QuantityUnitName string `json:"QuantityUnitName"`
}
