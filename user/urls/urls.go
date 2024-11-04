package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/Knofu/user/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/register", handlers.Register)
	group.POST("/login", handlers.Login)
}
