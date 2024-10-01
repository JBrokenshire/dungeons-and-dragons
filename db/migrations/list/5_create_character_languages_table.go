package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterLanguages struct{}

func (m *CreateCharacterLanguages) GetName() string {
	return "CreateCharacterLanguages"
}

func (m *CreateCharacterLanguages) Up(con *sqlx.DB) {
	table := builder.NewTable("character_languages", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("language").Type("ENUM('Abyssal','Celestial','Deep Speech','Draconic','Dwarvish','Elvish','Giant','Gnomish','Goblin','Halfling','Infernal','Orc','Primordial','Sylvan','Undercommon')").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterLanguages) Down(con *sqlx.DB) {
	builder.DropTable("character_languages", con).MustExec()
}
