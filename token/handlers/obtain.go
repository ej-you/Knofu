package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
    jwt "github.com/golang-jwt/jwt/v5"

	tokenErrors "github.com/Danil-114195722/Knofu/token/errors"
	"github.com/Danil-114195722/Knofu/token/services"
)


// эндпоинт для выпуска нового access токена по переданному в хедере refresh токену
//	@Summary		Obtain token
//	@Description	Obtain new accessToken with provided refreshToken in header
//	@Router			/token/obtain [post]
//	@ID				token-obtain
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Security		Refresh
func Obtain(context echo.Context) error {
	// достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return tokenErrors.InvalidTokenError
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return tokenErrors.GetTokenClaimsError
    }

    // достаём из map'а токена id юзера
    userIdFloat, ok := tokenClaims["id"].(float64)
	if !ok {
	    return tokenErrors.GetTokenUserIdError
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
