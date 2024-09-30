package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validSenses = []string{
	"Blindsight",
	"Darkvision",
	"Truesight",
}

type CharacterSense struct {
	ID          int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID int    `gorm:"not null" json:"character_id"`
	SenseName   string `gorm:"not null" json:"sense_name"`
	Distance    int    `gorm:"not null" json:"distance"`
}

func (c *CharacterSense) BeforeCreate(_ *gorm.DB) error {
	if c.SenseName == "" {
		return fmt.Errorf("sense name is required")
	}

	if !utils.SliceContains(validSenses, c.SenseName) {
		return fmt.Errorf("proficiency type '%s' is not valid", c.SenseName)
	}

	if c.Distance <= 0 {
		return fmt.Errorf("distance needs to be a positive number")
	}

	return nil
}
