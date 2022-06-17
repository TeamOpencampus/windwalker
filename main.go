package main

import (
	"log"
	"windwalker/middleware"
	"windwalker/models"

	"go.uber.org/zap"

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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	suger := logger.Sugar()
	database, err := models.SetupDatabase("windwalker.sqlite")
	if err != nil {
		return err
	}

	router := gin.Default()
	server := NewServer(router, database, suger)

	return server.Run()
}

type Server struct {
	engine   *gin.Engine
	database *gorm.DB
	logger   *zap.SugaredLogger
}

func (s *Server) Run() error {

	s.engine.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // disable in production
	}))
	s.engine.Use(middleware.WithDatabase(s.database))
	s.engine.Use(middleware.ErrorHandler(s.logger))
	s.Routes()
	return s.engine.Run()
}

func NewServer(engine *gin.Engine, database *gorm.DB, logger *zap.SugaredLogger) *Server {
	return &Server{engine: engine, database: database, logger: logger}
}
