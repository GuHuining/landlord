package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Mysql struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type DatabaseConfig struct {
	Mysql `yaml:"mysql"`
	Redis `yaml:"redis"`
}

func GetDatabaseConfig() DatabaseConfig {
	content, err := os.ReadFile("./config/database.yaml")
	if err != nil {
		panic(err)
	}
	var config DatabaseConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		panic(err)
	}
	return config
}
