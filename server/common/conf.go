package common

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
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
		fmt.Println(err, "viper read config error")
	}
}

// GetConfStr Get Yaml file string value
func GetConfStr(key string) string {
	confVal := viper.GetString(key)
	return confVal
}

// GetConfArr Get yaml file slice value
func GetConfArr(key string) []string {
	confVal := viper.GetStringSlice(key)
	return confVal
}

// GetConfInt Get yaml file int value
func GetConfInt(key string) int {
	confVal := viper.GetInt(key)
	return confVal
}
