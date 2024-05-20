package config

import (
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var (
	configInstance Configurator
	once           sync.Once
)

type Configurator interface {
	Get(key string) string
	Set(key, value string)
	GetAll() map[string]string
}

type Config struct {
	vars map[string]string
}

func GetInstance() Configurator {
	once.Do(func() {
		configInstance = &Config{
			vars: SetVars(),
		}
		println("config init success")
	})
	return configInstance
}

func SetVars() map[string]string {
	vars, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return vars
}

func (c *Config) Get(key string) string {
	return c.vars[key]
}

func (c *Config) Set(key, value string) {
	c.vars[key] = value
}

func (c *Config) GetAll() map[string]string {
	return c.vars
}
