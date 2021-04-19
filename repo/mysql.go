package repo

import (
	"database/sql"

	"github.com/Nithinparam/CleanArchitecture/entity"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type repos struct{}

func NewMysqlRepo() PostRepo {
	db, _ = sql.Open("mysql", "root:param356@tcp(localhost:3306)/myblog?charset=utf8")
	// defer db.Close()
	return &repos{}
}

func (*repos) Save(post *entity.Post) (*entity.Post, error) {
	_, err := db.Exec(`INSERT INTO posts VALUES (?,?,?)`, post.ID, post.Title, post.Text)

	return post, err
}

func (*repos) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	rows, err := db.Query(`select * from posts`)
	defer rows.Close()
	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.ID, &post.Title, &post.Text)
		posts = append(posts, post)
	}
	return posts, err
}
