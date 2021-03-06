package server

import (
	"back/internal/adapters/api/bank"
	"back/internal/adapters/api/mortgage"
	"back/internal/adapters/middlewares"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  *Config
	Engine  *gin.Engine
	storage *mysqlClient.MySQLClient
	Logger  logger.ILogger
}

func NewServer(config *Config, logger logger.ILogger) *Server {
	return &Server{
		config: config,
		Logger: logger,
		Engine: gin.Default(),
	}
}

func (s *Server) Run() (err error) {
	err = s.configureMySQLStorage()
	if err != nil {
		return err
	}

	s.initRoutes()
	s.Logger.Debug("Routes mounted successfully.")

	err = s.Engine.Run(s.config.BindAddr)
	if err != nil {
		return err
	}

	s.Logger.Debug("Server started successfully")
	return nil
}

func (s *Server) configureMySQLStorage() error {
	storage := mysqlClient.NewMySQLClient(s.config.Mysql)

	err := storage.Open()
	if err != nil {
		return err
	}

	s.storage = storage
	return nil
}

func (s *Server) initRoutes() {
	s.Engine.Use(middlewares.CORSMiddleware())
	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	bankHandler := bank.NewBankHandler(s.storage, s.Logger)
	bankHandler.Register(s.Engine)

	mortgageHandler := mortgage.NewMortgageHandler(s.storage, s.Logger)
	mortgageHandler.Register(s.Engine)
}
