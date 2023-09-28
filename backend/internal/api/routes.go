package api

import (
	"backend/internal/middleware"

	echo "github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {

	auth := e.Group("/auth")
	auth.POST("/register", a.RegisterUser)
	auth.POST("/login", a.LoginUser)
	auth.POST("/logout", a.Logout)

	users := e.Group("/users")
	users.Use(middleware.AuthenticationMiddleware)
	users.GET("/monthly-expenses/:dni", a.GetMonthlyTaxes)
}
