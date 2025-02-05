package config

import "time"

var configInstance = &Config{
	Server: Server{
		Port:    8000,
		Timeout: 5 * time.Second,
	},
	Datasource: Datasource{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "maas",
		Schema:   "public",
	},
}

func GetInstance() *Config {
	return configInstance
}

type Config struct {
	Server     Server
	Datasource Datasource
}

type Server struct {
	Port    int
	Timeout time.Duration
}

type Datasource struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Schema   string
}
