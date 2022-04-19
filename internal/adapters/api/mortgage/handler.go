package mortgage

import (
	companyMysql "back/internal/adapters/mysql/bank"
	"back/internal/domain/bank"
	"back/internal/domain/mortgage"
	"back/pkg/logger"
	"back/pkg/mysqlClient"
	"github.com/gin-gonic/gin"
)

const (
	calculatePath = "/calculate"
)

type Handler struct {
	mortgageService mortgageService
	logger          logger.ILogger
}

func NewMortgageHandler(client *mysqlClient.MySQLClient, logger logger.ILogger) *Handler {
	bankStorage := companyMysql.NewBankStorage(client)
	bankService := bank.NewBankService(bankStorage)
	mortgageService := mortgage.NewMortgageService(bankService)
	return &Handler{
		mortgageService: mortgageService,
		logger:          logger,
	}
}

func (h *Handler) Register(e *gin.Engine) {
	group := e.Group("/mortgage")
	group.POST(calculatePath, h.calculateMortgagePayment())
}
