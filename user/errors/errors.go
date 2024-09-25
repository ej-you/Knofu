package errors

import (
	echo "github.com/labstack/echo/v4"
)



// ошибка при создании юзера в БД с уже зарегистрированным email'ом
var UserAlreadyExistsError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"email": "User with given email already exists"})

// ошибка при нахождении юзера в БД по email'у
var UserDoesNotExistError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"email": "User with given email does not exist"})

// ошибка при сравнении введённого юзером пароля и хэша из БД
var InvalidPasswordError *echo.HTTPError = echo.NewHTTPError(400, map[string]string{"password": "Invalid password"})
