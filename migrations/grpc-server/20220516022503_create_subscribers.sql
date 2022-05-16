-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscribers (
                       id               bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                       user_id          bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL,
                       subscription_id  bigint REFERENCES subscriptions (id) ON DELETE CASCADE NOT NULL,
                       is_paid          boolean NOT NULL,
                       is_owner         boolean NOT NULL,
                       created_at       timestamp NOT NULL,
                       UNIQUE(user_id, subscription_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscribers;
-- +goose StatementEnd
