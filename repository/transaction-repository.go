package repository

import (
	"github.com/shwethadia/payment/entity"
	"gorm.io/gorm"
)

//AccountRepository
type TransactionRepository interface {
	MakeTransaction(b entity.Transaction) entity.Transaction
	AllTransaction() []entity.Transaction
	FindTransactionById(TransactionID uint64) entity.Transaction
}

type transactionConnection struct {
	connection *gorm.DB
}

//NewAccountRepository
func NewTransactionRepository(dbConn *gorm.DB) TransactionRepository {

	return &transactionConnection{

		connection: dbConn,
	}

}

func (db *transactionConnection) MakeTransaction(b entity.Transaction) entity.Transaction {

	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *transactionConnection) FindTransactionById(TransactionID uint64) entity.Transaction {

	var transaction entity.Transaction
	db.connection.Find(&transaction, TransactionID)
	return transaction
}


func (db *transactionConnection) AllTransaction() []entity.Transaction {

	var transactions []entity.Transaction
	db.connection.Find(&transactions)
	return transactions
}
