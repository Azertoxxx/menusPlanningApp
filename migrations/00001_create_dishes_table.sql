-- +goose Up
CREATE TABLE IF NOT EXISTS dishes
(
    id             UUID PRIMARY KEY,
    name           TEXT      NOT NULL,
    image_url      TEXT      NULL,
    ingredients    TEXT      NULL,
    description    TEXT      NULL,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL,
    deleted_at     TIMESTAMP NULL
);

INSERT INTO dishes (id, name, created_at, updated_at) 
VALUES ('5b598744-b38c-4aaa-beb5-0ef44158ede5', 'test', '2021-01-01 00:00:00'::TIMESTAMP, '2021-01-01 00:00:00'::TIMESTAMP);

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dishes;
-- +goose StatementEnd
