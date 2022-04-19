package bank

import (
	"back/internal/domain/bank"
	_ "back/internal/domain/bank"
	"back/internal/httpHelpers/httpResponse"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

const logLocation = "BANK CONTROLLER:"

func (h *Handler) getAllBanks() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		banks, err := h.bankService.GetAll(ctx)
		if err != nil {
			httpResponse.RequestErrCustomMessage(ctx, err, httpResponse.REQ_ERR_USER_NOT_FOUND)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, banks)
	}
}

func (h *Handler) createBank() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto bank.CreateBankInputDTO
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

		id, err := h.bankService.Create(ctx, dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessData(ctx, map[string]int64{"id": id})
	}
}

func (h *Handler) updateBank() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto bank.UpdateBankInputDTO
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

		err = h.bankService.Update(ctx, dto)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessOK(ctx)
	}
}

func (h *Handler) removeBank() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		value, err := strconv.Atoi(id)
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
		}
		err = h.bankService.Remove(ctx, int64(value))
		if err != nil {
			httpResponse.ErrorByType(ctx, err)
			h.logger.Error(logLocation + err.Error())
			return
		}

		httpResponse.SuccessOK(ctx)
	}
}
