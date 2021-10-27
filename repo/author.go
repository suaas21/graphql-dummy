package repo

import (
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"golang.org/x/net/context"
)

type Author struct {
	db  infra.DB
	ctx context.Context
	log logger.StructLogger
}

func NewAuthor(ctx context.Context, db infra.DB, lgr logger.StructLogger) *Author {
	return &Author{
		ctx: ctx,
		db:  db,
		log: lgr,
	}
}

func (a *Author) CreateAuthorDocument(author *model.Author) error {
	return a.db.CreateDocument(a.ctx, author)
}

func (a *Author) GetAuthorDocument(key string) (*model.Author, error) {
	var author model.Author
	err := a.db.ReadDocument(a.ctx, key, &author)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *Author) GetAuthorDocuments(keys []string) ([]model.Author, error) {
	authors := make([]model.Author, len(keys))
	err := a.db.ReadDocuments(a.ctx, keys, authors)
	if err != nil {
		return nil, err
	}
	return authors, nil
}