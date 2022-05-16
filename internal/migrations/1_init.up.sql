CREATE TABLE author (
    id SERIAL,
    name VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE book (
    id SERIAL,
    title VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE authoring (
    author_id BIGINT UNSIGNED NOT NULL,
    book_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author(id),
    FOREIGN KEY (book_id) REFERENCES book(id),
    UNIQUE (author_id, book_id)
);

CREATE VIEW book_list AS
SELECT
    author.id AS author_id,
    author.name AS author_name,
    book.id AS book_id,
    book.title AS book_title,
    book.content AS book_content
FROM authoring
JOIN author ON authoring.author_id=author.id
JOIN book ON authoring.book_id=book.id;