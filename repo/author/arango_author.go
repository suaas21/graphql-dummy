package author

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

type authorArangoRepository struct {
	db         infra.DB
	collection string
	ctx        context.Context
	log        logger.StructLogger
}

func NewArangoAuthorRepository(ctx context.Context, db infra.DB, collection string, lgr logger.StructLogger) repo.AuthorRepository {
	return &authorArangoRepository{
		ctx:        ctx,
		db:         db,
		collection: collection,
		log:        lgr,
	}
}

func (a *authorArangoRepository) CreateAuthor(author model.Author) error {
	panic("implement me")
}

func (a *authorArangoRepository) UpdateAuthor(author model.Author) error {
	panic("implement me")
}

func (a *authorArangoRepository) DeleteAuthor(id uint) error {
	panic("implement me")
}
