package main

import (
	"github.com/realjoni17/Cdlsiet/database"
	"github.com/realjoni17/Cdlsiet/server"
)

func main() {
	database.StartDatabase()
	server := server.NewServer()
	server.Run()
}
