package apiInputReader

type PointTransaction struct {
	PointTransactionHeader    *PointTransactionHeader
}

type PointTransactionHeader struct {
	PointTransaction    	int		`json:"PointTransaction"`
	PointTransactionType	*string	`json:"PointTransactionType"`
	SenderObjectType		*string	`json:"SenderObjectType"`
	SenderObject			*int	`json:"SenderObject"`
	ReceiverObjectType		*string	`json:"ReceiverObjectType"`
	ReceiverObject			*int	`json:"ReceiverObject"`
	IsCancelled    			*bool	`json:"IsCancelled"`
}
