package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"
)

func main() {
	// Load environment variables
	env := config.LoadEnv()
	// Connect to the database
	db := config.ConnectDB(env)
	// Start the server
	server.StartServer(db, env)

}
