package controller

import (
	"net/http"

	"github.com/asishshaji/freshFarm/app/usecase"
	"github.com/labstack/echo"
)

type Message struct {
	message string
}

type EchoController struct {
	usecase usecase.UsecaseInterface
}

func NewEchoController(uc usecase.UsecaseInterface) ControllerInterface {
	return EchoController{
		usecase: uc,
	}
}

// For superAdmin
func (ec EchoController) CreateAdmin(c echo.Context) error {

	superAdminPassword := c.FormValue("sup_password")

	err := ec.usecase.CheckIfSuperUser(c.Request().Context(), superAdminPassword)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	// Check superadmin password

	username := c.FormValue("username")
	password := c.FormValue("password")

	// imageURL :=  save image

	err = ec.usecase.CreateAdmin(c.Request().Context(), username, password, "ASD")

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
