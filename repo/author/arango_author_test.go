package author_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/suaas21/graphql-dummy/infra/mocks"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo/author"
	"testing"
)

func TestAuthorArangoRepository_CreateAuthor(t *testing.T) {
	authorData := model.Author{
		ID:      "1",
		Name:    "Dangerous Book",
		BookIDs: nil,
	}

	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("CreateDocument", mock.Anything, mock.AnythingOfType("string"), &authorData).Return(nil)

	arangoRepo := author.NewArangoAuthorRepository(arangoDB, nil)

	if err := arangoRepo.CreateAuthor(context.Background(), authorData); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)
}

func TestAuthorArangoRepository_UpdateAuthor(t *testing.T) {
	authorData := model.Author{
		ID:      "1",
		Name:    "Dangerous Book",
		BookIDs: nil,
	}

	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("UpdateDocument", mock.Anything, mock.AnythingOfType("string"), authorData.ID, &authorData).Return(nil)

	arangoRepo := author.NewArangoAuthorRepository(arangoDB, nil)

	if err := arangoRepo.UpdateAuthor(context.Background(), authorData); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)
}

func TestAuthorArangoRepository_DeleteAuthor(t *testing.T) {
	arangoDB := new(mocks.ArangoDB)
	arangoDB.On("RemoveDocument", mock.Anything, mock.AnythingOfType("string"), "1").Return(nil).Once()

	arangoRepo := author.NewArangoAuthorRepository(arangoDB, nil)

	if err := arangoRepo.DeleteAuthor(context.Background(), 1); err != nil {
		t.Fatal(err)
	}

	arangoDB.AssertExpectations(t)
}
