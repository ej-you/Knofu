package urls

import (
	echo "github.com/labstack/echo/v4"

	coreMiddlewares "github.com/ej-you/Knofu/core/middlewares"
	"github.com/ej-you/Knofu/token/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/verify", handlers.Verify, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsAccessMiddleware)
	group.POST("/obtain", handlers.Obtain, coreMiddlewares.ValidateJWTMiddleware, coreMiddlewares.TokenIsRefreshMiddleware)
}
