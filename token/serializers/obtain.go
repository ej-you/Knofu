package serializers

import (
	echo "github.com/labstack/echo/v4"
    jwt "github.com/golang-jwt/jwt/v5"

	tokenErrors "github.com/Danil-114195722/Knofu/token/errors"
	"github.com/Danil-114195722/Knofu/token/services"
)


// валидация содержимого токена из контекста и возврат id юзера
func GetUserId(context echo.Context) (uint64, error) {
	var userId uint64

	// достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return userId, tokenErrors.InvalidTokenError
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return userId, tokenErrors.GetTokenClaimsError
    }

    // достаём из map'а токена id юзера
    userIdFloat, ok := tokenClaims["id"].(float64)
	if !ok {
	    return userId, tokenErrors.GetTokenUserIdError
	}

	return uint64(userIdFloat), nil
}

// структура для выходных данных выпуска нового access токена
type ObtainTokenOut struct {
	Status 		string `json:"status" example:"ok"`
	AccessToken string `json:"accessToken" example:"c73gnetfhigcsi.gaes4inva4a.gcawinxbwi4"`
}

// формирование структуры для ответа
func GetObtainOutStruct(userId uint64) (ObtainTokenOut, error) {
	// создаём новый access токен для юзера по его id
	accessToken, err := services.GetAccessToken(userId)
	if err != nil {
		return ObtainTokenOut{}, err
	}

	return ObtainTokenOut{
		Status: "ok",
		AccessToken: accessToken,
	}, nil
}
