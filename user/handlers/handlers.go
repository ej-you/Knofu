package handlers

import (
	"fmt"

	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/serializers"

	"github.com/Danil-114195722/Knofu/settings/constants"
)


// эндпоинт для регистрации юзера
func Register(context echo.Context) error {
	var err error
	var userData serializers.RegisterUserIn

	// парсинг JSON-body
	if err = context.Bind(&userData); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = userData.Validate(); err != nil {
		return err
	}

	fmt.Println(userData)

	return context.String(constants.Status200, "Register endpoint")
}

// эндпоинт для входа юзера
func Login(context echo.Context) error {
	return context.String(constants.Status200, "Login endpoint")
}
