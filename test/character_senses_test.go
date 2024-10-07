package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterSenses(t *testing.T) {
	ts.ClearTable("character_proficient_skills")
	ts.ClearTable("characters")

	ts.SetupDefaultRaces()
	ts.SetupDefaultClasses()

	character := &models.Character{
		ID: 1,
	}
	factories.NewCharacter(ts.S.Db, character)

	noSenses := &models.Character{
		ID: 2,
	}
	factories.NewCharacter(ts.S.Db, noSenses)

	sense := &models.CharacterSense{
		CharacterID: character.ID,
		SenseName:   "Darkvision",
		Distance:    60,
	}
	factories.NewCharacterSense(ts.S.Db, sense)

	cases := []helpers.TestCase{
		{
			TestName: "Can get all senses for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/senses",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"character_id":%v`, character.ID),
					fmt.Sprintf(`"sense_name":"%s"`, sense.SenseName),
					fmt.Sprintf(`"distance":%v`, sense.Distance),
				},
			},
		},
		{
			TestName: "Empty response for character with no senses",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/senses",
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
				URL:    "/characters/invalid-id/senses",
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
