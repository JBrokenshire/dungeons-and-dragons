package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewWeapon(db *gorm.DB, weapon *models.Weapon) {
	fillWeaponDetails(weapon)
	db.Create(weapon)
}

func fillWeaponDetails(weapon *models.Weapon) {
	if weapon.Type == "" {
		weapon.Type = "Melee Weapon"
	}
	if weapon.Ability == "" {
		weapon.Ability = "STR"
	}
	if weapon.Damage == "" {
		weapon.Damage = "1d4"
	}
	if weapon.DamageType == "" {
		weapon.DamageType = "Piercing"
	}
	if weapon.ShortRange == 0 {
		weapon.ShortRange = 5
	}
}
