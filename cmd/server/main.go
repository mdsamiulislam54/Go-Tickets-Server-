package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gotickets/internal/model"
	"net/http"
	"os"

	// "github.com/labstack/echo/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

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
	err = db.AutoMigrate(model.User{})
	if err != nil {
		panic(`"failed to migrate database", "error", ${err}`)
	}
	fmt.Println("Connected to the database successfully!")
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World New")
	})
	e.GET("/users", func(c *echo.Context) error {
		var users []model.User
		if err := db.Find(&users).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "Failed to get user",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Get all users",
			"data":    users,
		})
	})
	e.GET("/users/:id", func(c *echo.Context) error {
		var user model.User
		if err := db.First(&user, c.Param("id")).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"error":   err.Error(),
				"message": "User not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Get user by ID",
			"data":    user,
		})
	})

	e.DELETE("/users/:id", func(c *echo.Context) error {
		var user model.User
		if err := db.Delete(&user, c.Param("id")).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"error":   err.Error(),
				"message": "User not found",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "User deleted successfully",
			"data":    user,
		})
	})

	e.POST("/users", func(c *echo.Context) error {
		dto := new(UserDTO)
		if err := c.Bind(dto); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		if err := c.Validate(dto); err != nil {

			validationErrors := make(map[string]string)

			if errs, ok := err.(validator.ValidationErrors); ok {
				for _, e := range errs {
					fmt.Println("Validation error:", e.Field(), e.Tag(), e.Param())
					switch e.Tag() {
					case "required":
						validationErrors[e.Field()] = e.Field() + " is required"
					case "email":
						validationErrors[e.Field()] = "Invalid email address"
					case "min":
						validationErrors[e.Field()] = e.Field() + " must be at least " + e.Param() + " characters"
					default:
						validationErrors[e.Field()] = "Invalid value"
					}
				}
			}

			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err.Error(),
				"message": validationErrors,
			})
		}

		user := model.User{
			Name:     dto.Name,
			Email:    dto.Email,
			Password: dto.Password,
		}
		// save user to database or perform other operations here
		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err.Error(),
				"message": "Failed to create user",
			})
		}

		return c.JSON(http.StatusOK, user)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
