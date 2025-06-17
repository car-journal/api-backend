-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profiles (
    user_id UUID NOT NULL UNIQUE,
    gender boolean,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    date_of_birth TIMESTAMPTZ,
    picture_url TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE user_profiles
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profiles;
-- +goose StatementEnd