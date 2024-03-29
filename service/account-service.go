package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/shwethadia/payment/dto"
	"github.com/shwethadia/payment/entity"
	"github.com/shwethadia/payment/repository"
)

//AccountService
type AccountService interface {
	Insert(b dto.AccountCreateDTO) entity.Account
	Update(b dto.AccountUpdateDTO) entity.Account
	Delete(b entity.Account)
	All() []entity.Account
	FindByID(accountID uint64) entity.Account
	IsAllowedToEdit(userID string, accountID uint64) bool
}

type accountService struct {
	accountRepository repository.AccountRepository
}

//NewAccountService
func NewAccountService(accountRepo repository.AccountRepository) AccountService {

	return &accountService{

		accountRepository: accountRepo,
	}
}

func (service *accountService) Insert(b dto.AccountCreateDTO) entity.Account {

	account := entity.Account{}
	err := smapping.FillStruct(&account, smapping.MapFields(&b))
	fmt.Println(account)
	if err != nil {

		log.Fatalf("Failed map %v", err)

	}
	res := service.accountRepository.InsertAccount(account)
	return res
}

func (service *accountService) Update(b dto.AccountUpdateDTO) entity.Account {

	account := entity.Account{}
	err := smapping.FillStruct(&account, smapping.MapFields(&b))
	if err != nil {

		log.Fatalf("Failed map %v", err)

	}

	res := service.accountRepository.UpdateAccount(account)
	return res

}

func (service *accountService) Delete(b entity.Account) {

	service.accountRepository.DeleteAccount(b)
}

func (service *accountService) All() []entity.Account {

	return service.accountRepository.AllAccount()
}

func (service *accountService) FindByID(accountID uint64) entity.Account {

	return service.accountRepository.FindAccountByID(accountID)
}

func (service *accountService) IsAllowedToEdit(userID string, accountID uint64) bool {

	b := service.accountRepository.FindAccountByID(accountID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id

}
