package db

import (
	"context"

	"github.com/ravjotsingh9/discussionForum-Web-Service/schema"
)

type Repository interface {
	Close()
	InsertComment(ctx context.Context, comment schema.Comment) error
	GetComment(ctx context.Context, comment schema.Comment) ([]schema.Comment, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertComment(ctx context.Context, comment schema.Comment) error {
	return impl.InsertComment(ctx, comment)
}

func GetComment(ctx context.Context, comment schema.Comment) ([]schema.Comment, error) {
	return impl.GetComment(ctx, comment)
}
