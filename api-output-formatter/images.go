package apiOutputFormatter

type Images struct {
	BusinessPartner               *BusinessPartnerImage          `json:"BusinessPartner,omitempty"`
	Event                         *EventImage                    `json:"Event,omitempty"`
	Article                       *ArticleImage                  `json:"Article,omitempty"`
	Site                          *SiteImage                     `json:"Site,omitempty"`
	Shop                          *ShopImage                     `json:"Shop,omitempty"`
	Participation                 *ParticipationImage            `json:"Participation,omitempty"`
	Post                          *PostImage                     `json:"Post,omitempty"`
	Equipment                     *ProductImage                  `json:"Equipment,omitempty"`
	Product                       *ProductImage                  `json:"Product,omitempty"`
	Barcord                       *BarcordImage                  `json:"Barcode,omitempty"`
	QRCode                        *QRCodeImage                   `json:"QRCode,omitempty"`
	ProductionVersion             *ProductImage                  `json:"ProductionVersion,omitempty"`
	Operations                    *ProductImage                  `json:"Operations,omitempty"`
	DocumentImageBusinessPartner  *DocumentImageBusinessPartner  `json:"DocumentImageBusinessPartner,omitempty"`
	DocumentImageEvent            *DocumentImageEvent            `json:"DocumentImageEvent,omitempty"`
	DocumentImageArticle          *DocumentImageArticle          `json:"DocumentImageArticle,omitempty"`
	DocumentImageSite             *DocumentImageSite             `json:"DocumentImageSite,omitempty"`
	DocumentImageShop             *DocumentImageShop             `json:"DocumentImageShop,omitempty"`
	DocumentImageOrders           *DocumentImageOrders           `json:"DocumentImageOrders,omitempty"`
	DocumentImageDeliveryDocument *DocumentImageDeliveryDocument `json:"DocumentImageDeliveryDocument,omitempty"`
	DocumentImageBillOfMaterial   *DocumentImageBillOfMaterial   `json:"DocumentImageBillOfMaterial,omitempty"`
	DocumentImageInspectionLot    *DocumentImageInspectionLot    `json:"DocumentImageInspectionLot,omitempty"`
}

type BusinessPartnerImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type EventImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type ArticleImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type SiteImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type ShopImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type ParticipationImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type ProductImage struct {
	BusinessPartnerID int    `json:"BusinessPartnerID"`
	DocID             string `json:"DocID"`
	FileExtension     string `json:"FileExtension"`
}

type PostImage struct {
	Post          int    `json:"Post"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type MessageImage struct {
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

type DocumentImageBusinessPartner struct {
	BusinessPartner int    `json:"BusinessPartner"`
	DocType         string `json:"DocType"`
	DocID           string `json:"DocID"`
	FileExtension   string `json:"FileExtension"`
}

type DocumentImageEvent struct {
	Event         int    `json:"Event"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageArticle struct {
	Article       int    `json:"Article"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageSite struct {
	Site          int    `json:"Site"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageShop struct {
	Shop          int    `json:"Shop"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageOrders struct {
	OrderID       int    `json:"OrderID"`
	OrderItem     int    `json:"OrderItem"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}

type DocumentImageDeliveryDocument struct {
	DeliveryDocument     int    `json:"DeliveryDocument"`
	DeliveryDocumentItem int    `json:"DeliveryDocumentItem"`
	DocType              string `json:"DocType"`
	DocID                string `json:"DocID"`
	FileExtension        string `json:"FileExtension"`
}

type DocumentImageBillOfMaterial struct {
	BillOfMaterial     int    `json:"BillOfMaterial"`
	BillOfMaterialItem int    `json:"BillOfMaterialItem"`
	DocType            string `json:"DocType"`
	DocID              string `json:"DocID"`
	FileExtension      string `json:"FileExtension"`
}

type DocumentImageInspectionLot struct {
	InspectionLot int    `json:"InspectionLot"`
	DocType       string `json:"DocType"`
	DocID         string `json:"DocID"`
	FileExtension string `json:"FileExtension"`
}
