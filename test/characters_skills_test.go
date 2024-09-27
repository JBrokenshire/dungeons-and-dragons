package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterProficientSkills(t *testing.T) {
	ts.ClearTable("character_proficient_skills")
	ts.ClearTable("characters")

	ts.SetupDefaultRaces()
	ts.SetupDefaultClasses()

	character := &models.Character{
		ID: 1,
	}
	factories.NewCharacter(ts.S.Db, character)

	noProficiencies := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noProficiencies)

	proficientSkill := &models.CharacterProficientSkill{CharacterID: character.ID, SkillName: "Athletics", ProficiencyType: "Half-Proficiency"}
	factories.NewCharacterProficientSkill(ts.S.Db, proficientSkill)

	cases := []helpers.TestCase{
		{
			TestName: "Can get all proficient skills for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/proficient-skills",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"character_id":%v`, proficientSkill.CharacterID),
					fmt.Sprintf(`"skill_name":"%s"`, proficientSkill.SkillName),
					fmt.Sprintf(`"proficiency_type":"%s"`, proficientSkill.ProficiencyType),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"character_id":%v`, noProficiencies.ID),
				},
			},
		},
		{
			TestName: "empty response on invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient-skills",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
