package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/pictures/service"
	"github.com/julienschmidt/httprouter"
)

//PictureHandler ..
type PictureHandler struct {
	picSrv *service.PictureService
}

//NewPictureHandler returns  new UserHandler
func NewPictureHandler(srv *service.PictureService) *PictureHandler {
	return &PictureHandler{picSrv: srv}
}

//GetSinglePicture returns a single user
func (picHdlr *PictureHandler) GetSinglePicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idRaw := ps.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := picHdlr.picSrv.GetSinglePicture(id)
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

//GetManyPictures returns many users
func (picHdlr *PictureHandler) GetManyPictures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idRaw := ps.ByName("placeid")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := picHdlr.picSrv.GetManyPictures(id)
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

//AddPicture adds a new user to the database
func (picHdlr *PictureHandler) AddPicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.Picture{}

	errs := json.Unmarshal(body, user)

	if errs != nil {
		// log.Println("Adding a new user")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		//w.Write([]byte("Unmarshal error"))
		return
	}

	medicine, err := picHdlr.picSrv.AddPicture(user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/host/picture/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

//UpdatePicture updates user
func (picHdlr *PictureHandler) UpdatePicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := picHdlr.picSrv.GetSinglePicture(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &user)

	userUp, err := picHdlr.picSrv.UpdatePicture(user)

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

//DeletePicture deletes a single user
func (picHdlr *PictureHandler) DeletePicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("id ", id)

	errs := picHdlr.picSrv.DeletePicture(id)

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
