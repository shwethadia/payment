package dto

//UserUpdateDTO
type UserUpdateDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Password    string `json:"password,omitempty" form:"password,omitempty"`
	City        string `json:"city" form:"city" binding:"required"`
	DateofBirth string `json:"dob" form:"dob" binding:"required"`
}
