package apiInputReader

type ArticleTypeGlobal struct {
	ArticleType     *ArticleType
	ArticleTypeText *ArticleTypeText
}

type ArticleType struct {
	ArticleType         string `json:"ArticleType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ArticleTypeText struct {
	ArticleType         string `json:"ArticleType"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
