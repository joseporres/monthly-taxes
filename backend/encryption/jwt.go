package encryption

import (
	"backend/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var revokedTokens = make(map[string]bool)

func RevokeToken(token string) {
	revokedTokens[token] = true
}

func IsTokenRevoked(token string) bool {
	return revokedTokens[token]
}

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(key))
}

func ParseLoginJWT(value string) (jwt.MapClaims, error) {
	// Check if the token has been revoked
	if IsTokenRevoked(value) {
		return nil, fmt.Errorf("el token ha sido revocado")
	}

	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	// Verifica tiempo exp
	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, fmt.Errorf("el campo 'exp' en el token no es válido")
		}

		// Obtiene el tiempo actual
		now := time.Now().Unix()

		// Comprueba si el token ha expirado
		if float64(now) > exp {
			return nil, fmt.Errorf("el token ha expirado")
		}

		// Si el token es válido y no ha expirado, devuelve las reclamaciones
		return claims, nil
	}

	return token.Claims.(jwt.MapClaims), nil
}

func Logout(e echo.HandlerFunc) {
	// Revoke the token associated with the session
	tokenString := c.Request().Header.Get("Authorization")
	RevokeToken(tokenString)
	return c.JSON(http.StatusOK, responseMessage{Message: "Logout successful"})
}
