package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cambiar_api/internal"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

func main() {
	fmt.Println("===Hello, Cambiar!===")
	fmt.Println("Fetching env....")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("APP_PORT")
	dbDriver := os.Getenv("DATABASE_DRIVER")
	fmt.Println("Fetching env done successfully")
	if dbDriver != "mongo" {
		fmt.Printf("Database driver %s is not supported yet\n", dbDriver)
	}
	database := setUpMongo()
	container := setUpContainer(database)
	e := echo.New()
	setUpRouter(container, e)
	serveApp(e, port)
}

func setUpMongo() *mongo.Database {
	uri := os.Getenv("DATABASE_URI")
	database := os.Getenv("DATABASE_NAME")
	fmt.Printf("Seting up database... URI: %s, Database: %s\n", uri, database)

	if uri == "" {
		log.Fatal("You must set your 'DATABASE_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if database == "" {
		log.Fatal("You must set your 'DATABASE_NAME' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
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
	fmt.Println("Seting up database done successfully")
	return client.Database(database)
}

func setUpContainer(database *mongo.Database) *dig.Container {
	fmt.Println("Seting up container....")
	container, err := internal.BuildContainer(database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seting up container done successfully")
	return container
}

func setUpRouter(container *dig.Container, e *echo.Echo) {
	fmt.Println("Seting up route...")
	err := internal.RegisterRoutes(e, container)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seting up route done successfully")
}

func serveApp(e *echo.Echo, port string) {
	fmt.Printf("App running in :%s\n", port)
	err := e.Start(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
