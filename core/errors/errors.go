package errors

import (
	echo "github.com/labstack/echo/v4"
)


// ошибка подключения к БД
var DBConnectError *echo.HTTPError = echo.NewHTTPError(500, map[string]string{"dbConnect": "Failed to connect to DB"})
