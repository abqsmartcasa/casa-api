package models

// SpecificTag model for specific tags
type SpecificTag struct {
	ID         *int32  `json:"id"`
	Value      *string `json:"value"`
	CategoryID *int32  `json:"category_id,omitempty"`
}

// CategoryTag model for category tags
type CategoryTag struct {
	ID           *int32         `json:"id"`
	Value        *string        `json:"value"`
	SpecificTags []*SpecificTag `json:"specific_tags,omitempty"`
}
