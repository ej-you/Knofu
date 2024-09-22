package serializers

import (
	"fmt"
	"errors"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
)


// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Email 		string `json:"email" validate:"required|email"`
	FirstName 	string `json:"firstName" validate:"required"`
	LastName 	string `json:"lastName" validate:"required"`
	Password 	string `json:"password" validate:"required|min:8"`
}

// базовая валидация полей по тегам
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {
	baseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *RegisterUserIn) Validate() error {
	// базовая валидация полей по тегам
	validateErrors := validate.Validate(self)
	if len(validateErrors.Errors) > 0 {
		var errMessage string

		for key, value := range validateErrors.Errors {
			errMessage += fmt.Sprintf("%s: %s -- ", key, value[0])
		}
		errMessage = strings.TrimSuffix(errMessage, " -- ")
		// добавляем к сообщению код ошибки
		errorWithCode := "400||" + string(errMessage)
		return errors.New(errorWithCode)
	}

	return nil
}


// структура для выходных данных регистрации юзера
type RegisterUserOut struct {
	ID			uint64 `json:"id"`
	Email 		string `json:"email"`
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
}
