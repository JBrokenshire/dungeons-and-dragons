package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateRacesTable struct{}

func (m *CreateRacesTable) GetName() string {
	return "CreateRacesTable"
}

func (m *CreateRacesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("races", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("short_description").Type("MEDIUMTEXT").Nullable()
	table.Column("long_description").Type("LONGTEXT").Nullable()
	table.Integer("base_walking_speed").NotNull()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateRacesTable) Down(con *sqlx.DB) {
	builder.DropTable("races", con).MustExec()
}
