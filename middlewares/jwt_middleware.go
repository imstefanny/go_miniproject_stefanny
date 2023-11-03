package middlewares

import (
	"miniproject/constants"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(username, password, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}

func IsAdmin(e echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user.Valid {
			claims := user.Claims.(jwt.MapClaims)
			role := claims["role"].(string)
			if role == "admin" {
				return e(c)
			}
		}
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "Access Denied",
		})
	}
}
