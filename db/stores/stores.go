package stores

import "github.com/jinzhu/gorm"

type Stores struct {
	Db                       *gorm.DB
	Character                *GormCharacterStore
	Class                    *GormClassStore
	Race                     *GormRaceStore
	CharacterSkills          *GormCharacterSkillsStore
	CharacterSenses          *GormCharacterSensesStore
	CharacterProficiencies   *GormCharacterProficienciesStore
	CharacterDefensesStore   *GormCharacterDefensesStore
	CharacterConditionsStore *GormCharacterConditionsStore
}

func NewStores(db *gorm.DB) *Stores {
	return &Stores{
		Db:                       db,
		Character:                NewGormCharacterStore(db),
		Class:                    NewGormClassStore(db),
		Race:                     NewGormRaceStore(db),
		CharacterSkills:          NewGormCharacterSkillsStore(db),
		CharacterSenses:          NewGormCharacterSensesStore(db),
		CharacterProficiencies:   NewGormCharacterProficienciesStore(db),
		CharacterDefensesStore:   NewGormCharacterDefensesStore(db),
		CharacterConditionsStore: NewGormCharacterConditionsStore(db),
	}
}
