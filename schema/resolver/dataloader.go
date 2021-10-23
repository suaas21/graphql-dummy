package resolver

import (
	"context"
	"errors"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/model"
	"log"
	"strconv"
)

func GetAuthorsBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}

	var authorIDs []uint
	for _, key := range keys {
		id, err := strconv.ParseUint(key.String(), 10, 32)
		if err != nil {
			return handleError(err)
		}
		authorIDs = append(authorIDs, uint(id))
	}

	var authorsMap = make(map[uint]model.Author, len(authorIDs))
	for _, author := range infra.ListAuthor {
		for _, authorID := range authorIDs {
			if author.ID == authorID {
				authorsMap[authorID] = author
			}
		}
	}

	var results []*dataloader.Result
	for _, authorID := range authorIDs {
		author, ok := authorsMap[authorID]
		if !ok {
			err := errors.New(fmt.Sprintf("author not found, "+
				"author_id: %d", authorID))
			return handleError(err)
		}
		result := dataloader.Result{
			Data:  author,
			Error: nil,
		}
		results = append(results, &result)
	}
	log.Printf("[GetAuthorsBatchFn] batch size: %d", len(results))
	return results
}

func GetBooksBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}

	var bookIDs []uint
	for _, key := range keys {
		id, err := strconv.ParseUint(key.String(), 10, 32)
		if err != nil {
			return handleError(err)
		}
		bookIDs = append(bookIDs, uint(id))
	}

	var booksMap = make(map[uint]model.Book, len(bookIDs))
	for _, book := range infra.ListBook {
		for _, bookID := range bookIDs {
			if book.ID == bookID {
				booksMap[bookID] = book
			}
		}
	}

	var results []*dataloader.Result
	for _, bookID := range bookIDs {
		book, ok := booksMap[bookID]
		if !ok {
			err := errors.New(fmt.Sprintf("book not found, "+
				"book_id: %d", bookID))
			return handleError(err)
		}
		result := dataloader.Result{
			Data:  book,
			Error: nil,
		}
		results = append(results, &result)
	}
	log.Printf("[GetBooksBatchFn] batch size: %d", len(results))
	return results
}
