package dto

//AccountCreateDTO
type TransactionCreateDTO struct {

	TransactionType string  `json:"transaction_type" form:"transaction_type" binding:"required"`
	TransactionDate string  `json:"transaction_date" form:"transaction_date" binding:"required"`
	Amount          float64 `json:"amount" form:"amount" binding:"required"`
	Balance         float64 `json:"balance" form:"balance,omitempty"`
	AccountID       uint64  `json:"account_id" form:"account_id"`
	
}
