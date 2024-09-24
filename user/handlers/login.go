package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/serializers"
)


// эндпоинт для входа юзера
func Login(context echo.Context) error {
	var err error
	var dataIn serializers.LoginUserIn

	// парсинг JSON-body
	if err = context.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры и получение юзера из БД (если валидация успешна)
	userFromDB, err := dataIn.Validate()
	if err != nil {
		return err
	}
	// формирование структуры для ответа
	dataOut, err := serializers.GetLoginOutStruct(userFromDB)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, dataOut)
}
