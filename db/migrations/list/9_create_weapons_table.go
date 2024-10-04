package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateWeaponsTable struct{}

func (m *CreateWeaponsTable) GetName() string {
	return "CreateWeaponsTable"
}

func (m *CreateWeaponsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("weapons", con)
	table.Integer("item_id").Unique().NotNull().NotAutoincrement()
	table.PrimaryKey("item_id")
	table.ForeignKey("item_id").Reference("items").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.String("type", 255).NotNull()
	table.Integer("short_range").NotNull()
	table.Integer("long_range").Nullable()
	table.String("damage", 8).NotNull()
	table.String("alt_damage", 8).Nullable()
	table.Column("damage_type").Type("ENUM('Acid','Bludgeoning','Cold','Fire','Force','Lightning','Necrotic','Piercing','Poison','Psychic','Radiant','Slashing','Thunder')").NotNull()
	table.Column("ability").Type("ENUM('STR','DEX','CON','INT','WIS','CHA')").NotNull()
	table.MustExec()
}

func (m *CreateWeaponsTable) Down(con *sqlx.DB) {
	builder.DropTable("weapons", con).MustExec()
}
