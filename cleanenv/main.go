package main

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Port string `yaml:"port" env:"port" env-default:"5432" env-description:"server port"`
	Host string `yaml:"host" env:"host" env-default:"host" env-description:"server host"`
	Name string `yaml:"name" env:"name" env-default:"name" env-description:"database name"`
	User string `yaml:"user" env:"user" env-default:"user" env-description:"database user"`
	Pass string `yaml:"pass" env:"pass" env-default:"pass" env-description:"database pass"`
}

var cfg ConfigDatabase

func main() {
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(cfg)

	err = cleanenv.ReadConfig("dev.env", &cfg)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(cfg)

	help, err := cleanenv.GetDescription(&cfg, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(help)
}
