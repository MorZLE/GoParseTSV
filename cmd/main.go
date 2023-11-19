package main

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/controller"
	"github.com/MorZLE/ParseTSVBiocad/internal/repository"
	"github.com/MorZLE/ParseTSVBiocad/internal/service"
	"github.com/MorZLE/ParseTSVBiocad/internal/watcher"
	"github.com/MorZLE/ParseTSVBiocad/logger"
)

func main() {
	logger.Initialize()

	conf := config.NewConfig()
	rep, err := repository.NewRepositoryImpl(conf)
	if err != nil {
		logger.Fatal("ошибка при создании репозитория:", err)
	}
	watch := watcher.NewWatcher(conf)

	logic := service.NewServiceImpl(conf, rep, watch)

	hand := controller.NewHandler(logic)

	logic.Scan()
	hand.Start()

}
