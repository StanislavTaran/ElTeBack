package bank

import (
	companyMysql "back/internal/adapters/mysql/bank"
	"back/internal/domain/bank"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	getAllPath = "/list"
)

type Handler struct {
	bankService bankService
	logger      logger.ILogger
}

func NewBankHandler(client *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	bankStorage := companyMysql.NewBankStorage(client)
	bankService := bank.NewBankService(bankStorage)
	return &Handler{
		bankService: bankService,
		logger:      logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	group := e.Group("/bank")
	group.GET(getAllPath, h.getAllBanks())
	group.POST("", h.createBank())
	group.PUT("", h.updateBank())
	group.DELETE("/:id", h.removeBank())
}
