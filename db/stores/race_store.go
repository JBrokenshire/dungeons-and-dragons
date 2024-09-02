package stores

import (
	"dungeons-and-dragons/db/models"
	"github.com/jinzhu/gorm"
)

type RaceStore interface {
	GetAll() ([]*models.Race, error)
	Get(id int) (*models.Race, error)
}

type GormRaceStore struct {
	DB *gorm.DB
}

func NewGormRaceStore(db *gorm.DB) *GormRaceStore {
	return &GormRaceStore{
		DB: db,
	}
}

func (s *GormRaceStore) GetAll() ([]*models.Race, error) {
	var races []*models.Race
	if err := s.DB.Find(&races).Error; err != nil {
		return nil, err
	}

	return races, nil
}

func (s *GormRaceStore) Get(id int) (*models.Race, error) {
	var race models.Race
	if err := s.DB.Where("id = ?", id).First(&race).Error; err != nil {
		return nil, err
	}
	return &race, nil
}
