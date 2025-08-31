package namespace

import "github.com/gofiber/fiber/v2"

func namespaceMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	namespace := app.Group("/namespace", namespaceMiddleware)

	namespace.Post("/apply", ApplyNamespace)
}
