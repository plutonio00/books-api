-- +goose Up
-- +goose StatementBegin
CREATE TABLE `author`
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(40) NOT NULL,
    last_name  VARCHAR(40) NOT NULL,
    is_male    TINYINT(1)   NOT NULL,
    birth_date DATE    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `author`;
-- +goose StatementEnd
