-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS maintenance_categories (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description varchar(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS maintenance_categories;
-- +goose StatementEnd
