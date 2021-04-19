package repo

import "github.com/Nithinparam/CleanArchitecture/entity"

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
