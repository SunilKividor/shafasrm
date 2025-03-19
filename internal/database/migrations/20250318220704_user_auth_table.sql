-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    refresh_token VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, refresh_token)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth;
-- +goose StatementEnd
