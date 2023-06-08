package dpfm_api_output_formatter

type DeliveryDocumentDetail struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	DeliveryDocumentItem    int     `json:"DeliveryDocumentItem"`
	PlannedGoodsReceiptDate *string `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime *string `json:"PlannedGoodsReceiptTime"`
	ActualGoodsReceiptDate  *string `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime  *string `json:"ActualGoodsReceiptTime"`

	ProductName string `json:"ProductName"`
	ProductCode string `json:"ProductCode"`
	Images      Images `json:"Images"`

	OrderQuantityInDelivery OrderQuantityInDelivery `json:"OrderQuantityInDelivery"`
	OrderQuantityInBase     OrderQuantityInBase     `json:"OrderQuantityInBase"`

	ProductTag []ProductTag `json:"ProductTag"`
	Stock      Stock        `json:"Stock"`

	StorageLocationFullName *string `json:"StorageLocationFullName"`
	StorageBin              string  `json:"StorageBin"`
	BestByDate              *string `json:"BestByDate"`
	ExpirationDate          *string `json:"ExpirationDate"`

	ProductInfo []ProductInfo `json:"ProductInfo"`

	BatchMgmtPolicyInDeliverToPlant      *string `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                  *string `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate *string `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime *string `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate   *string `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime   *string `json:"DeliverToPlantBatchValidityEndTime"`

	BatchMgmtPolicyInDeliverFromPlant      *string `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                  *string `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate *string `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime *string `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate   *string `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime   *string `json:"DeliverFromPlantBatchValidityEndTime"`

	PlannedGoodsIssueDate            *string                 `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime            *string                 `json:"PlannedGoodsIssueTime"`
	PlannedGoodsIssueQuantity        OrderQuantityInDelivery `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit   OrderQuantityInDelivery `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity      OrderQuantityInDelivery `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit OrderQuantityInDelivery `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate             *string                 `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime             *string                 `json:"ActualGoodsIssueTime"`
	ActualGoodsIssueQuantity         OrderQuantityInDelivery `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit    OrderQuantityInDelivery `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity       OrderQuantityInDelivery `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit  OrderQuantityInDelivery `json:"ActualGoodsReceiptQtyInBaseUnit"`
	OrderDetailJumpReq               OrderDetailJumpReq      `json:"OrdersDetailJumpReq"`
}
