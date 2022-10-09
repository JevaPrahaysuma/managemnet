package models

type Menu struct {
	ID                  int64  `json:"id,omitempty" field:"id" bson:"id,omitempty" gorm:"primary_key"`
	Name                string `json:"name,omitempty" field:"name" bson:"name,omitempty"`
	Sponsor             string `json:"sponsor,omitempty" field:"sponsor" bson:"sponsor,omitempty"`
	Event               string `json:"event,omitempty" field:"event" bson:"event,omitempty"`
	Vanue               string `json:"vanue,omitempty" field:"vanue" bson:"vanue,omitempty"`
	Place               string `json:"place,omitempty" field:"place" bson:"place,omitempty"`
	PhysicalDescription string `json:"physical_description,omitempty" field:"physical_description" bson:"physical_description,omitempty"`
	Occasion            string `json:"occasion,omitempty" field:"occasion" bson:"occasion,omitempty"`
	Notes               string `json:"notes,omitempty" field:"notes" bson:"notes,omitempty"`
	CallNumber          string `json:"call_number,omitempty" field:"call_number" bson:"call_number,omitempty"`
	Keywords            string `json:"keywords,omitempty" field:"keywords" bson:"keywords,omitempty"`
	Language            string `json:"language,omitempty" field:"language" bson:"language,omitempty"`
	Date                string `json:"date,omitempty" field:"date" bson:"date,omitempty"`
	Location            string `json:"location,omitempty" field:"location" bson:"location,omitempty"`
	LocationType        string `json:"location_type,omitempty" field:"location_type" bson:"location_type,omitempty"`
	Currency            string `json:"currency,omitempty" field:"currency" bson:"currency,omitempty"`
	CurrencySymbol      string `json:"currency_symbol,omitempty" field:"currency_symbol" bson:"currency_symbol,omitempty"`
	Status              string `json:"status,omitempty" field:"status" bson:"status,omitempty"`
	PageCount           int64  `json:"page_count,omitempty" field:"page_count" bson:"page_count,omitempty"`
	DishCount           int64  `json:"dish_count,omitempty" field:"dish_count" bson:"dish_count,omitempty"`
}
