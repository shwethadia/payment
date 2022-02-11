package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/shwethadia/payment/dto"
	"github.com/shwethadia/payment/entity"
	"github.com/shwethadia/payment/repository"
	"golang.org/x/crypto/bcrypt"
)

//AuthService
type AuthService interface {
	VerifyCredentials(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}


type authService struct {

	userRepository repository.UserRepository

}


func NewAuthService(userRep repository.UserRepository) AuthService {

	return &authService{

		userRepository: userRep,
	}
}


func (service *authService) VerifyCredentials(email string, password string) interface{} {

	res := service.userRepository.VerifyCredentials(email, password)
	
	if v, ok := res.(entity.User); ok {

		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {

			return res
		}
	}
	return res
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {

	userCreate := entity.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {

		log.Fatalf("Failed map %v", err)
	}

	res := service.userRepository.InsertUser(userCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {

	return service.userRepository.FindByEmail(email)

}

func (service *authService) IsDuplicateEmail(email string) bool {

	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)

}

func comparePassword(hashPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {

		log.Println(err)
		return false
	}

	return true
}
