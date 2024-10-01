package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validTools = []string{
	"Lute",
	"Thieves' Tools",
	"Playing Card Set",
	"Smith's Tools",
	"Vehicles (Land)",
	"Alchemist's Supplies",
	"Drum",
	"Viol",
}

type CharacterProficientTool struct {
	ID          int    `gorm:"primary_key auto_increment" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	Tool        string `gorm:"not null" json:"tool"`
}

func (c *CharacterProficientTool) BeforeCreate(_ *gorm.DB) error {
	if c.Tool == "" {
		c.Tool = validTools[0]
	}

	if !utils.SliceContains(validTools, c.Tool) {
		return fmt.Errorf("tool '%s' is not valid", c.Tool)
	}

	return nil
}
