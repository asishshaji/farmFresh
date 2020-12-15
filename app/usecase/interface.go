package usecase

import (
	"context"

	"github.com/asishshaji/freshFarm/app/models"
)

type UsecaseInterface interface {
	// For superAdmin
	CreateAdmin(ctx context.Context, username, password, imageURL string) error
	CheckIfSuperUser(ctx context.Context, superAdminPassword string) error
	LoginAdmin(ctx context.Context, username, password string) error
	ChangeFarmerState(ctx context.Context, farmerID, state string) error
	GetAdmins(ctx context.Context) ([]models.Admin, error)
	SignupFarmer(ctx context.Context, password, firstname, lastname, link string, age int) error
	LoginFarmer(ctx context.Context, username, password string) error

	// // For Admin
	// LoginAdmin(ctx context.Context, username, password string)
	// ApproveFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	// SuspendFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	// DeleteFarmer(ctx context.Context, farmerID primitive.ObjectID) error

	// // For Farmer
	// SignupFarmer(ctx, username, firstName, lastName, password, about string) error
	// LoginFarmer(ctx context.Context, username, password string) (models.Farmer, error)

}
