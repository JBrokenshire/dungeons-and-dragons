package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateCharacterProficientSkillsTable struct{}

func (m *CreateCharacterProficientSkillsTable) GetName() string {
	return "CreateCharacterProficientSkillsTable"
}

func (m *CreateCharacterProficientSkillsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("character_proficient_skills", con)
	table.PrimaryKey("id")
	table.Integer("character_id").NotNull()
	table.ForeignKey("character_id").Reference("characters").On("id").OnUpdate("cascade").OnDelete("cascade")
	table.Column("skill_name").Type("ENUM('Acrobatics','Animal Handling','Arcana','Athletics','Deception','History','Insight','Intimidation','Investigation','Medicine','Nature','Perception','Performance','Persuasion','Religion','Sleight of Hand','Stealth','Survival')").NotNull()
	table.Column("proficiency_type").Type("ENUM('Proficiency','Half-Proficiency','Expertise')").NotNull().Default("Proficiency")
	table.WithTimestamps()
	table.MustExec()
}

func (m *CreateCharacterProficientSkillsTable) Down(con *sqlx.DB) {
	builder.DropTable("character_proficient_skills", con).MustExec()
}
