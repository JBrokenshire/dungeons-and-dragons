package models

type Character struct {
	ID                int    `gorm:"autoIncrement;primary_key" json:"id"`
	Name              string `gorm:"not null" json:"name"`
	Level             int    `gorm:"not null" json:"level"`
	ProfilePictureURL string `json:"profile_picture_url"`
	ClassID           int    `json:"class_id"`
	RaceID            int    `json:"race_id"`
	Strength          int    `gorm:"not null" json:"strength"`
	Dexterity         int    `gorm:"not null" json:"dexterity"`
	Constitution      int    `gorm:"not null" json:"constitution"`
	Intelligence      int    `gorm:"not null" json:"intelligence"`
	Wisdom            int    `gorm:"not null" json:"wisdom"`
	Charisma          int    `gorm:"not null" json:"charisma"`

	Class Class `json:"class"`
	Race  Race  `json:"race"`
}
