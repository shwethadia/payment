package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shwethadia/payment/config"
	"github.com/shwethadia/payment/controller"
	"github.com/shwethadia/payment/middleware"
	"github.com/shwethadia/payment/repository"
	"github.com/shwethadia/payment/service"
	"gorm.io/gorm"
)

var (
	db                    *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository        repository.UserRepository        = repository.NewUserRepository(db)
	accountRepository     repository.AccountRepository     = repository.NewAccountRepository(db)
	transactionRepository repository.TransactionRepository = repository.NewTransactionRepository(db)
	jwtService            service.JWTService               = service.NewJWTService()
	userService           service.UserService              = service.NewUserService(userRepository)
	accountService        service.AccountService           = service.NewAccountService(accountRepository)
	transactionService    service.TransactionService       = service.NewTransactionService(transactionRepository)
	authService           service.AuthService              = service.NewAuthService(userRepository)
	authController        controller.AuthController        = controller.NewAuthController(authService, jwtService)
	userController        controller.UserController        = controller.NewUserController(userService, jwtService)
	accountController     controller.AccountController     = controller.NewAccountController(accountService, jwtService)
	transactionController controller.TransactionController = controller.NewTransactionController(transactionService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))

	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	accountRoutes := r.Group("api/accounts", middleware.AuthorizeJWT(jwtService))
	{
		accountRoutes.GET("/", accountController.All)
		accountRoutes.POST("/", accountController.Insert)
		accountRoutes.GET("/:id", accountController.FindByID)
		accountRoutes.PUT("/:id", accountController.Update)
		accountRoutes.DELETE("/:id", accountController.Delete)
	}

	transactionRoutes := r.Group("api/transactions", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.GET("/", transactionController.All)
		transactionRoutes.POST("/withdraw", transactionController.Insert)
		transactionRoutes.POST("/deposit", transactionController.Insert)
		transactionRoutes.GET("/:id", transactionController.FindById)
	}

	r.Run()
}
