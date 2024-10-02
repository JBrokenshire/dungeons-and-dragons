package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterCondition(db *gorm.DB, condition *models.CharacterCondition) {
	fillConditionDetails(condition)
	db.Create(condition)
}

func fillConditionDetails(defense *models.CharacterCondition) {
	if defense.ConditionName == "" {
		defense.ConditionName = "Blinded"
	}
}
