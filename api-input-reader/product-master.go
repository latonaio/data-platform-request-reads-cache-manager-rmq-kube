package apiInputReader

type ProductMaster struct {
	ProductMasterGeneral     *ProductMasterGeneral
}

type ProductMasterGeneral struct {
	Product                    string    `json:"Product"`
	IsMarkedForDeletion        *bool     `json:"IsMarkedForDeletion"`
}
