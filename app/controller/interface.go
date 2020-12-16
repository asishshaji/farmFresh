package controller

import "github.com/labstack/echo"

type ControllerInterface interface {
	CreateAdmin(c echo.Context) error // For SuperAdmin
	GetAdmins(c echo.Context) error

	LoginAdmin(c echo.Context) error
	AddProduct(c echo.Context) error
	ChangeStateFarmer(c echo.Context) error

	SignupFarmer(c echo.Context) error
	LoginFarmer(c echo.Context) error

	LoginUser(c echo.Context) error
	SignupUser(c echo.Context) error
	GetProductsByCategory(c echo.Context) error
}
