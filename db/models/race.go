package models

type Race struct {
	ID               int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
}
