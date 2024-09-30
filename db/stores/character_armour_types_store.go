package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterArmourTypesStore interface {
	GetProficientArmourTypesByCharacterID(id interface{}) ([]*models.CharacterProficientArmourType, error)
}

type GormCharacterArmourTypesStore struct {
	DB *gorm.DB
}

func NewGormCharacterArmourTypesStore(db *gorm.DB) *GormCharacterArmourTypesStore {
	return &GormCharacterArmourTypesStore{
		DB: db,
	}
}

func (g *GormCharacterArmourTypesStore) GetProficientArmourTypesByCharacterID(id interface{}) ([]*models.CharacterProficientArmourType, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var armourTypes []*models.CharacterProficientArmourType
	err := g.DB.Where("character_id = ?", id).Find(&armourTypes).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proficient armour types with character id: %q could not be found", id))
	}

	return armourTypes, nil
}
