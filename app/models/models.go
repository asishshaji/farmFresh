package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	Username        string             `json:"username" bson:"username"`
	Password        string             `json:"-" bson:"password"`
	ProfileImageURL string             `json:"image_url" bson:"image_url"`
	JoinedOn        primitive.DateTime `json:"joined_on" bson:"joined_on"`
}

type Farm struct {
	OwnerID             string   `json:"owner_id"`
	LocationCoordinates string   `json:"coordinates"`
	ImageUrls           []string `json:"image_urls"`
	About               string   `json:"farm_details"`
	Rating              float32  `json:"rating"`
	Reviews             []Review `json:"reviews"`
	State               string   `json:"state"` // can be active or suspended

}

type Farmer struct {
	Username        string             `json:"username"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Password        string             `json:"-"`
	Age             int                `json:"age"`
	About           string             `json:"about"`
	Farms           []Farm             `json:"farms"`
	JoinedOn        primitive.DateTime `json:"joined_on"`
	Rating          float32            `json:"rating"`
	Score           float32            `json:"score"`
	ProfileImageURL string             `json:"image_url"`
	State           string             `json:"state"` // can be under review, active or suspended
	Reviews         []Review           `json:"reviews"`
	Profit          float64            `json:"profit" bson:"profit"`
}

type Review struct {
	PostedByName string `json:"posted_by"`
	Content      string `json:"content"`
}

type User struct {
	Username        string             `json:"username"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	Password        string             `json:"-"`
	JoinedOn        primitive.DateTime `json:"joined_on"`
	ProfileImageURL string             `json:"image_url"`
	State           string             `json:"state"` // can be active or suspended
	FavoriteFarmers []Farmer           `json:"fav_farmers"`
	FavoriteFarms   []Farm             `json:"fav_farms"`
	UserCart        []CartItem         `json:"cart" bson:"cart"`
}

type CartItem struct {
	ProductID string `json:"product_id"`
	Count     int    `json:"count"`
}

type Order struct {
	OrderedBy  string     `json:"ordered_by"`
	TotalPrice float64    `json:"total_price"`
	Products   []CartItem `json:"products"`
	Status     string     `json:"status"`
}

type Nutrition struct {
	Type  string  `json:"type"`
	Score float32 `json:"score"`
}

type Product struct {
	Name              string             `json:"name"`
	ImageURLS         []string           `json:"image_urls"`
	CreatedAt         primitive.DateTime `json:"created_at"`
	OwnerID           string             `json:"owner_id"`
	Price             float64            `json:"vk_price"`
	MarketPrice       float64            `json:"mk_price"`
	Reviews           []Review           `json:"reviews"`
	NutritionalValues []Nutrition        `json:"nutrition_values"`
	About             string             `json:"about"`
	State             string             `json:"state"`
	TypeOfMeasurement string             `json:"measurement_type"`
	ProductCount      int                `json:"count"`
	Category          Category           `json:"category"`
}

type Category struct {
	CategoryName string `json:"category_name" bson:"category_name"`
}
