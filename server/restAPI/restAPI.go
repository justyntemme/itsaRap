package restAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"

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
	Title         string          `json:"post_Title" bson:"post_Title"`
	PostID        bson.ObjectId   `json:"post_Id" bson:"post_Id"`
	Comments      []bson.ObjectId `json:"comment" bson:"comments"`
	SoundcloudURL string          `json:"soundcloud_URL" bson:"soundcloud_url"`
}

///LyricReply struct holds the information for all recordings to user recording reply
type LyricReply struct {
	UserID       bson.ObjectId `json:"user_Id" bson:"user_Id"`
	FileID       bson.ObjectId `json:"file_Id" bson:"file_Id"`
	FilePath     string        `json:"file_Path" bson:"file_Path"`
	Comments     []string      `json:"comments" bson:"comments"`
	Username     string        `json:"username"  bson:"username"`
	ParentPostID bson.ObjectId `json:"parent_Post_Id" bson:"parent_Post_Id"`
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
	r.POST("/login", uc.Login)
	r.GET("/posts", uc.GetAllPosts)
	r.POST("/lreply", uc.uploadReply)

	fmt.Println(http.ListenAndServe("localhost:8080", handler))

}

func optionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Param) {
}

//GetReplys shows replys to a spacific post and returns json serialized string TODO
func (us UserController) GetReplys(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

//uploadReply uploads user reply to an original post TODO
func (uc UserController) uploadReply(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Well we at least hit the handler...")
	l := LyricReply{}
	fmt.Println(r.MultipartForm.File)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadDir := path.Join(currentDirectory(), "uploads")
	filename := path.Join(uploadDir, handler.Filename)
	outfile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()

	if _, err = io.Copy(outfile, file); err != nil {
		fmt.Println(err)
		return

		fmt.Println(filename)
	}

	l.FileID = bson.NewObjectId()
	l.FilePath = filename
	uc.session.DB("its-a-rap-db").C("lreply").Insert(l)
	//w.Header().Set("Content-type", "application-json")
	//w.Header().Set("Access-Control-Allow-Origin", "true")
	fmt.Fprintf(w, "200OK")

}

// Returns the current directory we are running in.
func currentDirectory() string {

	// Locate the current directory for the site.
	_, fn, _, _ := runtime.Caller(1)
	return path.Dir(fn)
}

//GetAllPosts retreives all new posts from mongodb
func (uc UserController) GetAllPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	posts := []Ipost{}
	_ = uc.session.DB("its-a-rap-db").C("iposts").Find(bson.M{}).All(&posts)
	fmt.Println(posts)

	pj, err := json.Marshal(posts)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-type", "application-json")
	fmt.Fprintf(w, "%s\n", pj)
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
func (uc UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}
	result := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	err := uc.session.DB("its-a-rap-db").C("users").Find(bson.M{"user_name": u.Name}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	if result.Pass == u.Pass {
		rj, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "%s\n", rj)
	} else {
		fmt.Println("incorrect user/pass: user Not Found")
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
	result := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)

	//Create bson ID
	u.ID = bson.NewObjectId()

	//store the user in mongodb
	err := uc.session.DB("its-a-ra-db").C("users").Find(bson.M{"user_name": u.Name}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	if result.Name == u.Name {
		w.WriteHeader(http.StatusBadRequest)
	} else {
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
