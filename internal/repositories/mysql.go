// Mysql books repository implementation
package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MysqlBooksRepository is a database repository struct implementing BooksRepository, holding database connection
type MysqlBooksRepository struct {
	db *sql.DB
}

// NewMysqlRepository connects to database and create new repository from it.
// If connection fails it retries up to 10 times. Then it runs migrations to update database schema
func NewMysqlRepository(dataSourceName string) (*MysqlBooksRepository, error) {
	// Connect to database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Retry if couldn't ping database
	ctx := context.Background()
	const maxAttempts = 10
	attempts := 0
	for {
		log.Printf("Trying to connect database, attempt #%d", attempts)
		if err = db.PingContext(ctx); err != nil {
			log.Printf("Failed to connect database, attempt #%d: %v", attempts, err)
			attempts++
			if attempts > maxAttempts {
				return nil, err
			}
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	// Migrate database schema
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://./internal/migrations", "mysql", driver)
	if err != nil {
		return nil, err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	// Create books repository
	return &MysqlBooksRepository{
		db: db,
	}, nil
}

// readBookRows reads list of books from database response into Book slice
func readBookRows(rows *sql.Rows) ([]Book, error) {
	var (
		authors     []byte
		bookTitle   string
		bookContent string
		authorsList []string
	)
	res := make([]Book, 0)
	for rows.Next() {
		// Scan another row
		rows.Scan(&authors, &bookTitle, &bookContent)
		// Authors list is JSON list, unmarshall it
		if err := json.Unmarshal(authors, &authorsList); err != nil {
			return nil, err
		}
		// Add to result
		res = append(res, Book{
			Authors: authorsList,
			Title:   bookTitle,
			Content: bookContent,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

// query runs SQL query that returns books list using req parameter
func (repo MysqlBooksRepository) query(ctx context.Context, req, sqlQuery string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, sqlQuery, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}

// SearchByAuthor searches all books with author names containing req as substring
func (repo MysqlBooksRepository) SearchByAuthor(ctx context.Context, req string) ([]Book, error) {
	return repo.query(ctx, req, `
		WITH authors AS (
			SELECT book_id
			FROM book_list
			WHERE author_name LIKE CONCAT('%', ?, '%')
		)
		SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
		FROM book_list
		GROUP BY book_id
		HAVING book_id IN (select * from authors);
	`)
}

// SearchByTitle searches all books with title containing req as substring
func (repo MysqlBooksRepository) SearchByTitle(ctx context.Context, req string) ([]Book, error) {
	return repo.query(ctx, req, `
		SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
		FROM book_list
		WHERE book_title LIKE CONCAT('%', ?, '%')
		GROUP BY book_id;
	`)
}

// SearchByContent searches all books with content containing req as substring
func (repo MysqlBooksRepository) SearchByContent(ctx context.Context, req string) ([]Book, error) {
	return repo.query(ctx, req, `
		SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
		FROM book_list
		WHERE book_content LIKE CONCAT('%', ?, '%')
		GROUP BY book_id;
	`)
}

// Close closes database conneection
func (repo MysqlBooksRepository) Close() error {
	return repo.db.Close()
}
