package controller

import (
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/asishshaji/freshFarm/app/usecase"
	"github.com/asishshaji/freshFarm/app/utils"
	"github.com/labstack/echo"
)

type EchoController struct {
	usecase usecase.UsecaseInterface
	bucket  *storage.BucketHandle
}

func NewEchoController(uc usecase.UsecaseInterface, bucket *storage.BucketHandle) ControllerInterface {
	return EchoController{
		usecase: uc,
		bucket:  bucket,
	}
}

// For superAdmin
func (ec EchoController) CreateAdmin(c echo.Context) error {

	// Checks if the request is from super user
	superAdminPassword := c.FormValue("sup_password")
	err := ec.usecase.CheckIfSuperUser(c.Request().Context(), superAdminPassword)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	// TODO Upload image to server
	username := c.FormValue("username")
	password := c.FormValue("password")

	image, err := c.FormFile("profile_pic")

	link, err := utils.UploadImage(image, ec.bucket)
	log.Println(link)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error uploading image"})
	}

	err = ec.usecase.CreateAdmin(c.Request().Context(), username, password, link)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
