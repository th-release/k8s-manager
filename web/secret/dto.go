package secret

import (
	types "k8s.io/api/core/v1"
)

type ListSecretRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailSecretRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateSecretRequestDto struct {
	Namespace string       `json:"namespace"`
	Secret    types.Secret `json:"secret"`
}

type UpdateSecretRequestDto struct {
	Namespace string       `json:"namespace"`
	Secret    types.Secret `json:"secret"`
}

type DeleteSecretRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
