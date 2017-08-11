package restAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//.Data

//User Struct holds all information about users"
type User struct {
	Name       string        `json:"user_Name" bson:"user_name"`
	Id         bson.ObjectId `json:"user_ID" bson:"User_ID"`
	Pass       string        `json:"user_Pass" bson:"user_Pass"`
	AvatarURL  string        `json:"avatar_URL" bson:"avatar_URL"`
	Songs      []string      `json:"user_Songs" bson:"user_Songs"`
	Comments   []string      `json:"user_comments" bson:"user_Comments"`
	Recordings []string      `json:"user_recordings" bson:"user_Recordings"`
}

//UserController holds the session value
type UserController struct {
	session *mgo.Session
}

//Run Starts Go REST API server
func Run() {

	uc := NewUserController(getSession())
	r := httprouter.New()
	r.POST("/user", uc.CreateUser)
	fmt.Println(http.ListenAndServe("localhost:8080", r))

}

//CreateUser creates a user entry in the mongoDB
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u.Name)

	//Create bson ID
	u.Id = bson.NewObjectId()

	//store the user in mongodb
	uc.session.DB("its-a-rap-db").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)

}

//NewUserController creates a new mongo UserController with session embedded
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

//getSessino connects to mongodb and returns session ID
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	//Check if connection err, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
