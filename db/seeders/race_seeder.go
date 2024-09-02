package seeders

import (
	"dungeons-and-dragons/db/models"
	"log"
)

func (s *Seeder) SetRaces() {
	races := []models.Race{
		{ID: 1, Name: "Aarakocra"},
		{ID: 2, Name: "Dragonborn"},
		{ID: 3, Name: "Hill Dwarf"},
		{ID: 4, Name: "Moutain Dwarf"},
		{ID: 5, Name: "Eladrin Elf"},
		{ID: 6, Name: "High Elf"},
		{ID: 7, Name: "Wood Elf"},
		{ID: 8, Name: "Air Genasi"},
		{ID: 9, Name: "Earth Genasi"},
		{ID: 10, Name: "Fire Genasi"},
		{ID: 11, Name: "Water Genasi"},
		{ID: 12, Name: "Rock Gnome"},
		{ID: 13, Name: "Deep Gnome"},
		{ID: 14, Name: "Goliath"},
		{ID: 15, Name: "Half-Elf"},
		{ID: 16, Name: "Half-Orc"},
		{ID: 17, Name: "Lightfoot Halfling"},
		{ID: 18, Name: "Stout Halfling"},
		{ID: 19, Name: "Human"},
		{ID: 20, Name: "Variant Human"},
		{ID: 21, Name: "Tiefling"},
		{ID: 22, Name: "Variant Aasimar"},
	}

	for _, race := range races {
		err := s.DB.Where("id = ?", race.ID).FirstOrCreate(&race).Error
		if err != nil {
			log.Printf("error creating race %s: %v", race.Name, err.Error())
		}
	}
}
