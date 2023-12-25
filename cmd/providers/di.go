package providers

import (
	"time"

	"github.com/Edwinfpirajan/server.git/internal/app"
	"github.com/Edwinfpirajan/server.git/internal/infra/adapters/db/implementation"
	"github.com/Edwinfpirajan/server.git/internal/infra/api/handlers"
	"github.com/Edwinfpirajan/server.git/internal/infra/api/router"
	"github.com/Edwinfpirajan/server.git/internal/infra/api/router/groups"
	"github.com/Edwinfpirajan/server.git/internal/src/db"
	"github.com/Edwinfpirajan/server.git/utils/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {

	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(func() http.HttpClient {
		return http.NewHTTPClient(3, 5*time.Second, 30*time.Second)
	})

	_ = Container.Provide(db.NewPostgresConnection)

	// Repository
	_ = Container.Provide(implementation.NewRepository)

	// App
	_ = Container.Provide(app.NewUserApp)
	_ = Container.Provide(app.NewRoleApp)

	// Handlers
	_ = Container.Provide(handlers.NewUserHandler)
	_ = Container.Provide(handlers.NewRoleHandler)

	// Groups
	_ = Container.Provide(groups.NewUserGroup)
	_ = Container.Provide(groups.NewRoleGroup)

	// Router
	_ = Container.Provide(router.NewRouter)

	return Container
}
