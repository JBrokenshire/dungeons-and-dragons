package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllWeapons(t *testing.T) {
	ts.ClearTable("items")
	ts.ClearTable("weapons")

	itemOne := &models.Item{
		Name: "Greataxe",
	}
	factories.NewItem(ts.S.Db, itemOne)

	weaponOne := &models.Weapon{
		ItemID:     itemOne.ID,
		Type:       "Melee Weapon",
		ShortRange: 5,
		Damage:     "1d12",
		DamageType: "Slashing",
		Ability:    "STR",
	}
	factories.NewWeapon(ts.S.Db, weaponOne)

	itemTwo := &models.Item{
		Name: "Crossbow, Light",
	}
	factories.NewItem(ts.S.Db, itemTwo)

	weaponTwo := &models.Weapon{
		ItemID:     itemTwo.ID,
		Type:       "Ranged Weapon",
		ShortRange: 80,
		LongRange:  320,
		Damage:     "1d8",
		DamageType: "Piercing",
		Ability:    "DEX",
	}
	factories.NewWeapon(ts.S.Db, weaponTwo)

	request := helpers.Request{
		Method: http.MethodGet,
		URL:    "/items/weapons",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can get all weapons",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"type":"%s"`, weaponOne.Type),
					fmt.Sprintf(`"damage_type":"%s"`, weaponOne.DamageType),
					fmt.Sprintf(`"damage":"%s"`, weaponTwo.Damage),
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
					fmt.Sprintf(`"meta":"%s"`, itemOne.Meta),
					fmt.Sprintf(`"type":"%s"`, weaponTwo.Type),
					fmt.Sprintf(`"damage_type":"%s"`, weaponTwo.DamageType),
					fmt.Sprintf(`"damage":"%s"`, weaponTwo.Damage),
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
					fmt.Sprintf(`"meta":"%s"`, itemTwo.Meta),
				},
			},
		},
		{
			TestName: "Can get no items (empty table)",
			Setup: func() {
				ts.ClearTable("items")
				ts.ClearTable("weapons")
			},
			Request: request,
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
