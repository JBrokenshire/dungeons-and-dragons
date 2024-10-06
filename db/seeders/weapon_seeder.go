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
			ItemID:     3,
			Type:       "ranged",
			ShortRange: 80,
			LongRange:  320,
			Damage:     "1d8",
			DamageType: "Piercing",
			Ability:    "DEX",
		},
		{
			ItemID:     17,
			Type:       "melee",
			ShortRange: 20,
			LongRange:  60,
			Damage:     "1d4",
			DamageType: "Piercing",
			Ability:    "DEX",
			Bonus:      1,
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
			ItemID:     24,
			Type:       "melee",
			ShortRange: 5,
			Damage:     "2d6",
			DamageType: "Slashing",
			Ability:    "STR",
		},
		{
			ItemID:     36,
			Type:       "melee",
			ShortRange: 20,
			LongRange:  60,
			Damage:     "1d6",
			DamageType: "Slashing",
			Ability:    "STR",
		},
		{
			ItemID:     18,
			Type:       "melee",
			ShortRange: 20,
			LongRange:  60,
			Damage:     "1d6",
			DamageType: "Slashing",
			Ability:    "STR",
			Bonus:      1,
		},
		{
			ItemID:     34,
			Type:       "ranged",
			ShortRange: 30,
			LongRange:  120,
			Damage:     "1d6",
			DamageType: "Piercing",
			Ability:    "STR",
		},
		{
			ItemID:     31,
			Type:       "melee",
			ShortRange: 5,
			Damage:     "1d8",
			DamageType: "Piercing",
			Ability:    "STR",
		},
		{
			ItemID:     27,
			Type:       "melee",
			ShortRange: 5,
			Damage:     "1d6",
			AltDamage:  "1d8",
			DamageType: "Bludgeoning",
			Ability:    "STR",
		},
		{
			ItemID:     19,
			Type:       "melee",
			ShortRange: 5,
			Damage:     "1d8",
			AltDamage:  "1d10",
			DamageType: "Bludgeoning",
			Ability:    "STR",
			Bonus:      1,
		},
	}

	for _, weapon := range weapons {
		err := s.DB.Where("item_id = ?", weapon.ItemID).FirstOrCreate(&weapon).Error
		if err != nil {
			log.Printf("error creating weapon with item_id '%v': %v", weapon.ItemID, err)
		}
	}
}
