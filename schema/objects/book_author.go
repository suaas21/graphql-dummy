package objects

import (
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/schema/resolver"
	"github.com/suaas21/graphql-dummy/utils"
	"strings"
)

func LoadSchemaObjects(loaders map[string]*dataloader.Loader, bookType, authorType *graphql.Object) {
	loaders[utils.BookAuthorIds] = dataloader.NewBatchedLoader(resolver.GetAuthorsBatchFn)
	loaders[utils.AuthorBookIds] = dataloader.NewBatchedLoader(resolver.GetBooksBatchFn)

	bookType.AddFieldConfig("authors", &graphql.Field{
		Type: graphql.NewList(authorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var (
				book, bookOk = p.Source.(model.Book)
				loaders      = p.Context.Value("loaders").(map[string]*dataloader.Loader)
				handleErrors = func(errors []error) error {
					if len(errors) == 0 {
						return nil
					}
					var errs []string
					for _, e := range errors {
						errs = append(errs, e.Error())
					}
					return fmt.Errorf(strings.Join(errs, "\n"))
				}
			)
			if !bookOk {
				return nil, nil
			}
			var keys dataloader.Keys
			for i := range book.AuthorIDs {
				keys = append(keys, utils.NewResolverKey(book.AuthorIDs[i]))
			}

			thunk := loaders[utils.BookAuthorIds].LoadMany(p.Context, keys)
			return func() (interface{}, error) {
				res, errors := thunk()
				return res, handleErrors(errors)
			}, nil
		},
	})

	authorType.AddFieldConfig("books", &graphql.Field{
		Type: graphql.NewList(bookType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var (
				author, authorOk = p.Source.(model.Author)
				loaders          = p.Context.Value("loaders").(map[string]*dataloader.Loader)
				handleErrors     = func(errors []error) error {
					if len(errors) == 0 {
						return nil
					}
					var errs []string
					for _, e := range errors {
						errs = append(errs, e.Error())
					}
					return fmt.Errorf(strings.Join(errs, "\n"))
				}
			)
			if !authorOk {
				return nil, nil
			}
			var keys dataloader.Keys
			for i := range author.BookIDs {
				keys = append(keys, utils.NewResolverKey(author.BookIDs[i]))
			}

			thunk := loaders[utils.AuthorBookIds].LoadMany(p.Context, keys)
			return func() (interface{}, error) {
				res, errors := thunk()
				return res, handleErrors(errors)
			}, nil
		},
	})
}
