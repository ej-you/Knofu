package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
    jwt "github.com/golang-jwt/jwt/v5"

	"github.com/Danil-114195722/Knofu/token/services"
)


// эндпоинт для проверки access токена на валидность
func Verify(context echo.Context) error {
	// вся проверка происходит в middlewares, поэтому если дело доходит до сюда, то проверка была пройдена
	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"statusCode": http.StatusOK,
	})
}


// эндпоинт для выпуска нового access токена по переданному в хедере refresh токену
func Obtain(context echo.Context) error {
	// достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return echo.NewHTTPError(401, map[string]string{"token": "JWT token missing or invalid"})
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return echo.NewHTTPError(400, map[string]string{"token": "Failed to get claims from token"})
    }

    // достаём из map'а токена id юзера
    userIdFloat, ok := tokenClaims["id"].(float64)
	if !ok {
	    return echo.NewHTTPError(400, map[string]string{"token": "Failed to get user id from token"})
	}

    // создаём новый access токен для юзера по его id из refresh токена
	accessToken, err := services.GetAccessToken(uint64(userIdFloat))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"status": "ok",
		"accessToken": accessToken,
	})
}
