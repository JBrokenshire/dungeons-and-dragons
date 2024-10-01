package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterProficientWeapon(db *gorm.DB, weapon *models.CharacterProficientWeapon) {
	fillProficientWeaponDetails(weapon)
	db.Create(weapon)
}

func fillProficientWeaponDetails(weapon *models.CharacterProficientWeapon) {
	if weapon.Weapon == "" {
		weapon.Weapon = "Martial Weapons"
	}
}
