package configMap

import "github.com/gofiber/fiber/v2"

func ConfigMapMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	deployment := app.Group("/configMap", ConfigMapMiddleware)

	deployment.Get("/list", ListConfigMapRequest)
	deployment.Get("/detail/:name", DetailConfigMapRequest)
	deployment.Post("/create", CreateConfigMapRequest)
	deployment.Put("/update", UpdateConfigMapRequest)
	deployment.Delete("/delete", DeleteConfigMapRequest)
}
