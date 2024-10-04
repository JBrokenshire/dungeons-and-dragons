package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersInventory() {
	characterInventoryItems := []models.CharacterInventoryItem{
		{
			ID:          1,
			CharacterID: 1,
			ItemID:      1,
			Equipped:    true,
			Location:    "Equipment",
		},
		{
			ID:          2,
			CharacterID: 1,
			ItemID:      1,
			Location:    "Backpack",
		},
		{
			ID:          3,
			CharacterID: 1,
			ItemID:      3,
			Equipped:    false,
			Location:    "Equipment",
		},
		{
			ID:          4,
			CharacterID: 1,
			ItemID:      4,
			Equipped:    true,
			Location:    "Equipment",
		},
		{
			ID:          5,
			CharacterID: 1,
			ItemID:      5,
			Location:    "Equipment",
		},
		{
			ID:          6,
			CharacterID: 2,
			ItemID:      2,
			Equipped:    true,
			Location:    "Equipment",
		},
	}

	for _, item := range characterInventoryItems {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("error creating character inventory item with id: %v -- %s", item.ID, err.Error())
		}
	}
}
