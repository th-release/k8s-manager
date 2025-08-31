package pod

import (
	types "k8s.io/api/core/v1"
)

type ListPodRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailPodRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
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

type PodLogRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
	Name      string `json:"name" query:"name"`
	Container string `json:"container" query:"container"`
	lines     int64  `json:"lines" query:"lines"`
}
