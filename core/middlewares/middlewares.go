package middlewares

import (
	echo "github.com/labstack/echo/v4"
    echoJWT "github.com/labstack/echo-jwt/v4"

    jwt "github.com/golang-jwt/jwt/v5"

    coreErrorHandler "github.com/Danil-114195722/Knofu/core/error_handler"
    "github.com/Danil-114195722/Knofu/settings"
)


// middleware для распаковки содержимого токена в содержимое context'а запроса и валидации токена
var ValidateJWTMiddleware echo.MiddlewareFunc = echoJWT.WithConfig(echoJWT.Config{
    SigningKey: []byte(settings.SecretForJWT),
    ErrorHandler: coreErrorHandler.CustomJWTErrorHandler,
})

// проверка соответствия типа токена из контекста типу tokenType
func checkTokenType(context echo.Context, tokenType string) error {
    // достаём map значений JWT-токена из контекста context
    token, ok := context.Get("user").(*jwt.Token)
    if !ok {
        return echo.NewHTTPError(401, map[string]string{"token": "JWT token missing or invalid"})
    }
    tokenClaims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return echo.NewHTTPError(401, map[string]string{"token": "Failed to get claims from token"})
    }

    // приведение значенгия типа токена к string
    contextTokenType, ok := tokenClaims["type"].(string)
    if !ok {
        return echo.NewHTTPError(401, map[string]string{"token": "Invalid payload of JWT token"})
    }
    // проверка, что тип токена соответствует tokenType
    if contextTokenType != tokenType {
        return echo.NewHTTPError(401, map[string]string{"token": "Invalid JWT token"})
    }

    return nil
}


// middleware для проверки соответствия токена типу access
func TokenIsAccessMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(context echo.Context) error {
        // проверка типа токена
        if err := checkTokenType(context, "access"); err != nil {
            return err
        }
        return next(context)
    }
}

// middleware для проверки соответствия токена типу refresh
func TokenIsRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(context echo.Context) error {
        // проверка типа токена
        if err := checkTokenType(context, "refresh"); err != nil {
            return err
        }
        return next(context)
    }
}
