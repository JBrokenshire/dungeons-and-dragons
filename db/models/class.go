package models

type Class struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
