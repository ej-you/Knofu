package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/Knofu/user/handlers"
)


func UserRoutersGroup(userGroup *echo.Group) {
	userGroup.GET("/register", handlers.Register)
	userGroup.GET("/login", handlers.Login)
}
