package service

import (
	types "k8s.io/api/core/v1"
)

type ListServiceRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailServiceRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateServiceRequestDto struct {
	Namespace   string            `json:"namespace"`
	ServiceSpec types.ServiceSpec `json:"serviceSpec"`
}

type UpdateServiceRequestDto struct {
	Namespace   string            `json:"namespace"`
	ServiceSpec types.ServiceSpec `json:"serviceSpec"`
}

type DeleteServiceRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
