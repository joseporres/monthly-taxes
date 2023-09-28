package api

import (
	echo "github.com/labstack/echo/v4"

	"backend/internal/middleware"
)

func (a *API) RegisterRoutes(e *echo.Echo) {

	auth := e.Group("/auth")
	auth.POST("/register", a.RegisterUser)
	auth.POST("/login", a.LoginUser)

	users := e.Group("/users")
	users.Use(middleware.AuthenticationMiddleware)
	users.GET("/monthly-expenses/:dni", a.GetMonthlyTaxes)
}
