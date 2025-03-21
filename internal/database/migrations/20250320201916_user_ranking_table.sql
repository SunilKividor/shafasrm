-- +goose Up
-- +goose StatementBegin
CREATE TABLE ranking (
    rank int,
    user_id UUID PRIMARY KEY,
    points int NOT NULL DEFAULT 0,
    status VARCHAR(20) CHECK (status IN ('ranked','unranked')),
    CONSTRAINT unique_rank UNIQUE (rank),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)

    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP Table ranking;
-- +goose StatementEnd
