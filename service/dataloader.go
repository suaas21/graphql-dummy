package service

import (
	"fmt"
	"github.com/graph-gophers/dataloader"
	"golang.org/x/net/context"
	"log"
	"strings"
)

func (ba *BookAuthor) GetAuthorsBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	handleError := func(err error) []*dataloader.Result {
		var results []*dataloader.Result
		var result dataloader.Result
		result.Error = err
		results = append(results, &result)
		return results
	}

	query := fmt.Sprintf(`FOR x IN Author FILTER %s RETURN x`, CommonFilter(keys.Keys()))
	authors, err := ba.authorRepo.QueryAuthors(ctx, query, nil)
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

	query := fmt.Sprintf(`FOR x IN Book FILTER %s RETURN x`, CommonFilter(keys.Keys()))
	books, err := ba.bookRepo.QueryBooks(ctx, query, nil)
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

func CommonFilter(keys []string) string {
	var in []string
	for _, key := range keys {
		in = append(in, fmt.Sprintf("CONTAINS(x.id, %s)", key))
	}
	return strings.Join(in, " AND ")
}
