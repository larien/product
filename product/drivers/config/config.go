package config

import "os"

// Config contains the configuration fields used by the system
type Config struct {
	DB       *DB
	GRPC     *GRPC
}

// DB contains the configuration fields necessary for database connection
type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// GRPC contains the configuration fields necessary for gRPC connection
type GRPC struct {
	Port string
	Host string
}

// New obtains configurations from environment variables. For now it doesn't have any default values.
// Some package like viper might be applied to handle these cases better.
func New() *Config {
	return &Config{
		DB: &DB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		GRPC: &GRPC{
			Host: os.Getenv("GRPC_HOST"),
			Port: os.Getenv("GRPC_PORT"),
		},
	}
}
