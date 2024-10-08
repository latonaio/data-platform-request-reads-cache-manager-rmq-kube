package apiModuleRuntimesResponsesRank

type RankRes struct {
	Message RankGlobal `json:"message,omitempty"`
}

type RankGlobal struct {
	Rank    *[]Rank    `json:"Rank,omitempty"`
	Text    *[]Text    `json:"Text,omitempty"`
}

type Rank struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	Language          	string  `json:"Language"`
	RankName			string 	`json:"RankName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
