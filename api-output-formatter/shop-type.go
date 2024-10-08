package apiOutputFormatter

type ShopType struct {
	ShopTypeShopType    []ShopTypeShopType    `json:"ShopTypeShopType"`
	ShopTypeText        []ShopTypeText        `json:"ShopTypeText"`
	Accepter            []string              `json:"Accepter"`
}

type ShopTypeShopType struct {
	ShopType            string	`json:"ShopType"`
}

type ShopTypeText struct {
	ShopType            string `json:"ShopType"`
	Language            string `json:"Language"`
	ShopTypeName        string `json:"ShopTypeName"`
}
