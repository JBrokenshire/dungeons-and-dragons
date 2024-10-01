package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterProficientWeapons struct{}

func (m *CreateCharacterProficientWeapons) GetName() string {
	return "CreateCharacterProficientWeapons"
}

func (m *CreateCharacterProficientWeapons) Up(con *sqlx.DB) {
	table := builder.NewTable("character_proficient_weapons", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("weapon").Type("ENUM('Martial Weapons','Simple Weapons')").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterProficientWeapons) Down(con *sqlx.DB) {
	builder.DropTable("character_proficient_weapons", con).MustExec()
}
