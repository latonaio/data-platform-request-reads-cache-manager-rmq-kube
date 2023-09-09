package apiOutputFormatter

type Images struct {
	Equipment         *ProductImage `json:"Equipment,omitempty"`
	Product           *ProductImage `json:"Product,omitempty"`
	Barcord           *BarcordImage `json:"Barcode,omitempty"`
	QRCode            *QRCodeImage  `json:"QRCode,omitempty"`
	ProductionVersion *ProductImage `json:"ProductionVersion,omitempty"`
	Operations        *ProductImage `json:"Operations,omitempty"`
}

type ProductImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type BarcordImage struct {
	ID          string `json:"Id"`
	Barcode     string `json:"Barcode"`
	BarcodeType string `json:"BarcodeType"`
}

type QRCodeImage struct {
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}
