-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        varchar(255) NOT NULL,
    pwd         varchar(255) NOT NULL,
    created_at  timestamp NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
