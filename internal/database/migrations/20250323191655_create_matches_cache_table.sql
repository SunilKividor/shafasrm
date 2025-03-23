-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches_cache (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    matches UUID[],
    updated_at TIMESTAMP,
    PRIMARY KEY(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches_cache;
-- +goose StatementEnd
