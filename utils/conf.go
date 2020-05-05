package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	fmt.Println(">>>>>> :init config files")
	env := os.Getenv("APP_ENV")

	confName := "conf.dev"
	if env == "production" {
		confName = "conf.prod"
	} else if env == "test" {
		confName = "conf.test"
	}

	viper.AddConfigPath("./conf")
	viper.SetConfigName(confName)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(">>>>>> :viper read config error", err)
	}
}

// GetConfStr Get Yaml file value
func GetConfStr(key string) string {
	confVal := viper.GetString(key)
	return confVal
}
