-- +goose Up
-- +goose StatementBegin
CREATE TABLE movies(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    runtime INT NOT NULL,
    genres TEXT[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
-- +goose StatementEnd
