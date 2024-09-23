package urls

import (
	echo "github.com/labstack/echo/v4"

	userUrls "github.com/Danil-114195722/Knofu/user/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiUserGroup := echoApp.Group("/api/user")
	userUrls.UserRoutersGroup(apiUserGroup)
}
