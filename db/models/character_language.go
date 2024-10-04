package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validLanguages = []string{
	"Abyssal",
	"Celestial",
	"Deep Speech",
	"Draconic",
	"Dwarvish",
	"Elvish",
	"Giant",
	"Gnomish",
	"Goblin",
	"Halfling",
	"Infernal",
	"Orc",
	"Primordial",
	"Sylvan",
	"Undercommon",
}

type CharacterLanguage struct {
	ID          int    `gorm:"primary_key auto_increment" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Language    string `gorm:"not null" json:"language"`
}

func (c *CharacterLanguage) BeforeCreate(db *gorm.DB) error {
	var character Character
	err := db.Where("id = ?", c.CharacterID).Find(&character).Error
	if err != nil {
		return fmt.Errorf("character with id '%v' not found", c.CharacterID)
	}

	if !utils.SliceContains(validLanguages, c.Language) {
		return fmt.Errorf("language '%s' is not valid", c.Language)
	}

	return nil
}
