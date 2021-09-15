-- +goose Up
-- +goose StatementBegin
ALTER TABLE `books`
ADD CONSTRAINT fk_book_author
foreign key (author_id)
references authors (id)
ON UPDATE RESTRICT
ON DELETE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `books`
DROP FOREIGN KEY `fk_book_author`
-- +goose StatementEnd
