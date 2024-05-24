package config

import (
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"sync"
	"time"
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
		Logger.GetInstance().GetLogger().Info("the setting is configured",
			zap.Time("timestamp", time.Now()),
		)
	})
	return configInstance
}

func SetVars() map[string]string {
	vars, err := godotenv.Read()
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Error loading .env file: ",
			zap.Error(err),
			zap.Time("timestamp", time.Now()),
		)
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
