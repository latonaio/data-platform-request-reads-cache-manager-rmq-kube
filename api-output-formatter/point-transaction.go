package apiOutputFormatter

type PointTransaction struct {
	PointTransactionHeader []PointTransactionHeader `json:"PointTransactionHeader"`
	MountPath              *string                  `json:"mount_path"`
	Accepter               []string                 `json:"Accepter"`
}

type PointTransactionHeader struct {
	PointTransaction                      int     `json:"PointTransaction"`
	PointTransactionType                  string  `json:"PointTransactionType"`
	PointTransactionTypeName              string  `json:"PointTransactionTypeName"`
	PointTransactionDate                  string  `json:"PointTransactionDate"`
	PointTransactionTime                  string  `json:"PointTransactionTime"`
	SenderObjectType					  string  `json:"SenderObjectType"`
	SenderObject						  int	  `json:"SenderObject"`
//	SenderName                            string  `json:"SenderName"`
	ReceiverObjectType				      string  `json:"ReceiverObjectType"`
	ReceiverObject						  int	  `json:"ReceiverObject"`
//	ReceiverName                          string  `json:"ReceiverName"`
	PointSymbol                           string  `json:"PointSymbol"`
	PlusMinus                             string  `json:"PlusMinus"`
	PointTransactionAmount                float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            string  `json:"PointTransactionObjectType"`
	PointTransactionObjectTypeName        string  `json:"PointTransactionObjectTypeName"`
	PointTransactionObject                int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance							  *int	  `json:"Attendance"`
	Participation						  *int	  `json:"Participation"`
	CreationDate                          string  `json:"CreationDate"`
	CreationTime                          string  `json:"CreationTime"`
	IsCancelled                           *bool   `json:"IsCancelled"`
	Images                                Images  `json:"Images"`
}
