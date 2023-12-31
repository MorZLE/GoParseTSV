package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MorZLE/GoParseTSV/constants"
	"github.com/MorZLE/GoParseTSV/internal/model"
	"github.com/MorZLE/GoParseTSV/internal/service"
	"github.com/gofiber/fiber/v2"
)

// NewHandler создает новый экземпляр Handler
func NewHandler(s service.Service) *Handler {
	return &Handler{s}
}

type Handler struct {
	s service.Service
}

// Route Роут контроллера
func (h *Handler) Route(app *fiber.App) {
	app.Post("/", h.GetGuid)
}

// GetGuid API отправляет guids по уникальному номеру
func (h *Handler) GetGuid(c *fiber.Ctx) error {
	var req model.RequestGetGuid
	err := c.BodyParser(&req)
	if err != nil {
		return ErrorHandler(c, err)
	}

	body, err := h.s.GetAllGuid(req)

	if err != nil {
		return ErrorHandler(c, err)
	}
	return c.JSON(body)
}

// ErrorHandler обработчик ошибок
func ErrorHandler(c *fiber.Ctx, err error) error {
	var unmarshalTypeError *json.UnmarshalTypeError
	if err != nil {
		if errors.Is(err, unmarshalTypeError) {
			return c.Status(400).JSON(fiber.Map{
				"message": err,
			})
		}
		if errors.Is(err, constants.ErrEnabledData) {
			return c.Status(400).JSON(fiber.Map{
				"message": err,
			})
		}
		if errors.Is(err, constants.ErrNotFound) {
			err := fmt.Sprintf("not data %s", err)
			return c.Status(409).JSON(fiber.Map{
				"message": err,
			})
		}
		err := fmt.Sprintf("error get guid %s", err)
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.Status(200).SendString("success")
}
