-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_details DROP COLUMN name;
ALTER TABLE user_details DROP COLUMN username;
ALTER TABLE user_details DROP COLUMN email;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_details ADD COLUMN name;
ALTER TABLE user_details ADD COLUMN username;
ALTER TABLE user_details ADD COLUMN email;
-- +goose StatementEnd
