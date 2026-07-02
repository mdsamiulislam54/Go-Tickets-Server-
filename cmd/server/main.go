package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gotickets/internal/user"
	"net/http"
	"os"
	// "github.com/labstack/echo/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	// e.Use(middleware.RequestLogger())
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(user.UserDTO{})
	if err != nil {
		panic(`"failed to migrate database", "error", ${err}`)
	}
	fmt.Println("Connected to the database successfully!")
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World ")
	})

	e.POST("/user", userHandler.CreateUser)
	



	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
