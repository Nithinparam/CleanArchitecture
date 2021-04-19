package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Nithinparam/CleanArchitecture/entity"
	"github.com/Nithinparam/CleanArchitecture/error"
	"github.com/Nithinparam/CleanArchitecture/service"
)

var (
	serve service.PostService
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPosts(w http.ResponseWriter, r *http.Request)
}
type controller struct{}

func NewPostController(service service.PostService) PostController {
	serve = service
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	posts, err := serve.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error getting the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application:json")
	var post entity.Post = entity.Post{Title: "Title 233", Text: "Text 3332"}
	err1 := serve.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: err1.Error()})
		return
	}
	res, err2 := serve.Create(&post)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: err2.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
