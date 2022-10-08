package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/JevaPrahaysuma/managemnet.git/collection"
	"github.com/JevaPrahaysuma/managemnet.git/config"
	"github.com/JevaPrahaysuma/managemnet.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStorage struct {
	Database *mongo.Database
	Timeout  time.Duration
}

func GetMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//ctx := context.Background()
	var DB = config.ConnectDB()

	//var menu []models.MenuItem
	var menu []bson.M
	//postId := c.Param("id")

	//fmt.Println(postId)

	defer cancel()

	qry := bson.D{
		{
			"$lookup", bson.D{
				// Define the tags collection for the join.
				{"from", "menu_page"},
				{"localField", "menu_page_id"},
				{"foreignField", "id"},
				{"pipeline", []bson.D{
					// Sort tags by their name field in asc. -1 = desc
					{
						{"from", "menu"},
						{"localField", "menu_id"},
						{"foreignField", "id"},
						{"as", "menus"},
						{"$sort", bson.D{
							{"id", 1},
						}},
					},
				}},
				{"as", "menu_pages"},
			},

			// Use tags as the field name to match struct field.
		},
		// {
		// 	"pipeline", bson.D{
		// 		{"$lookup", bson.D{
		// 			{"from", "dish"},
		// 			{"localField", "dish_id"},
		// 			{"foreignField", "$id"},
		// 			{"pipeline", []bson.D{
		// 				// Sort tags by their name field in asc. -1 = desc
		// 				{
		// 					{"$sort", bson.D{
		// 						{"id", 1},
		// 					}},
		// 				},
		// 			}},
		// 			{"as", "dishes"},
		// 		}},
		// 	},
		// },
	}
	// qry := []bson.M{
	// 	{
	// 		"$lookup": bson.M{
	// 			// Define the tags collection for the join.
	// 			"from":         "menu_page",
	// 			"localField":   "menu_page_id",
	// 			"foreignField": "id",
	// 			"pipeline": []bson.M{
	// 				// Sort tags by their name field in asc. -1 = desc
	// 				{
	// 					"$sort": bson.M{
	// 						"id": 1,
	// 					},
	// 				},
	// 			},
	// 		},
	// 		"as": "menu_pages",
	// 	},
	// 	// Use tags as the field name to match struct field.
	// 	// {
	// 	// 	"$lookup": bson.M{
	// 	// 		"from":         "dish",
	// 	// 		"localField":   "dish_id",
	// 	// 		"foreignField": "id",
	// 	// 	},
	// 	// 	"as": "dishes",
	// 	// },
	// 	{
	// 		"$limit": 10,
	// 	},
	// }
	qryDishies := bson.D{{
		"$lookup", bson.D{
			{"from", "dish"},
			{"localField", "dish_id"},
			{"foreignField", "id"},
			{"pipeline", []bson.D{
				// Sort tags by their name field in asc. -1 = desc
				{
					{"$sort", bson.D{
						{"id", 1},
					}},
				},
			}},
			{"as", "dishes"},
		},
	}}
	limit := bson.D{{
		"$limit", 10,
		// Use tags as the field name to match struct field.
	}}
	// select * from menu_page
	// join menu_item on menu_page.menu_item_id = menu_item.id
	// join dish on menu_page.dish_id = dish.id
	// join menu on menu.id = menu_item.menu_id
	// where

	//unwindStage := bson.D{{"$unwind", bson.D{{"path", "$menu_page_id"}, {"preserveNullAndEmptyArrays", false}}}}

	postCollection, err := collection.GetCollection(DB, "menu_item").Aggregate(ctx, mongo.Pipeline{qry, qryDishies, limit})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	//objId, _ := primitive.ObjectIDFromHex(postId)
	// var postCollection = collection.GetCollection(DB, "menu_item").Aggregate(ctx, qry)

	//err = postCollection.All(ctx, &menu)
	err = postCollection.All(ctx, &menu)

	fmt.Println(menu)

	//err := postCollection.All(ctx, &menu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)
	defer postCollection.Close(ctx)
	fmt.Println(menu)
	// for _, singleMenu := range menu {
	// 	//var singleMenu menu.Menu
	// 	// if err = postCollection.Decode(&singleMenu); err != nil {
	// 	// 	c.JSON(http.StatusInternalServerError, models.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	// 	// }
	// 	temp := singleMenu.Date
	// 	temp1 := temp.Format("2006-01-02")
	// 	singleMenu.Date, _ = time.Parse("2006-01-02", temp1)
	// 	//fmt.Println(singleMenu)

	// 	//menu = append(menu, singleMenu)
	// }

	c.JSON(http.StatusOK,
		models.ApiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": menu}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
}
