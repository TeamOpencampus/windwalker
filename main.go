package main

import (
	"context"
	"log"
	"time"
	"windwalker/controllers"
	"windwalker/ent"
	"windwalker/middleware"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := SetupRouter(client).Run(); err != nil {
		log.Fatalf("failed to start server: %v ", err)
	}
}

func SetupRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		MaxAge:          12 * time.Hour,
	}))
	router.Use(middleware.ErrorHandler())
	controller := controllers.Controller{Client: client}
	controller.SetupRoutes(router)

	return router
}
