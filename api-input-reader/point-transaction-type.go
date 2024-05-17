package apiInputReader

type PointTransactionTypeGlobal struct {
	PointTransactionType     *PointTransactionType
	PointTransactionTypeText *PointTransactionTypeText
}

type PointTransactionType struct {
	PointTransactionType  string `json:"PointTransactionType"`
	IsMarkedForDeletion   *bool  `json:"IsMarkedForDeletion"`
}

type PointTransactionTypeText struct {
	PointTransactionType   string `json:"PointTransactionType"`
	Language               string `json:"Language"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}
