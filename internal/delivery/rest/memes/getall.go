package memes

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ger9000/memes-as-a-service/internal/delivery/rest"
	router "github.com/ger9000/memes-as-a-service/internal/shared/http"
	"github.com/ger9000/memes-as-a-service/internal/usecases/memes"
)

type controller struct {
	action *memes.GetAllAction
}

func (c *controller) Invoke(w http.ResponseWriter, r *http.Request) {
	var latitude, longitude float64
	query := r.URL.Query().Get("query")

	if latitudeParam := r.URL.Query().Get("latitude"); latitudeParam != "" {
		latitudeParsed, err := strconv.ParseFloat(latitudeParam, 64)
		if err != nil {
			router.BadRequest(w, r, errors.New("malformed latitude query param"))
			return
		}
		latitude = latitudeParsed
	}

	if longitudeParam := r.URL.Query().Get("longitude"); longitudeParam != "" {
		longitudeParsed, err := strconv.ParseFloat(longitudeParam, 64)
		if err != nil {
			router.BadRequest(w, r, errors.New("malformed longitude query param"))
			return
		}
		longitude = longitudeParsed
	}

	memes := c.action.Do(r.Context(), latitude, longitude, query)

	router.Success(w, r, Response{
		Count: len(memes),
		Data:  memes,
	})
}

func NewGetAllController(action *memes.GetAllAction) rest.IController {
	return &controller{
		action,
	}
}
