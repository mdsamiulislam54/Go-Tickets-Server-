package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	e.POST("/user", userHandler.CreateUser)
}
