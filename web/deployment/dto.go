package deployment

import (
	types "k8s.io/api/apps/v1"
)

type ListDeploymentRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailDeploymentRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateDeploymentRequestDto struct {
	Namespace      string               `json:"namespace"`
	DeploymentSpec types.DeploymentSpec `json:"deploymentSpec"`
}

type UpdateDeploymentRequestDto struct {
	Namespace      string               `json:"namespace"`
	DeploymentSpec types.DeploymentSpec `json:"deploymentSpec"`
}

type DeleteDeploymentRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type ScaleDeploymentRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Replicas  int32  `json:"replicas"`
}
