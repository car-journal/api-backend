-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS oauth_access_tokens (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    user_id UUID NOT NULL,
    revoked boolean NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE oauth_access_tokens
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS oauth_access_tokens;
-- +goose StatementEnd
