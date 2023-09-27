package api

import echo "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {

	users := e.Group("/users")
	users.POST("/register", a.RegisterUser)
	users.POST("/login", a.LoginUser)
	users.GET("/monthly-expenses/{dni}", a.GetMonthlyTaxes)
}
