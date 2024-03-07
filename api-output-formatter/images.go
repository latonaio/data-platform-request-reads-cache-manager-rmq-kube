package apiOutputFormatter

type Images struct {
	Equipment                     *ProductImage                  `json:"Equipment,omitempty"`
	Product                       *ProductImage                  `json:"Product,omitempty"`
	Barcord                       *BarcordImage                  `json:"Barcode,omitempty"`
	QRCode                        *QRCodeImage                   `json:"QRCode,omitempty"`
	ProductionVersion             *ProductImage                  `json:"ProductionVersion,omitempty"`
	Operations                    *ProductImage                  `json:"Operations,omitempty"`
	DocumentImageOrders           *DocumentImageOrders           `json:"DocumentImageOrders,omitempty"`
	DocumentImageDeliveryDocument *DocumentImageDeliveryDocument `json:"DocumentImageDeliveryDocument,omitempty"`
	DocumentImageInspectionLot    *DocumentImageInspectionLot    `json:"DocumentImageInspectionLot,omitempty"`
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

type DocumentImageOrders struct {
	DocType       string `json:"DocType"`
	OrdersID      int    `json:"OrdersID"`
	OrdersItem    int    `json:"OrdersItem"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageDeliveryDocument struct {
	DocType              string `json:"DocType"`
	DeliveryDocument     int    `json:"DeliveryDocument"`
	DeliveryDocumentItem int    `json:"DeliveryDocumentItem"`
	DocID                string `json:"DocID"`
	FileExtension        string `json:"FileExtension"`
}

type DocumentImageInspectionLot struct {
	DocType       string `json:"DocType"`
	InspectionLot int    `json:"InspectionLot"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}
