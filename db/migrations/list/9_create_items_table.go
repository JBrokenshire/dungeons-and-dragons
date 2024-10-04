package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateItemsTable struct{}

func (m *CreateItemsTable) GetName() string {
	return "CreateItemsTable"
}

func (m *CreateItemsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("items", con)
	table.PrimaryKey("id")
	table.String("name", 255).NotNull()
	table.Column("meta").Type("MEDIUMTEXT").NotNull()
	table.Column("weight").Type("FLOAT").Nullable()
	table.Column("cost").Type("FLOAT").Nullable()
	table.Column("notes").Type("MEDIUMTEXT").Nullable()
	table.Column("rarity").Type("ENUM('Common','Uncommon','Rare','Very Rare','Legendary')")
	table.MustExec()
}

func (m *CreateItemsTable) Down(con *sqlx.DB) {
	builder.DropTable("items", con).MustExec()
}
