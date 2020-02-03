package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/places/service"
	"github.com/julienschmidt/httprouter"
)

//PlacePHandler ..
type PlacePHandler struct {
	placeSrv *service.PlaceService
}

//NewPlacePHandler returns  new UserHandler
func NewPlacePHandler(srv *service.PlaceService) *PlacePHandler {
	return &PlacePHandler{placeSrv: srv}
}

//GetSinglePlace returns a single user
func (placeHdlr *PlacePHandler) GetSinglePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idRaw := ps.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := placeHdlr.placeSrv.GetSinglePlace(id)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
	return

}

//GetManyPlaces returns many users
func (placeHdlr *PlacePHandler) GetManyPlaces(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idRaw := ps.ByName("userid")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := placeHdlr.placeSrv.GetManyPlaces(id)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
	return
}

//AddPlace adds a new user to the database
func (placeHdlr *PlacePHandler) AddPlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.Place{}

	errs := json.Unmarshal(body, user)

	if errs != nil {
		// log.Println("Adding a new user")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		//w.Write([]byte("Unmarshal error"))
		return
	}

	medicine, err := placeHdlr.placeSrv.AddPlace(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/host/place/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

//UpdatePlace updates user
func (placeHdlr *PlacePHandler) UpdatePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := placeHdlr.placeSrv.GetSinglePlace(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &user)

	userUp, err := placeHdlr.placeSrv.UpdatePlace(user)

	if err != nil {
		fmt.Println("Error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(userUp, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//DeletePlace deletes a single user
func (placeHdlr *PlacePHandler) DeletePlace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("id ", id)

	errs := placeHdlr.placeSrv.DeletePlace(id)

	if errs != nil {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
