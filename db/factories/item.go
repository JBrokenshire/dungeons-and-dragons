package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/random"
)

func NewItem(db *gorm.DB, item *models.Item) {
	fillItemDetails(item)
	db.Create(item)
}

func fillItemDetails(item *models.Item) {
	if item.Name == "" {
		item.Name = random.String(16)
	}
	if item.Meta == "" {
		item.Meta = random.String(64)
	}
}
