package main

import (
	"net/http"
	"project_4/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/globalsign/mgo"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27018")
	if err != nil {
		panic(err)
	}
	return s
}
