package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateCharactersTableBaseArmourClass struct{}

func (m *UpdateCharactersTableBaseArmourClass) GetName() string {
	return "UpdateCharactersTableBaseArmourClass"
}

func (m *UpdateCharactersTableBaseArmourClass) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Integer("base_armour_class").NotNull().Default("10")
	table.Column("armour_class_add_dexterity").Type("BOOLEAN").Default("1")
	table.MustExec()
}

func (m *UpdateCharactersTableBaseArmourClass) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("base_armour_class")
	table.DropColumn("armour_class_add_dexterity")
	table.MustExec()
}
