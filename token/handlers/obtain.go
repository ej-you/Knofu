package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/Knofu/token/serializers"
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
//	@Success		200	{object}	serializers.ObtainTokenOut
//	@Failure		400	{object}	error	"BadRequest (See ErrorsDeafultSchema in README.md)"
//	@Failure		401	{object}	error	"Unauthorized (See ErrorsDeafultSchema in README.md)"
func Obtain(context echo.Context) error {
    // достаём из контекста из содержимого токена id юзера
    userId, err := serializers.GetUserId(context)
	if err != nil {
	    return err
	}

	// формирование структуры для ответа
    dataOut, err := serializers.GetObtainOutStruct(userId)
    if err != nil {
	    return err
	}

	return context.JSON(http.StatusOK, dataOut)
}
