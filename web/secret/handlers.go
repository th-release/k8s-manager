package secret

import "github.com/gofiber/fiber/v2"

func SecretMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	secret := app.Group("/secret", SecretMiddleware)

	secret.Get("/list", ListSecretRequest)
	secret.Get("/detail/:name", DetailSecretRequest)
	secret.Post("/create", CreateSecretRequest)
	secret.Put("/update", UpdateSecretRequest)
	secret.Delete("/delete", DeleteSecretRequest)
}
