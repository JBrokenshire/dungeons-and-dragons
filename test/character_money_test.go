package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterMoney(t *testing.T) {
	ts.ClearTable("character_money")
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	noMoneyCharacter := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, noMoneyCharacter)

	characterOneMoney := &models.CharacterMoney{CharacterID: characterOne.ID, Money: "gold", Amount: 10}
	factories.NewCharacterMoney(ts.S.Db, characterOneMoney)

	characterTwoMoney := &models.CharacterMoney{CharacterID: characterTwo.ID, Money: "silver", Amount: 5}
	factories.NewCharacterMoney(ts.S.Db, characterTwoMoney)

	cases := []helpers.TestCase{
		{
			TestName: "Can get money for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory/money",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"money":"%s"`, characterOneMoney.Money),
					fmt.Sprintf(`"amount":%v`, characterOneMoney.Amount),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"money":"%s"`, characterTwoMoney.Money),
					fmt.Sprintf(`"amount":%v`, characterTwoMoney.Amount),
				},
			},
		},
		{
			TestName: "Empty response for character with no money",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/3/inventory/money",
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
