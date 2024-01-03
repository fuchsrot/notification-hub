-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS channels
(
    id             UUID PRIMARY KEY,
    title          TEXT      NOT NULL,
    description    TEXT      NULL,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS channels;
-- +goose StatementEnd
