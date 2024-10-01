package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersProficientTools() {
	charactersProficientTools := []models.CharacterProficientTool{
		{
			CharacterID: 1,
			Tool:        "Alchemist's Supplies",
		},
		{
			CharacterID: 1,
			Tool:        "Viol",
		},
		{
			CharacterID: 2,
			Tool:        "Playing Card Set",
		},
		{
			CharacterID: 2,
			Tool:        "Smith's Tools",
		},
		{
			CharacterID: 2,
			Tool:        "Vehicles (Land)",
		},
		{
			CharacterID: 3,
			Tool:        "Alchemist's Supplies",
		},
		{
			CharacterID: 3,
			Tool:        "Drum",
		},
		{
			CharacterID: 4,
			Tool:        "Lute",
		},
		{
			CharacterID: 4,
			Tool:        "Thieves' Tools",
		},
	}

	for _, tool := range charactersProficientTools {
		err := s.DB.Where("character_id = ? AND tool = ?", tool.CharacterID, tool.Tool).FirstOrCreate(&tool).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", tool.CharacterID, tool.Tool, err.Error())
		}
	}
}
