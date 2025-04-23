package pgrepo

import (
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/google/uuid"
)

func (dbClient *PGRepo) Swipe(swipe models.Swipe) error {
	db := dbClient.PostgresDBClient

	smt := `INSERT INTO swipes(swiper_id,swiped_id,action) VALUES ($1,$2,$3)`

	_, err := db.Exec(smt, swipe.SwiperID, swipe.SwipedID, swipe.Action)
	return err
}

func (dbClient *PGRepo) SwipeFeed(id uuid.UUID) ([]models.SwipeFeed, error) {
	db := dbClient.PostgresDBClient

	smt := `
			SELECT 
			u.id,u.name,
			ud.phone,ud.gender,ud.birthday,ud.location,
			ud.religion,ud.department,ud.stream,ud.degree,
			up.photo_key 
			FROM users u 
			INNER JOIN user_details ud ON ud.user_id = u.id 
			INNER JOIN user_photos up ON up.user_id = u.id AND up.is_primary = true 
			WHERE u.id = ANY (
				SELECT UNNEST(matches) FROM swipes_feed WHERE user_id = $1
			)
			`

	var feeds []models.SwipeFeed
	rows, err := db.Query(smt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var feed models.SwipeFeed
		err := rows.Scan(
			&feed.Id, &feed.Name,
			&feed.Phone, &feed.Gender, &feed.Birthday, &feed.Location,
			&feed.Religion, &feed.Department, &feed.Stream, &feed.Degree,
			&feed.PhotoUrl,
		)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
