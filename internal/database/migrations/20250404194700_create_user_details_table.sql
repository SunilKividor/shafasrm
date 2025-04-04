-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_details (
    id UUID DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    gender varchar(255) NOT NULL,
    birthday date NOT NULL,
    location varchar(255) NOT NULL,
    religion varchar(255) NOT NULL,
    department varchar(255) NOT NULL,
    stream varchar(255) NOT NULL,
    degree varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
