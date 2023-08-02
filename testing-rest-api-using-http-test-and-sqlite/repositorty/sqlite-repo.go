package repositorty

import (
	"database/sql"
	"log"

	"github.com/kuma-coffee/go-crash-course/testing-rest-api-using-http-test-and-sqlite/entity"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) PostRepository {
	return &sqliteRepo{db}
}
func (u *sqliteRepo) CreateTableSqlit(post *entity.Post) {
	statement, err := u.db.Prepare("CREATE TABLE IF NOT EXIST users (id INTEGER PRIMARY KEY, title VARCHAR(64), text VARCHAR(64))")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table users!")
	}
	statement.Exec()
}

func (u *sqliteRepo) SaveSqlite(post *entity.Post) (*entity.Post, error) {
	statement, _ := u.db.Prepare("INSERT INTO users (title, text) VALUES (?,?)")
	statement.Exec(post.Title, post.Text)
	return post, nil
}

func (u *sqliteRepo) FindAllSqlite() ([]entity.Post, error) {
	var result []entity.Post

	rows, _ := u.db.Query("SELECT * FROM users")
	var post entity.Post
	for rows.Next() {
		rows.Scan(&post.ID, &post.Text, &post.Title)
		result = append(result, post)
	}
	return result, nil
}

func (u *sqliteRepo) DeleteSqlite(id int) error {
	statement, _ := u.db.Prepare("DELETE FROM users WHERE id = ?")
	statement.Exec(id)
	return nil
}
