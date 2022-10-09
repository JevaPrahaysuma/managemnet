package models

type Dish struct {
	ID            int64  `json:"id,omitempty" field:"id" bson:"id,omitempty" gorm:"primary_key"`
	Name          string `json:"name,omitempty" field:"name" bson:"name,omitempty"`
	Description   string `json:"description,omitempty" field:"description" bson:"description,omitempty"`
	MenusAppeared int64  `json:"menus_appeared,omitempty" field:"menus_appeared" bson:"menus_appeared,omitempty"`
	TimesAppeared int64  `json:"times_appeared,omitempty" field:"times_appeared" bson:"times_appeared,omitempty"`
	FirstAppeared int64  `json:"first_appeared,omitempty" field:"first_appeared" bson:"first_appeared,omitempty"`
	LastAppeared  int64  `json:"last_appeared,omitempty" field:"last_appeared" bson:"last_appeared,omitempty"`
	LowestPrice   string `json:"lowest_price,omitempty" field:"lowest_price" bson:"lowest_price,omitempty"`
	HighestPrice  string `json:"highest_price,omitempty" field:"highest_price" bson:"highest_price,omitempty"`
}
