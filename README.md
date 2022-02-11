# Simple Payment System

This repository for [Building Simple Payment Application  with Go using Data Driven Architecture]



- Built in Go version 1.16
- Uses the [HTTP Web Framework gin](github.com/gin-gonic/gin)
- Uses [dgrijalva Token management](github.com/dgrijalva/jwt-go/v4)
- Uses [smapping to map struct to interface](github.com/mashingan/smapping)
- Uses [crypto to encrypt and decrypt data](golang.org/x/crypto)
- Uses [mysqlDatabse](gorm.io/driver/mysql)
- Uses [ORMforGolang](gorm.io/gorm)

To run the application you have to define the environment variables, default values of the variables are defined inside the .env file

- DB_USER
- DB_PASSWARD
- DB_ADDR
- DB_PORT
- DB_NAME


You can use any one of the fallowing procedure to make a install dependencies,

- Run go mod init
- Run go mod tidy
- RUN go mod vendor[Option]
- docker-compose.yaml file.This contains the script to run the instance of applicaton 