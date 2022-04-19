package mortgage

import (
	_ "back/internal/domain/bank"
	"back/internal/domain/mortgage"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const logLocation = "MORTGAGE CONTROLLER:"

func (h *Handler) calculateMortgagePayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto mortgage.CalculateMortgagePaymentInputDTO
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = json.Unmarshal(body, &dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		err = dto.Validate()
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.VALIDATION_ERR)
			h.logger.Error(logLocation + err.Error())
			return
		}

		data, err := h.mortgageService.CalculateMonthlyPayment(ctx, dto)
		if err != nil {
			httpResponse.RequestErr(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, data)
	}
}
