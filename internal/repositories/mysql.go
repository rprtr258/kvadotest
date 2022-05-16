package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

type MysqlBooksRepository struct {
	db *sql.DB
}

func NewMysqlRepository(dataSourceName string) (*MysqlBooksRepository, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &MysqlBooksRepository{
		db: db,
	}, nil
}

func readBookRows(rows *sql.Rows) ([]Book, error) {
	res := make([]Book, 0)
	var (
		authors     []byte
		bookTitle   string
		bookContent string
		authorsList []string
	)
	for rows.Next() {
		rows.Scan(&authors, &bookTitle, &bookContent)
		if err := json.Unmarshal(authors, &authorsList); err != nil {
			return nil, err
		}
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

func (repo MysqlBooksRepository) SearchByAuthor(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
		WITH authors AS (
			SELECT book_id
			FROM book_list
			WHERE author_name LIKE CONCAT('%', ?, '%')
		)
		SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
		FROM book_list
		GROUP BY book_id
		HAVING book_id IN (select * from authors);
	`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}
func (repo MysqlBooksRepository) SearchByTitle(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			WHERE book_title LIKE CONCAT('%', ?, '%')
			GROUP BY book_id;
		`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}
func (repo MysqlBooksRepository) SearchByContent(ctx context.Context, req string) ([]Book, error) {
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = repo.db.QueryContext(ctx, `
			SELECT JSON_ARRAYAGG(author_name) AS authors, book_title, book_content
			FROM book_list
			WHERE book_content LIKE CONCAT('%', ?, '%')
			GROUP BY book_id;
		`, req)
	if err != nil {
		return nil, err
	}
	return readBookRows(rows)
}

func (repo MysqlBooksRepository) Close() error {
	return repo.db.Close()
}
