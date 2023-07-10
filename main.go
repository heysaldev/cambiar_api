package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cambiar_api/internal/core/feature/model"
	"github.com/cambiar_api/internal/core/feature/repository/adapter"
	"github.com/cambiar_api/internal/core/feature/usecase"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()
	db := client.Database(database)
	repository := adapter.NewMongo(db)
	service := usecase.NewFeatureUsecase(repository)
	flag := "feature-flag-1"
	test := service.GetAllWithQuery(ctx, model.GetAllWithQuerySpec{Name: &flag})
	fmt.Printf("%+v\n", test)
}
