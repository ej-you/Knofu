package errors


import (
	"net/http"
	"strings"
	"strconv"
	"time"

	echo "github.com/labstack/echo/v4"

	baseSettings "github.com/Danil-114195722/Knofu/settings/base"
)


// разделение строки ошибки на код и само сообщение (например, 400||"error -- message")
func splitStatusAndMessage(errorStr string) (int, string) {
	errorToSlice := strings.Split(errorStr, "||")

	// если код ошибки не указан, то ставим по умолчанию 500
	if len(errorToSlice) == 1 {
		return 500, errorToSlice[0]
	} else {
		statusCode, _ := strconv.Atoi(errorToSlice[0])
		return statusCode, errorToSlice[1]
	}
}


// настройка обработчика ошибок
func UserHTTPErrorHandler(err error, context echo.Context) {
	// срез с кодом и описанием ошибки
	errorStatus, errorString := splitStatusAndMessage(err.Error())

	// поле timestamp
	strTime := time.Now().Format(baseSettings.TimeFmt)
	// поле path
	requestPath := context.Path()
	// поле message
	errMessage := map[string]interface{}{
		"status": errorStatus,
		"path": requestPath,
		"timestamp": strTime,
		"message": errorString,
	}

	// отправка ответа
	var respErr error
	if !context.Response().Committed {
		// если метод HEAD
		if context.Request().Method == http.MethodHead {
			respErr = context.NoContent(errorStatus)
		} else {
			respErr = context.JSON(errorStatus, errMessage)
		}

		if respErr != nil {
			context.Echo().Logger.Error(respErr)
		}
	}
}
