package event

type ListEventDto struct {
	Namespace    string `json:"namespace" query:"namespace"`
	ResourceName string `json:"resourceName" query:"resourceName"`
}
