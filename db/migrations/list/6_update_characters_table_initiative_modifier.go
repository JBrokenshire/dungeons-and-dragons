package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersInitiativeModifier struct{}

func (m *UpdateCharactersInitiativeModifier) GetName() string {
	return "UpdateCharactersInitiativeModifier"
}

func (m *UpdateCharactersInitiativeModifier) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("initiative_modifier").NotNull().Default("0")
	table.MustExec()
}

func (m *UpdateCharactersInitiativeModifier) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("initiative_modifier")
	table.MustExec()
}
