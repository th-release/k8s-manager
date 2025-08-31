package pod

import (
	types "k8s.io/api/core/v1"
)

type ListPodRequestDto struct {
	Namespace string `json:"namespace"`
}

type DetailPodRequestDto struct {
	Namespace string `json:"namespace"`
}

type CreatePodRequestDto struct {
	Namespace string        `json:"namespace"`
	PodSpec   types.PodSpec `json:"podSpec"`
}

type UpdatePodRequestDto struct {
	Namespace string        `json:"namespace"`
	PodSpec   types.PodSpec `json:"podSpec"`
}

type DeletePodRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
