package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterInventoryItem(db *gorm.DB, item *models.CharacterInventoryItem) {
	fillCharacterInventoryItemDetails(item)
	db.Create(item)
}

func fillCharacterInventoryItemDetails(item *models.CharacterInventoryItem) {
	if item.Location == "" {
		item.Location = "Equipment"
	}
	if item.Quantity == 0 {
		item.Quantity = 1
	}
}
