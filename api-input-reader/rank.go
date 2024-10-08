package apiInputReader

type RankGlobal struct {
	Rank     *Rank
	RankText *RankText
}

type Rank struct {
	RankType             string `json:"RankType"`
    Rank                 int    `json:"Rank"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type RankText struct {
	RankType            string `json:"RankType"`
    Rank                int    `json:"Rank"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
