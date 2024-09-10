package helpers

import (
	"bytes"
	"dungeons-and-dragons/db/seeders"
	"dungeons-and-dragons/db/stores"
	"dungeons-and-dragons/server"
	"dungeons-and-dragons/server/routes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type TestServer struct {
	S      *server.Server
	seeder *seeders.Seeder
}

func NewTestServer() *TestServer {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
	)

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	ts := &TestServer{
		S: &server.Server{
			Echo:   echo.New(),
			Db:     db,
			Stores: stores.NewStores(db),
		},
		seeder: seeders.NewSeeder(db),
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
	req.Header.Set(echo.HeaderXRealIP, "127.0.0.0")
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

// ClearTable Clear a table and reset the autoincrement
func (ts *TestServer) ClearTable(tableName string) {
	err := ts.S.Db.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}
	err = ts.S.Db.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		log.Fatalf("Error setting autoincrement. Err: %v", err)
	}
}

func (ts *TestServer) SetupDefaultClasses() {
	ts.ClearTable("characters")
	ts.ClearTable("classes")

	ts.seeder.SetClasses()
}

func (ts *TestServer) SetupDefaultRaces() {
	ts.ClearTable("characters")
	ts.ClearTable("races")

	ts.seeder.SetRaces()
}

func (ts *TestServer) SetupDefaultCharacters() {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	ts.seeder.SetCharacters()
}
