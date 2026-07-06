package inventory

type CreateItemRequest struct {
	ItemNumber  string `json:"itemNumber" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category"`

	UnitCost  float64 `json:"unitCost"`
	UnitPrice float64 `json:"unitPrice"`

	MinimumStock int `json:"minimumStock"`
	SafetyStock  int `json:"safetyStock"`
}

type CreateTransactionRequest struct {
	ItemID uint `json:"itemId" binding:"required"`

	TransactionType TransactionType `json:"transactionType" binding:"required"`

	Direction TransactionDirection `json:"direction" binding:"required"`

	Quantity int `json:"quantity" binding:"required"`

	Reference string `json:"reference"`

	Notes string `json:"notes"`
}

type StockResponse struct {
	ItemID uint `json:"itemId"`

	CurrentStock int `json:"currentStock"`

	Purchases   int `json:"purchases"`
	Sales       int `json:"sales"`
	Returns     int `json:"returns"`
	Adjustments int `json:"adjustments"`
	Damaged     int `json:"damaged"`
}
