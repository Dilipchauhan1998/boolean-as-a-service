package models

//Boolean ... struct to store a boolean value
type Boolean struct {
	ID    string `json:"id" gorm:"primary_key"`
	Value *bool  `json:"value"  binding:"required" gorm:"not_null"`
	Key   string `json:"key,omitempty"`
}
