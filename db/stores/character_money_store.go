package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterMoneyStore interface {
	GetMoneyByCharacterID(id interface{}) ([]*models.CharacterMoney, error)
}

type GormCharacterMoneyStore struct {
	DB *gorm.DB
}

func NewGormCharacterMoneyStore(db *gorm.DB) *GormCharacterMoneyStore {
	return &GormCharacterMoneyStore{DB: db}
}

func (g *GormCharacterMoneyStore) GetMoneyByCharacterID(id interface{}) ([]*models.CharacterMoney, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterMoney []*models.CharacterMoney
	err := g.DB.Table("character_money").Where("character_id = ?", id).Find(&characterMoney).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("character money with character id '%v' could not be found", id))
	}

	return characterMoney, nil
}
