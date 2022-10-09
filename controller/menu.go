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
				{"as", "menu_pages"},
			},

			// Use tags as the field name to match struct field.
		},
	}

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

func GetMenuPostgres(c *gin.Context) {
	var qQuery string
	temp := 0
	// check search menu and dish
	if c.Query("q") != "" {
		qQuery += `(menus.name LIKE '%` + c.Query("q") + `%' OR dishes.name LIKE '%` + c.Query("q") + `%') `
		temp++
	}
	var qVanue string
	if c.Query("vanues") != "" {
		if temp != 0 {
			qVanue = "AND"
		}
		qVanue += ` menus.vanue like '%` + c.Query("vanues") + `%' `
		temp++
	}
	var qEvent string
	if c.Query("events") != "" {
		if temp != 0 {
			qEvent = "AND"
		}
		qEvent += ` menus.event LIKE '%` + c.Query("events") + `%' `
		temp++
	}
	var qPrice string
	if c.Query("startPrice") != "" || c.Query("endPrice") != "" {
		tempStart := c.DefaultQuery("startPrice", "0")
		tempEnd := c.DefaultQuery("endPrice", "0")
		if tempStart == "" {
			tempStart = "0"
		}
		if tempEnd == "" {
			tempEnd = "0"
		}
		if temp != 0 {
			qPrice = "AND"
		}
		qPrice += ` (menu_items.price >= ` + tempStart + ` AND menu_items.price <= ` + tempEnd + `)`
		temp++
	}

	var menuList []models.ListMenu
	if err := config.DB.Table("menu_items").Select(`menu_items.id as menu_item_id, dishes.name as name_dish, 
	dishes.description as description_dish, menu_items.price, menus.name as name_menu,
	menus.sponsor, menus.event, menus.vanue, menus.place, menus.physical_description,
	menus.occasion, menus.notes, menus.call_number, menus.keywords, menus.language,
	menus.date, menus.location, menus.location_type, menus.currency, menus.currency_symbol,
	menus.status, menu_items.created_at, menu_items.updated_at`).Joins(`LEFT JOIN menu_pages on 
	menu_items.menu_page_id = menu_pages.id`).Joins(`LEFT JOIN dishes on 
	menu_items.dish_id = dishes.id`).Joins(`LEFT JOIN menus on 
	menu_pages.menu_id = menus.id`).Where(qQuery + qVanue + qEvent + qPrice).Limit(10).Find(&menuList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	// if err := config.DB.Limit(10).Find(&menuPage).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, models.ApiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
	// 	return
	// }
	c.JSON(http.StatusOK,
		models.ApiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": menuList}},
	)

	//sqlDB.Close()
}
