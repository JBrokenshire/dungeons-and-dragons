package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterMoney struct{}

func (m *CreateCharacterMoney) GetName() string {
	return "CreateCharacterMoney"
}

func (m *CreateCharacterMoney) Up(con *sqlx.DB) {
	table := builder.NewTable("character_money", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("money").Type("ENUM('platinum','gold','electrum','silver','copper')").NotNull()
	table.Integer("amount").NotNull()
	table.MustExec()
}

func (m *CreateCharacterMoney) Down(con *sqlx.DB) {
	builder.DropTable("character_money", con).MustExec()
}
