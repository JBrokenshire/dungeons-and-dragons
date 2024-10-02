package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterConditionsStore interface {
	GetConditionsByCharacterID(id interface{}) ([]*models.CharacterCondition, error)
}

type GormCharacterConditionsStore struct {
	DB *gorm.DB
}

func NewGormCharacterConditionsStore(db *gorm.DB) *GormCharacterConditionsStore {
	return &GormCharacterConditionsStore{
		DB: db,
	}
}

func (g *GormCharacterConditionsStore) GetConditionsByCharacterID(id interface{}) ([]*models.CharacterCondition, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterConditions []*models.CharacterCondition
	err := g.DB.Where("character_id = ?", id).Find(&characterConditions).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("conditions with character id: %q could not be found", id))
	}

	return characterConditions, nil
}
