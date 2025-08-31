package configMap

import (
	types "k8s.io/api/core/v1"
)

type ListConfigMapRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailConfigMapRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateConfigMapRequestDto struct {
	Namespace string          `json:"namespace"`
	ConfigMap types.ConfigMap `json:"configMap"`
}

type UpdateConfigMapRequestDto struct {
	Namespace string          `json:"namespace"`
	ConfigMap types.ConfigMap `json:"configMap"`
}

type DeleteConfigMapRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
