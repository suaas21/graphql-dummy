package schema

import (
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/repo"
)

type BookAuthor struct {
	bookRepo   repo.BookRepository
	authorRepo repo.AuthorRepository
	log        logger.StructLogger
}

func NewBookAuthor(bookRepo repo.BookRepository, authorRepo repo.AuthorRepository, lgr logger.StructLogger) *BookAuthor {
	return &BookAuthor{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
		log:        lgr,
	}
}
