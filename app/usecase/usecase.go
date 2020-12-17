package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/asishshaji/freshFarm/app/models"
	"github.com/asishshaji/freshFarm/app/repository"
	"github.com/asishshaji/freshFarm/app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return Usecase{
		repo,
	}
}

func (usecase Usecase) CheckIfSuperUser(ctx context.Context, superAdminPassword string) error {
	err := usecase.repo.CheckIfSuperUser(ctx, superAdminPassword)

	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) CreateAdmin(ctx context.Context, username, password, imageURL string) error {
	admin := models.Admin{
		Username:        username,
		Password:        password,
		ProfileImageURL: imageURL,
		JoinedOn:        primitive.NewDateTimeFromTime(time.Now()),
	}

	err := usecase.repo.CreateAdmin(ctx, admin)

	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) LoginAdmin(ctx context.Context, username, password string) error {
	err := usecase.repo.LoginAdmin(ctx, username, password)
	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) ChangeFarmerState(ctx context.Context, farmerID, state string) error {
	err := usecase.repo.ChangeFarmerState(ctx, farmerID, state)
	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) AddProduct(ctx context.Context, product models.Product) error {
	err := usecase.repo.CreateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) GetAdmins(ctx context.Context) ([]models.Admin, error) {
	admins, err := usecase.repo.GetAdmins(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (usecase Usecase) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := usecase.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (usecase Usecase) CreateCategory(ctx context.Context, categoryName string) error {
	err := usecase.repo.CreateCategory(ctx, categoryName)
	if err != nil {
		return err
	}
	return nil
}

func (usecase Usecase) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	orderID, err := usecase.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}
	return orderID, nil
}

func (usecase Usecase) SignupFarmer(ctx context.Context, password, firstname, lastname, link string, age int) (models.Farmer, error) {
	// hashPassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Farmer{}, err
	}

	username := utils.MakeTimestamp()

	farmer := models.Farmer{
		Username:        username,
		FirstName:       firstname,
		LastName:        lastname,
		Password:        string(hashedPassword),
		Age:             age,
		ProfileImageURL: link,
		About:           "",
		Farms:           []models.Farm{},
		JoinedOn:        primitive.NewDateTimeFromTime(time.Now()),
		Rating:          0.0,
		Score:           0.0,
		State:           "review",
		Reviews:         []models.Review{},
		Profit:          0.0,
	}

	err = usecase.repo.CreateFarmer(ctx, farmer)

	if err != nil {
		return models.Farmer{}, err
	}
	return farmer, nil

}

func (usecase Usecase) LoginFarmer(ctx context.Context, username, password string) error {
	farmer, err := usecase.repo.GetFarmerWithUsername(ctx, username)
	if err != nil {
		return errors.New("No such user exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(farmer.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (usecase Usecase) SignupUser(ctx context.Context, firstname, lastname, link, password string) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	username := utils.MakeTimestamp()

	user := models.User{
		Username:        username,
		FirstName:       firstname,
		LastName:        lastname,
		Password:        string(hashedPassword),
		JoinedOn:        primitive.NewDateTimeFromTime(time.Now()),
		ProfileImageURL: link,
		State:           "active",
		FavoriteFarmers: []models.Farmer{},
		FavoriteFarms:   []models.Farm{},
	}

	err = usecase.repo.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, err

	}
	return user, nil

}

func (usecase Usecase) LoginUser(ctx context.Context, username, password string) error {
	user, err := usecase.repo.GetUserWithUsername(ctx, username)
	if err != nil {
		return errors.New("No such user exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (usecase Usecase) GetProductsByCategory(ctx context.Context, category string) ([]models.Product, error) {
	products, err := usecase.repo.GetProductsByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (usecase Usecase) ChangeItemInCart(ctx context.Context, action, productID, username string) error {

	err := usecase.repo.ChangeItemInCart(ctx, action, username, productID)
	if err != nil {
		return err
	}

	return nil

}
