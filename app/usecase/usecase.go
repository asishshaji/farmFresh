package usecase

import (
	"context"
	"time"

	"github.com/asishshaji/freshFarm/app/models"
	"github.com/asishshaji/freshFarm/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
