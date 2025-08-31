package pod

import "github.com/gofiber/fiber/v2"

func podMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(app fiber.Router) {
	pod := app.Group("/pod", podMiddleware)

	pod.Get("/log", PodLogRequest)
	pod.Get("/list", ListPodRequest)
	pod.Get("/detail/:name", DetailPodRequest)
	pod.Post("/create", CreatePodRequest)
	pod.Put("/update", UpdatePodRequest)
	pod.Delete("/delete", DeletePodRequest)
}
