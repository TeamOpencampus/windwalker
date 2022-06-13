package main

import (
	"log"
	"os"
)

type Config struct {
	Port string
	Dsn  string
	Mode string
}

func LoadConfig() *Config {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	mode, ok := os.LookupEnv("MODE")
	if !ok {
		mode = "PROD"
	}
	dsn, ok := os.LookupEnv("DSN")
	if !ok && mode == "PROD" {
		log.Fatalln("Environment variable DSN is undefined. Refusing to start.")
	}
	return &Config{Port: port, Dsn: dsn, Mode: mode}
}
