package settings

import (
	"strings"

	echo "github.com/labstack/echo/v4"

	userErrors "github.com/Danil-114195722/Knofu/user/errors"
)


// настройка распределителя обработчиков ошибок на разные микроприложения
func CustomHTTPErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, context echo.Context) {
		// путь запроса
		requestPath := context.Path()

		if strings.HasPrefix(requestPath, "/api/user") {
			// вызываем обработчик ошибок для микроприложения user
			userErrors.UserHTTPErrorHandler(err, context)
		} else {
			// вызываем дефолтный обработчик ошибок
			echoApp.DefaultHTTPErrorHandler(err, context)
		}
	}
}
