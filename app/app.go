package app

import (
	"github.com/asishshaji/freshFarm/app/controller"
	custommiddlewares "github.com/asishshaji/freshFarm/app/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// App creates the starting point of the server
type App struct {
	e    *echo.Echo
	port string
}

// NewApp creates new app
func NewApp(port string, controller controller.ControllerInterface) *App {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(custommiddlewares.RateLimitWithConfig(custommiddlewares.RateLimitConfig{
		Limit: 2,
		Burst: 2,
	}))

	// Superadmin group
	sp := e.Group("/super")
	sp.POST("/create", controller.CreateAdmin)

	// Admin group
	// ad := e.Group("/admin")
	// ad.GET("/",)

	// Farmer group
	fm := e.Group("/farmer")
	fm.Use(middleware.JWT([]byte("secret")))

	return &App{
		e:    e,
		port: port,
	}

}

// RunServer starts the server
func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))

}
