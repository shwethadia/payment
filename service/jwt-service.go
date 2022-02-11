package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtservice can do
type JWTService interface {
	GenerateToken(UserId string) string
	ValidateToken(token string) (*jwt.Token, error)
}


type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}


type jwtService struct {
	secretKey string
	issuer    string
}


//NewJWTService
func NewJWTService() JWTService {

	return &jwtService{

		secretKey: getSecretKey(),
		issuer:    "shwetha",
	}

}


func getSecretKey() string {

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "hammilton"
	}
	return secretKey

}

func (j *jwtService) GenerateToken(UserId string) string {

	claims := &jwtCustomClaim{

		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}


func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {

		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected siging method %v", t_.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}
