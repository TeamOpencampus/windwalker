package main

import (
	"fmt"
	"log"
	"windwalker/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("Failed to start: ", err)
	}
}

func run() error {
	config := LoadConfig()
	db, err := models.SetupDatabase(config.Dsn)
	if err != nil {
		return err
	}

	router := gin.Default()
	server := NewServer(router, db)

	return server.Run(config.Port)
}

type Server struct {
	engine   *gin.Engine
	database *gorm.DB
}

func (s *Server) Run(port string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // disable in production
	}))
	s.engine.Use(WithDatabase(s.database))
	s.Routes()
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
