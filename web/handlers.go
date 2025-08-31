package web

import (
	"cth.release/common/utils"
	"cth.release/web/deployment"
	"cth.release/web/pod"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App    *fiber.App
	Config utils.Config
}

func InitServer(config *utils.Config) *ServerConfig {
	app := fiber.New()

	if config == nil {
		return nil
	}

	server := &ServerConfig{
		App:    app,
		Config: *config,
	}

	server.SetupRoutes(app)
	return server
}

func (s *ServerConfig) SetupRoutes(app *fiber.App) *fiber.App {
	api := app.Group("/api", EmptyMiddleware)
	api.Get("/health", HealthHandler)

	pod.SetupRoutes(api)
	deployment.SetupRoutes(api)

	return app
}

func EmptyMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func HealthHandler(c *fiber.Ctx) error {
	return c.SendString("200")
}
