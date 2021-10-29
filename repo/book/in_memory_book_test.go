package book_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo/book"
	"testing"
)

func TestBookInMemoryRepository_CRUD(t *testing.T) {
	bookData := model.Book{
		ID:          1,
		Name:        "Dangerous Book",
		Description: "A very dangerous description",
		AuthorIDs:   []uint{1, 2},
	}

	inMemoryBookRepo := book.NewInMemoryBookRepository(nil)

	t.Run("create", func(t *testing.T) {
		// should not return error
		err := inMemoryBookRepo.CreateBook(context.Background(), bookData)
		assert.NoError(t, err)

		// duplicate create, should return error
		err = inMemoryBookRepo.CreateBook(context.Background(), bookData)
		assert.Error(t, err)
	})

	t.Run("update", func(t *testing.T) {
		bookData.Name = "A Good Book"

		// should not return error
		err := inMemoryBookRepo.UpdateBook(context.Background(), bookData)
		assert.NoError(t, err)

		nonExistingBookData := model.Book{
			ID:          2,
			Name:        "A Damn Good Book",
			Description: "A Short Description about the book",
			AuthorIDs:   []uint{1, 2},
		}

		// nonExistingBookData doesn't exist, should return error
		err = inMemoryBookRepo.UpdateBook(context.Background(), nonExistingBookData)
		assert.Error(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		// should not return error
		err := inMemoryBookRepo.DeleteBook(context.Background(), 1)
		assert.NoError(t, err)

		// duplicate delete, should return error
		err = inMemoryBookRepo.DeleteBook(context.Background(), 1)
		assert.Error(t, err)
	})
}
