package apiOutputFormatter

type ProductionOrder struct {
	ProductionOrderHeader []ProductionOrderHeader `json:"Header"`
	ProductionOrderItem   []ProductionOrderItem   `json:"Item"`
}

type ProductionOrderHeader struct {
	ProductionOrder                            int      `json:"ProductionOrder"`
	MRPArea                                    *string  `json:"MRPArea"`
	Product                                    string   `json:"Product"`
	ProductDescription                         string   `json:"ProductDescription"`
	OwnerProductionPlantBusinessPartner        int      `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantBusinessPartnerName    string   `json:"OwnerProductionPlantBusinessPartnerName"`
	OwnerProductionPlant                       string   `json:"OwnerProductionPlant"`
	OwnerProductionPlantName                   string   `json:"OwnerProductionPlantName"`
	ProductionOrderQuantityInBaseUnit          float32  `json:"ProductionOrderQuantityInBaseUnit"`
	IsReleased                                 *bool    `json:"IsReleased"`
	IsPartiallyConfirmed                       *bool    `json:"IsPartiallyConfirmed"`
	isConfirmed                                *bool    `json:"isConfirmed"`
	IsCancelled                                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                        *bool    `json:"IsMarkedForDeletion"`
	Images                                     Images   `json:"Images"`
}

type ProductionOrderItem struct {
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	MRPArea                              *string  `json:"MRPArea"`
	Product                              string   `json:"Product"`
	ProductDescription                   string   `json:"ProductDescription"`
	ProductionPlantBusinessPartner       int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantBusinessPartnerName   string   `json:"ProductionPlantBusinessPartnerName"`
	ProductionPlant                      string   `json:"ProductionPlant"`
	ProductionPlantName                  string   `json:"ProductionPlantName"`
	ProductionOrderQuantityInBaseUnit    float32  `json:"ProductionOrderQuantityInBaseUnit"`
    ConfirmedYieldQuantityInBaseUnit     *float32 `json:"ConfirmedYieldQuantityInBaseUnit"`
	IsPartiallyConfirmed                 *bool    `json:"IsPartiallyConfirmed"`
	isConfirmed                          *bool    `json:"isConfirmed"`
	IsCancelled                          *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                  *bool    `json:"IsMarkedForDeletion"`
	Images                               Images   `json:"Images"`
}
