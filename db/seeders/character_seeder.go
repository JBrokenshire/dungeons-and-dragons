package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacters() {
	characters := []models.Character{
		{
			ID:           1,
			Name:         "Faelan Haversham",
			Level:        3,
			ClassID:      3,
			RaceID:       18,
			Strength:     12,
			Dexterity:    16,
			Constitution: 14,
			Intelligence: 10,
			Wisdom:       15,
			Charisma:     8,
		},
		{
			ID:                   2,
			Name:                 "PeeWee McAnkle-Biter",
			Level:                5,
			ClassID:              1,
			RaceID:               3,
			Strength:             16,
			Dexterity:            14,
			Constitution:         16,
			Intelligence:         12,
			Wisdom:               11,
			Charisma:             8,
			WalkingSpeedModifier: 10,
		},
		{
			ID:           3,
			Name:         "Zelphar Qinhice",
			Level:        6,
			ClassID:      8,
			RaceID:       6,
			Strength:     14,
			Dexterity:    10,
			Constitution: 10,
			Intelligence: 14,
			Wisdom:       12,
			Charisma:     17,
		},
		{
			ID:                4,
			Name:              "Iratham Veomakute",
			Level:             3,
			ClassID:           3,
			RaceID:            14,
			ProfilePictureURL: "https://www.dndbeyond.com/avatars/36645/678/1581111423-94761552.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
			Strength:          18,
			Dexterity:         13,
			Constitution:      14,
			Intelligence:      15,
			Wisdom:            11,
			Charisma:          11,
		},
		{
			ID:                5,
			Name:              "Kael Drakeshield",
			Level:             4,
			ClassID:           6,
			RaceID:            2,
			ProfilePictureURL: "https://www.dndbeyond.com/avatars/42817/335/1581111423-124892113.jpeg?width=150&height=150&fit=crop&quality=95&auto=webp",
			Strength:          19,
			Dexterity:         15,
			Constitution:      19,
			Intelligence:      6,
			Wisdom:            14,
			Charisma:          11,
		},
	}

	for _, character := range characters {
		err := s.DB.Where("id = ?", character.ID).FirstOrCreate(&character).Error
		if err != nil {
			log.Printf("error creating character %s: %s", character.Name, err.Error())
		}
	}
}
