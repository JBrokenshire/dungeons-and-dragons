package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
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

	otherArmourProficiencies := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherArmourProficiencies)

	proficientArmourType := &models.CharacterProficientArmourType{CharacterID: character.ID, ArmourType: "Light Armour"}
	factories.NewCharacterProficientArmourType(ts.S.Db, proficientArmourType)

	otherProficientArmourType := &models.CharacterProficientArmourType{CharacterID: otherArmourProficiencies.ID, ArmourType: "Heavy Armour"}
	factories.NewCharacterProficientArmourType(ts.S.Db, otherProficientArmourType)

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
					`"Light Armour"`,
				},
				BodyPartsMissing: []string{
					`"Heavy Armour"`,
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
