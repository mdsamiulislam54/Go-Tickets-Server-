package server

import (
	"fmt"
	"gotickets/internal/config"
	"gotickets/internal/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
func StartServer( db *gorm.DB, cfg *config.Config) {
	e :=echo.New()
	e.Validator = &CustomValidator{validator: *validator.New()}
	err := db.AutoMigrate(user.UserDTO{})
	if err != nil {
		panic(`"failed to migrate database", "error", ${err}`)
	}
	fmt.Println("Connected to the database successfully!")

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World ")
	})

	user.RegisterRoute(e, db,cfg)

	
	if err := e.Start(":" + cfg.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
