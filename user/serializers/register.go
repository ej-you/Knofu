package serializers

import (
	// "fmt"
	// "errors"
	// "strings"

	echo "github.com/labstack/echo/v4"

	validate "github.com/gobuffalo/validate/v3"

	coreValidator "github.com/Danil-114195722/Knofu/core/validator"
)


// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Email 		string `json:"email" validate:"required|email"`
	FirstName 	string `json:"firstName" validate:"required"`
	LastName 	string `json:"lastName" validate:"required"`
	Password 	string `json:"password" validate:"required|min:8|max:50"`
}

// базовая валидация полей по тегам
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *RegisterUserIn) Validate() error {
	// базовая валидация полей по тегам
	var validateErrors *validate.Errors = validate.Validate(self)

	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
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
