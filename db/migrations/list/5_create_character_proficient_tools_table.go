package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterProficientTools struct{}

func (m *CreateCharacterProficientTools) GetName() string {
	return "CreateCharacterProficientTools"
}

func (m *CreateCharacterProficientTools) Up(con *sqlx.DB) {
	table := builder.NewTable("character_proficient_tools", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("tool").Type(`ENUM("Lute","Thieves' Tools","Playing Card Set","Smith's Tools","Vehicles (Land)","Alchemist's Supplies","Drum","Viol")`).NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterProficientTools) Down(con *sqlx.DB) {
	builder.DropTable("character_proficient_tools", con).MustExec()
}
