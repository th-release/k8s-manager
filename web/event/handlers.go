package event

import "github.com/gofiber/fiber/v2"

func eventMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	event := app.Group("/event", eventMiddleware)

	event.Post("/apply", ListEvent)
}
