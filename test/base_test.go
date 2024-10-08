package test

import (
	"dnd-api/test/helpers"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	ts *helpers.TestServer
)

func TestMain(m *testing.M) {
	err := os.Setenv("ENVIRONMENT", "development")
	if err != nil {
		log.Printf("Error setting ENVIRONMENT: %v\n", err)
	}

	ts = helpers.NewTestServer()
	defer ts.S.Db.Close()

	// Run the test
	code := m.Run()

	os.Exit(code)
}

func RunTestCase(t *testing.T, test helpers.TestCase) {
	// Perform any setup needed for the test case
	if test.Setup != nil {
		test.Setup()
	}

	res := ts.ExecuteTestCase(&test)
	ValidateResults(t, test, res)
}

func ValidateResults(t *testing.T, test helpers.TestCase, res *httptest.ResponseRecorder) {
	if res.Code != 0 {
		assert.Equal(t, test.Expected.StatusCode, res.Code)
	}

	if test.Expected.BodyPart != "" {
		isIn(t, res.Body.String(), test.Expected.BodyPart)
	}

	if len(test.Expected.BodyParts) > 0 {
		for _, expectedText := range test.Expected.BodyParts {
			isIn(t, res.Body.String(), expectedText)
		}
	}

	if len(test.Expected.BodyPartsMissing) > 0 {
		for _, missingText := range test.Expected.BodyPartsMissing {
			isNotIn(t, res.Body.String(), missingText)
		}
	}
}

func isIn(t *testing.T, s, contains string, msgAndArgs ...interface{}) bool {
	t.Helper()

	ok := strings.Contains(s, contains)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("%#v is not in %#v", contains, s), msgAndArgs...)
	}

	return true
}

func isNotIn(t *testing.T, s, notContains string, msgAndArgs ...interface{}) bool {
	t.Helper()

	ok := !strings.Contains(s, notContains)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("%#v is not in %#v", notContains, s), msgAndArgs...)
	}

	return true
}
