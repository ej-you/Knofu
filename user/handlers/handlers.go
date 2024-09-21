package handlers

import (
	echo "github.com/labstack/echo/v4"

	// "github.com/Danil-114195722/Knofu/user/serializers"

	"github.com/Danil-114195722/Knofu/settings/constants"
)


// эндпоинт для регистрации юзера
func Register(context echo.Context) error {
	return context.String(constants.Status200, "Register endpoint")
}

// эндпоинт для входа юзера
func Login(context echo.Context) error {
	return context.String(constants.Status200, "Login endpoint")
}
