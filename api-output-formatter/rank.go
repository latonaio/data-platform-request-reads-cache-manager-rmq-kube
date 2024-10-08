package apiOutputFormatter

type Rank struct {
	RankRank    []RankRank    `json:"RankRank"`
	RankText    []RankText    `json:"RankText"`
	Accepter    []string      `json:"Accepter"`
}

type RankRank struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
}

type RankText struct {
	RankType			string	`json:"RankType"`
	Rank				int		`json:"Rank"`
	Language          	string  `json:"Language"`
	RankName			string 	`json:"RankName"`
}
