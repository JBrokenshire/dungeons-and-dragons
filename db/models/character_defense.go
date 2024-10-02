package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validDefenseTypes = []string{
	"Resistance",
	"Immunity",
	"Vulnerability",
}

type CharacterDefense struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	DamageType  string `gorm:"not null" json:"damage_type"`
	DefenseType string `gorm:"not null" json:"defense_type"`
}

func (c *CharacterDefense) BeforeCreate(_ *gorm.DB) error {
	if !utils.SliceContains(validDefenseTypes, c.DefenseType) {
		return fmt.Errorf("defense type '%s' is not valid", c.DefenseType)
	}

	return nil
}
