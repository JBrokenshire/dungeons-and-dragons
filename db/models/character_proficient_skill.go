package models

import (
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

func (c *CharacterProficientSkill) BeforeCreate(tx *gorm.DB) error {
	if c.ProficiencyType == "" {
		c.ProficiencyType = validProficiencyTypes[0]
	}

	if !sliceContains(validProficiencyTypes, c.ProficiencyType) {
		return fmt.Errorf("proficiency type '%s' is not valid", c.ProficiencyType)
	}

	if !sliceContains(validSkillNames, c.SkillName) {
		return fmt.Errorf("skill name '%s' is not valid", c.SkillName)
	}

	return nil
}

func sliceContains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
