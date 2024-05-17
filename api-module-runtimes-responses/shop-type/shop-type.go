package apiModuleRuntimesResponsesShopType

type ShopTypeRes struct {
	Message ShopTypeGlobal `json:"message,omitempty"`
}

type ShopTypeGlobal struct {
	ShopType     *[]ShopType    `json:"ShopType,omitempty"`
	Text         *[]Text        `json:"Text,omitempty"`
}

type ShopType struct {
	ShopType			string	`json:"ShopType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	ShopType     		string  `json:"ShopType"`
	Language          	string  `json:"Language"`
	ShopTypeName		string 	`json:"ShopTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
