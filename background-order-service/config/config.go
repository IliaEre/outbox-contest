package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Servers  string `yaml:"servers"`
	Group    string `yaml:"group"`
	Timeout  int    `yaml:"timeout"`
	Offset   string `yaml:"offset"`
	Security string `yaml:"security"`
	Poll     int    `yaml:"poll"`
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
