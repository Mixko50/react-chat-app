package endpoints

import (
	"backend/endpoints/home"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/", home.Home)
}
