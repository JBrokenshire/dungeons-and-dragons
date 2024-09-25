package models

type Class struct {
	ID               int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Colour           string `json:"colour"`
}
