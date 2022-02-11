package dto

//AccountUpdateDTO
type AccountUpdateDTO struct {
	ID          uint64  `json:"id" form:"id" binding:"required"`
	AccountType string  `json:"account_type" form:"account_type" binding:"required"`
	OpeningDate string  `json:"opening_date" form:"opening_date" binding:"required"`
	Amount      float64 `json:"amount" form:"amount" binding:"required"`
	UserID      uint64  `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//AccountCreateDTO
type AccountCreateDTO struct {
	AccountType string  `json:"account_type" form:"account_type" binding:"required"`
	OpeningDate string  `json:"opening_date" form:"opening_date" binding:"required"`
	Amount      float64 `json:"amount" form:"amount" binding:"required"`
	UserID      uint64  `json:"user_id,omitempty" form:"user_id,omitempty"`
}
