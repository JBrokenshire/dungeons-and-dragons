package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharactersTable struct{}

func (m *CreateCharactersTable) GetName() string {
	return "CreateCharactersTable"
}

func (m *CreateCharactersTable) Up(con *sqlx.DB) {
	table := builder.NewTable("characters", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Integer("level").NotNull().Default("1")
	table.Integer("class_id").NotNull()
	table.ForeignKey("class_id").Reference("classes").On("id").OnDelete("cascade").OnUpdate("cascade").SetKeyName("fk__character_class")
	table.Integer("race_id").NotNull()
	table.ForeignKey("race_id").Reference("races").On("id").OnDelete("cascade").OnUpdate("cascade").SetKeyName("fk__character_race")
	table.Column("profile_picture_url").Type("MEDIUMTEXT").Nullable()
	table.Integer("strength").NotNull().Default("10")
	table.Integer("dexterity").NotNull().Default("10")
	table.Integer("constitution").NotNull().Default("10")
	table.Integer("intelligence").NotNull().Default("10")
	table.Integer("wisdom").NotNull().Default("10")
	table.Integer("charisma").NotNull().Default("10")
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharactersTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
