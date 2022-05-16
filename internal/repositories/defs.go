// Repository for work with books storage definitions
//go:generate mockgen -destination=./mock.go -source=./defs.go -package=repositories

package repositories

import "context"

// Book POD
type Book struct {
	Title   string
	Content string
	Authors []string
}

// Repository that allows to search list of books by author, title or content.
// Also close when it is no longer needed.
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
