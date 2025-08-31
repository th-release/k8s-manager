package ingress

import "github.com/gofiber/fiber/v2"

func IngressMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	ingress := app.Group("/ingress", IngressMiddleware)

	ingress.Get("/list", ListIngressRequest)
	ingress.Get("/detail/:name", DetailIngressRequest)
	ingress.Post("/create", CreateIngressRequest)
	ingress.Put("/update", UpdateIngressRequest)
	ingress.Delete("/delete", DeleteIngressRequest)
}
