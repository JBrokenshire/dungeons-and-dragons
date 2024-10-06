package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateWeaponsAddBonus struct{}

func (m *UpdateWeaponsAddBonus) GetName() string {
	return "UpdateWeaponsAddBonus"
}

func (m *UpdateWeaponsAddBonus) Up(con *sqlx.DB) {
	table := builder.ChangeTable("weapons", con)
	table.Integer("bonus").Nullable()
	table.MustExec()
}

func (m *UpdateWeaponsAddBonus) Down(con *sqlx.DB) {
	table := builder.ChangeTable("weapons", con)
	table.DropColumn("bonus")
	table.MustExec()
}
