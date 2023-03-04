package config

import "caturandi-labs/golang-starter/utils"

type DatabaseConfig struct {
	Host 		string
	Name 		string
	Password 	string
	Port		string
	User		string
}

func NewDatabase() *DatabaseConfig {
	return &DatabaseConfig{
		Host: utils.GetIni("database", "DATABASE_HOST", ""),
		Name: utils.GetIni("database", "DATABASE_NAME", ""),
		Password: utils.GetIni("database", "DATABASE_PASSWORD", ""),
		Port: utils.GetIni("database", "DATABASE_PORT", ""),
		User: utils.GetIni("database", "DATABASE_USER", ""),
	}
}