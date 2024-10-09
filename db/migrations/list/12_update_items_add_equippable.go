package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type UpdateItemsAddEquippable struct{}

func (m *UpdateItemsAddEquippable) GetName() string {
	return "UpdateItemsAddEquippable"
}

func (m *UpdateItemsAddEquippable) Up(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.Column("equippable").Type("BOOLEAN").NotNull()
	table.MustExec()
}

func (m *UpdateItemsAddEquippable) Down(con *sqlx.DB) {
	table := builder.ChangeTable("items", con)
	table.DropColumn("equippable")
	table.MustExec()
}
