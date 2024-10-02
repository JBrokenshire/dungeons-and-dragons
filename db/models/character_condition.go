package models

import (
	"dnd-api/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validConditions = []string{
	"Blinded",
	"Charmed",
	"Deafened",
	"Frightened",
	"Grappled",
	"Incapacitated",
	"Invisible",
	"Paralysed",
	"Petrified",
	"Poisoned",
	"Prone",
	"Restrained",
	"Stunned",
	"Unconscious",
}

type CharacterCondition struct {
	ID            int    `gorm:"primary_key;auto_increment" json:"id"`
	CharacterID   int    `gorm:"not null" json:"character_id"`
	ConditionName string `gorm:"not null" json:"condition_name"`
}

func (c *CharacterCondition) BeforeCreate(_ *gorm.DB) error {
	if !utils.SliceContains(validConditions, c.ConditionName) {
		return errors.New(fmt.Sprintf("condition '%s' is not valid", c.ConditionName))
	}

	return nil
}
