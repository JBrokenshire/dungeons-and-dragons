package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterSkillsStore interface {
	GetProficientByCharacterID(id interface{}) ([]*models.CharacterProficientSkill, error)
}

type GormCharacterSkillsStore struct {
	DB *gorm.DB
}

func NewGormCharacterSkillsStore(db *gorm.DB) *GormCharacterSkillsStore {
	return &GormCharacterSkillsStore{
		DB: db,
	}
}

func (g *GormCharacterSkillsStore) GetProficientByCharacterID(id interface{}) ([]*models.CharacterProficientSkill, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterProficientSkills []*models.CharacterProficientSkill
	if err := g.DB.Where("character_id = ?", id).Find(&characterProficientSkills).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character proficient skills with character id: %q could not be found", id))
	}

	return characterProficientSkills, nil
}
