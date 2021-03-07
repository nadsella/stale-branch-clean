package partner

import (
	"fmt"
	"github.com/spf13/viper"
)

type Partner struct {
	Name string
	BaseURL string
}

func Init() {
	for _, config := range viper.GetStringSlice("apis") {

		p := Partner{
			Name: viper.GetString(config + "name"),
			BaseURL: viper.GetString(config + "baseUrl"),
		}

		fmt.Printf("%v", p)
	}
}