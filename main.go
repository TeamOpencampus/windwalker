package main

import (
	"log"
	"time"
	"windwalker/middleware"
	"windwalker/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db, err := models.SetupDatabase("windwalker.sqlite")
	if err != nil {
		log.Fatalln("Failed to connect to database: ", err)
	}
	if err := SetupRouter(db).Run(); err != nil {
		log.Fatalln("Failed to start: ", err)
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		MaxAge:          12 * time.Hour,
	}))
	router.Use(middleware.WithDatabase(db))
	router.Use(middleware.ErrorHandler())
	SetupRoutes(router)

	return router
}
