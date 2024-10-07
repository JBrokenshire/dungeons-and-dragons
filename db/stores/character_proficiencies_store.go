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
	GetProficientToolsByCharacterID(id interface{}) ([]*models.CharacterProficientTool, error)
	GetLanguagesByCharacterID(id interface{}) ([]*models.CharacterLanguage, error)
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

	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", id).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", id, err))
	}

	var armourTypes []*models.CharacterProficientArmourType
	err = g.DB.
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

	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", id).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", id, err))
	}

	var weapons []*models.CharacterProficientWeapon
	err = g.DB.
		Where("character_id = ?", id).
		Find(&weapons).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proficient weapons with character id: %d could not be found", id))
	}

	return weapons, nil
}

func (g *GormCharacterProficienciesStore) GetProficientToolsByCharacterID(id interface{}) ([]*models.CharacterProficientTool, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", id).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", id, err))
	}

	var tools []*models.CharacterProficientTool
	err = g.DB.
		Where("character_id = ?", id).
		Find(&tools).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proficient tools with character id: %d could not be found", id))
	}

	return tools, nil
}

func (g *GormCharacterProficienciesStore) GetLanguagesByCharacterID(id interface{}) ([]*models.CharacterLanguage, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var character models.Character
	err := g.DB.Table("characters").Where("id = ?", id).Find(&character).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error getting character with id %v: %v", id, err))
	}

	var languages []*models.CharacterLanguage
	err = g.DB.
		Where("character_id = ?", id).
		Find(&languages).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("languages with character id: %d could not be found", id))
	}

	return languages, nil
}
