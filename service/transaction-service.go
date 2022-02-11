package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/shwethadia/payment/dto"
	"github.com/shwethadia/payment/entity"
	"github.com/shwethadia/payment/repository"
)

//AccountService
type TransactionService interface {
	Insert(b dto.TransactionCreateDTO) entity.Transaction
	All() []entity.Transaction
	FindById(transactionID uint64) entity.Transaction
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

//NewAccountService
func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {

	return &transactionService{

		transactionRepository: transactionRepo,
	}
}

func (service *transactionService) Insert(b dto.TransactionCreateDTO) entity.Transaction {

	transaction := entity.Transaction{}
	err := smapping.FillStruct(&transaction, smapping.MapFields(&b))
	if err != nil {

		log.Fatalf("Failed map %v", err)

	}
	res := service.transactionRepository.MakeTransaction(transaction)
	return res
}

func (service *transactionService) All() []entity.Transaction {

	return service.transactionRepository.AllTransaction()
}


func (service *transactionService) FindById(transactionID uint64) entity.Transaction {

	return service.transactionRepository.FindTransactionById(transactionID)
}
