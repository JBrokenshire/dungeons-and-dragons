package test

import (
	"bytes"
	"dungeons-and-dragons/controller"
	"dungeons-and-dragons/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubCharacterStore struct {
	characters []*models.Character
}

func (s *StubCharacterStore) Create(character *models.Character) error {
	s.characters = append(s.characters, character)
	return nil
}

func (s *StubCharacterStore) GetAll() ([]*models.Character, error) {
	return s.characters, nil
}

func (s *StubCharacterStore) Get(id int) (*models.Character, error) {
	for _, character := range s.characters {
		if character.ID == id {
			return character, nil
		}
	}

	return nil, echo.NewHTTPError(http.StatusNotFound)
}

func (s *StubCharacterStore) Update(character *models.Character) error {
	for i, char := range s.characters {
		if char.ID == character.ID {
			s.characters[i] = character
			return nil
		}
	}

	return echo.NewHTTPError(http.StatusNotFound)
}

func (s *StubCharacterStore) LevelUp(character *models.Character) error {
	character.Level++
	return nil
}

func (s *StubCharacterStore) Delete(id int) error {
	for i, char := range s.characters {
		if char.ID == id {
			s.characters = append(s.characters[:i], s.characters[i+1:]...)
			return nil
		}
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

func TestPOSTCharacter(t *testing.T) {
	t.Run("POST /characters creates a new character", func(t *testing.T) {
		store := &StubCharacterStore{[]*models.Character{}}
		characterController := controller.CharacterController{Store: store}

		reqBody, _ := json.Marshal(models.Character{
			Name:  "Test",
			Level: 1,
		})
		request := httptest.NewRequest(http.MethodPost, "/characters", bytes.NewReader(reqBody))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)

		err := characterController.Create(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, response.Code)

		characters, err := characterController.Store.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(characters))
	})
}

func TestGETAllCharacters(t *testing.T) {
	t.Run("GET /characters should return all characters", func(t *testing.T) {
		store := &StubCharacterStore{[]*models.Character{
			{
				Name:  "Test",
				Level: 1,
			},
			{
				Name:  "Test2",
				Level: 2,
			},
		}}
		characterController := controller.CharacterController{Store: store}

		request := httptest.NewRequest(http.MethodGet, "/characters", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)

		err := characterController.GetAll(ctx)
		assert.NoError(t, err)

		var got []*models.Character
		err = json.Unmarshal(response.Body.Bytes(), &got)
		assert.NoError(t, err)

		assert.Equal(t, len(store.characters), len(got))
	})
}

func TestGETCharacter(t *testing.T) {
	store := &StubCharacterStore{[]*models.Character{
		{
			ID:    1,
			Name:  "Test",
			Level: 1,
		},
		{
			ID:    2,
			Name:  "Test2",
			Level: 1,
		},
	}}

	characterController := controller.CharacterController{Store: store}

	for i, c := range store.characters {
		t.Run(fmt.Sprintf("GET /character/%v should return %v", c.ID, c.Name), func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/characters/%v", c.ID), nil)
			response := httptest.NewRecorder()
			ctx := ts.S.Echo.NewContext(request, response)
			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(c.ID))

			err := characterController.Get(ctx)
			assert.NoError(t, err)

			var got *models.Character
			err = json.Unmarshal(response.Body.Bytes(), &got)
			assert.NoError(t, err)
			assert.Equal(t, got, store.characters[i])
		})
	}

	t.Run("returns bad request status when id is NaN", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/character/NaN", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("NaN")

		err := characterController.Get(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, ctx.Response().Status)
	})

	t.Run("returns not found status for invalid character id", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/character/100", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("100")

		err := characterController.Get(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
	})
}

func TestPUTCharacter(t *testing.T) {
	updatedCharacter := models.Character{
		ID:    1,
		Name:  "This is a new name",
		Level: 2,
	}
	store := &StubCharacterStore{[]*models.Character{
		{
			ID:    1,
			Name:  "Test",
			Level: 1,
		},
	}}
	characterController := controller.CharacterController{Store: store}

	reqBody, _ := json.Marshal(updatedCharacter)

	t.Run("PUT /character/:id should update character with that id", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/characters/1", bytes.NewReader(reqBody))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := characterController.Update(ctx)
		assert.NoError(t, err)

		var got models.Character
		err = json.Unmarshal(response.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, updatedCharacter, got)
	})

	t.Run("PUT /character/:id returns bad request status when id is NaN", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/character/NaN", bytes.NewReader(reqBody))
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("NaN")

		err := characterController.Update(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, ctx.Response().Status)
	})

	t.Run("PUT /character/:id returns not found status when id doesn't exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/character/100", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("100")

		err := characterController.Update(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
	})
}

func TestDELETECharacter(t *testing.T) {
	store := &StubCharacterStore{[]*models.Character{
		{
			ID: 1,
		},
	}}
	characterController := controller.CharacterController{Store: store}

	t.Run("DELETE /character/:id should delete character with that id", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/characters/%v", 1), nil)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		err := characterController.Delete(ctx)
		assert.NoError(t, err)
		assert.Equal(t, len(store.characters), 0)
	})

	t.Run("DELETE /character/:id returns bad request status when id is NaN", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/character/NaN", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("NaN")

		err := characterController.Delete(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, ctx.Response().Status)
	})

	t.Run("DELETE /character/:id returns not found status when id doesn't exist", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/character/100", nil)
		response := httptest.NewRecorder()
		ctx := ts.S.Echo.NewContext(request, response)
		ctx.SetParamNames("id")
		ctx.SetParamValues("100")

		err := characterController.Delete(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, ctx.Response().Status)
	})
}

func TestLevelUpCharacter(t *testing.T) {
	store := &StubCharacterStore{[]*models.Character{
		{
			ID:    1,
			Level: 1,
		},
	}}
	characterController := controller.CharacterController{Store: store}

	for _, c := range store.characters {
		t.Run(fmt.Sprintf("PUT /character/%v/level-up should level up the character with ID %q", c.ID, c.ID), func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/characters/%v/level-up", c.ID), nil)
			response := httptest.NewRecorder()
			ctx := ts.S.Echo.NewContext(request, response)
			ctx.SetParamNames("id")
			ctx.SetParamValues(strconv.Itoa(c.ID))

			err := characterController.LevelUp(ctx)
			assert.NoError(t, err)

			var got *models.Character
			err = json.Unmarshal(response.Body.Bytes(), &got)
			assert.NoError(t, err)
			assert.Equal(t, 2, got.Level)
		})

	}
}
