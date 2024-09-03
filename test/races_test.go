package test

import (
	"dungeons-and-dragons/test/helpers"
	"net/http"
	"testing"
)

func TestGetAllRaces(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get list of races (populated table)",
			Setup: func() {
				ts.ClearTable("characters") // Have to clear characters first because of foreign key constraint
				ts.ClearTable("races")
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/races"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Aarakocra"`,
					`"name":"Stout Halfling"`,
					`"name":"Variant Aasimar"`,
				},
			},
		},
		{
			TestName: "can get list of races (empty table)",
			Setup: func() {
				ts.ClearTable("characters") // Have to clear characters first because of foreign key constraint
				ts.ClearTable("races")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/races"},
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

func TestGetRace(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get race by id",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("races")
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/races/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Aarakocra"`},
			},
		},
		{
			TestName: "get /races/:id returns 404 not found on race id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("races")
				ts.SetupDefaultRaces()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/races/100",
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
