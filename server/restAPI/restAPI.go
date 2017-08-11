package restAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

//.Data

//User Struct holds all information about users"
type User struct {
	Name       string   `json:"user_Name"`
	ID         string   `json:"user_ID"`
	Pass       string   `json:"user_Pass"`
	AvatarURL  string   `json:"avatar_URL"`
	Songs      []string `json:"user_Songs"`
	Comments   []string `json:"user_comments"`
	Recordings []string `json:"user_recordings"`
}

//UserController holds the session value
type UserController struct {
	sssion *mgo.Session
}

//Run Starts Go REST API server
func Run() {

	go func() {
		fmt.Print("goRoutine Started!")
		r := httprouter.New()
		r.POST("/user", createUser)
		http.ListenAndServe("localhost:8080", r)

	}()
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := new(User)
	json.NewDecoder(r.Body).Decode(&u)

}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
