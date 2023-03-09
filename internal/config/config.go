package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Postgres struct {
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type Server struct {
	PortREST string `yaml:"portREST"`
	PortGRPC string `yaml:"portGRPC"`
}

type RabbitMQ struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Postgres Postgres `yaml:"postgres"`
	Server   Server   `yaml:"server"`
	SaveType string   `yaml:"saveType"`
	RabbitMQ RabbitMQ `yaml:"rabbitMQ"`
}

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func Init() (*Config, error) {
	filePath := flag.String("c", "etc/config.yml", "Path to configuration file")
	flag.Parse()
	config = &Config{}
	data, err := os.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}
