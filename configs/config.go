package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Encryption EncryptionConfig `yaml:"encryption"`
}

type EncryptionConfig struct {
	Salt      uint8  `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DBConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnetcion     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	if database_host := os.Getenv("DATABASE_HOST"); database_host != "" {
		Cfg.DB.Host = database_host
	}

	if database_port := os.Getenv("DATABASE_PORT"); database_port != "" {
		Cfg.DB.Port = database_port
	}

	if database_name := os.Getenv("DATABASE_NAME"); database_name != "" {
		Cfg.DB.Name = database_name
	}

	if database_username := os.Getenv("DATABASE_USERNAME"); database_username != "" {
		Cfg.DB.User = database_username
	}

	if database_password := os.Getenv("DATABASE_PASSWORD"); database_password != "" {
		Cfg.DB.Password = database_password
	}

	return yaml.Unmarshal(configByte, &Cfg)
}