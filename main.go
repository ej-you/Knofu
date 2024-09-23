package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	coreErrorHandler "github.com/Danil-114195722/Knofu/core/error_handler"
	coreUrls "github.com/Danil-114195722/Knofu/core/urls"
	"github.com/Danil-114195722/Knofu/core/db"
	"github.com/Danil-114195722/Knofu/settings"
)


func main() {
	echoApp := echo.New()
	echoApp.HideBanner = true

	// если при запуске указан аргумент "dev" или "migrate"
	args := os.Args
	if len(args) > 1 {
		// запуск в dev режиме
		if args[1] == "dev" {
			echoApp.Debug = true
			echoApp.HideBanner = false
		// проведение миграций БД без запуска самого приложения
		} else if args[1] == "migrate" {
			db.Migrate()
			return
		}
	}

	// кастомизация логирования
	echoApp.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: settings.LogFmt,
	}))
	// отлавливание паник для беспрерывной работы сервиса
	echoApp.Use(echoMiddleware.Recover())

	// настройка кастомного обработчика ошибок
	coreErrorHandler.CustomHTTPErrorHandler(echoApp)
	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(echoApp)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
