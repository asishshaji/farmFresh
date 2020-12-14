package main

import (
	"log"
	"os"

	"github.com/asishshaji/freshFarm/app"
	"github.com/asishshaji/freshFarm/app/controller"
	"github.com/asishshaji/freshFarm/app/repository"
	"github.com/asishshaji/freshFarm/app/usecase"
	"github.com/asishshaji/freshFarm/app/utils"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	log.Println("Loaded .env file ")
}

func main() {
	loadEnv()
	port := os.Getenv("PORT")

	dbName := os.Getenv("DB")
	mongodbURL := os.Getenv("MONGODB_URL")

	// Initialising database and cloud storage
	db := utils.InitDB(mongodbURL, dbName)

	repo := repository.NewMongoRepository(db)
	usecase := usecase.NewUsecase(repo)
	controller := controller.NewEchoController(usecase)

	app := app.NewApp(port, controller)

	app.RunServer()

}
