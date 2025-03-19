-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD CONSTRAINT unique_user UNIQUE (username, email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT unique_user;
-- +goose StatementEnd
