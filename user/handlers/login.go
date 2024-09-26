package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/serializers"
)


// эндпоинт для входа юзера
//	@Summary		Login user
//	@Description	Login existing user by email and password
//	@Router			/user/login [post]
//	@ID				user-login
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		serializers.LoginUserIn	true	"Login params"
//	@Success		200				{object}	serializers.LoginUserOut
//	@Failure		400				{object}	error	"BadRequest (See ErrorsDeafultSchema in README.md)"
//	@Failure		500				{object}	error	"InternalServerError (See ErrorsDeafultSchema in README.md)"
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
