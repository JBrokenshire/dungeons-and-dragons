package stores

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

type ItemStore interface {
	GetAll() ([]*models.Item, error)
}

type GormItemStore struct {
	DB *gorm.DB
}

func NewGormItemStore(db *gorm.DB) *GormItemStore {
	return &GormItemStore{
		DB: db,
	}
}

func (s *GormItemStore) GetAll() ([]*models.Item, error) {
	var items []*models.Item
	if err := s.DB.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
