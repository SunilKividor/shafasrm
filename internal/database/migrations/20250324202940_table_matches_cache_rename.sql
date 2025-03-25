-- +goose Up
-- +goose StatementBegin
ALTER TABLE matches_cache RENAME TO swipes_feed;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE swipes_feed RENAME TO matches_cache;
-- +goose StatementEnd
