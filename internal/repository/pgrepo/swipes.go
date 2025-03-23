package pgrepo

import "github.com/SunilKividor/shafasrm/internal/models"

func (dbClient *PGRepo) AddSwipeAction(swipe models.Swipe) error {
	db := dbClient.PostgresDBClient

	smt := `INSERT INTO swipes(swiper_id,swiped_id,action) VALUES ($1,$2,$3)`

	_, err := db.Exec(smt, swipe.SwiperID, swipe.SwipedID, swipe.Action)
	return err
}
