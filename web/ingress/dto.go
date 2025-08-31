package ingress

import (
	types "k8s.io/api/networking/v1"
)

type ListIngressRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type DetailIngressRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateIngressRequestDto struct {
	Namespace   string            `json:"namespace"`
	IngressSpec types.IngressSpec `json:"ingressSpec"`
}

type UpdateIngressRequestDto struct {
	Namespace   string            `json:"namespace"`
	IngressSpec types.IngressSpec `json:"ingressSpec"`
}

type DeleteIngressRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
