package entity

type Account struct {
	ID          uint64  `gorm:"primary_key:auto_increment" json:"id"`
	AccountType string  `gorm:"type:varchar(255)" json:"account_type"`
	OpeningDate string  `gorm:"type:varchar(255)" json:"opening_date"`
	Amount      float64 `gorm:"type:decimal(7,6)" json:"amount"`
	UserID      uint64  `gorm:"not null" json:"-"`
	User        User    `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
