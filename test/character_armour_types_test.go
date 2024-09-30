package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterProficientArmourTypes(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noArmourProficiencies := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noArmourProficiencies)

	proficientArmourType := &models.CharacterProficientArmourType{CharacterID: character.ID, ArmourType: "Light Armour"}
	factories.NewCharacterProficientArmourType(ts.S.Db, proficientArmourType)

	cases := []helpers.TestCase{
		{
			TestName: "Can get armour proficiencies for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/proficient-armour",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"character_id":%v`, character.ID),
					`"armour_type":"Light Armour"`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"character_id":%v`, noArmourProficiencies.ID),
				},
			},
		},
		{
			TestName: "Empty response for no armour proficiencies",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/proficient-armour",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "[]",
			},
		},
		{
			TestName: "Empty response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient-armour",
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
