package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func About(c *fiber.Ctx) error {
	return c.SendString("About")
}
