package usecase

import (
	"context"

	"github.com/asishshaji/freshFarm/app/models"
)

type UsecaseInterface interface {
	// For superAdmin
	CreateAdmin(ctx context.Context, username, password, imageURL string) error
	CheckIfSuperUser(ctx context.Context, superAdminPassword string) error
	GetAdmins(ctx context.Context) ([]models.Admin, error)

	// // For Admin
	LoginAdmin(ctx context.Context, username, password string) error
	ChangeFarmerState(ctx context.Context, farmerID, state string) error
	AddProduct(ctx context.Context, product models.Product) error

	// // For Farmer
	SignupFarmer(ctx context.Context, password, firstname, lastname, link string, age int) (models.Farmer, error)
	LoginFarmer(ctx context.Context, username, password string) error

	// For user
	SignupUser(ctx context.Context, firstname, lastname, link, password string) error
	LoginUser(ctx context.Context, username, password string) error
	GetProductsByCategory(ctx context.Context, category string) ([]models.Product, error)
}
