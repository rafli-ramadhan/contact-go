package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapStructure:"port"`
	Storage string `mapStructure:"storage"`
}

func LoadConfig(name string, path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(name)

	// search defined path of file
	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err)
		_, ok := err.(viper.ConfigFileNotFoundError)
		log.Print(ok)
		if ok {
			return nil, errors.New(".env tidak ditemukan")
		}
		return nil, fmt.Errorf("fatal error config file %s", err)
	}

	// unmarshal parameter file .env to struct
	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("fatal error decode : %s", err)
	}

	return &config, nil
}
