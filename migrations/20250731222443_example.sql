-- +goose Up
-- +goose StatementBegin
CREATE TABLE example_table (
    id BIGINT PRIMARY KEY,
    formed_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE example_table;
-- +goose StatementEnd
