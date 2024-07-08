package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var CorsOptions = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:     []string{"http://localhost:5173"},
	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	AllowCredentials: true,
})
