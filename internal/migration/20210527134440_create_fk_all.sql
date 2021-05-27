-- +goose Up
-- +goose StatementBegin
ALTER TABLE `book`
ADD CONSTRAINT fk_book_author
foreign key (author_id)
references author (id)
ON UPDATE RESTRICT
ON DELETE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `book`
DROP FOREIGN KEY `fk_book_author`
-- +goose StatementEnd
