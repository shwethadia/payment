package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/shwethadia/payment/dto"
	"github.com/shwethadia/payment/entity"
	"github.com/shwethadia/payment/repository"
)


//UserService
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}


type userService struct {
	userRepository repository.UserRepository
}



func NewUserService(userRepo repository.UserRepository) UserService {

	return &userService{

		userRepository: userRepo,
	}
}


func (service *userService) Update(user dto.UserUpdateDTO) entity.User {

	userToUpdate := entity.User{}

	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser

}


func (service *userService) Profile(userID string) entity.User {

	return service.userRepository.ProfileUser((userID))
}


