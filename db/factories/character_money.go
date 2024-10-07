package factories

import (
	"dnd-api/db/models"
	"github.com/jinzhu/gorm"
)

func NewCharacterMoney(db *gorm.DB, money *models.CharacterMoney) {
	fillCharacterMoneyDetails(money)
	db.Table("character_money").Create(money)
}

func fillCharacterMoneyDetails(characterMoney *models.CharacterMoney) {
	if characterMoney.Amount == 0 {
		characterMoney.Amount = 1
	}
	if characterMoney.Money == "" {
		characterMoney.Money = "gold"
	}
}
