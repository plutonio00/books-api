-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user`
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    email         VARCHAR(200) NOT NULL,
    password      VARCHAR(200) NOT NULL,
    constraint user_email_uindex unique (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user`;
-- +goose StatementEnd
