package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project_4/models"
	"github.com/julienschmidt/httprouter"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
 return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err!=nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err !=nil{
		w.WriteHeader(500)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err!=nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete user", oid, "\n")
}