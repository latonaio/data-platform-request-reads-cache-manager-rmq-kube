package apiModuleRuntimesResponsesArticleType

type ArticleTypeRes struct {
	Message ArticleTypeGlobal `json:"message,omitempty"`
}

type ArticleTypeGlobal struct {
	ArticleType  *[]ArticleType  `json:"ArticleType,omitempty"`
	Text         *[]Text         `json:"Text,omitempty"`
}

type ArticleType struct {
	ArticleType			string	`json:"ArticleType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	ArticleType     	string  `json:"ArticleType"`
	Language          	string  `json:"Language"`
	ArticleTypeName		string 	`json:"ArticleTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
