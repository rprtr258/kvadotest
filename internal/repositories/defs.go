// Repository for work with books storage definitions
//go:generate mockgen -destination=./mock.go -source=./defs.go -package=repositories

package repositories

import "context"

// Book is a POD of book data
type Book struct {
	Title   string
	Content string
	Authors []string
}

// BooksRepository is a repository that allows to search list of books by author, title or content.
// Also it must be closed as soon as it is no longer needed.
type BooksRepository interface {
	// Search list of books by author
	SearchByAuthor(context.Context, string) ([]Book, error)
	// Search list of books by title
	SearchByTitle(context.Context, string) ([]Book, error)
	// Search list of books by content
	SearchByContent(context.Context, string) ([]Book, error)
	// Close repository
	Close() error
}
