package http

import (
	"github.com/gin-gonic/gin"
	"github.com/uralgolang/account-balance-control/internal/usecase"
)

func InitRouter(router *gin.Engine, accountUseCase usecase.IAccountUseCase) {
	groupRoutes := router.Group("/account")
	newAccountRoutes(groupRoutes, accountUseCase)
}
