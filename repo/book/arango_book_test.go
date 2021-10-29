package book_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/suaas21/graphql-dummy/infra/mocks"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo/book"

	"testing"
)

func TestBookArangoRepository_CreateBook(t *testing.T) {
	t.Parallel()

	bookData := model.Book{
		ID:          1,
		Name:        "Dangerous Book",
		Description: "A very dangerous description",
		AuthorIDs:   []uint{1, 2},
	}

	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("CreateDocument", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(nil)
	arangoDB.On("DocumentExists", mock.Anything, "book_author_edges", mock.AnythingOfType("string")).Return(false, nil)

	arangoRepo := book.NewArangoBookRepository(context.Background(), arangoDB, "Book", nil)

	if err := arangoRepo.CreateBook(bookData); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)

	arangoDB.AssertCalled(t, "CreateDocument", mock.Anything, "Book", &bookData)
	arangoDB.AssertCalled(t, "CreateDocument", mock.Anything, "book_author_edges", mock.Anything)
}
