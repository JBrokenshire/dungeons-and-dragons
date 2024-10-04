package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterInventory struct{}

func (c *CreateCharacterInventory) GetName() string {
	return "CreateCharacterInventory"
}

func (c *CreateCharacterInventory) Up(con *sqlx.DB) {
	table := builder.NewTable("character_inventory_items", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Integer("item_id").NotNull()
	table.ForeignKey("item_id").Reference("items").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("equipped").Type("BOOLEAN").Nullable()
	table.Integer("quantity").Nullable()
	table.Column("location").Type("ENUM('Equipment','Backpack')").NotNull()
	table.MustExec()
}

func (c *CreateCharacterInventory) Down(con *sqlx.DB) {
	builder.DropTable("character_inventory", con).MustExec()
}
