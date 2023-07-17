package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	db "github.com/zubairmh/blapi/internal/database"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meta_id := r.Header.Get("meta_id")
	user := db.CreateUser(meta_id, meta_id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func createPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	body, _ := ioutil.ReadAll(r.Body)
	var post db.Post
	json.Unmarshal(body, &post)
	// post := db.CreateBounty(
	// 	r.Form.Get("meta_id"),
	// 	bounty,
	// 	r.Form.Get("title"),
	// 	r.Form.Get("description"),
	// 	time.Now(),
	// )
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts := db.GetBounties()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
