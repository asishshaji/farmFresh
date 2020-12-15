package repository

import (
	"context"

	"github.com/asishshaji/freshFarm/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryInterface interface {
	CheckIfSuperUser(ctx context.Context, superAdminPassword string) error

	// SuperAdmin features ????????
	CreateAdmin(ctx context.Context, admin models.Admin) error

	// // Admin start
	LoginAdmin(ctx context.Context, username, password string) error
	GetAdmin(ctx context.Context, AdminUsername string) (models.Admin, error)
	ApproveFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	SuspendFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	DeleteFarmer(ctx context.Context, farmerID primitive.ObjectID) error
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

	// Farmer start
	CreateFarmer(ctx context.Context, farmer models.Farmer) error
	// Farmer end

	// User start
	// end

}
