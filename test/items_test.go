package test

import (
	"dnd-api/db/factories"
	"dnd-api/db/models"
	"dnd-api/test/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestGetAllItems(t *testing.T) {
	ts.ClearTable("items")

	itemOne := &models.Item{
		Name: "Test Item One",
	}
	factories.NewItem(ts.S.Db, itemOne)

	itemTwo := &models.Item{
		Name: "Test Item Two",
	}
	factories.NewItem(ts.S.Db, itemTwo)

	request := helpers.Request{
		Method: http.MethodGet,
		URL:    "/items",
	}

	cases := []helpers.TestCase{
		{
			TestName: "Can get all items",
			Request:  request,
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%s"`, itemOne.Name),
					fmt.Sprintf(`"meta":"%s"`, itemOne.Meta),
					fmt.Sprintf(`"name":"%s"`, itemTwo.Name),
					fmt.Sprintf(`"meta":"%s"`, itemTwo.Meta),
				},
				BodyPartsMissing: nil,
			},
		},
		{
			TestName: "Can get no items (empty table)",
			Setup:    func() { ts.ClearTable("items") },
			Request:  request,
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
