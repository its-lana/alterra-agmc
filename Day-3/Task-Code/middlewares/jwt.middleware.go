package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //Token expires after 1 hour
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(encodedToken string) (int, error) {
	signatureKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(encodedToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signatureKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("token invalid")
	}

}
