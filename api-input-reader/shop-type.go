package apiInputReader

type ShopTypeGlobal struct {
	ShopType     *ShopType
	ShopTypeText *ShopTypeText
}

type ShopType struct {
	ShopType            string `json:"ShopType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ShopTypeText struct {
	ShopType            string `json:"ShopType"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
