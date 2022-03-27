package main

type Config struct {
	SQLConnectionString string `yaml:"sqlConnectionString"`
	Port                int    `yaml:"port"`
}
