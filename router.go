package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/Nithinparam/CleanArchitecture/entity"
	"github.com/Nithinparam/CleanArchitecture/repo"
)

var (
	repos repo.PostRepo = repo.NewPostRepo()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	posts, err := repos.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting postg"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	var post entity.Post = entity.Post{Title: "Title 233", Text: "Text 3332"}
	post.ID = rand.Int63()
	repos.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}
