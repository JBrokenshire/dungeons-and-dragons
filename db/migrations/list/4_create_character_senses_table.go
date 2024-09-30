package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterSensesTable struct{}

func (m *CreateCharacterSensesTable) GetName() string {
	return "CreateCharacterSensesTable"
}

func (m *CreateCharacterSensesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_senses", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("sense_name").Type("ENUM('Darkvision','Blindsight','Truesight')").NotNull()
	table.Integer("distance").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterSensesTable) Down(con *sqlx.DB) {
	builder.DropTable("character_senses", con).MustExec()
}
