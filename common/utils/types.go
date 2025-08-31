package utils

type BasicResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Config struct {
	Port       int    `json:"PORT"`
	KubeConfig string `json:"KUBE-CONFIG"`
}

type Meta struct {
	Total int64 `json:"total"`
	Links Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

func InterfaceToGeneric[T any](i interface{}) T {
	v, ok := i.(T)
	if !ok {
		var zero T
		return zero
	}
	return v
}
