package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"

	"github.com/labstack/echo/v5"
	// "github.com/labstack/echo/v5/middleware"
)

func main() {

	env := config.LoadEnv()
	db := config.ConnectDB(env)
	e := echo.New()
	server.StartServer(e, db, env)

	if err := e.Start(":" + env.Port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
