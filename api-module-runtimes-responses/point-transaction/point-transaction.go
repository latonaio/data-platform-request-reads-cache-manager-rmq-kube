package apiModuleRuntimesResponsesPointTransaction

type PointTransactionRes struct {
	Message PointTransaction `json:"message,omitempty"`
}

type PointTransaction struct {
	Header    *[]Header    `json:"Header"`
}

type Header struct {
	PointTransaction						int		`json:"PointTransaction"`
	PointTransactionType					string	`json:"PointTransactionType"`
	PointTransactionDate					string	`json:"PointTransactionDate"`
	PointTransactionTime					string	`json:"PointTransactionTime"`
	SenderObjectType						string	`json:"SenderObjectType"`
	SenderObject							int		`json:"SenderObject"`
	ReceiverObjectType						string	`json:"ReceiverObjectType"`
	ReceiverObject							int		`json:"ReceiverObject"`
	PointSymbol								string	`json:"PointSymbol"`
	PlusMinus								string	`json:"PlusMinus"`
	PointTransactionAmount					float32	`json:"PointTransactionAmount"`
	PointTransactionObjectType				string	`json:"PointTransactionObjectType"`
	PointTransactionObject					int		`json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction		float32	`json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction		float32	`json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction	float32	`json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction	float32	`json:"ReceiverPointBalanceAfterTransaction"`
	Attendance								*int	`json:"Attendance"`
	Participation							*int	`json:"Participation"`
	CreationDate							string	`json:"CreationDate"`
	CreationTime							string	`json:"CreationTime"`
	IsCancelled								*bool	`json:"IsCancelled"`
}
