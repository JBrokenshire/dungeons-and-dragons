package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateClassesTable struct{}

func (m *CreateClassesTable) GetName() string {
	return "CreateClassesTable"
}

func (m *CreateClassesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("classes", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("short_description").Type("MEDIUMTEXT").Nullable()
	table.Column("long_description").Type("LONGTEXT").Nullable()
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateClassesTable) Down(con *sqlx.DB) {
	builder.DropTable("classes", con).MustExec()
}
