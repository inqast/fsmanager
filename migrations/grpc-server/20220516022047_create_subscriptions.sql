-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
       id                      bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
       chat_id                 bigint,
       service_name            varchar(255) NOT NULL,
       capacity                integer NOT NULL,
       price_in_centi_units    integer NOT NULL,
       payment_day             integer NOT NULL,
       created_at              timestamp NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subscriptions;
-- +goose StatementEnd
