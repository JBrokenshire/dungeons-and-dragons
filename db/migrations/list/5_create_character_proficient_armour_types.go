package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterProficientArmourTypes struct{}

func (m *CreateCharacterProficientArmourTypes) GetName() string {
	return "CreateCharacterProficientArmourTypes"
}

func (m *CreateCharacterProficientArmourTypes) Up(con *sqlx.DB) {
	table := builder.NewTable("character_proficient_armour_types", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("armour_type").Type("ENUM('Light Armour','Medium Armour','Heavy Armour','Shields')").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterProficientArmourTypes) Down(con *sqlx.DB) {
	builder.DropTable("character_proficient_armour_types", con).MustExec()
}
