package initialize

import (
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/utils"
	"strings"
)

func GetBookAuthorObject() (*graphql.Object, *graphql.Object) {
	// book and author objects....
	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	// Add extra field for book and author....
	AddBookAuthorFieldConfig(bookType, authorType)

	return bookType, authorType
}

func AddBookAuthorFieldConfig(bookType, authorType *graphql.Object) {
	bookType.AddFieldConfig("authors", &graphql.Field{
		Type: graphql.NewList(authorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var (
				book, bookOk = p.Source.(*model.Book)
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
				author, authorOk = p.Source.(*model.Author)
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
