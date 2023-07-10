package main

import (
	"context"
	"log"
	"os"

	"github.com/cambiar_api/internal/core/feature/repository/adapter"
	"github.com/cambiar_api/internal/core/feature/usecase"
	"github.com/cambiar_api/internal/handler/feature"
	"github.com/cambiar_api/internal/handler/feature/controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	db := client.Database(database)
	repository := adapter.NewFeatureMongo(db)
	service := usecase.NewFeatureUsecase(repository)
	ctrl := controller.NewFeatureController(service)

	e := echo.New()

	// Initialize the router
	router := feature.NewRouter(ctrl)
	router.RegisterRoutes(e)

	// Register the application instance with the Echo context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", ctrl)
			return next(c)
		}
	})

	// Start the server
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
