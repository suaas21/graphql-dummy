package book_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/suaas21/graphql-dummy/infra/mocks"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo/book"

	"testing"
)

func TestBookArangoRepository_CreateBook(t *testing.T) {
	bookData := model.Book{
		ID:          1,
		Name:        "Dangerous Book",
		Description: "A very dangerous description",
		AuthorIDs:   []uint{1, 2},
	}

	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("CreateDocument", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(nil)
	arangoDB.On("DocumentExists", mock.Anything, "book_author_edges", mock.AnythingOfType("string")).Return(false, nil)

	arangoRepo := book.NewArangoBookRepository(arangoDB, "Book", nil)

	if err := arangoRepo.CreateBook(context.Background(), bookData); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)

	arangoDB.AssertCalled(t, "CreateDocument", mock.Anything, "Book", &bookData)
	arangoDB.AssertCalled(t, "CreateDocument", mock.Anything, "book_author_edges", mock.Anything)
}

func TestBookArangoRepository_UpdateBook(t *testing.T) {
	bookData := model.Book{
		ID:          1,
		Name:        "Dangerous Book",
		Description: "A very dangerous description",
		AuthorIDs:   []uint{1, 2},
	}

	arangoDB := new(mocks.ArangoDB)

	arangoDB.On("CreateDocument", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(nil)
	arangoDB.On("UpdateDocument", mock.Anything, mock.AnythingOfType("string"), mock.Anything, mock.Anything).Return(nil)
	arangoDB.On("DocumentExists", mock.Anything, "book_author_edges", mock.AnythingOfType("string")).Return(false, nil)

	arangoRepo := book.NewArangoBookRepository(arangoDB, "Book", nil)

	if err := arangoRepo.UpdateBook(context.Background(), bookData); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertCalled(t, "UpdateDocument", mock.Anything, "Book", fmt.Sprintf("%d", bookData.ID), &bookData)
	arangoDB.AssertNumberOfCalls(t, "UpdateDocument", 1)
	arangoDB.AssertCalled(t, "CreateDocument", mock.Anything, "book_author_edges", mock.Anything)

	arangoDB.AssertExpectations(t)
}

func TestBookArangoRepository_DeleteBook(t *testing.T) {
	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("RemoveDocument", mock.Anything, "Book", "1").Return(nil).Once()

	arangoRepo := book.NewArangoBookRepository(arangoDB, "Book", nil)

	if err := arangoRepo.DeleteBook(context.Background(), 1); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)
}
