package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/ej-you/Knofu/core/middlewares"
	"github.com/ej-you/Knofu/shorts/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/test", handlers.Test)
	group.GET("/auth", handlers.Auth, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
}
