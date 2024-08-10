package main

import (
	"casamento_api/src/config"
	"casamento_api/src/controller/routes"
	"casamento_api/src/model"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	conn := config.Connection()
	if err := conn.AutoMigrate(&model.Guest{}, &model.Present{}); err != nil {
		fmt.Printf("Erro ao sincronizar as tabelas, err: %v", err)
		panic(err)
	}

	router := gin.Default()
	configCors := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Create"},
		AllowCredentials: true,
	}

	router.Use(cors.New(configCors))
	routes.InitRoute(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
