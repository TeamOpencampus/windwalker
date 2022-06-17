package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"windwalker/middleware"
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
	db, err := models.SetupDatabase(config.Dsn, config.Mode)
	if err != nil {
		return err
	}

	router := gin.Default()
	if config.Mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := NewServer(router, db)

	return server.Run(config.Port)
}

type Server struct {
	engine   *gin.Engine
	database *gorm.DB
}

func (s *Server) Run(port string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	suger := logger.Sugar()
	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // disable in production
	}))
	s.engine.Use(middleware.WithDatabase(s.database), middleware.ErrorHandler(suger))
	s.Routes()
	return s.engine.Run(fmt.Sprintf(":%s", port))
}

func NewServer(engine *gin.Engine, database *gorm.DB) *Server {
	return &Server{engine: engine, database: database}
}
