-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fuels (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    brand VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS fuels_brand_name_unique
ON fuels (brand, name)
WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fuels;
-- +goose StatementEnd
