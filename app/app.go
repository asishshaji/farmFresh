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
		Limit: 200,
		Burst: 200,
	}))

	// Superadmin group
	sp := e.Group("/super")
	sp.POST("/create", controller.CreateAdmin)
	sp.GET("/admins", controller.GetAdmins)

	// Admin group
	ad := e.Group("/admin")
	ad.POST("/login", controller.LoginAdmin)
	ad.Use(middleware.JWT([]byte("adminSecret")))
	ad.POST("/product", controller.AddProduct)
	ad.POST("/farmer/state", controller.ChangeStateFarmer)

	// Farmer group
	fm := e.Group("/farmer")
	fm.POST("/signup", controller.SignupFarmer)
	fm.POST("/login", controller.LoginFarmer)
	fm.Use(middleware.JWT([]byte("farmerSecret")))

	// User group
	u := e.Group("/user")
	u.POST("/signup", controller.SignupUser)
	u.POST("/login", controller.LoginUser)
	u.Use(middleware.JWT([]byte("userSecret")))
	u.GET("/products/:category", controller.GetProductsByCategory)

	// get products by category, add pagination
	//cart
	// searcj product

	return &App{
		e:    e,
		port: port,
	}

}

// RunServer starts the server
func (a *App) RunServer() {
	a.e.Logger.Fatal(a.e.Start(a.port))

}
