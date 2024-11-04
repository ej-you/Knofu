package services

import (
	"time"

	echo "github.com/labstack/echo/v4"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/ej-you/Knofu/settings"
)


// создание access токена для юзера
func GetAccessToken(userId uint64) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": "access",
		"id": userId,
		"exp": time.Now().Add(settings.AccessTokenExpiredTime).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(settings.SecretForJWT))
	if err != nil {
		return "", echo.NewHTTPError(500, map[string]string{"token": err.Error()})
	}

	return tokenString, nil
}

// создание refresh токена для юзера
func GetRefreshToken(userId uint64) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": "refresh",
		"id": userId,
		"exp": time.Now().Add(settings.RefreshTokenExpiredTime).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(settings.SecretForJWT))
	if err != nil {
		return "", echo.NewHTTPError(500, map[string]string{"token": err.Error()})
	}

	return tokenString, nil
}
