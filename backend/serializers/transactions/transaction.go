package transactions

type CreateTransactionRequest struct {
	Description     string  `json:"description,omitempty"`
	Amount          float64 `json:"amount" validate:"required,gte=1"`
	Category        string  `json:"category" validate:"required"`
	TransactionDate string  `json:"transaction_date" validate:"required,date"`
}
