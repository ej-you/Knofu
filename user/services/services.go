package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Danil-114195722/Knofu/settings"
)


// кодирование пароля в хэш
func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		// лог об ошибке
		settings.ErrorLog.Println(err)

		// возвращаем 400, потому что скорее всего ошибка длины пароля
		newErrorMessage := "400||" + err.Error()
		return "", errors.New(newErrorMessage)
	}
	return string(hash), nil
}


// сравнение пароля и хэша
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
