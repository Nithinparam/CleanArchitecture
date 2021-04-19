package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Nithinparam/CleanArchitecture/entity"
	"github.com/Nithinparam/CleanArchitecture/service"
)

var (
	serve service.PostService = service.NewPostService()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	posts, err := serve.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	var post entity.Post = entity.Post{Title: "Title 233", Text: "Text 3332"}
	err := serve.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting postg"))
		return
	}
	serve.Create(&post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}
