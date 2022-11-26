package main

import (
	"github.com/uralgolang/account-balance-control/configs"
	"github.com/uralgolang/account-balance-control/internal/app"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal("config init error")
	}

	app.Run(cfg)
}
