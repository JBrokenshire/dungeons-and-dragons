package test

import (
	"dungeons-and-dragons/server/requests"
	"dungeons-and-dragons/test/helpers"
	"net/http"
	"testing"
)

func TestGetAllCharacters(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get list of characters (populated table)",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/characters"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Faelan Haversham"`,
					`"name":"PeeWee McAnkle-Biter"`,
				},
			},
		},
		{
			TestName: "can get list of characters (empty table)",
			Setup: func() {
				ts.ClearTable("characters")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/characters"},
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

func TestCreateCharacter(t *testing.T) {
	ts.ClearTable("characters")
	ts.SetupDefaultCharacters()

	request := helpers.Request{
		Method: http.MethodPost,
		URL:    "/characters",
	}

	cases := []helpers.TestCase{
		{
			TestName: "can create a new character from valid json in request body",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"id":3`, `"name":"Test"`, `"level":1`},
			},
		},
		{
			TestName: "post /character/:id 400 bad request on no request body",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid class id",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1,
				ClassID: 1000,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid race id",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1,
				ClassID: 1,
				RaceID:  1000,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "post /character/:id 400 internal server error on invalid level",
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "test",
				Level:   1000,
				ClassID: 1,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestGetCharacter(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get character by id",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":4`, `"race_id":18`},
			},
		},
		{
			TestName: "get /characters/:id returns 404 not found on character id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/characters/10",
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

func TestUpdateCharacter(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPut,
		URL:    "/characters/1",
	}

	cases := []helpers.TestCase{
		{
			TestName: "can update character by id with valid json in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test"`, `"level":1`, `"class_id":1`, `"race_id":1`},
			},
		},
		{
			TestName: "put /characters/:id returns 404 not found on character id not in database",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/10",
			},
			RequestBody: requests.CharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "put /characters/:id returns 400 bad request on empty request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character request body",
			},
		},
		{
			TestName:    "no update on empty character in request body",
			Setup:       ts.SetupDefaultCharacters,
			Request:     request,
			RequestBody: requests.CharacterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":4`, `"race_id":18`},
			},
		},
		{
			TestName: "update with only name in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name: "Test",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test"`, `"level":3`, `"class_id":4`, `"race_id":18`},
			},
		},
		{
			TestName: "update with only class id in request body",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				ClassID: 1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":3`, `"class_id":1`, `"race_id":18`},
			},
		},
		{
			TestName: "put /character/:id returns 400 bad request on invalid class id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/1",
			},
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1000,
				RaceID:  1,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character classID",
			},
		},
		{
			TestName: "put /character/:id returns 400 bad request on invalid race id",
			Setup:    ts.SetupDefaultCharacters,
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/1",
			},
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1000,
				Level:   1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character raceID",
			},
		},
		{
			TestName: "put /characters/:id returns 400 bad request on invalid level",
			Setup:    ts.SetupDefaultCharacters,
			Request:  request,
			RequestBody: requests.CharacterRequest{
				Name:    "Test",
				ClassID: 1,
				RaceID:  1,
				Level:   100,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "invalid character level",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}

func TestLevelUpCharacter(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can level up character by id",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/1/level-up",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Faelan Haversham"`, `"level":4`, `"class_id":4`, `"race_id":18`},
			},
		},
		{
			TestName: "put /characters/:id/level-up returns 404 not found on character id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/characters/10/level-up",
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

func TestDeleteCharacter(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can delete character by id",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodDelete,
				URL:    "/characters/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "character successfully deleted",
			},
		},
		{
			TestName: "delete /characters/:id returns 404 not found on character id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.SetupDefaultCharacters()
			},
			Request: helpers.Request{
				Method: http.MethodDelete,
				URL:    "/characters/10",
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
