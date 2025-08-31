package namespace

import "github.com/gofiber/fiber/v2"

func namespaceMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	namespace := app.Group("/namespace", namespaceMiddleware)
	namespace.Get("/list", ListNamespaceRequest)
	namespace.Get("/detail/:namespace", DetailNamespaceRequest)
	namespace.Post("/create", CreateNamespaceRequest)
	namespace.Delete("/delete", DeleteNamespaceRequest)

	namespace.Post("/apply", ApplyNamespace)
}
