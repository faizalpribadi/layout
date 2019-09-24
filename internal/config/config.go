package config

import (
	"errors"

	"github.com/spf13/viper"
)

type (
	Common struct {
		Name  string `yaml:"name"`
		Debug bool   `yaml:"debug"`
		Env   string `yaml:"env"`
	}
	DB struct {
		Providers *Providers
	}
	Log struct {
		File string `yaml:"file"`
	}
	Http struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	}

	Providers struct {
		MySQL struct {
			Hosts string `yaml:"hosts"`
			User  string `yaml:"user`
			Port  int    `yaml:"port"`
			Name  string `yaml:"name"`
			Pass  string `yaml:"pass"`
		}
		Redis struct {
			Hosts string `yaml:"hosts"`
			User  string `yaml:"user`
			Port  int    `yaml:"port"`
			Name  string `yaml:"name"`
			Pass  string `yaml:"pass"`
		}
	}
	CircuitBreaker struct {
		Timeout int `yaml:"timeout"`
	}

	Configuration struct {
		Common         *Common
		DB             *DB
		Log            *Log
		Http           *Http
		CircuitBreaker *CircuitBreaker
	}

	Config interface {
		ReadFileConfiguration(path string) (*Configuration, error)
	}
	config struct{}
)

func NewConfiguration() Config {
	return &config{}
}

func (c *config) ReadFileConfiguration(path string) (*Configuration, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("[config-error] config file not found")
	}

	config := &Configuration{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, errors.New("[error] un-marshalling config error")
	}

	return config, nil
}
