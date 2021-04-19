package repo

import (
	"github.com/Nithinparam/CleanArchitecture/entity"
)

type repos struct{}

func NewMysqlRepo() PostRepo {
	return &repos{}
}

func (*repos) Save(post *entity.Post) (*entity.Post, error) {
	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, projectId)
	// if err != nil {
	// 	log.Fatalf("Failed to cretae a forestore client:%v", err)
	// 	return nil, err
	// }
	// defer client.Close()
	// _, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
	// 	"ID":    post.ID,
	// 	"Title": post.Title,
	// 	"Text":  post.Text,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, err
	// }
	return post, nil
}

func (*repos) FindAll() ([]entity.Post, error) {
	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, projectId)
	// if err != nil {
	// 	log.Fatalf("Failed to cretae a forestore client:%v", err)
	// 	return nil, err
	// }
	// defer client.Close()
	// var posts []entity.Post
	// iter := client.Collection(collectionName).Documents(ctx)
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Failed to cretae a forestore client:%v", err)
	// 		return nil, err
	// 	}
	// 	post := entity.Post{
	// 		ID:    doc.Data()["ID"].(int64),
	// 		Title: doc.Data()["Title"].(string),
	// 		Text:  doc.Data()["Text"].(string),
	// 	}
	// 	posts = append(posts, post)
	// }
	return nil, nil
}
