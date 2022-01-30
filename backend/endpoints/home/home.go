package home

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	println("Home")
	return c.SendString("Hello")
}
