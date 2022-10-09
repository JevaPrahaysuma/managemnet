package models

type MenuItem struct {
	ID         int64    `json:"id,omitempty" field:"id" bson:"id,omitempty" gorm:"primary_key"`
	MenuPageId int64    `json:"menu_page_id,omitempty" field:"menu_page_id" bson:"menu_page_id,omitempty"`
	Price      string   `json:"price,omitempty" field:"price" bson:"price,omitempty"`
	HighPrice  string   `json:"high_price,omitempty" field:"high_price" bson:"high_price,omitempty"`
	DishId     int64    `json:"dish_id,omitempty" field:"dish_id" bson:"dish_id,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty" field:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty" field:"updated_at" bson:"updated_at,omitempty"`
	Xpos       float64  `json:"xpos,omitempty" field:"xpos" bson:"xpos,omitempty"`
	Ypos       float64  `json:"ypos,omitempty" field:"ypos" bson:"ypos,omitempty"`
	MenuPage   MenuPage `json:"menu_page,omitempty" bson:menu_page gorm:"references:id"`
	Dish       Dish     `json:"dish,omitempty" field:"dishes" bson:dish,omitempty gorm:"references:id"`
}
