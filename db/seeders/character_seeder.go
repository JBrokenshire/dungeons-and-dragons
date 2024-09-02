package seeders

import (
	"dungeons-and-dragons/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []models.Character{
		{
			ID:      1,
			Name:    "Faelan Haversham",
			Level:   3,
			ClassID: 4,
			RaceID:  18,
		},
		{
			ID:      2,
			Name:    "PeeWee McAnkle-Biter",
			Level:   5,
			ClassID: 2,
			RaceID:  3,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("error creating character %s: %s", character.Name, err.Error())
		}
	}
}
