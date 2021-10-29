package repo

import (
	"context"
	"github.com/suaas21/graphql-dummy/model"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author model.Author) error
	UpdateAuthor(ctx context.Context, author model.Author) error
	DeleteAuthor(ctx context.Context, id uint) error
}
