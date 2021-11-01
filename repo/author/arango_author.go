package author

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

type authorArangoRepository struct {
	db  infra.DB
	ctx context.Context
	log logger.StructLogger
}

func NewArangoAuthorRepository(ctx context.Context, db infra.DB, lgr logger.StructLogger) repo.AuthorRepository {
	return &authorArangoRepository{
		ctx: ctx,
		db:  db,
		log: lgr,
	}
}

func (a *authorArangoRepository) CreateAuthor(author *model.Author) error {
	return a.db.CreateDocument(a.ctx, author)
}

func (a *authorArangoRepository) GetAuthorByID(key string) (*model.Author, error) {
	var author model.Author
	err := a.db.ReadDocument(a.ctx, key, &author)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *authorArangoRepository) ListAuthorByIDs(keys []string) ([]model.Author, error) {
	authors := make([]model.Author, len(keys))
	err := a.db.ReadDocuments(a.ctx, keys, authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *authorArangoRepository) UpdateAuthor(author model.Author) error {
	panic("implement me")
}

func (a *authorArangoRepository) DeleteAuthor(id uint) error {
	panic("implement me")
}
