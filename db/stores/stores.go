package stores

import "github.com/jinzhu/gorm"

type Stores struct {
	Db        *gorm.DB
	Character *GormCharacterStore
	Class     *GormClassStore
	Race      *GormRaceStore
}

func NewStores(db *gorm.DB) *Stores {
	return &Stores{
		Db:        db,
		Character: NewGormCharacterStore(db),
		Class:     NewGormClassStore(db),
		Race:      NewGormRaceStore(db),
	}
}
