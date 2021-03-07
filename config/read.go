package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type EndpointResponse struct {
	Endpoints endpoint
}

type endpoint struct {
	Name string
	URL  string
}

// GetEndpoints reads the endpoints from the config.toml file and returns them
func GetEndpoints(n string) (EndpointResponse, error) {
	setupViper()

	r := EndpointResponse{}
	endpoints := viper.GetStringMapStringSlice(n + ".endpoints")

	for name, url := range endpoints {
		e := endpoint{
			Name: name,
			URL:  url[0],
		}

		fmt.Printf("Endpoint: %v", e)
		r.Endpoints = e
	}

	return r, nil
}

func setupViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
