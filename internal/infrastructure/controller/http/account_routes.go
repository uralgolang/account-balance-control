package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uralgolang/account-balance-control/internal/entity/DTO/request"
	"github.com/uralgolang/account-balance-control/internal/usecase"
	"net/http"
)

type AccountRoutes struct {
	AccountUseCase usecase.IAccountUseCase
}

func newAccountRoutes(router *gin.RouterGroup, accountUseCase usecase.IAccountUseCase) {
	accountRoutes := &AccountRoutes{accountUseCase}
	router.POST("/deposit-account", accountRoutes.DepositAccount)
}

func (accountRoutes *AccountRoutes) DepositAccount(c *gin.Context) {
	req := request.DepositAccountRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO обработка ошибок
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		c.Error(err)
		return
	}

	c.JSON(201, "ok")
}
