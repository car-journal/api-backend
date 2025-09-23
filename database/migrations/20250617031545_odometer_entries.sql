-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS odometer_entries (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    car_id UUID NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
    odometer_reading FLOAT NOT NULL,
    reading_unit VARCHAR(255) NOT NULL CHECK (reading_unit IN ('km', 'mi')),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_odometer_entries_car_id ON odometer_entries(car_id);
CREATE INDEX idx_odometer_entries_odometer_reading ON odometer_entries(odometer_reading);
CREATE INDEX idx_odometer_entries_reading_unit ON odometer_entries(reading_unit);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS odometer_entries;
-- +goose StatementEnd
