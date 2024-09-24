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
	}
}
