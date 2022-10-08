package models

type MenuPage struct {
	ID         int64  `json:"id,omitempty" field:"id" bson:"id,omitempty"`
	MenuId     int64  `json:"menu_id,omitempty" field:"menu_id" bson:"menu_id,omitempty"`
	PageNumber int64  `json:"page_number,omitempty" field:"page_number" bson:"page_number,omitempty"`
	ImageId    int64  `json:"image_id,omitempty" field:"image_id" bson:"image_id,omitempty"`
	FullHeight string `json:"full_height,omitempty" field:"full_height" bson:"full_height,omitempty"`
	FullWidth  string `json:"full_width,omitempty" field:"full_width" bson:"full_width,omitempty"`
	Uuid       string `json:"uuid,omitempty" field:"uuid" bson:"uuid,omitempty"`
	Menu       []Menu `json:"menus,omitempty" bson:menus,omitempty`
}
