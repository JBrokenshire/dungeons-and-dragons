package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"net/http"
	"testing"
)

func TestGetCharacterDefenses(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noDefenses := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noDefenses)

	otherDefenses := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherDefenses)

	characterDefense := &models.CharacterDefense{CharacterID: character.ID, DamageType: "Cold"}
	factories.NewCharacterDefense(ts.S.Db, characterDefense)

	otherCharacterDefense := &models.CharacterDefense{CharacterID: otherDefenses.ID, DamageType: "Poison"}
	factories.NewCharacterDefense(ts.S.Db, otherCharacterDefense)

	cases := []helpers.TestCase{
		{
			TestName: "Can get defenses for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/defenses",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"damage_type":"Cold"`,
					`"defense_type":"Resistance"`,
				},
				BodyPartsMissing: []string{
					`"damage_type":"Poison"`,
				},
			},
		},
		{
			TestName: "Empty response for no defenses",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/defenses",
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
				URL:    "/characters/invalid-id/defenses",
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
