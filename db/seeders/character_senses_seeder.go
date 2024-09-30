package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterSenses() {
	characterSenses := []models.CharacterSense{
		{
			CharacterID: 2,
			SenseName:   "Darkvision",
			Distance:    60,
		},
		{
			CharacterID: 4,
			SenseName:   "Blindsight",
			Distance:    10,
		},
	}

	for _, characterSense := range characterSenses {
		err := s.DB.Where("character_id = ? AND sense_name = ?", characterSense.CharacterID, characterSense.SenseName).FirstOrCreate(&characterSense).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", characterSense.CharacterID, characterSense.SenseName, err.Error())
		}
	}
}
