package memes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	"github.com/ger9000/memes-as-a-service/internal/usecases/memes"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllController(t *testing.T) {
	t.Run("OK response without params", func(t *testing.T) {
		// Given
		controller := Init()
		r := httptest.NewRequest("GET", "/memes", nil)
		w := httptest.NewRecorder()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, w.Code, http.StatusOK)

		response := Response{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, 10, response.Count)
	})

	t.Run("OK response with params", func(t *testing.T) {
		// Given
		controller := Init()
		r := httptest.NewRequest("GET", "/memes?latitude=0.1&longitude=0.2&query=test", nil)
		w := httptest.NewRecorder()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, w.Code, http.StatusOK)

		response := Response{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, 4, response.Count)
	})

	t.Run("Error response when bad latitude given", func(t *testing.T) {
		// Given
		controller := Init()
		r := httptest.NewRequest("GET", "/memes?latitude=test&longitude=0.2&query=test", nil)
		w := httptest.NewRecorder()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, w.Code, http.StatusBadRequest)

		response := map[string]interface{}{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, int(response["code"].(float64)))
	})

	t.Run("Error response when bad longitude given", func(t *testing.T) {
		// Given
		controller := Init()
		r := httptest.NewRequest("GET", "/memes?latitude=0.1&longitude=test&query=test", nil)
		w := httptest.NewRecorder()

		// When
		controller.Invoke(w, r)

		// Then
		assert.Equal(t, w.Code, http.StatusBadRequest)

		response := map[string]interface{}{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, int(response["code"].(float64)))
	})
}

func Init() rest.IController {
	action := memes.NewGetAllAction()
	return NewGetAllController(action)
}
