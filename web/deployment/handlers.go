package deployment

import "github.com/gofiber/fiber/v2"

func deploymentMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	deployment := app.Group("/deployment", deploymentMiddleware)

	deployment.Get("/list", ListDeploymentRequest)
	deployment.Get("/detail/:name", DetailDeploymentRequest)
	deployment.Post("/create", CreateDeploymentRequest)
	deployment.Put("/update", UpdateDeploymentRequest)
	deployment.Delete("/delete", DeleteDeploymentRequest)
	deployment.Put("/scale", ScaleDeployment)
}
