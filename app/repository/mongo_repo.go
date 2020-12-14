package repository

import (
	"context"
	"errors"
	"log"

	"github.com/asishshaji/freshFarm/app/models"
	"github.com/asishshaji/freshFarm/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	adminCollection   *mongo.Collection
	farmerCollection  *mongo.Collection
	farmCollection    *mongo.Collection
	userCollection    *mongo.Collection
	productCollection *mongo.Collection
	superAdmin        *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) RepositoryInterface {
	return mongoRepo{
		adminCollection:   db.Collection("admin"),
		farmerCollection:  db.Collection("farmer"),
		farmCollection:    db.Collection("farm"),
		userCollection:    db.Collection("user"),
		productCollection: db.Collection("product"),
		superAdmin:        db.Collection("super_admin"),
	}
}

func (repo mongoRepo) CheckIfSuperUser(
	ctx context.Context,
	superAdminPassword string) error {

	var superAdmin bson.M

	if err := repo.superAdmin.FindOne(ctx, bson.M{"password": superAdminPassword}).
		Decode(&superAdmin); err != nil {
		return err
	}

	return nil

}

// For superAdmin
func (repo mongoRepo) CreateAdmin(ctx context.Context, admin models.Admin) error {

	opts := options.Update().SetUpsert(true)

	up, err := utils.ToDoc(admin)
	if err != nil {
		return err
	}
	doc := bson.D{{"$set", up}}

	result, err := repo.adminCollection.UpdateOne(ctx, bson.M{"username": admin.Username}, doc, opts)

	if result.MatchedCount != 0 {
		return errors.New("Admin already exists")
	}

	if err != nil {
		return err
	}

	return nil
}

// Admin methods start
func (repo mongoRepo) GetAdmin(ctx context.Context, AdminUsername string) (models.Admin, error) {
	admin := models.Admin{}

	filter := bson.M{"username": AdminUsername}
	result := repo.adminCollection.FindOne(ctx, filter)
	err := result.Decode(&admin)
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}
func (repo mongoRepo) ApproveFarmer(ctx context.Context, farmerID primitive.ObjectID) error {

	farmer := models.Farmer{}

	filter := bson.M{"_id": farmerID}
	updation := bson.M{"$set": bson.M{"state": "approved"}}

	result := repo.farmerCollection.FindOneAndUpdate(ctx, filter, updation)
	err := result.Decode(&farmer)
	if err != nil {
		return err
	}

	return nil
}

func (repo mongoRepo) SuspendFarmer(ctx context.Context, farmerID primitive.ObjectID) error {
	farmer := models.Farmer{}

	filter := bson.M{"_id": farmerID}
	updation := bson.M{"$set": bson.M{"state": "suspended"}}

	result := repo.farmerCollection.FindOneAndUpdate(ctx, filter, updation)
	err := result.Decode(&farmer)
	if err != nil {
		return err
	}

	return nil
}
func (repo mongoRepo) DeleteFarmer(ctx context.Context, farmerID primitive.ObjectID) error {
	filter := bson.M{"_id": farmerID}
	result, err := repo.farmerCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	log.Println("Deleted a farmer : ", result)

	return nil
}

// Farmer methods
func (repo mongoRepo) CreateFarmer(ctx context.Context, farmer models.Farmer) error {
	result, err := repo.farmCollection.InsertOne(ctx, farmer)
	if err != nil {
		return err
	}

	log.Println("New farmer created :", result)

	return nil
}
