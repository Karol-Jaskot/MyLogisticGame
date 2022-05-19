package handlers

import (
	"MyLogisticGame/backend/entity"
	"MyLogisticGame/backend/routes"
	"MyLogisticGame/configs"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var e *echo.Echo

func setupTest() {
	// Setup
	e = echo.New()
	configs.SetTstConnection()
	routes.CreateLocationsAPI(e)

	defer e.Close()
}

func TestCreateLocation(t *testing.T) {
	setupTest()
	url := "/locations"

	var loc = entity.Location{Name: "Test"}
	jLoc, err := json.Marshal(loc)
	if err != nil {
		e.Logger.Error(err)
	}

	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(jLoc)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateLocation(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
