package test

import (
	"dungeons-and-dragons/test/helpers"
	"net/http"
	"testing"
)

func TestGetAllClasses(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get list of classes (populated table)",
			Setup: func() {
				ts.ClearTable("characters") // Have to clear characters first because of foreign key constraint
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/classes"},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name":"Barbarian"`,
					`"name":"Artificer"`,
					`"name":"Wizard"`,
				},
			},
		},
		{
			TestName: "can get list of classes (empty table)",
			Setup: func() {
				ts.ClearTable("classes")
			},
			Request: helpers.Request{Method: http.MethodGet, URL: "/classes"},
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
