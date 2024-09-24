package handlers

import (
	// "fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/serializers"
)


// эндпоинт для регистрации юзера
func Register(context echo.Context) error {
	var err error
	var dataIn serializers.RegisterUserIn

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = dataIn.Validate(); err != nil {
		return err
	}
	// создание нового юзера в БД
	newUser, err := dataIn.Create()
	if err != nil {
		return err
	}
	// формирование структуры для ответа
	dataOut, err := serializers.GetRegisterOutStruct(newUser)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, dataOut)
}
