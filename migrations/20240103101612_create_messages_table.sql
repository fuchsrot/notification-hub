-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages
(
    id             UUID PRIMARY KEY,
    title          TEXT      NOT NULL,
    content        TEXT      NOT NULL,
    type           TEXT      NOT NULL,
    status         TEXT      NOT NULL,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL,
    channel_id     TEXT      NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
