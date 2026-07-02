package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)
	api := e.Group("/api/v1/auth")
	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.loginUser)
}
