-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS maintenance_entries (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    car_id UUID NOT NULL,
    odometer_entry_id UUID NOT NULL REFERENCES odometer_entries(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES maintenance_categories(id) ON DELETE CASCADE,
    brand VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    performed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    notes VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS maintenance_entries;
-- +goose StatementEnd
