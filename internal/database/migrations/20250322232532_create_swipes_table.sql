-- +goose Up
-- +goose StatementBegin
CREATE TABLE swipes (
    swiper_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    swiped_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    action VARCHAR(8),
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (swiper_id,swiped_id),
    CONSTRAINT self_swipe CHECK (swiper_id != swiped_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE swipes;
-- +goose StatementEnd
