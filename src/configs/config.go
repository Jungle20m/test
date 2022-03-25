package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	config := viper.New()

	config.AddConfigPath(".")
	config.SetConfigName("env")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("failed to read configuration:", err.Error())
			os.Exit(1)
		}
	}
	Config = config
}
