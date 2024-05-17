package apiOutputFormatter

type AfterPointAcquisitionGlobal struct {
	AfterPointAcquisition []AfterPointAcquisition `json:"AfterPointAcquisition"`
	MountPath             *string                 `json:"mount_path"`
	Accepter              []string                `json:"Accepter"`
}

type AfterPointAcquisition struct {
	Event                          *int     `json:"Event"`
	BusinessPartner                *int     `json:"BusinessPartner"`
	PointSymbol                    *string  `json:"PointSymbol"`
	CurrentBalance                 *float32 `json:"CurrentBalance"`
	LimitBalance                   *float32 `json:"LimitBalance"`
	PointConditionRecord           *int     `json:"PointConditionRecord"`
	PointConditionSequentialNumber *int     `json:"PointConditionSequentialNumber"`
	Sender                         *int     `json:"Sender"`
	PointTransactionType           *string  `json:"PointTransactionType"`
	PointConditionType             *string  `json:"PointConditionType"`
	PointConditionRateValue        *float32 `json:"PointConditionRateValue"`
	PointConditionRatio            *float32 `json:"PointConditionRatio"`
	PlusMinus                      *string  `json:"PlusMinus"`
	Images                         Images   `json:"Images"`
}
