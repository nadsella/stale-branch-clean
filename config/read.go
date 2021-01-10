package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Code       string
	AccessCode string
}

//Read a file to be used
func Read() {
	//read file into variable
	dat, err := ioutil.ReadFile("./config.toml")

	check(err)

	var d tomlConfig

	if _, err := toml.Decode(string(dat), &d); err != nil {
		check(err)
	}

	fmt.Print(string(d.AccessCode))
	fmt.Print(string(d.Code))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
