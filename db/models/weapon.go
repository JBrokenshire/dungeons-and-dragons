package models

import (
	"dnd-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

var validDamageTypes = []string{
	"Acid",
	"Bludgeoning",
	"Cold",
	"Fire",
	"Force",
	"Lightning",
	"Necrotic",
	"Piercing",
	"Poison",
	"Psychic",
	"Radiant",
	"Slashing",
	"Thunder",
}

var validAbilities = []string{
	"STR",
	"DEX",
	"CON",
	"INT",
	"WIS",
	"CHA",
}

type Weapon struct {
	ItemID     int    `gorm:"primary_key" json:"item_id"`
	Type       string `json:"type"`
	ShortRange int    `json:"short_range"`
	LongRange  int    `json:"long_range"`
	Damage     string `json:"damage"`
	AltDamage  string `json:"alt_damage"`
	DamageType string `json:"damage_type"`
	Ability    string `json:"ability"`

	Item Item `json:"item"`
}

func (w *Weapon) BeforeCreate(db *gorm.DB) error {
	var item Item
	err := db.Where("id = ?", w.ItemID).Find(&item).Error
	if err != nil {
		return fmt.Errorf("item with id %v not found", w.ItemID)
	}

	if !utils.SliceContains(validDamageTypes, w.DamageType) {
		return fmt.Errorf("damage type '%s' is not valid", w.DamageType)
	}

	if !utils.SliceContains(validAbilities, w.Ability) {
		return fmt.Errorf("ability '%s' is not valid", w.Ability)
	}

	return nil
}
