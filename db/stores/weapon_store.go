package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type WeaponStore interface {
	GetAll() ([]*models.Weapon, error)
}

type GormWeaponStore struct {
	DB *gorm.DB
}

func NewGormWeaponsStore(db *gorm.DB) *GormWeaponStore {
	return &GormWeaponStore{
		DB: db,
	}
}

func (s *GormWeaponStore) GetAll() ([]*models.Weapon, error) {
	var weapons []*models.Weapon
	if err := s.DB.
		Preload("Item").
		Find(&weapons).Error; err != nil {
		return nil, err
	}

	return weapons, nil
}
