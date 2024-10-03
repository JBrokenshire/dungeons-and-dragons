package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharacterAttacksPerAction struct{}

func (m *UpdateCharacterAttacksPerAction) GetName() string {
	return "UpdateCharacterAttacksPerAction"
}

func (m *UpdateCharacterAttacksPerAction) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("attacks_per_action").NotNull().Default("1")
	table.MustExec()
}

func (m *UpdateCharacterAttacksPerAction) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("attacks_per_action")
	table.MustExec()
}
