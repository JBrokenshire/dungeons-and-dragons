package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetWeapons() {
	weapons := []models.Weapon{
		{
			ItemID:     1,
			Type:       "melee",
			ShortRange: 20,
			LongRange:  60,
			Damage:     "1d4",
			DamageType: "Piercing",
			Ability:    "DEX",
		},
		{
			ItemID:     2,
			Type:       "melee",
			ShortRange: 5,
			Damage:     "1d12",
			DamageType: "Slashing",
			Ability:    "STR",
		},
		{
			ItemID:     3,
			Type:       "ranged",
			ShortRange: 80,
			LongRange:  320,
			Damage:     "1d8",
			DamageType: "Piercing",
			Ability:    "DEX",
		},
	}

	for _, weapon := range weapons {
		err := s.DB.Where("item_id = ?", weapon.ItemID).FirstOrCreate(&weapon).Error
		if err != nil {
			log.Printf("error creating weapon with item_id '%v': %v", weapon.ItemID, err)
		}
	}
}
