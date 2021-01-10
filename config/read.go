package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title     string
	Endpoints endpoints
}

type endpoints struct {
	Code       string
	AccessCode string `toml:"access_token"`
}

//Read a file to be used
func Read() {
	var config tomlConfig

	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", config.Title)
	fmt.Printf("Code: %s\n", config.Endpoints.Code)
	fmt.Printf("Access code: %s\n", config.Endpoints.AccessCode)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
