package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterSense(db *gorm.DB, sense *models.CharacterSense) {
	fillSenseDetails(sense)
	db.Create(sense)
}

func fillSenseDetails(sense *models.CharacterSense) {
	if sense.SenseName == "" {
		sense.SenseName = "Darkvision"
	}
	if sense.Distance == 0 {
		sense.Distance = 60
	}
}
