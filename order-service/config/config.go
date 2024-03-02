package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Address  string `yaml:"address"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Database string `yaml:"database"`
	} `yaml:"database"`
}

func ReadConfig() Config {
	f, err := os.Open("config/app.yaml")
	if err != nil {
		log.Fatalf("error while oping file. %s", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("error while parsing config. %s", err)
	}

	return cfg
}
