package test

import (
	"dungeons-and-dragons/test/helpers"
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

	// Run the test
	code := m.Run()

	os.Exit(code)
}

func RunTestCase(t *testing.T, test helpers.TestCase) {
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
}

func isIn(t *testing.T, s, contains string, msgAndArgs ...interface{}) bool {
	t.Helper()

	ok := strings.Contains(s, contains)
	if !ok {
		return assert.Fail(t, fmt.Sprintf("%#v is not in %#v", contains, s), msgAndArgs...)
	}

	return true
}
