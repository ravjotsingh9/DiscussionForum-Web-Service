package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ravjotsingh9/discussionForum-Web-Service/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertComment(ctx context.Context, comment schema.Comment) error {
	_, err := r.db.Exec("INSERT INTO comment(id, content, pid,  tid) VALUES($1, $2, $3, $4)", comment.ID, comment.Content, comment.PID, comment.TID)
	return err
}

func (r *PostgresRepository) GetComment(ctx context.Context, comment schema.Comment) ([]schema.Comment, error) {

	rows, err := r.db.Query("SELECT * FROM comment where tid='" + comment.ID + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of User
	comments := []schema.Comment{}
	for rows.Next() {
		comment := schema.Comment{}
		if err = rows.Scan(&comment.ID, &comment.Content, &comment.PID, &comment.TID); err == nil {
			comments = append(comments, comment)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
