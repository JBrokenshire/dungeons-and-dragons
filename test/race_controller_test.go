package test

//type StubRaceStore struct {
//	races []*models.Race
//}
//
//func (s *StubRaceStore) GetAll() ([]*models.Race, error) {
//	return s.races, nil
//}
//
//func (s *StubRaceStore) Get(id int) (*models.Race, error) {
//	for _, race := range s.races {
//		if race.ID == id {
//			return race, nil
//		}
//	}
//
//	return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("race with id %q not found", id))
//}
//
//func TestGetAllRaces(t *testing.T) {
//	store := &StubRaceStore{races: []*models.Race{
//		{
//			ID:   1,
//			Name: "Test Race",
//		},
//		{
//			ID:   2,
//			Name: "Test Race 2",
//		},
//	}}
//	raceController := controllers.RaceController{Store: store}
//
//	e := echo.New()
//	request := httptest.NewRequest(http.MethodGet, "/races", nil)
//	response := httptest.NewRecorder()
//	ctx := e.NewContext(request, response)
//
//	err := raceController.GetAll(ctx)
//	assert.NoError(t, err)
//
//	var got []*models.Race
//	err = json.Unmarshal(response.Body.Bytes(), &got)
//	assert.NoError(t, err)
//	assert.Equal(t, len(got), len(store.races))
//}
//
//func TestGETRace(t *testing.T) {
//	store := &StubRaceStore{races: []*models.Race{
//		{
//			ID:   1,
//			Name: "Test Race",
//		},
//		{
//			ID:   2,
//			Name: "Test Race 2",
//		},
//	}}
//	raceController := controllers.RaceController{Store: store}
//	e := echo.New()
//
//	for _, race := range store.races {
//		t.Run(fmt.Sprintf("GET /races/%v should return %v", race.ID, race), func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/races/%v", race.ID), nil)
//			response := httptest.NewRecorder()
//			ctx := e.NewContext(request, response)
//			ctx.SetParamNames("id")
//			ctx.SetParamValues(strconv.Itoa(race.ID))
//
//			err := raceController.Get(ctx)
//			assert.NoError(t, err)
//
//			var got *models.Race
//			err = json.Unmarshal(response.Body.Bytes(), &got)
//			assert.NoError(t, err)
//			assert.Equal(t, race, got)
//		})
//	}
//}
