package provider

import (
	"fmt"
	"net/http"

	config "exam/config"
)

var cfg config.Config

// Provider is the interface to provide needed instance by main
type Provider interface {
	GetServer() *http.Server
}

type providerImpl struct{}

// GetProvider will return the instance of a provider
func GetProvider() Provider {
	return providerImpl{}
}

func getConfig() config.Config {
	fmt.Println(cfg.ServerCfg.ReadTimeout)
	if !cfg.IsValid {
		cfg = config.GetConfig()
	}
	return cfg
}
