package main

import (
	"github.com/MorZLE/ParseTSVBiocad/config"
	"github.com/MorZLE/ParseTSVBiocad/internal/controller"
	"github.com/MorZLE/ParseTSVBiocad/internal/repository"
	"github.com/MorZLE/ParseTSVBiocad/internal/service"
	"github.com/MorZLE/ParseTSVBiocad/internal/workers"
	"github.com/MorZLE/ParseTSVBiocad/logger"
	"github.com/gofiber/fiber/v2"
	lg "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	logger.Initialize()

	conf := config.NewConfig()
	rep, err := repository.NewRepositoryImpl(conf)
	if err != nil {
		logger.Fatal("ошибка при создании репозитория:", err)
	}
	wScaner := workers.NewWatcher(conf)
	wWriter := workers.NewWriter(conf)
	wParser := workers.NewParser(conf)

	logic := service.NewServiceImpl(rep, wScaner, wWriter, wParser)
	hand := controller.NewHandler(logic)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(lg.New())

	hand.Route(app)

	// Start App
	go logic.Scan()

	err = app.Listen("127.0.0.1:8080")
	if err != nil {
		logger.Fatal("ошибка при запуске api:", err)
	}

}
