package helpers

import (
	"bytes"
	"dungeons-and-dragons/db"
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/server/routes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	S *server.Server
}

func NewTestServer() *TestServer {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ts := &TestServer{
		S: &server.Server{
			Echo: echo.New(),
			Db:   db.Init(),
		},
	}

	routes.ConfigureRoutes(ts.S)

	return ts
}

func (ts *TestServer) ExecuteTestCase(testCase *TestCase) *httptest.ResponseRecorder {
	req := ts.GenerateRequest(testCase)
	return ts.ExecuteRequest(req)
}

func (ts *TestServer) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ts.S.Echo.ServeHTTP(rr, req)
	return rr
}

func (ts *TestServer) SetDefaultTestHeaders(req *http.Request) {
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//req.Header.Set(echo.HeaderXRealIP, "127.0.0.0")
}

func (ts *TestServer) GenerateRequest(testCase *TestCase) *http.Request {
	reqJSON, err := json.Marshal(testCase.RequestBody)
	if err != nil {
		log.Printf("There was an error marshalling the JSON: %v", err)
	}

	req, err := http.NewRequest(testCase.Request.Method, testCase.Request.URL, bytes.NewBuffer(reqJSON))

	ts.SetDefaultTestHeaders(req)

	return req
}
