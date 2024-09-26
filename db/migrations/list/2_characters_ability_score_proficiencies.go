package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CharacterAbilityScoreProficiencies struct{}

func (m *CharacterAbilityScoreProficiencies) GetName() string {
	return "CharacterAbilityScoreProficiencies"
}

func (m *CharacterAbilityScoreProficiencies) Up(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.Column("proficient_strength").Type("boolean").NotNull().Default("0")
	table.Column("proficient_dexterity").Type("boolean").NotNull().Default("0")
	table.Column("proficient_constitution").Type("boolean").NotNull().Default("0")
	table.Column("proficient_intelligence").Type("boolean").NotNull().Default("0")
	table.Column("proficient_wisdom").Type("boolean").NotNull().Default("0")
	table.Column("proficient_charisma").Type("boolean").NotNull().Default("0")
	table.MustExec()
}

func (m *CharacterAbilityScoreProficiencies) Down(con *sqlx.DB) {
	table := builder.ChangeTable("characters", con)
	table.DropColumn("proficient_strength")
	table.DropColumn("proficient_dexterity")
	table.DropColumn("proficient_constitution")
	table.DropColumn("proficient_intelligence")
	table.DropColumn("proficient_wisdom")
	table.DropColumn("proficient_charisma")
	table.MustExec()
}
