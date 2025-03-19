-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN password varchar(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN password;
-- +goose StatementEnd
