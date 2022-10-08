package collection

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("nypl").Collection(collectionName)

	fmt.Println(collection)
	return collection
}
