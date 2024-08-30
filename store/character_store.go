package store

import (
	"dungeons-and-dragons/db"
	"dungeons-and-dragons/models"
	"github.com/jinzhu/gorm"
)

type CharacterStore interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	Get(id int) (*models.Character, error)
	Update(character *models.Character) error
	LevelUp(character *models.Character) error
	Delete(id int) error
}

type GormCharacterStore struct {
	DB *gorm.DB
}

func NewGormCharacterStore() *GormCharacterStore {
	return &GormCharacterStore{
		DB: db.DB(),
	}
}

func (g *GormCharacterStore) Create(character *models.Character) error {
	return g.DB.Create(character).Error
}

func (g *GormCharacterStore) GetAll() ([]*models.Character, error) {
	var characters []*models.Character
	if err := g.DB.Preload("Class").Preload("Race").Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (g *GormCharacterStore) Get(id int) (*models.Character, error) {
	var character models.Character
	if err := g.DB.
		Preload("Class").
		Preload("Race").
		Where("characters.id = ?", id).
		First(&character).Error; err != nil {
		return nil, err
	}

	return &character, nil
}

func (g *GormCharacterStore) Update(character *models.Character) error {
	return g.DB.Preload("Class").Preload("Race").Save(character).Error
}

func (g *GormCharacterStore) LevelUp(character *models.Character) error {
	character.Level++
	return g.DB.Preload("Class").Preload("Race").Save(character).Error
}

func (g *GormCharacterStore) Delete(id int) error {
	return g.DB.Delete(&models.Character{}, "id = ?", id).Error
}
