package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JevaPrahaysuma/managemnet.git/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//const uri = "mongodb://user:pass@sample.host:2701/?maxPoolSize=20&w=majority"

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}

func ConnectDB() *mongo.Client {
	Mongo_URL := "mongodb://0.0.0.0:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return client
}

// func ConnectMongo() (*mongo.Client, context.Context,
// 	context.CancelFunc, error) {

// 	// ctx will be used to set deadline for process, here
// 	// deadline will of 30 seconds.
// 	ctx, cancel := context.WithTimeout(context.Background(),
// 		30*time.Second)

// 	// mongo.Connect return mongo.Client method
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	return client, ctx, cancel, err
// }
