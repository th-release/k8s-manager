package utils

import (
	"os"
)

func GetConfig() *Config {
	return &Config{
		Port:       ThreeTermString(len(os.Getenv("PORT")) >= 1, os.Getenv("PORT"), "8080"),
		KubeConfig: ThreeTermString(len(os.Getenv("KUBE_CONFIG")) >= 1, os.Getenv("KUBE_CONFIG"), "~/.kube/config"),
	}
}
