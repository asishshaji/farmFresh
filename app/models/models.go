package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	ProfileImageURL string             `bson:"image_url"`
	JoinedOn        primitive.DateTime `bson:"joined_on"`
}

type Farm struct {
	ID                  primitive.ObjectID `bson:"_id"`
	OwnerID             primitive.ObjectID `bson:"owner_id"`
	LocationCoordinates string             `bson:"coordinates"`
	ImageUrls           []string           `bson:"image_urls"`
	About               string             `bson:"farm_details"`
	Rating              float32            `bson:"rating"`
	Reviews             []Review           `bson:"reviews"`
	State               string             `bson:"state"` // can be active or suspended

}

type Farmer struct {
	ID              primitive.ObjectID `bson:"_id"`
	Username        string             `bson:"username"`
	FirstName       string             `bson:"first_name"`
	LastName        string             `bson:"last_name"`
	Password        string             `bson:"password"`
	Age             int                `bson:"age"`
	About           string             `bson:"about"`
	Farms           []Farm             `bson:"farms"`
	JoinedOn        primitive.DateTime `bson:"joined_on"`
	Rating          float32            `bson:"rating"`
	Score           float32            `bson:"score"`
	ProfileImageURL string             `bson:"image_url"`
	State           string             `bson:"state"` // can be under review, active or suspended
	Reviews         []Review           `bson:"reviews"`
}

type Review struct {
	ID           primitive.ObjectID `bson:"_id"`
	PostedByName string             `bson:"posted_by"`
	Content      string             `bson:"content"`
}

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Username        string             `bson:"username"`
	FirstName       string             `bson:"first_name"`
	LastName        string             `bson:"last_name"`
	Password        string             `bson:"password"`
	Age             int                `bson:"age"`
	JoinedOn        primitive.DateTime `bson:"joined_on"`
	ProfileImageURL string             `bson:"image_url"`
	State           string             `bson:"state"` // can be active or suspended
	FavoriteFarmers []Farmer           `bson:"fav_farmers"`
	FavoriteFarms   []Farm             `bson:"fav_farms"`

	// Implement cart
}

type Nutrition struct {
	ID    primitive.ObjectID `bson:"_id"`
	Type  string             `bson:"type"`
	Score float32            `bson:"score"`
}

type Product struct {
	ID                primitive.ObjectID `bson:"_id"`
	Name              string             `bson:"name"`
	ImageURLS         []string           `bson:"image_urls"`
	CreatedAt         primitive.DateTime `bson:"created_at"`
	OwnerID           primitive.ObjectID `bson:"owner_id"`
	Price             float32            `bson:"vk_price"`
	MarketPrice       float32            `bson:"mk_price"`
	Reviews           []Review           `bson:"reviews"`
	NutritionalValues []Nutrition        `bson:"nutrition_values"`
	About             string             `bson:"about"`
	State             string             `bson:"state"`
	TypeOfMeasurement string             `bson:"measurement_type"`
	ProductCount      int                `bson:"count"`
}
