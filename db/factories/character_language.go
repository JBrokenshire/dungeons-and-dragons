package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterLanguage(db *gorm.DB, language *models.CharacterLanguage) {
	fillLanguageDetails(language)
	db.Create(language)
}

func fillLanguageDetails(language *models.CharacterLanguage) {
	if language.Language == "" {
		language.Language = "Abyssal"
	}
}
