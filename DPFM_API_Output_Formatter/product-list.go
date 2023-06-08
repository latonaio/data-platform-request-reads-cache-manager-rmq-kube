package dpfm_api_output_formatter

type ProductList struct {
	Products []Product `json:"Products"`
}

type Product struct {
	Product string `json:"Product"`
	// ProductName        *string `json:"ProductName"`
	ProductDescription *string `json:"ProductDescription"`
	ProductGroup       *string `json:"ProductGroup"`
	ProductGroupName   *string `json:"ProductGroupName"`

	BaseUnit            *string `json:"BaseUnit"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Images              Images  `json:"Images"`
}
