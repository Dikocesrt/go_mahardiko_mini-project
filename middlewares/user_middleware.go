package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token) // Mendapatkan token pengguna dari konteks (asumsi menggunakan JWT)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "user" {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "Only users are allowed to access this endpoint",
			})
		}

		return next(c)
	}
}