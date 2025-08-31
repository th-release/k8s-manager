package service

import (
	"strings"

	"cth.release/common/kubernetes"
	"cth.release/common/utils"
	"github.com/gofiber/fiber/v2"
	v1 "k8s.io/api/core/v1"
)

func CreateServiceRequest(c *fiber.Ctx) error {
	var dto CreateServiceRequestDto
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if len(strings.Trim(dto.Namespace, "")) <= 0 {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Please enter a valid namespace value.",
		})
	}

	client, err := kubernetes.NewK8sClient()

	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "A problem occurred while connecting to Kubernetes.",
		})
	}

	service, err := client.CreateService(dto.Namespace, &v1.Service{
		Spec: dto.ServiceSpec,
	})

	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(201).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
		Data:    service,
	})
}
