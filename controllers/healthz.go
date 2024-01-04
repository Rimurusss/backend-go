package controllers

import "github.com/gofiber/fiber/v2"

type healthzController interface {
	GetHealthz(c *fiber.Ctx) error
}

type HealthzController struct{}

func NewHealthzController() healthzController {
	return &HealthzController{}
}

func (h *HealthzController) GetHealthz(c *fiber.Ctx) error {
	return c.Status(200).SendString("OK")
}