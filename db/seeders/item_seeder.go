package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetItems() {
	items := []models.Item{
		{
			ID:     1,
			Name:   "Dagger",
			Meta:   "Melee Weapon",
			Weight: 1,
			Cost:   2,
			Notes:  "Simple, Finesse, Light, Thrown, Range(20/60)",
		},
		{
			ID:     2,
			Name:   "Greataxe",
			Meta:   "Melee Weapon",
			Weight: 7,
			Cost:   30,
			Notes:  "Martial, Heavy, Two-Handed",
		},
		{
			ID:     3,
			Name:   "Crossbow, Light",
			Meta:   "Ranged Weapon",
			Weight: 5,
			Cost:   25,
			Notes:  "Simple, Ammunition, Loading, Range, Two-Handed, Slow, Range(80/320)",
		},
		{
			ID:     4,
			Name:   "Studded Leather",
			Meta:   "Light Armour",
			Weight: 13,
			Cost:   45,
			Notes:  "AC 12",
		},
		{
			ID:    5,
			Name:  "Potion of Healing",
			Meta:  "Gear • Potion",
			Cost:  100,
			Notes: "Healing, Combat, Consumable",
		},
	}

	for _, item := range items {
		err := s.DB.Where("id = ?", item.ID).FirstOrCreate(&item).Error
		if err != nil {
			log.Printf("error creating item %s: %v", item.Name, err)
		}
	}
}
