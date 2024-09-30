package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterProficientArmourType(db *gorm.DB, armourType *models.CharacterProficientArmourType) {
	fillProficientArmourTypeDetails(armourType)
	db.Create(armourType)
}

func fillProficientArmourTypeDetails(armourType *models.CharacterProficientArmourType) {
	if armourType.ArmourType == "" {
		armourType.ArmourType = "Light Armour"
	}
}
