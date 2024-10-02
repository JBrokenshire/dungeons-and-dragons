package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterDefenses struct{}

func (c *CreateCharacterDefenses) GetName() string {
	return "CreateCharacterDefenses"
}

func (c *CreateCharacterDefenses) Up(con *sqlx.DB) {
	table := builder.NewTable("character_defenses", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.String("damage_type", 255).NotNull()
	table.Column("defense_type").Type("ENUM('Resistance','Immunity','Vulnerability')").NotNull()
	table.MustExec()
}

func (c *CreateCharacterDefenses) Down(con *sqlx.DB) {
	builder.DropTable("character_defenses", con).MustExec()
}
