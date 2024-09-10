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
			ClassID: 3,
			RaceID:  18,
		},
		{
			ID:      2,
			Name:    "PeeWee McAnkle-Biter",
			Level:   5,
			ClassID: 1,
			RaceID:  3,
		},
		{
			ID:      3,
			Name:    "Zelphar Qinhice",
			Level:   6,
			ClassID: 8,
			RaceID:  18,
		},
		{
			ID:                4,
			Name:              "Iratham Veomakute",
			Level:             3,
			ClassID:           3,
			RaceID:            3,
			ProfilePictureURL: "https://www.dndbeyond.com/avatars/36645/678/1581111423-94761552.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("error creating character %s: %s", character.Name, err.Error())
		}
	}
}
