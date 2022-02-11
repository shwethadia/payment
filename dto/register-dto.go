package dto

//RegisterDTO
type RegisterDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email" `
	Password    string `json:"password" form:"password" binding:"required"`
	City        string `json:"city" form:"city" binding:"required"`
	DateofBirth string `json:"dob" form:"dob" binding:"required"`
}
