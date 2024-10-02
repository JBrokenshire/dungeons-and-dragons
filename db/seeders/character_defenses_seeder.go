package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharactersDefenses() {
	charactersDefenses := []models.CharacterDefense{
		{
			CharacterID: 1,
			DamageType:  "Poison",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Poison",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Acid",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Bludgeoning",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Cold",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Fire",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Force",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Lightning",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Necrotic",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Piercing",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Radiant",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Slashing",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 2,
			DamageType:  "Thunder",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 3,
			DamageType:  "Cold",
			DefenseType: "Resistance",
		},
		{
			CharacterID: 4,
			DamageType:  "Fire",
			DefenseType: "Resistance",
		},
	}

	for _, defense := range charactersDefenses {
		err := s.DB.Where("character_id = ? AND damage_type = ? AND defense_type = ?", defense.CharacterID, defense.DamageType, defense.DefenseType).FirstOrCreate(&defense).Error
		if err != nil {
			log.Printf("error creating character defense for CharacterID: %q, Damage Type: %s, Defense Type: %s -- %s", defense.CharacterID, defense.DamageType, defense.DefenseType, err.Error())
		}
	}
}
