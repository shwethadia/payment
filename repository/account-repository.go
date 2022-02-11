package repository

import (
	"fmt"

	"github.com/shwethadia/payment/entity"
	"gorm.io/gorm"
)

//AccountRepository
type AccountRepository interface {
	InsertAccount(b entity.Account) entity.Account
	UpdateAccount(b entity.Account) entity.Account
	DeleteAccount(b entity.Account)
	AllAccount() []entity.Account
	FindAccountByID(AccountID uint64) entity.Account
}


type accountConnection struct {
	connection *gorm.DB
}


//NewAccountRepository
func NewAccountRepository(dbConn *gorm.DB) AccountRepository {

	return &accountConnection{

		connection: dbConn,
	}

}



func (db *accountConnection) InsertAccount(b entity.Account) entity.Account {

	fmt.Println(b)
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}



func (db *accountConnection) UpdateAccount(b entity.Account) entity.Account {

	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}



func (db *accountConnection) DeleteAccount(b entity.Account) {

	db.connection.Delete(&b)

}



func (db *accountConnection) FindAccountByID(AccountID uint64) entity.Account {

	var account entity.Account
	db.connection.Preload("User").Find(&account, AccountID)
	return account
}



func (db *accountConnection) AllAccount() []entity.Account {

	var accounts []entity.Account
	db.connection.Preload("User").Find(&accounts)
	return accounts
}
