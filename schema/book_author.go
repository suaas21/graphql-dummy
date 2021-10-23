package schema

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/schema/objects"
)

func (ba *BookAuthor) BookAuthorSchema(incomingReq string) (*graphql.Result, error) {
	// loaders...
	var loaders = make(map[string]*dataloader.Loader)

	// book and author objects....
	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
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
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	objects.LoadSchemaObjects(loaders, bookType, authorType)

	// query and mutation....
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:        bookType,
				Description: "get book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, book := range infra.ListBook {
							if book.ID == uint(id) {
								return book, nil
							}
						}
					}
					return nil, nil
				},
			},
			"author": &graphql.Field{
				Type:        authorType,
				Description: "get author by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, author := range infra.ListAuthor {
							if author.ID == uint(id) {
								return author, nil
							}
						}
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
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"author_ids": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.Int)),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, idOk := p.Args["id"].(int)
					name, nameOk := p.Args["name"].(string)
					description, desOk := p.Args["description"].(string)
					authorIds, authorIdsOk := p.Args["author_ids"].([]interface{})
					book := model.Book{}
					if idOk {
						book.ID = uint(id)
					}
					if nameOk {
						book.Name = name
					}
					if desOk {
						book.Description = description
					}
					if authorIdsOk {
						var auids = make([]uint, 0)
						for _, aid := range authorIds {
							auids = append(auids, uint(aid.(int)))
						}
						book.AuthorIDs = auids
					}
					infra.ListBook = append(infra.ListBook, book) //for save data in memory
					//err := ba.bookRepo.CreateBookDocument(book)
					//if err != nil {
					//	return nil, err
					//}
					return book, nil
				},
			},
			"author": &graphql.Field{
				Description: "create new author",
				Type:        bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"book_ids": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, idOk := p.Args["id"].(int)
					name, nameOk := p.Args["name"].(string)
					bookIds, bookIdsOk := p.Args["book_ids"].([]interface{})
					author := model.Author{}
					if idOk {
						author.ID = uint(id)
					}
					if nameOk {
						author.Name = name
					}
					if bookIdsOk {
						var auids = make([]uint, 0)
						for _, aid := range bookIds {
							auids = append(auids, uint(aid.(int)))
						}
						author.BookIDs = auids
					}
					infra.ListAuthor = append(infra.ListAuthor, author)
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
		fmt.Printf("Failed to create new schema, error: %v\n", err)
	}
	ctx := context.WithValue(context.Background(), "loaders", loaders)

	return graphql.Do(graphql.Params{
		Context:       ctx,
		Schema:        schema,
		RequestString: incomingReq,
	}), nil

}
