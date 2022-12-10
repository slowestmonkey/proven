package main

import (
	"proven/internal/app"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app.Run()
}
