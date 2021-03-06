package repository

import (
	"context"

	"github.com/asishshaji/freshFarm/app/models"
)

type RepositoryInterface interface {
	CheckIfSuperUser(ctx context.Context, superAdminPassword string) error

	// SuperAdmin features ????????
	CreateAdmin(ctx context.Context, admin models.Admin) error

	// // Admin start
	LoginAdmin(ctx context.Context, username, password string) error
	GetAdmin(ctx context.Context, AdminUsername string) (models.Admin, error)
	GetAdmins(ctx context.Context) ([]models.Admin, error)

	ChangeFarmerState(ctx context.Context, farmerID, state string) error
	CreateProduct(ctx context.Context, product models.Product) error
	// // GetOrders(ctx context.Context) error
	// SearchFarmer(ctx context.Context, farmerUsername string) (models.Farmer, error)
	// SearchFarm(ctx context.Context, farmUsername string) (models.Farm, error)
	// // GetProfits(ctx context.Context) error
	// AddProduct(ctx context.Context, product models.Product) error
	// UpdateProduct(ctx context.Context, product models.Product) error
	// DeleteProduct(ctx context.Context, product models.Product) error
	// GetTopFarmers(ctx context.Context) ([]models.Farmer, error)
	// GetTopFarms(ctx context.Context) ([]models.Farm, error)
	// ApproveFarm(ctx context.Context, farmID primitive.ObjectID) error
	// DeleteFarm(ctx context.Context, farmID primitive.ObjectID) error
	// Admin end

	GetAllCategories(ctx context.Context) ([]models.Category, error)
	CreateCategory(ctx context.Context, categoryName string) error
	// Farmer start
	CreateFarmer(ctx context.Context, farmer models.Farmer) error
	GetFarmerWithUsername(ctx context.Context, username string) (models.Farmer, error)
	ChangeItemInCart(ctx context.Context, action, username, productID string) error

	// Farmer end

	// User start
	CreateUser(ctx context.Context, user models.User) error
	GetUserWithUsername(ctx context.Context, username string) (models.User, error)
	GetProductsByCategory(ctx context.Context, category string) ([]models.Product, error)
	CreateOrder(ctx context.Context, order models.Order) (string, error)

	// end

}
