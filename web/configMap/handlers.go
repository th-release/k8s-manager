package configMap

import "github.com/gofiber/fiber/v2"

func ConfigMapMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	configMap := app.Group("/configMap", ConfigMapMiddleware)

	configMap.Get("/list", ListConfigMapRequest)
	configMap.Get("/detail/:name", DetailConfigMapRequest)
	configMap.Post("/create", CreateConfigMapRequest)
	configMap.Put("/update", UpdateConfigMapRequest)
	configMap.Delete("/delete", DeleteConfigMapRequest)
}
