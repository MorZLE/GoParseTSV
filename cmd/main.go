package main

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/controller"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"github.com/MorZLE/ParseTSVBiocad/repository"
	"github.com/MorZLE/ParseTSVBiocad/service"
	"github.com/MorZLE/ParseTSVBiocad/watcher"
)

func main() {
	logger.Initialize()

	conf := config.NewConfig()
	rep := repository.NewRepositoryImpl(conf)
	logic := service.NewServiceImpl(conf, rep)
	hand := controller.NewHandler(logic)
	wt := watcher.NewWatcher(conf, logic)

	wt.Scan()
	hand.Start()

}
