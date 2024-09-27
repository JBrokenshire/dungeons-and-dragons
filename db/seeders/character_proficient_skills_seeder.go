package seeders

import (
	"dnd-api/db/models"
	"log"
)

func (s *Seeder) SetCharacterProficientSkills() {
	charactersProficientSkills := []models.CharacterProficientSkill{
		{
			CharacterID: 1,
			SkillName:   "Arcana",
		},
		{
			CharacterID: 1,
			SkillName:   "Athletics",
		},
		{
			CharacterID: 1,
			SkillName:   "History",
		},
		{
			CharacterID: 1,
			SkillName:   "Insight",
		},
		{
			CharacterID: 1,
			SkillName:   "Survival",
		},
		{
			CharacterID: 2,
			SkillName:   "Athletics",
		},
		{
			CharacterID: 2,
			SkillName:   "Intimidation",
		},

		{
			CharacterID: 2,
			SkillName:   "Perception",
		},

		{
			CharacterID: 2,
			SkillName:   "Survival",
		},
		{
			CharacterID: 3,
			SkillName:   "Athletics",
		},
		{
			CharacterID: 3,
			SkillName:   "Insight",
		},
		{
			CharacterID: 3,
			SkillName:   "Intimidation",
		},
		{
			CharacterID: 3,
			SkillName:   "Investigation",
		},
		{
			CharacterID: 3,
			SkillName:   "Perception",
		},
		{
			CharacterID: 3,
			SkillName:   "Survival",
		},
		{
			CharacterID: 4,
			SkillName:   "Insight",
		},
		{
			CharacterID: 4,
			SkillName:   "Intimidation",
		},
		{
			CharacterID: 4,
			SkillName:   "Perception",
		},
		{
			CharacterID: 4,
			SkillName:   "Stealth",
		},
	}

	for _, characterProficientSkill := range charactersProficientSkills {
		err := s.DB.Where("id = ?", characterProficientSkill.ID).FirstOrCreate(&characterProficientSkill).Error
		if err != nil {
			log.Printf("error creating character proficient skill for CharacterID: %q, Skill: %s: %s", characterProficientSkill.CharacterID, characterProficientSkill.SkillName, err.Error())
		}
	}
}
