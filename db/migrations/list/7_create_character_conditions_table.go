package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterConditions struct{}

func (c *CreateCharacterConditions) GetName() string {
	return "CreateCharacterConditions"
}

func (c *CreateCharacterConditions) Up(con *sqlx.DB) {
	table := builder.NewTable("character_conditions", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("condition_name").Type("ENUM('Blinded','Charmed','Deafened','Frightened','Grappled','Incapacitated','Invisible','Paralysed','Petrified','Poisoned','Prone','Restrained','Stunned','Unconscious')").NotNull()
	table.MustExec()
}

func (c *CreateCharacterConditions) Down(con *sqlx.DB) {
	builder.DropTable("character_conditions", con).MustExec()
}
