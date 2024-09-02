package stores

import (
	"dungeons-and-dragons/db/models"
	"github.com/jinzhu/gorm"
)

type ClassStore interface {
	GetAll() ([]*models.Class, error)
	Get(id int) (*models.Class, error)
	Update(class *models.Class) error
}

type GormClassStore struct {
	db *gorm.DB
}

func NewGormClassStore(db *gorm.DB) *GormClassStore {
	return &GormClassStore{db: db}
}

func (s *GormClassStore) GetAll() ([]*models.Class, error) {
	var classes []*models.Class
	if err := s.db.Find(&classes).Error; err != nil {
		return nil, err
	}

	return classes, nil
}

func (s *GormClassStore) Get(id int) (*models.Class, error) {
	var class models.Class
	if err := s.db.Where("id = ?", id).First(&class).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *GormClassStore) Update(class *models.Class) error {
	if err := s.db.Save(class).Error; err != nil {
		return err
	}

	return nil
}
