package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/users/service"
	"github.com/julienschmidt/httprouter"
)

//UserHandler ..
type UserHandler struct {
	usrSrv *service.UserServices
}

//NewUserHandler returns  new UserHandler
func NewUserHandler(srv *service.UserServices) *UserHandler {
	return &UserHandler{usrSrv: srv}
}

//GetSingleUser returns a single user
func (usrHdlr *UserHandler) GetSingleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idRaw := ps.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, err := usrHdlr.usrSrv.GetSingleUser(id)
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

//GetManyUsers returns many users
func (usrHdlr *UserHandler) GetManyUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, err := usrHdlr.usrSrv.GetManyUsers()
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

//AddUser adds a new user to the database
func (usrHdlr *UserHandler) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := entity.User{}

	errs := json.Unmarshal(body, user)

	if errs != nil {
		log.Println("Adding a new user")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, err := usrHdlr.usrSrv.AddUser(&user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/host/user/%d", medicine.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

//UpdateUser updates user
func (usrHdlr *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	medicine, errs := usrHdlr.usrSrv.GetSingleUser(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &medicine)

	medicine, errs = usrHdlr.usrSrv.UpdateUser(medicine)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(medicine, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//DeleteUser deletes a single user
func (usrHdlr *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	errs := usrHdlr.usrSrv.DeleteUser(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
