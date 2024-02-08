package main

import (
	"fmt"
	"log"

	"example.com/rest-project/db"
	"example.com/rest-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("kia ora, starting server...")

	db.InitDb()
	fmt.Println("db connected")

	server := gin.Default()

	routes.RegisterRoutes(server)

	log.Fatal(server.Run(":8080"))
}
