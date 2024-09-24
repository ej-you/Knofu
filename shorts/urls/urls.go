package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/Danil-114195722/Knofu/core/middlewares"
	"github.com/Danil-114195722/Knofu/shorts/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/test", handlers.Test)
	group.GET("/auth", handlers.Auth, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
}
