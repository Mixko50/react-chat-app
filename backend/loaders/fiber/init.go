package fiber

import (
	"backend/endpoints"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func Init() {
	App = fiber.New()

	endpoints.Init(App)

	err := App.Listen(":8080")
	if err != nil {
		return
	}
}
