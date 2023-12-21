package db

import (
	"fmt"
	"sync"

	"github.com/Edwinfpirajan/server.git/config"
	"github.com/labstack/gommon/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// This is a singleton pattern, it will create a single instance of the connection
var (
	connection *gorm.DB
	once       sync.Once
)

// This is a struct that will hold the database connection options
type pgOptions struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

// This is a method that will build the DSN string
func (p *pgOptions) buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", p.Host, p.User, p.Password, p.Dbname, p.Port)
}

// This is a method that will create a single instance of the connection
func NewPostgresConnection() *gorm.DB {
	// This is a singleton pattern, it will create a single instance of the connection
	once.Do(func() {
		connection = MainDbConnection()
	})
	return connection
}

// This is a method that will create a single instance of the connection
func MainDbConnection() *gorm.DB {
	// config.Enviroments()

	dsn := enviromentsDSN()

	// This is a method that will create a single instance of the connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(color.Green("Database connection established"))

	return db
}

// This is a method that will create a single instance of the connection
func enviromentsDSN() string {
	// This is a struct that will hold the database connection options
	options := pgOptions{
		Host:     config.Environments().MainDb.Host,
		User:     config.Environments().MainDb.User,
		Password: config.Environments().MainDb.Password,
		Dbname:   config.Environments().MainDb.DbName,
		Port:     fmt.Sprintf("%d", config.Environments().MainDb.Port),
	}

	return options.buildDSN()
}
