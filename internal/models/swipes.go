package models

import "github.com/google/uuid"

type Swipe struct {
	SwiperID uuid.UUID `json:"swiper_id"`
	SwipedID uuid.UUID `json:"swiped_id"`
	Action   string    `json:"action"`
}

type SwipeReq struct {
	SwipedID uuid.UUID `json:"swiped_id"`
	Action   string    `json:"action"`
}
