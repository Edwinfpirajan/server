package main

import (
	"fmt"
	"log"

	"github.com/Edwinfpirajan/server.git/cmd/providers"
	"github.com/Edwinfpirajan/server.git/config"
	"github.com/Edwinfpirajan/server.git/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverPort = config.Enviroments().Server.Port
)

func main() {

	container := providers.BuildContainer()

	if err := container.Invoke(func(server *echo.Echo, router *router.Router) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", serverPort)))
	}); err != nil {
		log.Panic(err)
	}
}
