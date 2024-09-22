package serializers

import (
	"fmt"

	validate "github.com/gobuffalo/validate/v3"
)


// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Email 		string `json:"email" validate:"required|email"`
	FirstName 	string `json:"firstName" validate:"required"`
	LastName 	string `json:"lastName" validate:"required"`
	Password 	string `json:"password" validate:"required"`
}

// базовая валидация полей по тегам
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {
	baseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *RegisterUserIn) Validate() error {
	errors := validate.Validate(self)
	fmt.Println(errors)
	return nil
}


// структура для выходных данных регистрации юзера
type RegisterUserOut struct {
	ID			uint64 `json:"id"`
	Email 		string `json:"email"`
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
}

