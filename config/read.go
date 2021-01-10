package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type endpointResponse struct {
	Endpoints endpoint
}

type endpoint struct {
	Name string
	URL  string
}

//GetEndpoints reads the endpoints from the config.toml file and returns them
func (r *endpointResponse) GetEndpoints(n string) {
	setupViper()

	endpoints := viper.GetStringMapStringSlice(n + ".endpoints")

	for name, url := range endpoints {
		e := endpoint{
			Name: name,
			URL:  url[0],
		}

		fmt.Printf("Endpoint: %v", e)
		r.Endpoints = append(r.Endpoints, e)
	}
}

func setupViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}
