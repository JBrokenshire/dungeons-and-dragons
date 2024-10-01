package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterProficientTool(db *gorm.DB, tool *models.CharacterProficientTool) {
	fillProficientToolDetails(tool)
	db.Create(tool)
}

func fillProficientToolDetails(tool *models.CharacterProficientTool) {
	if tool.Tool == "" {
		tool.Tool = "Lute"
	}
}
