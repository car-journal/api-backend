-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cars (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    user_id UUID NOT NULL,
    brand varchar(255) NOT NULL,
    model varchar(255) NOT NULL,
    manufacture_year SMALLINT,
    cylinder_capacity SMALLINT,
    vehicle_identity_number VARCHAR(255),
    engine_number VARCHAR(255),
    color varchar(255),
    fuel_type VARCHAR(255) NOT NULL,
    registration_year SMALLINT,
    vehicle_ownership_document_number VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE cars
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
-- +goose StatementEnd
