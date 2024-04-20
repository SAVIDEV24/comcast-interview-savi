package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stringinator-go/datastore"
	"stringinator-go/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_StringinatePost(t *testing.T) {

	tests := []struct {
		name                   string
		requestURI             string
		requestBody            []byte
		expectedResponseCode   int
		expectedresponseStruct []model.CharCount
	}{
		{
			name: "stringiNatePostSuccess",

			requestURI:             "/stringinate",
			requestBody:            []byte(`{"input":"hello world"}`),
			expectedResponseCode:   http.StatusOK,
			expectedresponseStruct: []model.CharCount{{Char: "l", Occurance: 3}},
		},
		{
			name:                   "stringiNatePostInvalidRequestBody",
			requestURI:             "/stringinate",
			requestBody:            []byte(`{"invalid":"invalid"}`),
			expectedResponseCode:   http.StatusBadRequest,
			expectedresponseStruct: nil,
		},
		{
			name:                   "stringiNatePostEmptyRequestBody",
			requestURI:             "/stringinate",
			requestBody:            []byte(`{}`),
			expectedResponseCode:   http.StatusBadRequest,
			expectedresponseStruct: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stringnatorstruct = NewStringinatorService(make(map[string]int), datastore.InMemoryStore{})
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/stringinate", bytes.NewBuffer(test.requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder to record the response
			rec := httptest.NewRecorder()

			// Call the handler function directly
			c := e.NewContext(req, rec)
			stringnatorstruct.Stringinate(c)

			assert.Equal(t, test.expectedResponseCode, rec.Code)
			// Check the response status code
			if rec.Code != test.expectedResponseCode {
				t.Errorf("expected status OK; got %d", rec.Code)
			}

			if rec.Body.String() != "null" {
				jsonstr, _ := json.Marshal(test.expectedresponseStruct)
				expected := string(jsonstr)
				assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
			}

		})
	}

}

func Test_StringinateGet(t *testing.T) {

	tests := []struct {
		name                   string
		requestURI             string
		expectedResponseCode   int
		expectedresponseStruct []model.CharCount
	}{
		{
			name:                   "stringiNateGetSuccess",
			requestURI:             "/stringinate?input=helloworld",
			expectedResponseCode:   http.StatusOK,
			expectedresponseStruct: []model.CharCount{{Char: "l", Occurance: 3}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stringnatorstruct = NewStringinatorService(make(map[string]int), datastore.InMemoryStore{})
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, test.requestURI, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			stringnatorstruct.Stringinate(c)

			assert.Equal(t, test.expectedResponseCode, rec.Code)

			if rec.Body.String() != "null" {
				jsonstr, _ := json.Marshal(test.expectedresponseStruct)
				expected := string(jsonstr)
				assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
			}

		})
	}

}

func Test_Stats(t *testing.T) {
	var stringnatorstruct = NewStringinatorService(make(map[string]int), datastore.InMemoryStore{})
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	if err := stringnatorstruct.Stats(c); err != nil {
		t.Fatalf("Stats failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var expected = model.Statistics{Mostoccurred: []string{"hiiw"}, LongestInput: "hello"}

	jsonstr, _ := json.Marshal(expected)

	assert.Equal(t, string(jsonstr), strings.TrimSpace(rec.Body.String()))
}
