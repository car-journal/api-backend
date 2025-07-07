-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fuel_entries (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    car_id UUID NOT NULL,
    fuel_type VARCHAR(255) NOT NULL,
    fuel_brand VARCHAR(255) NOT NULL,
    fuel_name VARCHAR(255) NOT NULL,
    fuel_price FLOAT NOT NULL,
    fuel_unit VARCHAR(255) NOT NULL,
    distance_traveled FLOAT NOT NULL,
    volume_filled FLOAT NOT NULL,
    total_price FLOAT NOT NULL,
    fuel_consumption_rate FLOAT NOT NULL,
    notes VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE fuel_entries
ADD CONSTRAINT fk_cars FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fuel_entries;
-- +goose StatementEnd
