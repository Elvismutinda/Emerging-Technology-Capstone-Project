package transactions

type CreateTransactionRequest struct {
	Description     string  `json:"description,omitempty"`
	Amount          float64 `json:"amount" validate:"required,gte=1"`
	Category        string  `json:"category" validate:"required"`
	TransactionDate string  `json:"transaction_date" validate:"required"`
}

type UpdateTransactionRequest struct {
	Description     string  `json:"description,omitempty"`
	Amount          float64 `json:"amount" validate:"omitempty,gte=1"`
	Category        string  `json:"category,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	Type            string  `json:"type,omitempty" validate:"omitempty,oneof=INCOME EXPENSE"`
}

type FinancesOverviewResponse struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}

type FinancesStatsByCategoryResponse struct {
	DailyIncomes  []float64 `json:"daily_incomes"`
	DailyExpenses []float64 `json:"daily_expenses"`
}
