package controller

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/asishshaji/freshFarm/app/usecase"
	"github.com/asishshaji/freshFarm/app/utils"
	"github.com/dgrijalva/jwt-go"
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

// Admin
func (ec EchoController) LoginAdmin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := ec.usecase.LoginAdmin(c.Request().Context(), username, password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Check username and password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
