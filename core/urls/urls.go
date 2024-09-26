package urls

import (
	echo "github.com/labstack/echo/v4"

	userUrls "github.com/Danil-114195722/Knofu/user/urls"
	tokenUrls "github.com/Danil-114195722/Knofu/token/urls"
	// shortsUrls "github.com/Danil-114195722/Knofu/shorts/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiUserGroup := echoApp.Group("/api/user")
	userUrls.RouterGroup(apiUserGroup)

	// apiShortsGroup := echoApp.Group("/api/shorts")
	// shortsUrls.RouterGroup(apiShortsGroup)

	apiTokenGroup := echoApp.Group("/api/token")
	tokenUrls.RouterGroup(apiTokenGroup)
}
