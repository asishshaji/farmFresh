package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	firebase "firebase.google.com/go"

	"cloud.google.com/go/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

// InitDB creates a connection to MongoDB instance
func InitDB(mongodbURL, dbName string) *mongo.Database {
	log.Printf("Starting connection to MongoDB at : %v", mongodbURL)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB. Make sure mongodb instance is running.")

	}

	log.Println("Connected to MongoDB")

	return client.Database("farm")
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

// UploadImage uploads images to server
func UploadImage(file *multipart.FileHeader, bucket *storage.BucketHandle) (string, error) {

	src, err := file.Open()

	if err != nil {
		return "", fmt.Errorf("Failed to run os.Open:  %v", err)
	}
	defer src.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	sw := bucket.Object(file.Filename).NewWriter(ctx)

	if _, err = io.Copy(sw, src); err != nil {
		log.Println("Error copying to storage : ", err)
		return "", fmt.Errorf("Error copying %v", err)
	}

	if err := sw.Close(); err != nil {
		log.Println("Error copying to storage : ", err)
		return "", fmt.Errorf("Error copying %v", err)
	}

	url := "https://storage.googleapis.com/golang-f2010.appspot.com/" + file.Filename

	return url, nil
}

// InitStorage initializes the google cloud storageG
func InitStorage(storageBucket, credentialFilePath string) *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: storageBucket,
	}

	opt := option.WithCredentialsFile(credentialFilePath)

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket

}
