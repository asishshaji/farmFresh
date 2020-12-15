package controller

import "github.com/labstack/echo"

type ControllerInterface interface {
	// Login(c echo.Context) error
	// Signup(c echo.Context) error
	CreateAdmin(c echo.Context) error // For SuperAdmin
	LoginAdmin(c echo.Context) error
}
