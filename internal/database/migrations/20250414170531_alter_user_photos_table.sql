-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_photos DROP CONSTRAINT valid_url;
ALTER TABLE user_photos RENAME COLUMN photo_url TO photo_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_photos ADD CONSTRAINT valid_url CHECK (photo_url ~ '^(https?|ftp)://[^\s/$.?#].[^\s]*$');
ALTER TABLE user_photos RENAME COLUMN photo_key TO photo_url;
-- +goose StatementEnd
