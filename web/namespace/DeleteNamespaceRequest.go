package namespace

import (
	"strings"

	"cth.release/common/kubernetes"
	"cth.release/common/utils"
	"github.com/gofiber/fiber/v2"
)

func DeleteNamespaceRequest(c *fiber.Ctx) error {
	var dto DeleteNamespaceRequestDto
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

	if len(strings.Trim(dto.Name, "")) <= 0 {
		return c.Status(400).JSON(utils.BasicResponse{
			Success: false,
			Message: "Please enter a valid name value.",
		})
	}

	client, err := kubernetes.NewK8sClient()

	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: "A problem occurred while connecting to Kubernetes.",
		})
	}

	err = client.DeleteNamespace(dto.Namespace)
	if err != nil {
		return c.Status(500).JSON(utils.BasicResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(utils.BasicResponse{
		Success: true,
		Message: "",
	})
}
