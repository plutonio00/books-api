-- +goose Up
-- +goose StatementBegin
CREATE TABLE `books`
(
    id  INT AUTO_INCREMENT PRIMARY KEY,
    title        VARCHAR(100) NOT NULL,
    release_year INT          NOT NULL,
    author_id    INT          NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `books`;
-- +goose StatementEnd
