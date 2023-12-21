package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
)

var (
	Cfg  Config
	Once sync.Once
)

type Config struct {
	Server Server `mapstructure:"server" validate:"required"`
	MainDb MainDB `mapstructure:"main_db" validate:"required"`
}

type Server struct {
	Port string `mapstructure:"port" validate:"required"`
}

type MainDB struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DbName   string `mapstructure:"name" validate:"required"`
}

func Environments() Config {
	Once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			color.Green("Database connection established")
			log.Panicf(color.Red("Error loading environment variables: %v"), err)
		}

		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing environment variables: %v", err)
		}
	})

	return Cfg
}
