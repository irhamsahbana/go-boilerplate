package bootstrap

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	if err != nil {
		panic("Cant Find File config.yaml")
	}

	if config.GetBool("app.debug") {
		log.Println(color.BlueString("Service RUN on DEBUG mode, Happy code :)"))
	} else {
		log.Println(color.RedString("Service RUN on PRODUCTION mode, watch out!!!"))
	}

	return config
}
