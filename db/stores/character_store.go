package stores

import (
	"dungeons-and-dragons/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterStore interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	Get(id interface{}) (*models.Character, error)
	Update(character *models.Character) error
	Delete(id interface{}) error
}

type GormCharacterStore struct {
	DB *gorm.DB
}

func NewGormCharacterStore(db *gorm.DB) *GormCharacterStore {
	return &GormCharacterStore{
		DB: db,
	}
}

func (g *GormCharacterStore) Create(character *models.Character) error {
	return g.DB.Create(character).Error
}

func (g *GormCharacterStore) GetAll() ([]*models.Character, error) {
	var characters []*models.Character
	if err := g.DB.
		Preload("Class").
		Preload("Race").
		Find(&characters).
		Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (g *GormCharacterStore) Get(id interface{}) (*models.Character, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var character models.Character
	if err := g.DB.
		Where("characters.id = ?", id).
		First(&character).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("character with id %q not found", id))
	}

	var class models.Class
	if err := g.DB.
		Where("classes.id = ?", character.ClassID).
		First(&class).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("class with id %q not found", id))
	}
	character.Class = &class

	var race models.Race
	if err := g.DB.
		Where("races.id = ?", character.RaceID).
		First(&race).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("race with id %q not found", id))
	}
	character.Race = &race

	return &character, nil
}

func (g *GormCharacterStore) Update(character *models.Character) error {
	return g.DB.Save(character).Error
}

func (g *GormCharacterStore) Delete(id interface{}) error {
	_, err := g.Get(id)
	if err != nil {
		return err
	}
	return g.DB.Delete(&models.Character{}, "id = ?", id).Error
}
