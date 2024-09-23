package services

import (
	"time"

	echo "github.com/labstack/echo/v4"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/Danil-114195722/Knofu/settings"
)


// кодирование пароля в хэш
func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// возвращаем 400, потому что скорее всего ошибка длины пароля
		errorMap := make(map[string]string, 1)
		errorMap["encodePassword"] = err.Error()
		return "", echo.NewHTTPError(400, errorMap)
	}
	return string(hash), nil
}


// сравнение пароля и хэша
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


// создание токена для юзера
func GetJWTToken(userId uint64) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
		"expired": time.Now().Add(settings.TokenExpiredTime).Unix(),
	})

	tokenString, err := tokenStruct.SignedString([]byte(settings.SecretForJWT))
	if err != nil {
		errorMap := make(map[string]string, 1)
		errorMap["getJwtToken"] = err.Error()
		return "", echo.NewHTTPError(500, errorMap)
	}

	return tokenString, nil
}