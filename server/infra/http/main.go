package main

import (
	"fmt"
	Api "fullVagas/infra/http/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes := Api.Target()
	routes.Init(server)

	err := server.Run("localhost:8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
