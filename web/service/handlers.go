package service

import "github.com/gofiber/fiber/v2"

func serviceMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	service := app.Group("/service", serviceMiddleware)

	service.Get("/list", ListServiceRequest)
	service.Get("/detail/:name", DetailServiceRequest)
	service.Post("/create", CreateServiceRequest)
	service.Put("/update", UpdateServiceRequest)
	service.Delete("/delete", DeleteServiceRequest)
}
