package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterSensesStore interface {
	GetSensesByCharacterID(id interface{}) ([]*models.CharacterSense, error)
}

type GormCharacterSensesStore struct {
	DB *gorm.DB
}

func NewGormCharacterSensesStore(db *gorm.DB) *GormCharacterSensesStore {
	return &GormCharacterSensesStore{
		DB: db,
	}
}

func (g *GormCharacterSensesStore) GetSensesByCharacterID(id interface{}) ([]*models.CharacterSense, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", id).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", id, err))
	}

	var characterSenses []*models.CharacterSense
	if err = g.DB.Where("character_id = ?", id).Find(&characterSenses).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character senses with character id: %q could not be found", id))
	}

	return characterSenses, nil
}
