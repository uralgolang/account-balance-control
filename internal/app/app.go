package app

import (
	"github.com/gin-gonic/gin"
	"github.com/uralgolang/account-balance-control/configs"
	"github.com/uralgolang/account-balance-control/internal/infrastructure/adapter/repository"
	"github.com/uralgolang/account-balance-control/internal/infrastructure/controller/http"
	"github.com/uralgolang/account-balance-control/internal/usecase"
	"github.com/uralgolang/account-balance-control/pkg/postgres"
	"log"
)

func Run(cfg *configs.Config) {
	postgres, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal("postgres init error")
		return
	}
	accountRepo := repository.NewAccountRepo(postgres)
	accountUseCase := usecase.NewAccountUseCase(accountRepo)

	router := gin.New()
	http.InitRouter(router, accountUseCase)
	//router := gin.Default()
	//router.Handlers()
	router.Run(cfg.Http)
}
