package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (dbClient *PGRepo) Swipe(swipe models.Swipe) error {
	db := dbClient.PostgresDBClient

	smt := `INSERT INTO swipes(swiper_id,swiped_id,action) VALUES ($1,$2,$3)`

	_, err := db.Exec(smt, swipe.SwiperID, swipe.SwipedID, swipe.Action)
	return err
}

func (dbClient *PGRepo) SwipeFeed(id uuid.UUID) (models.SwipeFeed, error) {
	db := dbClient.PostgresDBClient

	smt := `SELECT matches FROM matches_cache WHERE user_id = $1`

	var feed models.SwipeFeed
	err := db.QueryRow(smt, id).Scan(pq.Array(&feed.Matches))
	return feed, err
}
