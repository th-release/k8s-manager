package configMap

import "github.com/gofiber/fiber/v2"

func ConfigMapMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	secret := app.Group("/configMap", ConfigMapMiddleware)

	secret.Get("/list", ListConfigMapRequest)
	secret.Get("/detail/:name", DetailConfigMapRequest)
	secret.Post("/create", CreateConfigMapRequest)
	secret.Put("/update", UpdateConfigMapRequest)
	secret.Delete("/delete", DeleteConfigMapRequest)
}
