package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validWeapons = []string{
	"Martial Weapons",
	"Simple Weapons",
}

type CharacterProficientWeapon struct {
	ID          int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Weapon      string `gorm:"not null" json:"weapon"`
}

func (c *CharacterProficientWeapon) BeforeCreate(_ *gorm.DB) error {
	if !utils.SliceContains(validWeapons, c.Weapon) {
		return fmt.Errorf("weapon '%s' is not valid", c.Weapon)
	}

	return nil
}
