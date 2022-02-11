package entity

type Transaction struct {
	ID              uint64  `gorm:"primary_key:auto_increment" json:"id"`
	TransactionType string  `gorm:"type:varchar(255)" json:"transaction_type"`
	TransactionDate string  `gorm:"type:varchar(255)" json:"transaction_date"`
	Amount          float64 `gorm:"type:decimal(7,6)" json:"amount"`
	Balance         float64 `gorm:"type:decimal(7,6)" json:"-"`
	AccountID       uint64  `gorm:"not null" json:"account_id"`
}
