package dpfm_api_output_formatter

type ProductDetailList struct {
	Header     ProductDetailHeader `json:"Header"`
	Existences []Exconf            `json:"Existences"`
}

type Exconf struct {
	Content string      `json:"Content"`
	Exist   *bool       `json:"Exist"`
	Param   interface{} `json:"Param"`
}

type ProductDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
