package user

import (
	"gotickets/internal/auth"
	"gotickets/internal/config"
	"gotickets/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB, env *config.Config) {
	userRepo := NewUserRepository(db)
	jwtService := auth.NewJwtService(env.JwtSecret, 0)
	userService := NewUserService(userRepo, jwtService)
	userHandler := NewUserHandler(userService)
	api := e.Group("/api/v1/auth")
	
	api.GET("/user", userHandler.GetAllUser)
	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.loginUser)
	api.GET("/me",  userHandler.GetMe, middlewares.AuthMiddleware(jwtService))
}
