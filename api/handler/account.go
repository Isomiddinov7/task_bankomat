package handler

import (
	"bankomat/api/models"
	"bankomat/config"
	"bankomat/pkg/helpers"
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateAccount godoc
// @ID create_account
// @Router /accounts [POST]
// @Summary Create Account
// @Description Create Account
// @Tags Account
// @Accept json
// @Produce json
// @Param object body models.CreateAccount true "CreateAccountRequestBody"
// @Success 200 {object} Response{data=models.Account} "AccountBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAccount(c *gin.Context) {
	var createAccount models.CreateAccount

	err := c.ShouldBindJSON(&createAccount)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = h.strg.Account().Create(ctx, &createAccount)
		if err != nil {
			h.logger.Error("Create Account error", zap.Error(err))
			return
		}
	}()
	wg.Wait()

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, http.StatusCreated, "OK!")
}

// Deposit godoc
// @ID deposit
// @Router /accounts/{id}/deposit [PUT]
// @Summary Deposit Account
// @Description Deposit Account
// @Tags Deposit
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param object body models.Deposit true "DepositRequestBody"
// @Success 200 {object} Response{data=models.Deposit} "Account"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Deposit(c *gin.Context) {
	var deposit models.Deposit

	err := c.ShouldBindJSON(&deposit)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "ShouldBindJSON err: "+err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	deposit.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = h.strg.Account().Deposit(ctx, &deposit)
		if err != nil {
			h.logger.Error("Deposit error", zap.Error(err))
			return
		}
	}()

	wg.Wait()

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Deposit failed")
		return
	}

	handleResponse(c, http.StatusAccepted, "OK!")
}

// Withdraw godoc
// @ID withdraw
// @Router /accounts/{id}/withdraw [POST]
// @Summary  Withdraw Account
// @Description Withdraw Account
// @Tags Withdraw
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param object body models.Withdraw true "WithdrawRequestBody"
// @Success 200 {object} Response{data=models.Withdraw} "Account"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Withdraw(c *gin.Context) {

	var withdraw models.Withdraw

	err := c.ShouldBindJSON(&withdraw)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	withdraw.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := h.strg.Account().Withdraw(ctx, &withdraw)
		if err != nil {
			h.logger.Error("Withdraw error", zap.Error(err))
			return
		}
	}()

	wg.Wait()

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	} else {
		handleResponse(c, http.StatusAccepted, "OK!")
	}

}

// GetBalance godoc
// @ID get_balance
// @Router /accounts/{id}/balance [GET]
// @Summary Get Balance
// @Description Get Balance
// @Tags Balance
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "GetListBranchResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetBalance(c *gin.Context) {
	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	var (
		wg   sync.WaitGroup
		resp float64
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp = h.strg.Account().GetBalance(ctx, id)
	}()
	wg.Wait()
	handleResponse(c, http.StatusOK, resp)
}
