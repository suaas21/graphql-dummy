package author

import (
	"context"
	"fmt"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

type authorArangoRepository struct {
	db         infra.ArangoDB
	collection string
	log        logger.StructLogger
}

func NewArangoAuthorRepository(db infra.ArangoDB, collection string, lgr logger.StructLogger) repo.AuthorRepository {
	return &authorArangoRepository{
		db:         db,
		collection: collection,
		log:        lgr,
	}
}

func (a *authorArangoRepository) CreateAuthor(ctx context.Context, author model.Author) error {
	return a.db.CreateDocument(ctx, a.collection, &author)
}

func (a *authorArangoRepository) UpdateAuthor(ctx context.Context, author model.Author) error {
	return a.db.UpdateDocument(ctx, a.collection, fmt.Sprintf("%d", author.ID), &author)
}

func (a *authorArangoRepository) DeleteAuthor(ctx context.Context, id uint) error {
	return a.db.RemoveDocument(ctx, a.collection, fmt.Sprintf("%d", id))
}
