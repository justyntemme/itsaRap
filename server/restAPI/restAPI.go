package restAPI

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	//json.NewDecoder(r.Body).Decode(//Right Here put user struct)

}
