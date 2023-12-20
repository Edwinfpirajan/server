package config

import (
	"log"
	"sync"

	"github.com/andresxlp/gosuite/config"
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

func Enviroments() Config {
	Once.Do(func() {
		if err := config.SetEnvsFromFile("server", ".env"); err != nil {
			color.Green("Database connection established")
			log.Panicf(color.Red("Error loading enviroment variables: %v"), err)
		}
		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panicf("Error parsing enviroment variables: %v", err)
		}
	})

	return Cfg
}
