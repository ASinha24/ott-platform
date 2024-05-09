package http

import (
	"encoding/json"
	"net/http"

	ott "github.com/asinha24/ott-platform"
	"github.com/asinha24/ott-platform/api"
	"github.com/asinha24/ott-platform/movies"
	"github.com/asinha24/ott-platform/movies/model"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

type OTTHandler struct {
	Manager  ott.OttManager
	OTTStore movies.OTTPlatform
}

func (o *OTTHandler) IntallRouter(mux *mux.Router) {
	// It will Create a ott
	mux.Methods(http.MethodPost).Path("/ott").HandlerFunc(o.CreateNewOTTPlatform)
	// It will get all platforms
	mux.Methods(http.MethodGet).Path("/ott").HandlerFunc(o.GetOTTPlatforms)
	// It will get platform by name
	mux.Methods(http.MethodGet).Path("/ott/{ott_name}").HandlerFunc(o.GetPlatformByName)
	// List all movies of a ott
	mux.Methods(http.MethodGet).Path("/ott/{ott_name}/movies").HandlerFunc(o.Getmovies)
	// Create a new movie of super mart
	mux.Methods(http.MethodPost).Path("/ott/{ott_name}/movies").HandlerFunc(o.CreateNewMovie)
	// Update an existing movie of a mart
	mux.Methods(http.MethodPut).Path("/ott/{ott_name}/movies/{movieID}").HandlerFunc(o.UpdateMovie)
	// delete any movie of a mart
	mux.Methods(http.MethodDelete).Path("/ott/{ott_name}/movies/{movieID}").HandlerFunc(o.DeleteMovie)
}

func NewOTTHandler(manager ott.OttManager, OTTStore movies.OTTPlatform) *OTTHandler {
	return &OTTHandler{Manager: manager, OTTStore: OTTStore}
}

func (o *OTTHandler) CreateNewOTTPlatform(w http.ResponseWriter, r *http.Request) {
	createReq := &api.OTT{}
	if err := json.NewDecoder(r.Body).Decode(createReq); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	createReq.ID = uuid.New()
	if err := o.OTTStore.CreateOTT(r.Context(), &model.OTT{
		Name: createReq.Name,
		ID:   createReq.ID,
	}); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusCreated, createReq, w)
}

func (o *OTTHandler) GetOTTPlatforms(w http.ResponseWriter, r *http.Request) {
	otts, err := o.OTTStore.GetAllOTTPLatform(r.Context())
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusOK, otts, w)
}

func (o *OTTHandler) GetPlatformByName(w http.ResponseWriter, r *http.Request) {
	ottName := mux.Vars(r)["ott_name"]
	otts, err := o.OTTStore.GetOTTPlatformByName(r.Context(), ottName)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusOK, otts, w)
}

func (o *OTTHandler) Getmovies(w http.ResponseWriter, r *http.Request) {
	ottName := mux.Vars(r)["ott_name"]

	resp, err := o.Manager.GetMovies(r.Context(), ottName)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusOK, resp, w)
}

func (o *OTTHandler) CreateNewMovie(w http.ResponseWriter, r *http.Request) {
	ottname := mux.Vars(r)["ott_name"]
	createRq := &api.MovieRequest{}

	if err := json.NewDecoder(r.Body).Decode(createRq); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	resp, err := o.Manager.CreateMovie(r.Context(), ottname, createRq)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusCreated, resp, w)
}

func (o *OTTHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	ottname := mux.Vars(r)["ott_name"]
	movieID := mux.Vars(r)["movieID"]
	updateRq := &api.MovieRequest{}

	if err := json.NewDecoder(r.Body).Decode(updateRq); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	resp, err := o.Manager.UpdateMovie(r.Context(), ottname, movieID, updateRq)
	if err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}

	WriteResponse(http.StatusCreated, resp, w)

}

func (o *OTTHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	ottname := mux.Vars(r)["ott_name"]
	movieID := mux.Vars(r)["movieID"]

	if err := o.Manager.DeleteMovie(r.Context(), ottname, movieID); err != nil {
		WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
}
