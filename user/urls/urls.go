package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/register", handlers.Register)
	group.GET("/login", handlers.Login)
}
