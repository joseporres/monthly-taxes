package api

import (
	"fmt"
	"log"
	"net/http"

	"backend/encryption"
	"backend/internal/api/dtos"
	"backend/internal/service"

	echo "github.com/labstack/echo/v4"
)

type responseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	fmt.Print("RegisterUser")
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, responseMessage{Message: "User already exists"})
		}

		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	response := map[string]string{"message": "Usuario creado"}

	return c.JSON(http.StatusCreated, response)
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LoginUser{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, responseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusUnauthorized, responseMessage{Message: "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	// cookie := &http.Cookie{
	// 	Name:     "Authorization",
	// 	Value:    token,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteNoneMode,
	// 	HttpOnly: true,
	// 	Path:     "/",
	// }

	// c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"success": "true", "token": token})
}

func (a *API) GetMonthlyTaxes(c echo.Context) error {
	//get dni from /monthly-taxes/{dni}
	//call service
	dni := c.Param("dni")
	ctx := c.Request().Context()
	response, err := a.serv.GetMonthlyTaxes(ctx, dni)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "true", "response": response})
}

func (a *API) Logout(c echo.Context) error {

	tokenHeader := c.Request().Header.Get("Authorization")
	response, err := a.serv.Logout(c.Request().Context(), tokenHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "Internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "true", "response": response})
}
