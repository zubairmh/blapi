package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	db "github.com/zubairmh/blapi/internal/database"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := db.CreateUser("xyz", "walter")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func createPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	post := db.CreateBounty(
		"xyz",
		1.4,
		"Rust Developer Needed",
		"Need aspiring rust developer for big project",
		time.Now(),
	)
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
