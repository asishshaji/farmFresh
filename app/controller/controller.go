package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/asishshaji/freshFarm/app/models"
	"github.com/asishshaji/freshFarm/app/usecase"
	"github.com/asishshaji/freshFarm/app/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ec EchoController) AddProduct(c echo.Context) error {
	productName := c.FormValue("name")
	ownerID := c.FormValue("owner_id")
	price := c.FormValue("price")

	priceFloat, _ := strconv.ParseFloat(price, 32)
	marketPrice := c.FormValue("market_price")
	marketFloat, _ := strconv.ParseFloat(marketPrice, 32)

	productCount := c.FormValue("count")
	count, _ := strconv.Atoi(productCount)

	typeOfMeasurement := c.FormValue("type")
	about := c.FormValue("about")
	category := c.FormValue("category")

	product := models.Product{
		Name:              productName,
		ImageURLS:         []string{},
		CreatedAt:         primitive.NewDateTimeFromTime(time.Now()),
		OwnerID:           ownerID,
		Price:             priceFloat,
		MarketPrice:       marketFloat,
		Reviews:           []models.Review{},
		NutritionalValues: []models.Nutrition{},
		About:             about,
		State:             "active",
		TypeOfMeasurement: typeOfMeasurement,
		ProductCount:      count,
		Category:          category,
	}

	err := ec.usecase.AddProduct(c.Request().Context(), product)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return nil
}

func (ec EchoController) ChangeStateFarmer(c echo.Context) error {
	farmerID := c.FormValue("farmer_id")
	state := c.FormValue("state")
	err := ec.usecase.ChangeFarmerState(c.Request().Context(), farmerID, state)
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

// user
func (ec EchoController) SignupUser(c echo.Context) error {
	password := c.FormValue("password")
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")

	profilePic, err := c.FormFile("profile_pic")

	link, err := utils.UploadImage(profilePic, ec.bucket)
	log.Println(link)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error uploading image"})
	}

	err = ec.usecase.SignupUser(c.Request().Context(), firstname, lastname, link, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error registering farmer",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Saved user",
	})
}

func (ec EchoController) LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	err := ec.usecase.LoginUser(c.Request().Context(), username, password)

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

	t, err := token.SignedString([]byte("userSecret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (ec EchoController) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	log.Println(name, " requested for products in ", category)

	products, err := ec.usecase.GetProductsByCategory(c.Request().Context(), category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}
