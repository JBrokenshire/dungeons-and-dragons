package test

import (
	"dungeons-and-dragons/db/models"
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
				ts.ClearTable("characters")
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

func TestGetClass(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can get class by id",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/classes/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Artificer"`},
			},
		},
		{
			TestName: "get /classes/:id returns 404 not found on class id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodGet,
				URL:    "/classes/100",
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

func TestUpdateClass(t *testing.T) {
	cases := []helpers.TestCase{
		{
			TestName: "can update class by id with valid json in request body",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			RequestBody: models.Class{
				Name:        "Test Name",
				Description: "Test Description",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Test Name"`, `"description":"Test Description"`},
			},
		},
		{
			TestName: "put /classes/:id returns 404 not found on class id not in database",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/100",
			},
			RequestBody: models.Class{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			TestName: "put /classes/:id returns 400 bad request on empty request body",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			TestName: "put /classes/:id no update on empty class in request body",
			Setup: func() {
				ts.ClearTable("characters")
				ts.ClearTable("classes")
				ts.SetupDefaultClasses()
			},
			Request: helpers.Request{
				Method: http.MethodPut,
				URL:    "/classes/1",
			},
			RequestBody: models.Class{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts:  []string{`"id":1`, `"name":"Artificer"`},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			RunTestCase(t, test)
		})
	}
}
