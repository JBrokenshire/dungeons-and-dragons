package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterDefense(db *gorm.DB, defense *models.CharacterDefense) {
	fillDefenseDetails(defense)
	db.Create(defense)
}

func fillDefenseDetails(defense *models.CharacterDefense) {
	if defense.DefenseType == "" {
		defense.DefenseType = "Resistance"
	}
	if defense.DamageType == "" {
		defense.DamageType = "Fire"
	}
}
