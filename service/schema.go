package service

import (
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/service/initialize"
	"github.com/suaas21/graphql-dummy/utils"
	"golang.org/x/net/context"
)

func (ba *BookAuthor) BookAuthorSchema(incomingReq string) (*graphql.Result, error) {
	// loaders...
	var loaders = make(map[string]*dataloader.Loader)
	loaders[utils.BookAuthorIds] = dataloader.NewBatchedLoader(ba.GetAuthorsBatchFn)
	loaders[utils.AuthorBookIds] = dataloader.NewBatchedLoader(ba.GetBooksBatchFn)

	// initialize book, author object...
	bookType, authorType := initialize.GetBookAuthorObject()

	// query and mutation....
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:        bookType,
				Description: "get book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						return ba.bookRepo.GetBookByID(id)
					}
					return nil, nil
				},
			},
			"author": &graphql.Field{
				Type:        authorType,
				Description: "get author by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						return ba.authorRepo.GetAuthorByID(id)
					}
					return nil, nil
				},
			},
		},
	})
	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Description: "create new book",
				Type:        bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"author_ids": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, idOk := p.Args["id"].(string)
					name, nameOk := p.Args["name"].(string)
					description, desOk := p.Args["description"].(string)
					authorIds, authorIdsOk := p.Args["author_ids"].([]interface{})
					book := &model.Book{}
					if idOk {
						book.ID = id
						book.Xkey = id
					}
					if nameOk {
						book.Name = name
					}
					if desOk {
						book.Description = description
					}
					if authorIdsOk {
						var authorIdsStr = make([]string, 0)
						for _, aid := range authorIds {
							authorIdsStr = append(authorIdsStr, aid.(string))
						}
						book.AuthorIDs = authorIdsStr
					}
					err := ba.bookRepo.CreateBook(book)
					if err != nil {
						return nil, err
					}
					return book, nil
				},
			},
			"author": &graphql.Field{
				Description: "create new author",
				Type:        bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"book_ids": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, idOk := p.Args["id"].(string)
					name, nameOk := p.Args["name"].(string)
					bookIds, bookIdsOk := p.Args["book_ids"].([]interface{})
					author := &model.Author{}
					if idOk {
						author.ID = id
						author.Xkey = id
					}
					if nameOk {
						author.Name = name
					}
					if bookIdsOk {
						var bookIdsStr = make([]string, 0)
						for _, bId := range bookIds {
							bookIdsStr = append(bookIdsStr, bId.(string))
						}
						author.BookIDs = bookIdsStr
					}
					err := ba.authorRepo.CreateAuthor(author)
					if err != nil {
						return nil, err
					}
					return author, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		fmt.Printf("Failed to create new service, error: %v\n", err)
	}
	ctx := context.WithValue(context.Background(), "loaders", loaders)

	return graphql.Do(graphql.Params{
		Context:       ctx,
		Schema:        schema,
		RequestString: incomingReq,
	}), nil

}
