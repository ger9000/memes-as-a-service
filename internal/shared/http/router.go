package http

import (
	"net/http"

	"github.com/go-chi/render"
)

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, data)
}

func Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	code := http.StatusUnauthorized
	response := httpError{
		Code:    code,
		Message: "Bad Request",
	}

	render.Status(r, code)
	if err != nil {
		response.Message = err.Error()
		render.JSON(w, r, response)
		return
	}
	render.JSON(w, r, response)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	code := http.StatusInternalServerError
	response := httpError{
		Code:    code,
		Message: "Internal Server Error",
	}

	render.Status(r, code)
	if err != nil {
		response.Message = err.Error()
		render.JSON(w, r, response)
		return
	}
	render.JSON(w, r, response)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	code := http.StatusBadRequest
	response := httpError{
		Code:    code,
		Message: "Bad Request",
	}

	render.Status(r, code)
	if err != nil {
		response.Message = err.Error()
		render.JSON(w, r, response)
		return
	}
	render.JSON(w, r, response)
}

func NoContent(w http.ResponseWriter, r *http.Request) {
	render.NoContent(w, r)
}
