package models

type ListMenu struct {
	MenuItemID          int64   `json:"menu_item_id,omitempty" field:"menu_item_id" bson:"menu_item_id,omitempty"`
	NameDish            string  `json:"name_dish,omitempty" field:"name_dish" bson:"name_dish,omitempty"`
	DescriptionDish     string  `json:"description_dish,omitempty" field:"description_dish" bson:"description,omitempty"`
	Price               float64 `json:"price,omitempty" field:"price" bson:"price,omitempty"`
	NameMenu            string  `json:"name_menu,omitempty" field:"name_menu" bson:"name_menu,omitempty"`
	Sponsor             string  `json:"sponsor,omitempty" field:"sponsor" bson:"sponsor,omitempty"`
	Event               string  `json:"event,omitempty" field:"event" bson:"event,omitempty"`
	Vanue               string  `json:"vanue,omitempty" field:"vanue" bson:"vanue,omitempty"`
	Place               string  `json:"place,omitempty" field:"place" bson:"place,omitempty"`
	PhysicalDescription string  `json:"physical_description,omitempty" field:"physical_description" bson:"physical_description,omitempty"`
	Occasion            string  `json:"occasion,omitempty" field:"occasion" bson:"occasion,omitempty"`
	Notes               string  `json:"notes,omitempty" field:"notes" bson:"notes,omitempty"`
	CallNumber          string  `json:"call_number,omitempty" field:"call_number" bson:"call_number,omitempty"`
	Keywords            string  `json:"keywords,omitempty" field:"keywords" bson:"keywords,omitempty"`
	Language            string  `json:"language,omitempty" field:"language" bson:"language,omitempty"`
	Date                string  `json:"date,omitempty" field:"date" bson:"date,omitempty"`
	Location            string  `json:"location,omitempty" field:"location" bson:"location,omitempty"`
	LocationType        string  `json:"location_type,omitempty" field:"location_type" bson:"location_type,omitempty"`
	Currency            string  `json:"currency,omitempty" field:"currency" bson:"currency,omitempty"`
	CurrencySymbol      string  `json:"currency_symbol,omitempty" field:"currency_symbol" bson:"currency_symbol,omitempty"`
	Status              string  `json:"status,omitempty" field:"status" bson:"status,omitempty"`
	CreatedAt           string  `json:"created_at,omitempty" field:"created_at" bson:"created_at,omitempty"`
	UpdatedAt           string  `json:"updated_at,omitempty" field:"updated_at" bson:"updated_at,omitempty"`
}
