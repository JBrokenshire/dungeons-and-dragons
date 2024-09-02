package helpers

import (
	"bytes"
	models2 "dungeons-and-dragons/db/models"
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
	S *server.Server
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
	ts.ClearTable("classes")
	classes := []*models2.Class{
		{ID: 1, Name: "Artificer"},
		{ID: 2, Name: "Barbarian", Description: "A fierce warrior who can enter a battle rage"},
		{ID: 3, Name: "Bard", Description: "An inspiring magician whose power echoes the music of creation"},
		{ID: 4, Name: "Blood Hunter", Description: "Willing to suffer whatever it takes to achieve victory, these adept warriors have forged themselves into a potent force dedicated to protecting the innocent."},
		{ID: 5, Name: "Cleric", Description: "A priestly champion who wields divine magic in service of a higher power"},
		{ID: 6, Name: "Druid", Description: "A priest of the Old Faith, wielding the powers of nature and adopting animal forms"},
		{ID: 7, Name: "Fighter", Description: "A master of martial combat, skilled with a variety of weapons and armor"},
		{ID: 8, Name: "Monk", Description: "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection"},
		{ID: 9, Name: "Paladin", Description: "A holy warrior bound to a sacred oath"},
		{ID: 10, Name: "Ranger", Description: "A warrior who combats threats on the edges of civilization"},
		{ID: 11, Name: "Rogue", Description: "A scoundrel who uses stealth and trickery to overcome obstacles and enemies"},
		{ID: 12, Name: "Sorcerer", Description: "A spellcaster who draws on inherent magic from a gift or bloodline"},
		{ID: 13, Name: "Warlock", Description: "A wielder of magic that is derived from a bargain with an extraplanar entity"},
		{ID: 14, Name: "Wizard", Description: "A scholarly magic-user capable of manipulating the structures of reality"},
	}
	for _, class := range classes {
		ts.S.Db.Create(class)
	}
}

func (ts *TestServer) SetupDefaultRaces() {
	ts.ClearTable("races")
	races := []*models2.Race{
		{ID: 1, Name: "Aarakocra"},
		{ID: 2, Name: "Dragonborn"},
		{ID: 3, Name: "Hill Dwarf"},
		{ID: 4, Name: "Moutain Dwarf"},
		{ID: 5, Name: "Eladrin Elf"},
		{ID: 6, Name: "High Elf"},
		{ID: 7, Name: "Wood Elf"},
		{ID: 8, Name: "Air Genasi"},
		{ID: 9, Name: "Earth Genasi"},
		{ID: 10, Name: "Fire Genasi"},
		{ID: 11, Name: "Water Genasi"},
		{ID: 12, Name: "Rock Gnome"},
		{ID: 13, Name: "Deep Gnome"},
		{ID: 14, Name: "Goliath"},
		{ID: 15, Name: "Half-Elf"},
		{ID: 16, Name: "Half-Orc"},
		{ID: 17, Name: "Lightfoot Halfling"},
		{ID: 18, Name: "Stout Halfling"},
		{ID: 19, Name: "Human"},
		{ID: 20, Name: "Variant Human"},
		{ID: 21, Name: "Tiefling"},
		{ID: 22, Name: "Variant Aasimar"},
	}
	for _, race := range races {
		ts.S.Db.Create(race)
	}
}

func (ts *TestServer) SetupDefaultCharacters() {
	ts.ClearTable("characters")

	ts.SetupDefaultClasses()
	ts.SetupDefaultRaces()

	characters := []*models2.Character{
		{
			ID:      1,
			Name:    "Faelan Haversham",
			Level:   3,
			ClassID: 4,
			RaceID:  18,
		},
		{
			ID:      2,
			Name:    "PeeWee McAnkle-Biter",
			Level:   5,
			ClassID: 2,
			RaceID:  3,
		},
	}

	for _, character := range characters {
		ts.S.Db.Create(character)
	}

}
