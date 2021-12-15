package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port  string `env:"PORT" envDefault:"8080"`
	DbUrl string `env:"DB_URL" envDefault:"test.db"`
}

func ParseConfig() *Config {
	warn := godotenv.Load()

	if warn != nil {
		log.Println("No .env file found")
	}

	var config Config

	if err := env.Parse(&config); err != nil {
		log.Fatalln("Could not parse configuration")
	}

	return &config
}
