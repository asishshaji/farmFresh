package usecase

import (
	"context"
)

type UsecaseInterface interface {
	// For superAdmin
	CreateAdmin(ctx context.Context, username, password, imageURL string) error
	CheckIfSuperUser(ctx context.Context, superAdminPassword string) error

	// // For Admin
	// LoginAdmin(ctx context.Context, username, password string)
	// ApproveFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	// SuspendFarmer(ctx context.Context, farmerID primitive.ObjectID) error
	// DeleteFarmer(ctx context.Context, farmerID primitive.ObjectID) error

	// // For Farmer
	// SignupFarmer(ctx, username, firstName, lastName, password, about string) error
	// LoginFarmer(ctx context.Context, username, password string) (models.Farmer, error)

}
