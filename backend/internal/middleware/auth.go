package middleware

import (
	"backend/encryption"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtén el token del encabezado de autorización
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		_, err := encryption.ParseLoginJWT(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		// Si el token es válido, continúa con la siguiente manejador
		return next(c)
	}
}
