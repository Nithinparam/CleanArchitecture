package service

import (
	"errors"
	"math/rand"

	"github.com/Nithinparam/CleanArchitecture/entity"
	"github.com/Nithinparam/CleanArchitecture/repo"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repos repo.PostRepo = repo.NewFirestoreRepo()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The Post is Empty")
		return err
	}
	if post.Title == "" {
		return errors.New("The post title is empty")
	}
	return nil
}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repos.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repos.FindAll()
}
