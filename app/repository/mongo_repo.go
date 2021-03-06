package repository

import (
	"context"
	"errors"
	"log"

	"github.com/asishshaji/freshFarm/app/models"
	"github.com/asishshaji/freshFarm/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	adminCollection      *mongo.Collection
	farmerCollection     *mongo.Collection
	farmCollection       *mongo.Collection
	userCollection       *mongo.Collection
	productCollection    *mongo.Collection
	superAdmin           *mongo.Collection
	categoriesCollection *mongo.Collection
	orderCollection      *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) RepositoryInterface {
	return mongoRepo{
		adminCollection:      db.Collection("admin"),
		farmerCollection:     db.Collection("farmer"),
		farmCollection:       db.Collection("farm"),
		userCollection:       db.Collection("user"),
		productCollection:    db.Collection("product"),
		superAdmin:           db.Collection("super_admin"),
		categoriesCollection: db.Collection("categories"),
		orderCollection:      db.Collection("order"),
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

func (repo mongoRepo) GetAdmins(ctx context.Context) ([]models.Admin, error) {
	admins := []models.Admin{}
	cursor, err := repo.adminCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &admins); err != nil {
		return nil, err
	}

	return admins, nil
}

// Admin methods start
func (repo mongoRepo) LoginAdmin(ctx context.Context, username, password string) error {
	admin := models.Admin{}
	filter := bson.M{"username": username, "password": password}
	result := repo.adminCollection.FindOne(ctx, filter)
	err := result.Decode(&admin)

	if err != nil {
		return err
	}
	return nil
}

func (repo mongoRepo) CreateProduct(ctx context.Context, product models.Product) error {
	res, err := repo.productCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	log.Println("Inserted new product with ID : ", res.InsertedID)
	return nil
}

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
func (repo mongoRepo) ChangeFarmerState(ctx context.Context, farmerID, state string) error {

	farmer := models.Farmer{}

	filter := bson.M{"username": farmerID}
	log.Println("JERER")
	updation := bson.M{"$set": bson.M{"state": state}}

	result := repo.farmerCollection.FindOneAndUpdate(ctx, filter, updation)
	err := result.Decode(&farmer)
	if err != nil {
		return err
	}

	return nil
}

func (repo mongoRepo) CreateCategory(ctx context.Context, categoryName string) error {
	category := models.Category{
		CategoryName: categoryName,
	}

	opts := options.Update().SetUpsert(true)

	up, err := utils.ToDoc(category)
	if err != nil {
		return err
	}
	doc := bson.D{{"$set", up}}

	result, err := repo.categoriesCollection.UpdateOne(ctx, bson.M{"category_name": categoryName}, doc, opts)

	if result.MatchedCount != 0 {
		return errors.New("Category exists")
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo mongoRepo) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	categories := []models.Category{}
	cursor, err := repo.categoriesCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// Farmer methods
func (repo mongoRepo) CreateFarmer(ctx context.Context, farmer models.Farmer) error {
	opts := options.Update().SetUpsert(true)

	up, err := utils.ToDoc(farmer)
	if err != nil {
		return err
	}
	doc := bson.D{{"$set", up}}

	result, err := repo.farmerCollection.UpdateOne(ctx, bson.M{"firstname": farmer.FirstName, "lastname": farmer.LastName}, doc, opts)

	if result.MatchedCount != 0 {
		return errors.New("Farmer already exists")
	}

	if err != nil {
		return err
	}

	log.Println("New farmer created :", result)
	return nil

}

func (repo mongoRepo) GetFarmerWithUsername(ctx context.Context, username string) (models.Farmer, error) {
	farmer := models.Farmer{}
	filter := bson.M{"username": username}
	log.Println(username)
	result := repo.farmerCollection.FindOne(ctx, filter)
	err := result.Decode(&farmer)

	if err != nil {
		return models.Farmer{}, err
	}
	return farmer, nil
}

// user methods
func (repo mongoRepo) CreateUser(ctx context.Context, user models.User) error {

	opts := options.Update().SetUpsert(true)

	up, err := utils.ToDoc(user)
	if err != nil {
		return err
	}
	doc := bson.D{{"$set", up}}

	result, err := repo.userCollection.UpdateOne(ctx, bson.M{"firstname": user.FirstName, "lastname": user.LastName}, doc, opts)

	if result.MatchedCount != 0 {
		return errors.New("user already exists")
	}

	if err != nil {
		return err
	}

	log.Println("New farmer created :", result)
	return nil

}

func (repo mongoRepo) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	result, err := repo.orderCollection.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil
}

func (repo mongoRepo) ChangeItemInCart(ctx context.Context, action, username, productID string) error {

	if action == "increment" {
		res, err := repo.userCollection.UpdateOne(ctx, bson.M{"username": username, "cart.product_id": productID}, bson.M{
			"$inc": bson.M{"cart.$.count": 1},
		})

		if res.MatchedCount == 0 {
			repo.userCollection.UpdateOne(ctx, bson.M{"username": username}, bson.M{
				"$push": bson.M{"cart": models.CartItem{
					ProductID: productID,
					Count:     1,
				}},
			})
		}

		if err != nil {
			return err
		}

	} else {
		_, err := repo.userCollection.UpdateOne(ctx, bson.M{"username": username, "cart.product_id": productID}, bson.M{
			"$inc": bson.M{"cart.$.count": -1},
		})
		if err != nil {
			return err
		}
	}
	return nil

}

func (repo mongoRepo) GetUserWithUsername(ctx context.Context, username string) (models.User, error) {

	user := models.User{}
	filter := bson.M{"username": username}

	result := repo.userCollection.FindOne(ctx, filter)
	err := result.Decode(&user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (repo mongoRepo) GetProductsByCategory(ctx context.Context, category string) ([]models.Product, error) {
	products := []models.Product{}
	options := options.Find()

	options.SetLimit(10)

	cursor, err := repo.productCollection.Find(ctx, bson.M{}, options)

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}
