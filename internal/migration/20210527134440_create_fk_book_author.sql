-- +goose Up
-- +goose StatementBegin
ALTER TABLE `books`
ADD CONSTRAINT fk_authors_books
foreign key (author_id)
references authors (id)
ON UPDATE RESTRICT
ON DELETE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `books`
DROP FOREIGN KEY `fk_authors_books`
-- +goose StatementEnd
