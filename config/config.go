package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	DataBase struct {
		Dsn          string
		MaxIdleConns int
	}
	Redis struct {
		Address  string
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file : %v ", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct %v", err)
	}

	initDB()
	InitRedis()
}
