package test

import (
	"dungeons-and-dragons/models"
	"dungeons-and-dragons/test/helpers"
	"net/http"
	"testing"
)

func TestCharacterList(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "Can get list of characters",
			Request:  helpers.Request{Method: http.MethodGet, URL: "/characters"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Faelan Haversham"`,
					`"name":"PeeWee McAnkle-Biter"`,
				},
			},
		},
		{
			TestName: "Can create new character",
			Request:  helpers.Request{Method: http.MethodPost, URL: "/characters"},
			RequestBody: models.Character{
				ID:      3,
				Name:    "Test Character",
				ClassID: 1,
				RaceID:  1,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts:  []string{`"name":"Test Character"`, `"class_id":1`, `"race_id":1`},
			},
		},
		{
			TestName: "Can get a character by id",
			Request: helpers.Request{
				Method:    http.MethodGet,
				URL:       "/characters/3",
				PathParam: &helpers.PathParam{Name: "id", Value: "3"},
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"name":"Test Character"`, `"class_id":1`, `"race_id":1`},
			},
		},
		{
			TestName: "Can get update a character by id",
			Request: helpers.Request{
				Method:    http.MethodPut,
				URL:       "/characters/3",
				PathParam: &helpers.PathParam{Name: "id", Value: "3"},
			},
			RequestBody: models.Character{
				Name:    "New Test Character",
				ClassID: 2,
				RaceID:  2,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"name":"New Test Character"`, `"class_id":2`, `"race_id":2`},
			},
		},
		{
			TestName: "Can delete a character by id",
			Request: helpers.Request{
				Method:    http.MethodDelete,
				URL:       "/characters/3",
				PathParam: &helpers.PathParam{Name: "id", Value: "3"},
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   "character successfully deleted",
			},
		},
		{
			TestName: "Bad request on invalid id (NaN)",
			Request: helpers.Request{
				Method:    http.MethodGet,
				URL:       "/characters/invalid-id",
				PathParam: &helpers.PathParam{Name: "id", Value: "invalid-id"},
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "404 on character not found",
			Request: helpers.Request{
				Method:    http.MethodGet,
				URL:       "/characters/0",
				PathParam: &helpers.PathParam{Name: "id", Value: "0"},
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
