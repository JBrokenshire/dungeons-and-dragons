package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterDefensesStore interface {
	GetDefensesByCharacterID(id interface{}) ([]*models.CharacterDefense, error)
}

type GormCharacterDefensesStore struct {
	DB *gorm.DB
}

func NewGormCharacterDefensesStore(db *gorm.DB) *GormCharacterDefensesStore {
	return &GormCharacterDefensesStore{DB: db}
}

func (g *GormCharacterDefensesStore) GetDefensesByCharacterID(id interface{}) ([]*models.CharacterDefense, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterDefenses []*models.CharacterDefense
	err := g.DB.Where("character_id = ?", id).Find(&characterDefenses).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("defenses with character id: %q could not be found", id))
	}

	return characterDefenses, nil
}
