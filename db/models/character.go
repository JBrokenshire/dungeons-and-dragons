package models

type Character struct {
	ID      int    `gorm:"autoIncrement;primary_key" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	Level   int    `gorm:"not null" json:"level"`
	ClassID int    `json:"class_id"`
	RaceID  int    `json:"race_id"`

	Class *Class `json:"class" gorm:"foreignKey:ClassID;references:ID"`
	Race  *Race  `json:"race" gorm:"foreignKey:RaceID;references:ID"`
}
