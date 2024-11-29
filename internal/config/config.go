package config

import (
	"errors"
	"flag"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTPPort   int    `mapstructure:"HTTP"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbSSLMode  string `mapstructure:"DB_SSLMODE"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
}

// LoadConfig - initializing app config from yaml file
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigFile(mustGetConfigPathFromCMDLine())
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, err
}

func (cfg *AppConfig) Validate() error {
	return validateStructFields(cfg)
}

func validateStructFields(configField interface{}) error {
	valueOfConfigs := reflect.ValueOf(configField)
	typeOfConfigs := valueOfConfigs.Type()

	for i := 0; i < typeOfConfigs.NumField(); i++ {
		field := typeOfConfigs.Field(i)
		value := valueOfConfigs.Field(i)
		if len(value.String()) == 0 {
			return fmt.Errorf("[%s]: {%s} field value is invalid", valueOfConfigs.Type(), field.Name)
		}
	}

	return nil
}

func mustGetConfigPathFromCMDLine() (configPath string) {
	cfgPath := flag.String("config", "sample.env", "Default config")
	flag.Parse()
	configPath = *cfgPath

	if len(configPath) == 0 {
		panic(errors.New("config path is empty"))
	}
	return
}
