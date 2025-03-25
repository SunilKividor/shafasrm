-- +goose Up
-- +goose StatementBegin
CREATE TABLE matches (
    id UUID DEFAULT uuid_generate_v4(),
    user_id_1 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    user_id_2 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(id),
    CONSTRAINT no_self_match CHECK(user_id_1 != user_id_2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE matches;
-- +goose StatementEnd
