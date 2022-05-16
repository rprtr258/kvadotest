package repositories

import "context"

type Book struct {
	Title   string
	Content string
	Authors []string
}

type BooksRepository interface {
	SearchByAuthor(context.Context, string) ([]Book, error)
	SearchByTitle(context.Context, string) ([]Book, error)
	SearchByContent(context.Context, string) ([]Book, error)
	Close() error
}
