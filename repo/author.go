package repo

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
)

type Author struct {
	db  infra.DB
	collection string
	ctx context.Context
	log logger.StructLogger
}

func NewAuthor(ctx context.Context, db infra.DB, collection string, lgr logger.StructLogger) *Author {
	return &Author{
		ctx: ctx,
		db:  db,
		collection: collection,
		log: lgr,
	}
}
