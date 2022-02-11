package controller

import (
	_ "fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shwethadia/payment/dto"
	"github.com/shwethadia/payment/entity"
	"github.com/shwethadia/payment/helper"
	"github.com/shwethadia/payment/service"
)

type TransactionController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
}

type transactionController struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

//NewController create a new instances of BoookController
func NewTransactionController(transactionServ service.TransactionService, jwtServ service.JWTService) TransactionController {
	return &transactionController{
		transactionService: transactionServ,
		jwtService:         jwtServ,
	}
}

func (c *transactionController) All(context *gin.Context) {

	var transactions []entity.Transaction = c.transactionService.All()
	res := helper.BuildResponse(true, "OK", transactions)
	context.JSON(http.StatusOK, res)

}

func (c *transactionController) FindById(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var transaction entity.Transaction = c.transactionService.FindById(id)
	if (transaction == entity.Transaction{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", transaction)
		context.JSON(http.StatusOK, res)
	}

}

func (c *transactionController) Insert(context *gin.Context) {

	var transactionCreateDTO dto.TransactionCreateDTO
	errDTO := context.ShouldBind(&transactionCreateDTO)

	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
	
		if strings.ToUpper(transactionCreateDTO.TransactionType) == "WITHDRAW" {

			transactionCreateDTO.Balance = transactionCreateDTO.Balance - transactionCreateDTO.Amount

		} else if strings.ToUpper(transactionCreateDTO.TransactionType) == "DEPOSIT" {

			transactionCreateDTO.Balance = transactionCreateDTO.Balance + transactionCreateDTO.Amount
		} else {

			transactionCreateDTO.Balance = 0
		}

		result := c.transactionService.Insert(transactionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)

		context.JSON(http.StatusCreated, response)

	}
}

