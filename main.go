package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/Danil-114195722/Knofu/settings"
	baseSettings "github.com/Danil-114195722/Knofu/settings/base"
)


func main() {
	echoApp := echo.New()
	echoApp.HideBanner = true

	// если при запуске указан аргумент "dev"
	args := os.Args
	if len(args) > 1 {
		if args[1] == "dev" {
			echoApp.Debug = true
			echoApp.HideBanner = false
		}
	}

	// кастомизация логирования
	echoApp.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: baseSettings.LogFmt,
	}))
	// отлавливание паник для беспрерывной работы сервиса
	echoApp.Use(echoMiddleware.Recover())

	// настройка кастомного обработчика ошибок
	settings.CustomHTTPErrorHandler(echoApp)
	// настройка роутеров для эндпоинтов
	settings.InitUrlRouters(echoApp)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", baseSettings.Port)))
}
