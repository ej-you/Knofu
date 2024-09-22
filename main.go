package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/Danil-114195722/Knofu/settings"
)


func main() {
	echoApp := echo.New()

	// если при запуске указан аргумент "dev"
	args := os.Args
	if len(args) > 1 {
		if args[1] == "dev" {
			echoApp.Debug = true
		}
	}

	// кастомизация логирования
	echoApp.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: settings.LogFmt,
	}))

	// настройка роутеров для эндпоинтов
	settings.InitUrlRouters(echoApp)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
