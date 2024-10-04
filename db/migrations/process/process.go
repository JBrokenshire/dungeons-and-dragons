package process

import (
	"dnd-api/db/migrations/list"
	gm "github.com/ShkrutDenis/go-migrations"
	gmStore "github.com/ShkrutDenis/go-migrations/store"
)

func Run() {
	gm.Run(getMigrationsList())
}

func getMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		&list.CreateClassesTable{},
		&list.CreateRacesTable{},
		&list.CreateCharactersTable{},
		&list.CharacterAbilityScoreProficiencies{},
		&list.CreateCharacterProficientSkillsTable{},
		&list.CreateCharacterSensesTable{},
		&list.CreateCharacterProficientArmourTypes{},
		&list.CreateCharacterProficientWeapons{},
		&list.CreateCharacterProficientTools{},
		&list.CreateCharacterLanguages{},
		&list.UpdateCharactersInitiativeModifier{},
		&list.UpdateCharactersTableBaseArmourClass{},
		&list.CreateCharacterDefenses{},
		&list.CreateCharacterConditions{},
		&list.UpdateCharacterAttacksPerAction{},
		&list.CreateItemsTable{},
		&list.CreateWeaponsTable{},
	}
}
