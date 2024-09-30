package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterProficientArmourTypes() {
	characterProficientArmourTypes := []models.CharacterProficientArmourType{
		{
			CharacterID: 1,
			ArmourType:  "Light Armour",
		},
		{
			CharacterID: 1,
			ArmourType:  "Medium Armour",
		},
		{
			CharacterID: 1,
			ArmourType:  "Shields",
		},
		{
			CharacterID: 2,
			ArmourType:  "Light Armour",
		},
		{
			CharacterID: 2,
			ArmourType:  "Medium Armour",
		},
		{
			CharacterID: 2,
			ArmourType:  "Shields",
		},
		{
			CharacterID: 3,
			ArmourType:  "Light Armour",
		},
		{
			CharacterID: 3,
			ArmourType:  "Medium Armour",
		},
		{
			CharacterID: 3,
			ArmourType:  "Shields",
		},
		{
			CharacterID: 4,
			ArmourType:  "Light Armour",
		},
		{
			CharacterID: 4,
			ArmourType:  "Medium Armour",
		},
		{
			CharacterID: 4,
			ArmourType:  "Heavy Armour",
		},
		{
			CharacterID: 4,
			ArmourType:  "Shields",
		},
	}

	for _, proficientArmourType := range characterProficientArmourTypes {
		err := s.DB.Where("character_id = ? AND armour_type = ?", proficientArmourType.CharacterID, proficientArmourType.ArmourType).FirstOrCreate(&proficientArmourType).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", proficientArmourType.CharacterID, proficientArmourType.ArmourType, err.Error())
		}
	}
}
