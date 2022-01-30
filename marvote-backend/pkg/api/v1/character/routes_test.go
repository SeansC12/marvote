package character

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SeansC12/marvote/pkg/model"
	"github.com/SeansC12/marvote/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CharacterRouteTestSuite struct {
	suite.Suite
}

func (ts *CharacterRouteTestSuite) TestGetOneCharacterRoute() {
	marvelCharJSON := `{"id":0,"name":"Spiderman","aka":"Peter Parker"}`
	ci := model.CharacterInfo{
		Id:   0,
		Name: "Spiderman",
		Aka:  "Peter Parker",
	}
	mockService := new(MockedCharacterService)
	mockService.On("Get", 0).Return(ci, nil)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/character/:id")
	c.SetParamNames("id")
	c.SetParamValues("0")

	h := NewCharacterRoutes(mockService)
	if assert.NoError(ts.T(), h.Get(c)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), marvelCharJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}

}

func (ts *CharacterRouteTestSuite) TestFailToGetOneCharacterRoute() {

	mockService := new(MockedCharacterService)
	reqErr := &service.ErrorFailedToLoadData{}
	mockService.On("Get", 0).Return(model.CharacterInfo{}, reqErr)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/character/:id")
	c.SetParamNames("id")
	c.SetParamValues("0")

	h := NewCharacterRoutes(mockService)
	err := h.Get(c)
	if assert.NotNil(ts.T(), err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(ts.T(), http.StatusBadRequest, he.Code)
		}

	}

}

func (ts *CharacterRouteTestSuite) TestInvalidId() {

	mockService := new(MockedCharacterService)
	reqErr := &service.ErrorFailedToLoadData{}
	mockService.On("Get", 0).Return(model.CharacterInfo{}, reqErr)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/character/:id")
	c.SetParamNames("id")
	c.SetParamValues("A")

	h := NewCharacterRoutes(mockService)
	err := h.Get(c)
	if assert.NotNil(ts.T(), err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(ts.T(), http.StatusBadRequest, he.Code)
		}
	}

}
func (ts *CharacterRouteTestSuite) TestGetAllCharacterRoute() {
	marvelCharJSON := `[{"id":0,"name":"Spiderman","aka":"Peter Parker"}]`
	ci := model.CharacterInfo{
		Id:   0,
		Name: "Spiderman",
		Aka:  "Peter Parker",
	}
	all := make([]model.CharacterInfo, 0)
	mockService := new(MockedCharacterService)
	all = append(all, ci)
	mockService.On("GetAll").Return(all, nil)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/characters")

	h := NewCharacterRoutes(mockService)
	if assert.NoError(ts.T(), h.GetAllCharacters(c)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), marvelCharJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}

}

func (ts *CharacterRouteTestSuite) TestFailedGetAllCharacterRoute() {
	all := make([]model.CharacterInfo, 0)
	reqErr := &service.ErrorFailedToLoadData{}
	mockService := new(MockedCharacterService)
	mockService.On("GetAll").Return(all, reqErr)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/characters")

	h := NewCharacterRoutes(mockService)
	err := h.GetAllCharacters(c)
	if assert.NotNil(ts.T(), err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(ts.T(), http.StatusBadRequest, he.Code)
		}
	}

}

func (ts *CharacterRouteTestSuite) TestSuccessfulSave() {
	marvelCharJSON := `{"name":"Wolverine","aka":"Logan"}`
	ci := model.CharacterInfo{
		Name: "Wolverine",
		Aka:  "Logan",
	}

	responseCi := model.CharacterInfo{
		Id:   3,
		Name: "Wolverine",
		Aka:  "Logan",
	}
	mockService := new(MockedCharacterService)
	mockService.On("Save", ci).Return(responseCi, nil)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/", strings.NewReader(marvelCharJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/character")

	h := NewCharacterRoutes(mockService)
	if assert.NoError(ts.T(), h.Save(c)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
	}

}

func TestCharacterRoutes(t *testing.T) {
	suite.Run(t, new(CharacterRouteTestSuite))
}
