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
				URL:    "/characters/1/proficient/armour",
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
			TestName: "Null response for no armour proficiencies",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/proficient/armour",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "null",
			},
		},
		{
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient/armour",
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

func TestGetCharacterProficientWeapons(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noWeaponProficiencies := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noWeaponProficiencies)

	otherWeaponProficiencies := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherWeaponProficiencies)

	proficientWeapon := &models.CharacterProficientWeapon{CharacterID: character.ID, Weapon: "Martial Weapons"}
	factories.NewCharacterProficientWeapon(ts.S.Db, proficientWeapon)

	otherProficientWeapon := &models.CharacterProficientWeapon{CharacterID: otherWeaponProficiencies.ID, Weapon: "Simple Weapons"}
	factories.NewCharacterProficientWeapon(ts.S.Db, otherProficientWeapon)

	cases := []helpers.TestCase{
		{
			TestName: "Can get weapon proficiencies for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/proficient/weapons",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"Martial Weapons"`,
				},
				BodyPartsMissing: []string{
					`"Simple Weapons"`,
				},
			},
		},
		{
			TestName: "Null response for no weapon proficiencies",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/proficient/weapons",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "null",
			},
		},
		{
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient/armour",
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

func TestGetCharacterProficientTools(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noToolProficiencies := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noToolProficiencies)

	otherToolProficiencies := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherToolProficiencies)

	proficientTool := &models.CharacterProficientTool{CharacterID: character.ID, Tool: "Lute"}
	factories.NewCharacterProficientTool(ts.S.Db, proficientTool)

	otherProficientTool := &models.CharacterProficientTool{CharacterID: otherToolProficiencies.ID, Tool: "Drum"}
	factories.NewCharacterProficientTool(ts.S.Db, otherProficientTool)

	cases := []helpers.TestCase{
		{
			TestName: "Can get tool proficiencies for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/proficient/tools",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"Lute"`,
				},
				BodyPartsMissing: []string{
					`"Drum"`,
				},
			},
		},
		{
			TestName: "Null response for no tool proficiencies",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/proficient/tools",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "null",
			},
		},
		{
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient/armour",
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

func TestGetCharacterLanguages(t *testing.T) {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	character := &models.Character{ID: 1}
	factories.NewCharacter(ts.S.Db, character)

	noLanguages := &models.Character{ID: 2}
	factories.NewCharacter(ts.S.Db, noLanguages)

	otherLanguages := &models.Character{ID: 3}
	factories.NewCharacter(ts.S.Db, otherLanguages)

	language := &models.CharacterLanguage{CharacterID: character.ID, Language: "Dwarvish"}
	factories.NewCharacterLanguage(ts.S.Db, language)

	otherLanguage := &models.CharacterLanguage{CharacterID: otherLanguages.ID, Language: "Elvish"}
	factories.NewCharacterLanguage(ts.S.Db, otherLanguage)

	cases := []helpers.TestCase{
		{
			TestName: "Can get languages for character",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1/proficient/languages",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					"Common",
					"Dwarvish",
				},
				BodyPartsMissing: []string{
					"Elvish",
				},
			},
		},
		{
			TestName: "Only common for no languages",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/2/proficient/languages",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "Common",
			},
		},
		{
			TestName: "404 response for invalid character id",
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/invalid-id/proficient/armour",
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
