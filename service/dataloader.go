package service

import (
	"github.com/graph-gophers/dataloader"
	"golang.org/x/net/context"
	"log"
)

func (ba *BookAuthor) GetAuthorsBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}

	authors, err := ba.authorRepo.ListAuthorByIDs(keys.Keys())
	if err != nil {
		handleError(err)
	}

	var results []*dataloader.Result
	for _, author := range authors {
		result := dataloader.Result{
			Data:  author,
			Error: nil,
		}
		results = append(results, &result)
	}
	log.Printf("[GetAuthorsBatchFn] batch size: %d", len(results))
	return results
}

func (ba *BookAuthor) GetBooksBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}

	books, err := ba.bookRepo.ListBookByIDs(keys.Keys())
	if err != nil {
		handleError(err)
	}

	var results []*dataloader.Result
	for _, book := range books {
		result := dataloader.Result{
			Data:  book,
			Error: nil,
		}
		results = append(results, &result)
	}
	log.Printf("[GetBooksBatchFn] batch size: %d", len(results))
	return results
}
