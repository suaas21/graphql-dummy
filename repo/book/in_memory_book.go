package book

import (
	"context"
	"errors"
	"fmt"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
	"sync"
)

type bookInMemoryRepository struct {
	log logger.StructLogger

	dataStore map[uint]model.Book
	mu        sync.Mutex
}

func NewInMemoryBookRepository(lgr logger.StructLogger) repo.BookRepository {
	dataStore := make(map[uint]model.Book)
	return &bookInMemoryRepository{
		dataStore: dataStore,
		log:       lgr,
	}
}

func (b *bookInMemoryRepository) CreateBook(ctx context.Context, book model.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.dataStore[book.ID]; ok {
		return errors.New(fmt.Sprintf("ID: %v violates foraign key constrains", book.ID))
	}
	b.dataStore[book.ID] = book
	return nil
}

func (b *bookInMemoryRepository) UpdateBook(ctx context.Context, book model.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.dataStore[book.ID]; !ok {
		return errors.New(fmt.Sprintf("ID: %v not found", book.ID))
	}
	b.dataStore[book.ID] = book
	return nil
}

func (b *bookInMemoryRepository) DeleteBook(ctx context.Context, id uint) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.dataStore[id]; !ok {
		return errors.New(fmt.Sprintf("ID: %v not found", id))
	}
	delete(b.dataStore, id)
	return nil
}

func (b *bookInMemoryRepository) GetBook(ctx context.Context, id uint) (*model.Book, error) {
	if book, ok := b.dataStore[id]; !ok {
		return nil, errors.New(fmt.Sprintf("ID: %v not found", id))
	} else {
		return &book, nil
	}
}
