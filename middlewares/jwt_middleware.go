package middlewares

import (
	"miniproject/constants"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["password"] = password

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}
