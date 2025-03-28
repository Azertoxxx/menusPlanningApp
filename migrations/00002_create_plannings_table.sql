-- +goose Up
CREATE TABLE IF NOT EXISTS plannings
(
    id             UUID PRIMARY KEY,
    slot           TEXT      NOT NULL,
    dish_id        UUID REFERENCES dishes(id) ON DELETE CASCADE
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS plannings;
-- +goose StatementEnd
