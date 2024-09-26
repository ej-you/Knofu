package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
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
func Verify(context echo.Context) error {
	// вся проверка происходит в middlewares, поэтому если дело доходит до сюда, то проверка была пройдена
	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"statusCode": http.StatusOK,
	})
}
