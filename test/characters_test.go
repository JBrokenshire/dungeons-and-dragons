package test

import (
	"dungeons-and-dragons/test/helpers"
	"net/http"
	"testing"
)

func TestCharacterList(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodGet,
		URL:    "/characters",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can get list of characters",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"name": "Faelan Haversham"`,
					`"name": "PeeWee McAnkle-Biter"`,
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
