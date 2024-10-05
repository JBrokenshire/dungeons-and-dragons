package stores

import (
	"dnd-api/db/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

type CharacterInventoryStore interface {
	GetInventoryByCharacterID(id interface{}) ([]*models.CharacterInventoryItem, error)
	GetEquippedWeaponsByCharacterID(id interface{}) ([]*models.Weapon, error)
}

type GormCharacterInventoryStore struct {
	DB *gorm.DB
}

func NewGormCharacterInventoryStore(db *gorm.DB) *GormCharacterInventoryStore {
	return &GormCharacterInventoryStore{
		DB: db,
	}
}

func (g *GormCharacterInventoryStore) GetInventoryByCharacterID(id interface{}) ([]*models.CharacterInventoryItem, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterInventory []*models.CharacterInventoryItem
	err := g.DB.
		Preload("Item").
		Where("character_id = ?", id).
		Find(&characterInventory).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inventory items with character id: %q could not be found", id))
	}

	return characterInventory, nil
}

func (g *GormCharacterInventoryStore) GetEquippedWeaponsByCharacterID(id interface{}) ([]*models.Weapon, error) {
	if reflect.TypeOf(id).Kind() != reflect.String && reflect.TypeOf(id).Kind() != reflect.Int {
		return nil, errors.New("id should be a string or int")
	}

	var characterEquippedItems []*models.CharacterInventoryItem
	err := g.DB.
		Table("character_inventory_items").
		Where("character_id = ? AND location = 'Equipment' AND equipped = true", id).
		Find(&characterEquippedItems).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("inventory items with character id: %q could not be found", id))
	}

	var equippedItemIDs []int
	for _, equippedItem := range characterEquippedItems {
		equippedItemIDs = append(equippedItemIDs, equippedItem.ItemID)
	}

	var weapons []*models.Weapon
	for _, itemID := range equippedItemIDs {
		var weapon models.Weapon
		_ = g.DB.
			Preload("Item").
			Where("item_id = ?", itemID).
			First(&weapon).Error
		if weapon.ItemID != 0 {
			weapons = append(weapons, &weapon)
		}
	}

	return weapons, nil
}
