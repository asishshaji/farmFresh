package controller

import (
	"log"
	"net/http"
	"strconv"
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

func (ec EchoController) GetAdmins(c echo.Context) error {

	// Checks if the request is from super user
	superAdminPassword := c.FormValue("sup_password")
	err := ec.usecase.CheckIfSuperUser(c.Request().Context(), superAdminPassword)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	admins, err := ec.usecase.GetAdmins(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error getting admins",
		})
	}

	return c.JSON(http.StatusOK, admins)

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
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("adminSecret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}

func (ec EchoController) ApproveFarmer(c echo.Context) error {
	farmerID := c.FormValue("farmer_id")
	err := ec.usecase.ChangeFarmerState(c.Request().Context(), farmerID, "approve")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return nil
}
func (ec EchoController) SuspendFarmer(c echo.Context) error {
	farmerID := c.FormValue("farmer_id")
	err := ec.usecase.ChangeFarmerState(c.Request().Context(), farmerID, "suspend")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return nil
}

// Farmer
func (ec EchoController) SignupFarmer(c echo.Context) error {
	password := c.FormValue("password")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	age := c.FormValue("age")
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Check your age"})

	}
	profilePic, err := c.FormFile("profile_pic")

	link, err := utils.UploadImage(profilePic, ec.bucket)
	log.Println(link)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error uploading image"})
	}

	err = ec.usecase.SignupFarmer(c.Request().Context(), password, firstname, lastname, link, ageInt)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error registering farmer",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Saved farmer",
	})

}

func (ec EchoController) LoginFarmer(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := ec.usecase.LoginFarmer(c.Request().Context(), username, password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("farmerSecret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
