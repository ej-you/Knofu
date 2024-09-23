package error_handler

import (
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/settings"
)


// настройка обработчика ошибок
func CustomHTTPErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, context echo.Context) {
		httpErrorHandler(err, context)
	}
}


// настройка обработчика ошибок
func httpErrorHandler(err error, context echo.Context) {
	fmt.Println("[\nerr:", err, "\n]\n")

	// проверка, является ли ошибка err структурой *echo.HTTPError (приведение типов)
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		errorMap := make(map[string]string, 1)
		errorMap["unknown"] = err.Error()
		httpError = &echo.HTTPError{
			Code: 500,
			Message: errorMap,
		}
	}

	// поле timestamp
	strTime := time.Now().Format(settings.TimeFmt)
	// поле path
	requestPath := context.Path()

	errMessage := map[string]interface{}{
		"status": "error",
		"statusCode": httpError.Code,
		"path": requestPath,
		"timestamp": strTime,
		"errors": httpError.Message,
	}

	// отправка ответа
	var respErr error
	if !context.Response().Committed {
		// если метод HEAD
		if context.Request().Method == http.MethodHead {
			respErr = context.NoContent(httpError.Code)
		} else {
			respErr = context.JSON(httpError.Code, errMessage)
		}

		if respErr != nil {
			context.Echo().Logger.Error(respErr)
		}
	}
}