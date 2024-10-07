package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"net/http"
	"testing"
)

func TestGetCharacterConditions(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noDefenses := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noDefenses)

	otherDefenses := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherDefenses)

	characterCondition := &models.CharacterCondition{CharacterID: character.ID, ConditionName: "Blinded"}
	factories.NewCharacterCondition(ts.S.Db, characterCondition)

	otherCharacterCondition := &models.CharacterCondition{CharacterID: otherDefenses.ID, ConditionName: "Frightened"}
	factories.NewCharacterCondition(ts.S.Db, otherCharacterCondition)

	cases := []helpers.TestCase{
		{
			TestName: "Can get conditions for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/conditions",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"condition_name":"Blinded"`,
				},
				BodyPartsMissing: []string{
					`"condition_name":"Frightened"`,
				},
			},
		},
		{
			TestName: "Empty response for no conditions",
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
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/conditions",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
