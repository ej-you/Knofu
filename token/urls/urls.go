package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/Danil-114195722/Knofu/core/middlewares"
	"github.com/Danil-114195722/Knofu/token/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/verify", handlers.Verify, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	group.GET("/obtain", handlers.Obtain, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsRefreshMiddleware)
}
