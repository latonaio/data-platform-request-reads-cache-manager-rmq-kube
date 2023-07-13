package apiOutputFormatter

type ProductMaster struct {
	ProductMasterGeneral		[]ProductMasterGeneral 			`json:"General"`
	ProductMasterDetailGeneral	[]ProductMasterDetailGeneral	`json:"General"`
}

type ProductMasterGeneral struct {
	Product             string  `json:"Product"`
	ProductName         *string `json:"ProductName"`
	ProductGroup        *string `json:"ProductGroup"`
	BaseUnit            string  `json:"BaseUnit"`
	ValidityStartDate   string  `json:"ValidityStartDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Images              Images  `json:"Images"`
}

type ProductMasterDetailGeneral struct {
	ProductType						string   `json:"ProductType"`
	GrossWeight						*float32 `json:"GrossWeight"`
	NetWeight						*float32 `json:"NetWeight"`
	WeightUnit						*string  `json:"WeightUnit"`
	InternalCapacityQuantity		*float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit	*string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText				*string  `json:"SizeOrDimensionText"`
	ProductStandardID				*string  `json:"ProductStandardID"`
	IndustryStandardName			*string  `json:"IndustryStandardName"`
	CountryOfOrigin					*string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage			*string  `json:"CountryOfOriginLanguage"`
	BarcodeType						*string  `json:"BarcodeType"`
	ProductAccountAssignmentGroup	*string  `json:"ProductAccountAssignmentGroup"`
	CreationDate					string   `json:"CreationDate"`
	LastChangeDate					string   `json:"LastChangeDate"`
	IsMarkedForDeletion				*bool    `json:"IsMarkedForDeletion"`
	Images							Images   `json:"Images"`
}
