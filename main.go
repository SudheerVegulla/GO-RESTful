package main

import (
	"example.com/event-bboking/db"
	"example.com/event-bboking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // this method configure the http server behind the screen

}
