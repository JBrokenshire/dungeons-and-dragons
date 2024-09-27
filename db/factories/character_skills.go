package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterProficientSkill(db *gorm.DB, skill *models.CharacterProficientSkill) {
	fillSkillDetails(skill)
	db.Create(skill)
}

func fillSkillDetails(skill *models.CharacterProficientSkill) {
	if skill.SkillName == "" {
		skill.SkillName = "Acrobatics"
	}
	if skill.CharacterID == 0 {
		skill.CharacterID = 1
	}
	if skill.ProficiencyType == "" {
		skill.ProficiencyType = "Proficiency"
	}
}
