package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const configFilePath = "./config.yml"

var cfg *Config

func GetConfigFromFile() *Config {
	cfg := GetConfig()
	if !cfg.loaded {
		if err := ReadConfig(configFilePath); err != nil {
			panic("Could not load configuration file")
		}
	}

	return cfg
}

type GrpcConfig struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	MaxConnectionIdle int    `yaml:"maxConnectionIdle"`
	Timeout           int    `yaml:"timeout"`
	MaxConnectionAge  int    `yaml:"maxConnectionAge"`
}

func (t GrpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

type GrpcGatewayConfig struct {
	Port string `yaml:"port"`
}

func (t GrpcGatewayConfig) Address() string {
	return fmt.Sprintf(":%s", t.Port)
}

type DatabaseConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Pass           string `yaml:"password"`
	Name           string `yaml:"name"`
	Migrations     string `yaml:"migrations"`
	SSLmode        string `yaml:"sslmode"`
	Driver         string `yaml:"driver"`
	ConnectRetries int    `yaml:"connectRetries"`
}

func (dbconf *DatabaseConfig) GetConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbconf.Host, dbconf.Port, dbconf.User, dbconf.Pass, dbconf.Name, dbconf.SSLmode)
}

type Config struct {
	loaded      bool
	Grpc        GrpcConfig        `yaml:"grpc"`
	GrpcGateway GrpcGatewayConfig `yaml:"grpc-gateway"`
	Database    DatabaseConfig    `yaml:"database"`
}

func ReadConfig(filePath string) error {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	if cfg == nil {
		cfg = &Config{
			loaded: false,
		}
	}

	return cfg
}
