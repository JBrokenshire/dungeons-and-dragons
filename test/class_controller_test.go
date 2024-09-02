package test

import (
	"bytes"
	"dungeons-and-dragons/db/models"
	"dungeons-and-dragons/server/controller"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubClassStore struct {
	classes []*models.Class
}

func (s *StubClassStore) GetAll() ([]*models.Class, error) {
	return s.classes, nil
}

func (s *StubClassStore) Get(id int) (*models.Class, error) {
	for _, class := range s.classes {
		if class.ID == id {
			return class, nil
		}
	}

	return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("class with id %q not found", id))
}

func (s *StubClassStore) Update(class *models.Class) error {
	for i, c := range s.classes {
		if c.ID == class.ID {
			s.classes[i] = class
			return nil
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("class with id %q not found", class.ID))
}

func TestGetAllClasses(t *testing.T) {
	store := &StubClassStore{classes: []*models.Class{
		{
			ID:   1,
			Name: "Test Class",
		},
		{
			ID:   2,
			Name: "Test Class 2",
		},
	}}
	classController := controller.ClassController{Store: store}

	request := httptest.NewRequest(http.MethodGet, "/classes", nil)
	response := httptest.NewRecorder()
	ctx := ts.S.Echo.NewContext(request, response)

	err := classController.GetAll(ctx)
	assert.NoError(t, err)

	var got []*models.Class
	err = json.Unmarshal(response.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, len(got), len(store.classes))
}

func TestGetClass(t *testing.T) {
	store := &StubClassStore{classes: []*models.Class{
		{
			ID:   1,
			Name: "Test Class",
		},
	}}
	classController := controller.ClassController{Store: store}

	request := httptest.NewRequest(http.MethodGet, "/class/test-class", nil)
	response := httptest.NewRecorder()
	ctx := ts.S.Echo.NewContext(request, response)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(1))

	err := classController.Get(ctx)
	assert.NoError(t, err)

	var got *models.Class
	err = json.Unmarshal(response.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, store.classes[0], got)
}

func TestPUTClass(t *testing.T) {
	updatedClass := models.Class{
		ID:   1,
		Name: "New Test Class",
	}
	store := &StubClassStore{classes: []*models.Class{
		{
			ID:   1,
			Name: "Test Class",
		},
	}}
	classController := controller.ClassController{Store: store}
	reqBody, _ := json.Marshal(updatedClass)
	request := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/classes/1"), bytes.NewReader(reqBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	ctx := ts.S.Echo.NewContext(request, response)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(updatedClass.ID))

	err := classController.Update(ctx)
	assert.NoError(t, err)

	var got *models.Class
	err = json.Unmarshal(response.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, &updatedClass, got)
}
