-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
       id                      bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
       owner_id                bigint REFERENCES users (id) ON DELETE CASCADE NOT NULL,
       service_name            varchar(255) NOT NULL,
       capacity                integer NOT NULL,
       price_in_centi_units    integer NOT NULL,
       payment_date            timestamp NOT NULL,
       created_at              timestamp NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscriptions;
-- +goose StatementEnd
