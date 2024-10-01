package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterProficienciesStore interface {
	GetProficientArmourTypesByCharacterID(id interface{}) ([]*models.CharacterProficientArmourType, error)
	GetProficientWeaponsByCharacterID(id interface{}) ([]*models.CharacterProficientWeapon, error)
}

type GormCharacterProficienciesStore struct {
	DB *gorm.DB
}

func NewGormCharacterProficienciesStore(db *gorm.DB) *GormCharacterProficienciesStore {
	return &GormCharacterProficienciesStore{
		DB: db,
	}
}

func (g *GormCharacterProficienciesStore) GetProficientArmourTypesByCharacterID(id interface{}) ([]*models.CharacterProficientArmourType, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var armourTypes []*models.CharacterProficientArmourType
	err := g.DB.
		Where("character_id = ?", id).
		Find(&armourTypes).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proficient armour types with character id: %q could not be found", id))
	}

	return armourTypes, nil
}

func (g *GormCharacterProficienciesStore) GetProficientWeaponsByCharacterID(id interface{}) ([]*models.CharacterProficientWeapon, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var weapons []*models.CharacterProficientWeapon
	err := g.DB.
		Where("character_id = ?", id).
		Find(&weapons).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proficient weapons with character id: %d could not be found", id))
	}

	return weapons, nil
}
