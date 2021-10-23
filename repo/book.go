package repo

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
)

type Book struct {
	db  infra.DB
	collection string
	ctx context.Context
	log logger.StructLogger
}

func NewBook(ctx context.Context, db infra.DB, collection string, lgr logger.StructLogger) *Book {
	return &Book{
		ctx: ctx,
		db:  db,
		collection: collection,
		log: lgr,
	}
}

//func (b *Book) CreateBookDocument(book model.Book) error {
//	return b.db.CreateDocument(b.ctx, b.collection, book)
//}
