package controller

import "github.com/labstack/echo"

type ControllerInterface interface {
	CreateAdmin(c echo.Context) error // For SuperAdmin
	GetAdmins(c echo.Context) error

	LoginAdmin(c echo.Context) error

	SignupFarmer(c echo.Context) error
	LoginFarmer(c echo.Context) error

	ApproveFarmer(c echo.Context) error
	SuspendFarmer(c echo.Context) error
}
