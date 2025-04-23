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

type SwipeFeed struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	Gender     string    `json:"gender"`
	Birthday   string    `json:"birthday"`
	Location   string    `json:"location"`
	Religion   string    `json:"religion"`
	Department string    `json:"department"`
	Stream     string    `json:"stream"`
	Degree     string    `json:"degree"`
	PhotoUrl   string    `json:"photo_url"`
}
