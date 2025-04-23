-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_profile
DROP COLUMN wild_place ,         -- 1-4
DROP COLUMN zombie_days ,        -- 1-4
DROP COLUMN ghost_reason ,       -- 1-4
DROP COLUMN late_food ,         -- 1-4
DROP COLUMN breakup_power ,     -- 1-4
DROP COLUMN weak_spot ,
DROP COLUMN into_signal ,       -- 1-4
DROP COLUMN dumb_line ,
DROP COLUMN trouble_sign ,      -- 1-4
DROP COLUMN campus_rumor;

ALTER TABLE user_profile
ALTER COLUMN guilty_song TYPE INT USING flirt_rating::INTEGER,
ALTER COLUMN chaotic_love TYPE INT USING flirt_rating::INTEGER;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_profile
ADD COLUMN wild_place INT,         -- 1-4
ADD COLUMN zombie_days INT,        -- 1-4
ADD COLUMN ghost_reason INT,       -- 1-4
ADD COLUMN late_food INT,         -- 1-4
ADD COLUMN breakup_power INT,     -- 1-4
ADD COLUMN weak_spot TEXT,
ADD COLUMN into_signal INT,       -- 1-4
ADD COLUMN dumb_line TEXT,
ADD COLUMN trouble_sign INT,      -- 1-4
ADD COLUMN campus_rumor TEXT;
-- +goose StatementEnd
