package restAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//.Data

//User Struct holds all information about users"
type User struct {
	Name       string        `json:"user_Name" bson:"user_name"`
	ID         bson.ObjectId `json:"user_ID" bson:"User_ID"`
	Pass       string        `json:"user_Pass" bson:"user_Pass"`
	AvatarURL  string        `json:"avatar_URL" bson:"avatar_URL"`
	Songs      []string      `json:"user_Songs" bson:"user_Songs"`
	Comments   []string      `json:"user_comments" bson:"user_Comments"`
	Recordings []string      `json:"user_recordings" bson:"user_Recordings"`
}

//Ipost struct is the structure that includes the data for all posts
type Ipost struct {
	UserID        bson.ObjectId   `json:"user_Id" bson:"user_Id"`
	PostID        bson.ObjectId   `json:"post_Id" bson:"post_Id"`
	Comments      []bson.ObjectId `json:"comment" bson:"comments"`
	SoundcloudURL string          `json:"soundcloud_url" bson:"soundcloud_url"`
}

//UserController holds the session value
type UserController struct {
	session *mgo.Session
}

//Run Starts Go REST API server
func Run() {

	uc := NewUserController(getSession())
	r := httprouter.New()
	handler := cors.Default().Handler(r)
	r.POST("/user", uc.CreateUser)
	r.POST("/posts", uc.CreateIPost)

	fmt.Println(http.ListenAndServe("localhost:8080", handler))

}

func optionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Param) {
}

//GetUser grabs user by id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//grab id
	id := p.ByName("id")
	//Verify id is ObjectId hex rep

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	oid := bson.ObjectIdHex(id)

	//composite literal
	u := User{}

	//Fetch user by id
	if err := uc.session.DB("its-a-rap-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	//marshal provided interface
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

//Login compares hashed passwords to username and returns complete user info if correct
func (uc UserController) Login(w http.ResponseWriter, r *http.Response, p httprouter.Params) {
	u := User{}
	result := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	err := uc.session.DB("its-a-ra-db").C("users").Find(bson.M{"user_name": u.Name}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	if result.Pass == u.Pass {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", result)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
	}
}

//CreateIPost creates new Ipost entry in the mongoDB
func (uc UserController) CreateIPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ip := Ipost{}

	json.NewDecoder(r.Body).Decode(&ip)
	fmt.Println(ip)

	//create bson id
	ip.PostID = bson.NewObjectId()

	//store the post in mongoDB
	uc.session.DB("its-a-rap-db").C("iposts").Insert(ip)

	ipj, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", ipj)
}

//CreateUser creates a user entry in the mongoDB
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	//Create bson ID
	u.ID = bson.NewObjectId()

	//store the user in mongodb
	uc.session.DB("its-a-rap-db").C("users").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Content-Type", "application/json")
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
