package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCharacterInventory(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("items")

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	noItemsCharacter := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, noItemsCharacter)

	itemOne := &models.Item{
		ID:   1,
		Name: "Test Item One",
	}
	factories.NewItem(ts.S.Db, itemOne)

	itemTwo := &models.Item{
		ID:   2,
		Name: "Test Item Two",
	}
	factories.NewItem(ts.S.Db, itemTwo)

	characterInventoryItemOne := &models.CharacterInventoryItem{
		CharacterID: characterOne.ID,
		ItemID:      itemOne.ID,
	}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItemOne)

	characterInventoryItemTwo := &models.CharacterInventoryItem{
		CharacterID: characterTwo.ID,
		ItemID:      itemTwo.ID,
	}
	factories.NewCharacterInventoryItem(ts.S.Db, characterInventoryItemTwo)

	cases := []helpers.TestCase{
		{
			TestName: "Can get all inventory items for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
				},
			},
		},
		{
			TestName: "Empty list for character with no inventory items",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/3/inventory",
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

func TestGetCharacterEquippedWeapons(t *testing.T) {
	ts.ClearTable("character_inventory_items")
	ts.ClearTable("characters")
	ts.ClearTable("items")

	characterOne := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, characterOne)

	characterTwo := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, characterTwo)

	characterThree := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, characterThree)

	itemOne := &models.Item{ID: 1, Name: "Test Item One"}
	factories.NewItem(ts.S.Db, itemOne)
	weaponOne := &models.Weapon{ItemID: 1}
	factories.NewWeapon(ts.S.Db, weaponOne)

	itemTwo := &models.Item{ID: 2, Name: "Test Item Two"}
	factories.NewItem(ts.S.Db, itemTwo)
	weaponTwo := &models.Weapon{ItemID: 2}
	factories.NewWeapon(ts.S.Db, weaponTwo)

	itemThree := &models.Item{ID: 3, Name: "Test Item Three"}
	factories.NewItem(ts.S.Db, itemThree)
	weaponThree := &models.Weapon{ItemID: 3}
	factories.NewWeapon(ts.S.Db, weaponThree)

	itemFour := &models.Item{ID: 4, Name: "Test Item Four"}
	factories.NewItem(ts.S.Db, itemFour)

	equippedWeapon := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemOne.ID, Equipped: true}
	factories.NewCharacterInventoryItem(ts.S.Db, equippedWeapon)

	unequippedWeapon := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemTwo.ID, Equipped: false}
	factories.NewCharacterInventoryItem(ts.S.Db, unequippedWeapon)

	notWeapon := &models.CharacterInventoryItem{CharacterID: characterOne.ID, ItemID: itemFour.ID}
	factories.NewCharacterInventoryItem(ts.S.Db, notWeapon)

	characterTwoEquippedWeapon := &models.CharacterInventoryItem{CharacterID: characterTwo.ID, ItemID: itemThree.ID, Equipped: true}
	factories.NewCharacterInventoryItem(ts.S.Db, characterTwoEquippedWeapon)

	cases := []helpers.TestCase{
		{
			TestName: "Can get equipped weapons for character - not unequipped, non-weapons or other characters",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/inventory/equipped-weapons",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
					fmt.Sprintf(`"item_id":%v`, itemOne.ID),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
					fmt.Sprintf(`"name":"%s"`, itemThree.Name),
					fmt.Sprintf(`"name":"%s"`, itemFour.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
