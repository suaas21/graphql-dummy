package book

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

type bookArangoRepository struct {
	db         infra.DB
	collection string
	ctx        context.Context
	log        logger.StructLogger
}

func NewArangoBookRepository(ctx context.Context, db infra.DB, collection string, lgr logger.StructLogger) repo.BookRepository {
	return &bookArangoRepository{
		ctx:        ctx,
		db:         db,
		collection: collection,
		log:        lgr,
	}
}

func (b *bookArangoRepository) CreateBook(book model.Book) error {
	panic("implement me")
}

func (b *bookArangoRepository) UpdateBook(book model.Book) error {
	panic("implement me")
}

func (b *bookArangoRepository) DeleteBook(id uint) error {
	panic("implement me")
}
