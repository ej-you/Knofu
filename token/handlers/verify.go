package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/Knofu/token/serializers"
)


// эндпоинт для проверки access токена на валидность
//	@Summary		Verify token
//	@Description	Verify accessToken sended in header
//	@Router			/token/verify [post]
//	@ID				token-verify
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Success		200	{object}	serializers.VerifyTokenOut
//	@Failure		400	{object}	error	"BadRequest (See ErrorsDeafultSchema in README.md)"
//	@Failure		401	{object}	error	"Unauthorized (See ErrorsDeafultSchema in README.md)"
func Verify(context echo.Context) error {
	// вся проверка происходит в middlewares, поэтому если дело доходит до сюда, то проверка была пройдена
	return context.JSON(http.StatusOK, serializers.VerifyTokenOut{
		Status: "ok",
		StatusCode: http.StatusOK,
	})
}
