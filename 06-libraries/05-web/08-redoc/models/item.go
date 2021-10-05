package models

// Item defines the structure for an API item
// swagger:model
type Item struct {
	// the id for the item
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the item

	// the name for this item
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this item
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
}
