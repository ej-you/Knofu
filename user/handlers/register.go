package handlers

import (
	// "fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/Knofu/user/serializers"
)


// эндпоинт для регистрации юзера
//	@Summary		Register user
//	@Description	Register new user with form
//	@Router			/user/register [post]
//	@ID				user-register
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		serializers.RegisterUserIn	true	"Register params"
//	@Success		200				{object}	serializers.RegisterUserOut
//	@Failure		400				{object}	error	"BadRequest (See ErrorsDeafultSchema in README.md)"
//	@Failure		500				{object}	error	"InternalServerError (See ErrorsDeafultSchema in README.md)"
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
