package namespace

type ApplyNamespaceDto struct {
	Namespace string `json:"namespace"`
}

type ListNamespaceRequestDto struct {
	Namespace string `json:"namespace" query:"namespace"`
}

type CreateNamespaceRequestDto struct {
	Namespace string `json:"namespace"`
}

type DeleteNamespaceRequestDto struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
