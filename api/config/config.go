package config

import (
	"os"
	"strconv"
)

//AppConfig Application Configuration Structure
type AppConfig struct {
	DBhost         string
	DBname         string
	DBport         int
	ConnectionPool int
	AppSecret      string
	AppServerPort  int
}

//GetAppConfig Return pointer to APPConfig for DEV/PROD
func GetAppConfig() *AppConfig {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "PROD":
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT_PROD"))
		conPool, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_POOL_PROD"))
		serverPort, _ := strconv.Atoi(os.Getenv("APP_SERVER_PORT_PROD"))

		return &AppConfig{
			DBhost:         os.Getenv("DB_HOST_PROD"),
			DBname:         os.Getenv("DB_NAME_PROD"),
			DBport:         dbPort,
			ConnectionPool: conPool,
			AppSecret:      os.Getenv("APP_SECRET_PROD"),
			AppServerPort:  serverPort,
		}
	default:
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT_DEV"))
		conPool, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_POOL_DEV"))
		serverPort, _ := strconv.Atoi(os.Getenv("APP_SERVER_PORT_DEV"))

		return &AppConfig{
			DBhost:         os.Getenv("DB_HOST_DEV"),
			DBname:         os.Getenv("DB_NAME_DEV"),
			DBport:         dbPort,
			ConnectionPool: conPool,
			AppSecret:      os.Getenv("APP_SECRET_DEV"),
			AppServerPort:  serverPort,
		}

	}
}

func (cfg *AppConfig) GetDatabaseHostname() string {
	return cfg.DBhost
}
func (cfg *AppConfig) GetDatabaseName() string {
	return cfg.DBname
}
func (cfg *AppConfig) GetAppSecret() string {
	return cfg.AppSecret
}
func (cfg *AppConfig) GetDatabasePort() string {
	return strconv.Itoa(cfg.DBport)
}
func (cfg *AppConfig) GetConnectionPool() int {
	return cfg.ConnectionPool
}
func (cfg *AppConfig) GetAppServerPort() string {
	return strconv.Itoa(cfg.AppServerPort)
}
