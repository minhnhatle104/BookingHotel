package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	MySQL MySQLConfig
	Port  string
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig() (*AppConfig, *error) {
	v := viper.New()
	v.AddConfigPath("./config") // path to look for the config file in
	v.SetConfigName("config")   // name of config file (without extension)
	v.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalln("error in loading config file: ", err)
			return nil, &err
		} else {
			// Config file was found but another error was produced
			log.Fatalln("error in loading config file but unknown error: ", err)
			return nil, &err
		}
	}

	var appConfig AppConfig
	if err := v.Unmarshal(&appConfig); err != nil {
		log.Fatalln("error binding config: ", err)
		return nil, &err
	}

	return &appConfig, nil
}
