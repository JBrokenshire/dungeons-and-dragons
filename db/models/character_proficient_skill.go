package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validProficiencyTypes = []string{
	"Proficiency",
	"Half-Proficiency",
	"Expertise",
}

var validSkillNames = []string{
	"Acrobatics",
	"Animal Handling",
	"Arcana",
	"Athletics",
	"Deception",
	"History",
	"Insight",
	"Intimidation",
	"Investigation",
	"Medicine",
	"Nature",
	"Perception",
	"Performance",
	"Persuasion",
	"Religion",
	"Sleight of Hand",
	"Stealth",
	"Survival",
}

type CharacterProficientSkill struct {
	ID              int    `gorm:"autoIncrement;primary_key" json:"id"`
	CharacterID     int    `gorm:"not null" json:"character_id"`
	SkillName       string `gorm:"not null" json:"skill_name"`
	ProficiencyType string `gorm:"not null default:'Proficiency'" json:"proficiency_type"`
}

func (c *CharacterProficientSkill) BeforeCreate(_ *gorm.DB) error {
	if c.ProficiencyType == "" {
		c.ProficiencyType = validProficiencyTypes[0]
	}

	if !utils.SliceContains(validProficiencyTypes, c.ProficiencyType) {
		return fmt.Errorf("proficiency type '%s' is not valid", c.ProficiencyType)
	}

	if !utils.SliceContains(validSkillNames, c.SkillName) {
		return fmt.Errorf("skill name '%s' is not valid", c.SkillName)
	}

	return nil
}
