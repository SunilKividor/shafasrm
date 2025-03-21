-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_profile (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    party_move INT,          -- 1-4
    wild_place INT,         -- 1-4
    zombie_days INT,        -- 1-4
    guilty_song TEXT,
    ghost_reason INT,       -- 1-4
    shots_confess INT,      -- 1-4
    late_food INT,         -- 1-4
    chaotic_love TEXT,
    breakup_power INT,     -- 1-4
    weak_spot TEXT,
    flirt_rating INT,      -- 1-10
    into_signal INT,       -- 1-4
    dumb_line TEXT,
    trouble_sign INT,      -- 1-4
    campus_rumor TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_profile
-- +goose StatementEnd
