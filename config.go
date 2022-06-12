package main

import (
	"log"
	"os"
)

type Config struct {
	Port string
	Dsn  string
}

func LoadConfig() *Config {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		log.Fatalln("Environment variable DSN is undefined. Refusing to start.")
	}
	return &Config{Port: port, Dsn: dsn}
}
